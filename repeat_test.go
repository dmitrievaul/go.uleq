package iteration

import "testing"

const repeatCount = 5

func Repeat(character string) string {
	var repeated string
	for i := 0; i < repeatCount; i++ {
		repeated += character
	}
	return repeated
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("")
	}
}

func TestRepeat(t *testing.T) {
	repeated := Repeat("")
	expected := ""

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}
