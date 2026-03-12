package loggers

import (
	"context"
	"log/slog"
	"os"
)

func log() {
	// Тест на спецсимволы:
	slog.Debug("start linter!") // want "log msg shouldn't have specifical symbols and emojis"

	l := slog.New(slog.NewTextHandler(os.Stdout, nil))
	// Тест на Upper case:
	l.Info("Test1", "234") // want "first symbol of log msg shouldn't be in upper case"

	// Тест на проверку множества аргументов и чувствительные данные (смотреть в файлы types.go и banwords.txt):
	l.InfoContext(context.Background(), "zasoV1!2_-", "creds") // want "log msg shouldn't have specifical symbols and emojis" "args shouldn't be had sensitive data: creds"
}
