package playground

import "testing"

func TestGreet(t *testing.T) {
	if got := Greet("Jay"); got != "Hello, Jay!" {
		t.Fatalf("Greet = %q", got)
	}
}

func TestWhisper(t *testing.T) {
	if got := Whisper("Bob"); got != "hello, bob!" {
		t.Fatalf("Whisper = %q", got)
	}
}
