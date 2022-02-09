package interpreter_test

import (
	"testing"

	"github.com/info700700/calc-go/interpreter"
)

func TestExec(t *testing.T) {
	cases := []struct {
		in   string
		want uint32
	}{
		{"1", 1},
		{"1+1", 2},
	}

	for _, c := range cases {
		got, err := interpreter.Exec(c.in)
		if err != nil {
			t.Error("Unexpected error")
		}

		if got != c.want {
			t.Errorf("Exec(%q) == %d, want %d", c.in, got, c.want)
		}
	}
}
