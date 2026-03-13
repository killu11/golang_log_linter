package good

import "go.uber.org/zap"

func foo() {
	s := zap.S()
	s.Warn("warn message", "bad request")
	s.Infof("info message %v", "good request2026")

}

func bar() {
	l := zap.L()
	l.Info("test", zap.Int("normal value", 1)) //
}
