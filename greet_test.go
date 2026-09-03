package playground

import "testing"

func TestGreet(t *testing.T) {
	if got := Greet("Jay"); got != "Hello, Jay!" {
		t.Fatalf("Greet = %q", got)
	}
}
