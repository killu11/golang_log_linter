package specSymbols

import (
	"go.uber.org/zap"
)

func sc() {
	//Та же логика, что и с отдельными функциями из пакета работает

	s := zap.L()

	s.Warn("test method with suffix w", zap.String("😭", "sad smile")) // want "log msg shouldn't have specifical symbols"
	s.Info("!@#5453")                                                 // want "log msg shouldn't have specifical symbols and emojis"
	s.Debug("hi surovo :)")                                           // want "log msg shouldn't have specifical symbols and emojis"
}
