package english

import (
	"go.uber.org/zap"
)

func e() {
	s := zap.S()

	s.Debug("россия священная наша держава") // want "log msg must be in english only"
	s.Info("аԥсуа бызшәа ара иҟоуп")         // want "log msg must be in english only"
	s.Error("lógica española")               // want "log msg must be in english only"
	s.Error("eng", "рус")                    // want "log msg must be in english only"
}
