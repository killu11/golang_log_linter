package analyzer

import (
	"flag"
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

var bwPath string

var LogAnalyzer = &analysis.Analyzer{
	Name: "log_checker",
	Doc:  "log messages should be start with lower case",
	Run:  run,
	Flags: func() flag.FlagSet {
		fs := flag.NewFlagSet("fs", flag.ExitOnError)
		fs.StringVar(&bwPath, "path", "", "usage to set banwords.txt file")
		return *fs
	}(),
}

func run(pass *analysis.Pass) (any, error) {
	loadBanWords()

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
func init() {
	LogAnalyzer.Flags.String("p", "", "set path to banword file, use `.txt` format")
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
			startInspect(pass, call.Args, 0, false)
		case slices.Contains(ContextCmd, sel.Sel.Name):
			startInspect(pass, call.Args, 1, false)
		default:
			return
		}
	}

	if pt.IsZapSugar() && containsSubCmd(sel.Sel.Name) {
		switch {
		case strings.HasSuffix(sel.Sel.Name, "f"):
			startInspect(pass, call.Args, 0, true)
		default:
			startInspect(pass, call.Args, 0, false)
		}
	}

	if pt.IsZapClassic() && containsSubCmd(sel.Sel.Name) {
		for i, arg := range call.Args {
			if i == 0 {
				if err := checkMsg(pass, arg, false); err != nil {
					return
				}
				continue
			}

			checkField(pass, arg)
		}
	}

	return
}

// startInspect - анализирует аргументы внутри функции, проверяя ее на правила
// fsStatus - статус форматированной строки, он необходим, чтобы проверка на спецсимволы не выдавала репорт
func startInspect(pass *analysis.Pass, args []ast.Expr, msgPos int, fsStatus bool) {
	for i, arg := range args {
		if msgPos > i {
			continue
		}

		if i == msgPos {
			if err := checkMsg(pass, arg, fsStatus); err != nil {
				return
			}
			continue
		}
		checkArg(pass, arg)
	}
}

func checkMsg(pass *analysis.Pass, expr ast.Expr, fsStatus bool) error {
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

	if err = pkg.SpecSymbols(msg, fsStatus); err != nil {
		pass.Reportf(expr.Pos(), "%v", err)
	}

	return nil
}

// checkArg проверяет аргументы для slog.Logger и zap.SugarLogger
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
	if err = pkg.SpecSymbols(val, false); err != nil {
		pass.Reportf(expr.Pos(), "%v", err)
	}
}

// checkField проверяет ключи-значения для zap.Logger
func checkField(pass *analysis.Pass, expr ast.Expr) {
	var val string
	call, ok := expr.(*ast.CallExpr)

	if !ok {
		id, ok := expr.(*ast.Ident)
		if !ok {
			log.Println("unknown expr parse field")
			return
		}

		obj := pass.TypesInfo.ObjectOf(id)
		if obj == nil {
			log.Println("nil obj")
			return
		}

		val = obj.Name()
		if isSensitive(val) {
			pass.Reportf(expr.Pos(), "args shouldn't be had sensitive data: %v", val)
		}
		return
	}

	tv := pass.TypesInfo.Types[call]
	if tv.Type == nil {
		return
	}

	if tv.Type.String() != "go.uber.org/zap.Field" {
		return
	}
	for _, arg := range call.Args {
		val, err := valueFromExpr(pass, arg)
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
		if err = pkg.SpecSymbols(val, false); err != nil {
			pass.Reportf(expr.Pos(), "%v", err)
		}
	}

}

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
