package english

import (
	"log/slog"
	"os"
)

var l = slog.New(slog.NewTextHandler(os.Stdout, nil))

func e() {
	// Test Upper case:
	slog.Debug("россия священная наша держава") // want "log msg must be in english only"
	slog.Info("аԥсуа бызшәа ара иҟоуп")         // want "log msg must be in english only"
	slog.Error("lógica española")               // want "log msg must be in english only"

	l.Debug("россия священная наша держава") // want "log msg must be in english only"
	l.Info("аԥсуа бызшәа ара иҟоуп")         // want "log msg must be in english only"
	l.Error("lógica española")               // want "log msg must be in english only"
}
