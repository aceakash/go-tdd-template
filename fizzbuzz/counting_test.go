package fizzbuzz_test

import (
	"github.com/aceakash/go-tdd-template/fizzbuzz"
	"testing"
)

func TestCounting(t *testing.T) {
	t.Run("when given zero, returns empty string", func(t *testing.T) {
		got := fizzbuzz.Count(0)
		want := "1"

		if got != want {
			t.Errorf("wanted %v but got '%v'", want, got)
		}
	})
}
