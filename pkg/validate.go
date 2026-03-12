package pkg

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

func IsEmpty(msg string) error {
	if utf8.RuneCountInString(msg) == 0 {
		return fmt.Errorf("empty log")
	}

	return nil
}

func UpperSymbol(msg string) error {
	first := []rune(msg)[0]
	if unicode.IsUpper(first) {
		return fmt.Errorf("first symbol of log msg shouldn't be in upper case")
	}

	return nil
}

func OnlyLatinAndNumSymbols(arg string) error {
	for _, r := range arg {
		switch {
		case (r >= 'A' && r <= 'z') || (r >= '0' && r <= '9'):
			continue
		case r == ' ':
			continue
		case unicode.In(r, unicode.Symbol, unicode.Sc, unicode.Sm, unicode.So) ||
			unicode.In(r, unicode.Punct, unicode.Pc, unicode.Pd, unicode.Po):
			continue
		default:
			return fmt.Errorf("log msg must be in english only")
		}
	}
	return nil
}

func SpecSymbols(arg string) error {
	for _, r := range arg {
		if unicode.In(r, unicode.Symbol, unicode.Sc, unicode.Sm, unicode.So) ||
			unicode.In(r, unicode.Punct, unicode.Pc, unicode.Pd, unicode.Po) {
			return fmt.Errorf("log msg shouldn't have specifical symbols and emojis")
		}
	}
	return nil
}
