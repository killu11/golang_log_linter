package analyzer

import (
	"go/ast"
	"go/types"

	"golang.org/x/tools/go/analysis"
)

func isSlogPkgSelector(pass *analysis.Pass, ident *ast.Ident) bool {
	obj := pass.TypesInfo.ObjectOf(ident)
	if obj == nil {
		return false
	}

	if pkgName, ok := obj.(*types.PkgName); ok {
		return pkgName.Imported().Path() == "log/slog"
	}

	return false
}

func isSlogLoggerSelector(pass *analysis.Pass, ident *ast.Ident) bool {
	t := pass.TypesInfo.TypeOf(ident)
	if t == nil {
		return false
	}

	if ptr, ok := t.(*types.Pointer); ok {
		return ptr.Elem().String() == "log/slog.Logger"
	}

	return false
}
