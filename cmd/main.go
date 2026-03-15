package main

import (
	"github.com/killu11/golang_log_linter/analyzer"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(analyzer.LogAnalyzer)
}
