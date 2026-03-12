package main

import (
	"linter/analyzer"

	"golang.org/x/tools/go/analysis"
)

func New(conf any) ([]*analysis.Analyzer, error) {
	return []*analysis.Analyzer{analyzer.LogAnalyzer}, nil
}
