package str

import (
	"testing"
)

func TestStr(t *testing.T) {
	want := 1
	if got := Str("b", "a"); got != want {
		t.Errorf("Str() = %q, want %q", got, want)
	}
}
