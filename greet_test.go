package playground

import "testing"

func TestGreet(t *testing.T) {
	if got := Greet("Jay"); got != "Hello, Jay!" {
		t.Fatalf("Greet = %q", got)
	}
}

func TestGreetInKorean(t *testing.T) {
	if got := GreetIn("Jay", LangKO); got != "안녕하세요, Jay님!" {
		t.Fatalf("GreetIn ko = %q", got)
	}
}

func TestGreetInUnknownLangFallsBackToEnglish(t *testing.T) {
	if got := GreetIn("Jay", Lang("ja")); got != "Hello, Jay!" {
		t.Fatalf("GreetIn ja = %q", got)
	}
}

func TestGreetInZeroLangFallsBackToEnglish(t *testing.T) {
	var lang Lang
	if got := GreetIn("Jay", lang); got != "Hello, Jay!" {
		t.Fatalf("GreetIn zero = %q", got)
	}
}
