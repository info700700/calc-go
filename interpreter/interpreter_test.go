package interpreter_test

import (
	"testing"

	"github.com/info700700/calc-go/interpreter"
)

func TestExec(t *testing.T) {
	testCases := []struct {
		str  string
		want uint32
	}{
		{"1", 1},
		{" 1", 1},  // Пробел в начале строки.
		{"1 ", 1},  // Пробел в конце строки.
		{" 1 ", 1}, // Пробел в начале и в конце строки.
		{"1+1", 2},
		{"1+1+1", 3},
	}

	for _, c := range testCases {
		got, err := interpreter.Exec(c.str)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}

		if got != c.want {
			t.Errorf("Exec(%q) == %d, want %d", c.str, got, c.want)
		}
	}
}
