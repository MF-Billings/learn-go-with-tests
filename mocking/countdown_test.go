package main

import (
	"bytes"
	"reflect"
	"testing"
)

func TestCountdown(t *testing.T) {
	const nDesiredSleeperCalls = 4

	t.Run("prints 3 to Go!", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		spySleeper := &SpySleeper{}

		Countdown(buffer, spySleeper)

		got := buffer.String()
		want := `3
2
1
Go!`

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}

		if spySleeper.Calls != nDesiredSleeperCalls {
			t.Errorf("not enough calls do sleeper, got %d wanted %d", spySleeper.Calls, nDesiredSleeperCalls)
		}
	})

	t.Run("sleep before every print", func(t *testing.T) {
		spySleepPrinter := &CountdownOperationsSpy{}
		Countdown(spySleepPrinter, spySleepPrinter)

		want := []string{
			"sleep", "write",
			"sleep", "write",
			"sleep", "write",
			"sleep", "write",
		}

		if !reflect.DeepEqual(spySleepPrinter.Calls, want) {
			t.Errorf("wanted calls %v, got %v", want, spySleepPrinter.Calls)
		}
	})
}
