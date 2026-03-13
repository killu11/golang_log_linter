package analyzer

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"
)

// Установите сюда свой путь до банвордов
var bwpath = "/home/egorsslv/Desktop/golang_log_linter/banwords.txt"

func TestGoodCases(t *testing.T) {

	LogAnalyzer.Flags.Set("path", bwpath)
	analysistest.Run(t, analysistest.TestData(), LogAnalyzer, "good")
}

func TestSlogUpperCase(t *testing.T) {

	LogAnalyzer.Flags.Set("path", bwpath)
	analysistest.Run(t, analysistest.TestData(), LogAnalyzer, "errors/slog/upper")
}

func TestSlogEnglishCase(t *testing.T) {

	LogAnalyzer.Flags.Set("path", bwpath)
	analysistest.Run(t, analysistest.TestData(), LogAnalyzer, "errors/slog/english")
}

func TestSlogSpecSymbolsCase(t *testing.T) {

	LogAnalyzer.Flags.Set("path", bwpath)
	analysistest.Run(t, analysistest.TestData(), LogAnalyzer, "errors/slog/specSymbols")
}

func TestSlogSensitiveDataCase(t *testing.T) {

	LogAnalyzer.Flags.Set("path", bwpath)
	analysistest.Run(t, analysistest.TestData(), LogAnalyzer, "errors/slog/sensitive")
}

func TestZapSugarUpperCase(t *testing.T) {

	LogAnalyzer.Flags.Set("path", bwpath)
	analysistest.Run(t, analysistest.TestData(), LogAnalyzer, "errors/zap/sugar/upper")
}
func TestZapSugarEnglishCase(t *testing.T) {

	LogAnalyzer.Flags.Set("path", bwpath)
	analysistest.Run(t, analysistest.TestData(), LogAnalyzer, "errors/zap/sugar/english")
}
func TestZapSugarSpecSymbolsCase(t *testing.T) {

	LogAnalyzer.Flags.Set("path", bwpath)
	analysistest.Run(t, analysistest.TestData(), LogAnalyzer, "errors/zap/sugar/specSymbols")
}
func TestZapSugarSensitiveDataCase(t *testing.T) {

	LogAnalyzer.Flags.Set("path", bwpath)
	analysistest.Run(t, analysistest.TestData(), LogAnalyzer, "errors/zap/sugar/sensitive")
}

func TestZapLoggerUpperCase(t *testing.T) {

	LogAnalyzer.Flags.Set("path", bwpath)
	analysistest.Run(t, analysistest.TestData(), LogAnalyzer, "errors/zap/logger/upper")
}
func TestZapLoggerEnglishCase(t *testing.T) {

	LogAnalyzer.Flags.Set("path", bwpath)
	analysistest.Run(t, analysistest.TestData(), LogAnalyzer, "errors/zap/logger/english")
}
func TestZapLoggerSpecSymbolsCase(t *testing.T) {

	LogAnalyzer.Flags.Set("path", bwpath)
	analysistest.Run(t, analysistest.TestData(), LogAnalyzer, "errors/zap/logger/specSymbols")
}
func TestZapLoggerSensitiveDataCase(t *testing.T) {

	LogAnalyzer.Flags.Set("path", bwpath)
	analysistest.Run(t, analysistest.TestData(), LogAnalyzer, "errors/zap/logger/sensitive")
}
