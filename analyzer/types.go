package analyzer

import (
	"slices"
	"strings"
)

var Cmd = []string{"Info", "Debug", "Warn", "Error", "Fatal", "Panic"}
var ContextCmd = []string{"InfoContext", "DebugContext", "WarnContext", "ErrorContext"}

var banWords []string

type PkgType string

func (t PkgType) IsZapClassic() bool {
	return strings.EqualFold(string(t), "zap/classic")
}
func (t PkgType) IsZapSugar() bool {
	return strings.EqualFold(string(t), "zap/sugar")
}
func (t PkgType) IsSlog() bool {
	return strings.EqualFold(string(t), "log/slog")
}

func isSensitive(name string) bool {
	return slices.Contains(banWords, strings.ToLower(name))
}
func containsSubCmd(command string) bool {
	for _, c := range Cmd {
		if strings.Contains(command, c) {
			return true
		}
	}
	return false
}

func loadBanWords() {
	// Здесь должна была быть подгрузка бан вордов из дока, но увы, в комбинации с custom golangci-lint
	// Нет прямой возможности указывать флаги (в моей реализации через флаг передавался бы путь к бан вордам)
	// Еще был вариант с конфигурационным файлом, но просто не хватило времени разобраться :)
	const bwTemplate = "apikey,env,password,authtoken,creds,credentials"

	bw := strings.Split(bwTemplate, ",")
	for _, word := range bw {
		banWords = append(banWords, strings.TrimSpace(word))
	}
}
