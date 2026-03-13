package specSymbols

import (
	"context"
	"log/slog"
	"os"
)

func sc() {
	//Та же логика, что и с отдельными функциями из пакета работает

	l := slog.New(slog.NewTextHandler(os.Stdout, nil))

	l.DebugContext(context.Background(), "zasoV1!2_-")  // want "log msg shouldn't have specifical symbols and emojis"
	l.InfoContext(context.Background(), "!@#5453")      // want "log msg shouldn't have specifical symbols and emojis"
	l.DebugContext(context.Background(), "surovo :)))") // want "log msg shouldn't have specifical symbols and emojis"
}
