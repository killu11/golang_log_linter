package analyzer

import (
	"fmt"
	"go/ast"
	"go/token"
	"linter/pkg"
	"log"
	"slices"
	"strconv"
	"strings"

	"golang.org/x/tools/go/analysis"
)

var LogAnalyzer = &analysis.Analyzer{
	Name: "invalid_log",
	Doc:  "log messages should be start with lower case",
	Run:  run,
}

func run(pass *analysis.Pass) (any, error) {
	loadOnce.Do(loadBanWords)
	log.Println(banWords)
	for _, file := range pass.Files {
		ast.Inspect(file, func(node ast.Node) bool {
			call, ok := node.(*ast.CallExpr)
			if !ok {
				return true
			}

			analyze(pass, call)

			return true
		})
	}
	return nil, nil
}

func analyze(pass *analysis.Pass, call *ast.CallExpr) {
	sel, ok := call.Fun.(*ast.SelectorExpr)
	if !ok {
		return
	}

	id, ok := sel.X.(*ast.Ident)
	if !ok {
		return
	}

	pt, ok := valid(pass, id)
	if !ok {
		return
	}

	if len(call.Args) == 0 {
		return
	}

	if pt.IsSlog() {
		switch {
		case slices.Contains(Cmd, sel.Sel.Name):
			startInspect(pass, call.Args, 0)
		case slices.Contains(ContextCmd, sel.Sel.Name):
			startInspect(pass, call.Args, 1)
		default:
			return
		}
	}

	if pt.IsZapSugar() && containsSubCmd(sel.Sel.Name) {
		switch {
		case strings.HasSuffix(sel.Sel.Name, "f"):
		case strings.HasSuffix(sel.Sel.Name, "w"):
		case strings.HasSuffix(sel.Sel.Name, "ln"):
		default:
			startInspect(pass, call.Args, 0)
		}
	}
	// TODO Дописать логику анализатора с классическим логгером и разновидносятми
	if pt.IsZapClassic() && containsSubCmd(sel.Sel.Name) {

	}

	return
}

func startInspect(pass *analysis.Pass, args []ast.Expr, msgPos int) {
	for i, arg := range args {
		if msgPos > i {
			continue
		}

		if i == msgPos {
			if err := checkMsg(pass, arg); err != nil {
				return
			}
			continue
		}
		checkArg(pass, arg)
	}
}

func checkMsg(pass *analysis.Pass, expr ast.Expr) error {
	msg, err := valueFromExpr(pass, expr)
	if err != nil {
		return err
	}

	if err = pkg.IsEmpty(msg); err != nil {
		return err
	}

	if err = pkg.UpperSymbol(msg); err != nil {
		pass.Reportf(expr.Pos(), "%v", err)
	}

	if err = pkg.OnlyLatinAndNumSymbols(msg); err != nil {
		pass.Reportf(expr.Pos(), "%v", err)
	}

	if err = pkg.SpecSymbols(msg); err != nil {
		pass.Reportf(expr.Pos(), "%v", err)
	}

	return nil
}

// checkArg проверяет аргументы для slog.Logger
func checkArg(pass *analysis.Pass, expr ast.Expr) {
	val, err := valueFromExpr(pass, expr)
	if err != nil {
		return
	}

	if isSensitive(val) {
		pass.Reportf(expr.Pos(), "args shouldn't be had sensitive data: %v", val)
		return
	}

	if err = pkg.OnlyLatinAndNumSymbols(val); err != nil {
		pass.Reportf(expr.Pos(), "%v", err)
	}
	if err = pkg.SpecSymbols(val); err != nil {
		pass.Reportf(expr.Pos(), "%v", err)
	}
}

// checkField проверяет ключи-значения для zap.SugarLogger
//func checkField(pass *analysis.Pass, exp ast.Expr) error {
//
//}

func valid(pass *analysis.Pass, ident *ast.Ident) (PkgType, bool) {
	if isSlogPkgSelector(pass, ident) || isSlogLoggerSelector(pass, ident) {
		return "log/slog", true
	}

	if isClassic(pass, ident) {
		return "zap/classic", true
	}

	if isSugar(pass, ident) {
		return "zap/sugar", true
	}

	return "", false
}

func valueFromExpr(pass *analysis.Pass, expr ast.Expr) (string, error) {
	switch v := expr.(type) {
	case *ast.BasicLit:
		s, err := strconv.Unquote(v.Value)
		if err != nil {
			return "", fmt.Errorf("unquote literal: %w", err)
		}
		return s, nil
	// Рассудил, что переменная - это всегда аргумент, описание (msg) идет всегда литералом, поэтому отдаю название переменной
	// для проверки на чувствительные данные по ключевым словам
	case *ast.Ident:
		obj := pass.TypesInfo.ObjectOf(v)
		if obj == nil {
			return "", fmt.Errorf("nil object")
		}

		// При конкатенации чувствительных данных линтер обрывает проверку на этапе нахождения этого кейса,
		// После того очистки чувствительных данных из лога остальные кейсы должны успешно выполнится
		// Смотреть в testdata/src/sensitive.go

		//Неправильно смешивать логику, но увы без костылей не обошлось :)

		if isSensitive(obj.Name()) {
			pass.Reportf(expr.Pos(), "args shouldn't be had sensitive data: %v", obj.Name())
			return "", fmt.Errorf("find sensitive data name: %s", obj.Name())
		}
		return obj.Name(), nil

	case *ast.BinaryExpr:
		if v.Op != token.ADD {
			return "", fmt.Errorf("invalid operation")
		}

		left, err := valueFromExpr(pass, v.X)
		if err != nil {
			return "", err
		}

		right, err := valueFromExpr(pass, v.Y)
		if err != nil {
			return "", err
		}
		return left + right, nil

	default:
		return "", fmt.Errorf("unknown type")
	}
}
