package sensitive

import (
	"go.uber.org/zap"
)

func sc() {
	s := zap.L()
	s.Error("some msg", zap.String("Creds", "sddpsdpsfo fd[f[df[d")) // want "args shouldn't be had sensitive data: Creds"
	apikey := "stg9iqte99ags9g9as9g"
	s.Info("show api key", zap.String("data", apikey)) // want  "args shouldn't be had sensitive data: apikey"
}
