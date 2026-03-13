package upper

import (
	"log/slog"
	"os"
)

func u() {
	// Test Upper case:
	slog.Debug("Test") // want "first symbol of log msg shouldn't be in upper case"
	slog.Info("Test2") // want "first symbol of log msg shouldn't be in upper case"
	slog.Error("Show") // want "first symbol of log msg shouldn't be in upper case"

	l := slog.New(slog.NewTextHandler(os.Stdout, nil))

	l.Debug("Test") // want "first symbol of log msg shouldn't be in upper case"
	l.Info("Test2") // want "first symbol of log msg shouldn't be in upper case"
	l.Error("Show") // want "first symbol of log msg shouldn't be in upper case"
}
