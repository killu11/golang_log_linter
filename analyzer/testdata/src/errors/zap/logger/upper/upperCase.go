package upper

import (
	"go.uber.org/zap"
)

func u() {
	s := zap.L()
	s.Debug("Test") // want "first symbol of log msg shouldn't be in upper case"
	s.Info("Test2") // want "first symbol of log msg shouldn't be in upper case"
	s.Error("Show") // want "first symbol of log msg shouldn't be in upper case"
}
