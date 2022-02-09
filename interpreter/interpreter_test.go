package interpreter_test

import (
	"testing"

	"github.com/info700700/calc-go/interpreter"
)

func TestExec(t *testing.T) {
	got, err := interpreter.Exec("1+1")

	if err != nil {
		t.Error("Unexpected error")
	}

	const want = 2
	if got != want {
		t.Errorf("Exec(\"1+1\") = %d; want %d", got, want)
	}
}
