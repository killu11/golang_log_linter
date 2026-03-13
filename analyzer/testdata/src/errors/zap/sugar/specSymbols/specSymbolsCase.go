package specSymbols

import (
	"go.uber.org/zap"
)

func sc() {
	//Та же логика, что и с отдельными функциями из пакета работает

	s := zap.S()

	s.Warnw("test method with suffix w", "😭", "sad smile") // want "log msg shouldn't have specifical symbols"
	s.Info("somedata", "!@#5453")                          // want "log msg shouldn't have specifical symbols and emojis"
	s.Debugf("hi %v", "surovo :)))")                       // want "log msg shouldn't have specifical symbols and emojis"
}
