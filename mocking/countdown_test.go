package main

import (
	"bytes"
	"testing"
)

func TestCountdown(t *testing.T) {
	const nDesiredSleeperCalls = 4
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
}
