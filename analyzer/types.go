package analyzer

import (
	"log"
	"os"
	"slices"
	"strings"
	"sync"
)

const banWordsPath = "../banwords.txt"
const bwTemplate = "apikey,env,password,authtoken,creds,credentials"

var Cmd = []string{"Info", "Debug", "Warn", "Error"}
var ContextCmd = []string{"InfoContext", "DebugContext", "WarnContext", "ErrorContext"}

var (
	banWords []string
	loadOnce sync.Once
)

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
	if _, err := os.Stat(banWordsPath); err != nil {
		f, err := os.OpenFile(banWordsPath, os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Println("create banwords.txt:", err)
			return
		}

		defer f.Close()

		if _, err := f.WriteString(bwTemplate); err != nil {
			log.Println("write banwords template:", err)
			return
		}
	}

	data, err := os.ReadFile(banWordsPath)
	if err != nil {
		log.Fatalln("read banwords.txt:", err)
	}

	content := strings.TrimSpace(string(data))
	if content == "" {
		banWords = []string{}
		return
	}

	banWords = strings.Split(content, ",")

	for i, word := range banWords {
		banWords[i] = strings.TrimSpace(word)
	}

}
func init() {

}
