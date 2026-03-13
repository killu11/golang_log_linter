package upper

import (
	"go.uber.org/zap"
)

func u() {
	s := zap.S()
	s.Debug("Test")            // want "first symbol of log msg shouldn't be in upper case"
	s.Infoln("Test2")          // want "first symbol of log msg shouldn't be in upper case"
	s.Errorw("Show", "2", "3") // want "first symbol of log msg shouldn't be in upper case"
}
