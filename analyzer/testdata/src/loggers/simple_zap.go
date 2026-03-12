package loggers

import "go.uber.org/zap"

func sugar() {
	s := zap.S()
	s.Info("Invalid data!") // want "log msg shouldn't have specifical symbols and emojis" "first symbol of log msg shouldn't be in upper case"
	s.Infof("hello %v", "apikey")
}

func classic() {
	l := zap.L()
	l.Info("")
}
