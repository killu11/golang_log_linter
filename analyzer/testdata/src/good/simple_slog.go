package good

import (
	"context"
	"log/slog"
	"os"
)

func log() {

	slog.Debug("test", "message")

	l := slog.New(slog.NewTextHandler(os.Stdout, nil))

	l.Info("test logger", "my message32313")

	l.DebugContext(context.Background(), "zasoV1")

	const somedata = "dewqeweq"
	l.WarnContext(context.Background(), "data:"+somedata)

	l.InfoContext(context.Background(), "hi", "user", "Egor")
}
