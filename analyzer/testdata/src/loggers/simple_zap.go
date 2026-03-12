package loggers

import "go.uber.org/zap"

func sugar() {
	s := zap.S()
	s.Info("Invalid data!")
}

func classic() {
	l := zap.L()
	l.Info("")
	l.Warn
}
