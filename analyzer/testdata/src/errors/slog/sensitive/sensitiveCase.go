package sensitive

import (
	"context"
	"log/slog"
	"os"
)

func sc() {
	//Та же логика, что и с отдельными функциями из пакета работает

	l := slog.New(slog.NewTextHandler(os.Stdout, nil))

	l.Error("some msg", "creds", "egor", "1123") // want "args shouldn't be had sensitive data: creds"
	apikey := "stg9iqte99ags9g9as9g"
	l.InfoContext(context.Background(), "show api key", apikey) // want  "args shouldn't be had sensitive data: apikey"
}
