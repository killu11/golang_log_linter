package analyzer

import (
	"go/ast"
	"go/types"

	"golang.org/x/tools/go/analysis"
)

func isSugar(pass *analysis.Pass, ident *ast.Ident) bool {
	t := pass.TypesInfo.TypeOf(ident)

	if t == nil {
		return false
	}
	if ptr, ok := t.(*types.Pointer); ok {
		return ptr.Elem().String() == "go.uber.org/zap.SugaredLogger"
	}
	return false
}

func isClassic(pass *analysis.Pass, ident *ast.Ident) bool {
	t := pass.TypesInfo.TypeOf(ident)

	if t == nil {
		return false
	}
	if ptr, ok := t.(*types.Pointer); ok {
		return ptr.Elem().String() == "go.uber.org/zap.Logger"
	}
	return false
}
