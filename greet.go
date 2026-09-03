package playground

import (
	"fmt"
	"strings"
)

// Greet returns a greeting for name.
func Greet(name string) string {
	return fmt.Sprintf("Hello, %s!", name)
}

// Whisper returns the greeting for name in all lowercase.
func Whisper(name string) string {
	return strings.ToLower(Greet(name))
}
