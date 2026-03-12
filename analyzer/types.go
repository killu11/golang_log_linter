package analyzer

import (
	"bufio"
	"io"
	"log"
	"os"
	"slices"
	"strings"
)

const banWordsPath = "../banwords.txt"
const bwTemplate = `apikey, env, password, authtoken, creds, credentials`

var Cmd = []string{"Info", "Debug", "Warn", "Error"}
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
func init() {
	if _, err := os.Stat(banWordsPath); err != nil {
		f, err := os.Create(banWordsPath)
		if err != nil {
			log.Fatalln("Create banwords.txt:", err)
		}

		if _, err := f.Write([]byte(bwTemplate)); err != nil {
			log.Fatalln("Write banwords template:", err)
		}
		return
	}

	f, err := os.Open(banWordsPath)
	if err != nil {
		log.Fatalln("Open banwords.txt:", err)
	}
	rd := bufio.NewReader(f)
	s, err := rd.ReadString('\n')

	if err != io.EOF {
		log.Fatalln(err)
	}
	banWords = strings.Split(s, ",")
}
