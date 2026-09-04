package playground

import (
	"fmt"
	"strings"
)

// Lang identifies the language of a greeting.
type Lang string

const (
	LangEN Lang = "en"
	LangKO Lang = "ko"
)

var greetings = map[Lang]string{
	LangEN: "Hello, %s!",
	LangKO: "안녕하세요, %s님!",
}

// GreetIn returns a greeting for name in lang.
// Languages without a greeting, including the zero value, fall back to English.
func GreetIn(name string, lang Lang) string {
	format, ok := greetings[lang]
	if !ok {
		format = greetings[LangEN]
	}
	return fmt.Sprintf(format, name)
}

// Greet returns an English greeting for name.
func Greet(name string) string {
	return GreetIn(name, LangEN)
}

// Whisper returns the greeting for name in all lowercase.
func Whisper(name string) string {
	return strings.ToLower(Greet(name))
}
