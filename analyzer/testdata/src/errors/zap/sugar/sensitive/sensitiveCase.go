package sensitive

import (
	"go.uber.org/zap"
)

func sc() {
	s := zap.S()
	s.Error("some msg", "creds", "egor", "1123") // want "args shouldn't be had sensitive data: creds"
	apikey := "stg9iqte99ags9g9as9g"
	s.Infof("show api key %v", apikey) // want  "args shouldn't be had sensitive data: apikey"
}
