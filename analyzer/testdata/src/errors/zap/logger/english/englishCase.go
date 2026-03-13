package english

import (
	"go.uber.org/zap"
)

func e() {
	s := zap.L()
	s.Debug("россия священная наша держава")        // want "log msg must be in english only"
	s.Info("аԥсуа бызшәа ара иҟоуп")                // want "log msg must be in english only"
	s.Error("lógica española", zap.Int("число", 9)) // want "log msg must be in english only" "log msg must be in english only"
	s.Error("eng", zap.String("rus", "россия"))     // want "log msg must be in english only"
}
