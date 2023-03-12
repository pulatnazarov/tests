package mocking

import (
	"bytes"
	"testing"
	"time"
)

type SpySleeper struct {
	count int
	args  []time.Duration
}

func (s *SpySleeper) Sleep(d time.Duration) {
	s.count++
	s.args = append(s.args, d)
}

func TestCountdown(t *testing.T) {
	buf := &bytes.Buffer{}

	s := &SpySleeper{}
	Countdown(buf, s)

	want := `3
2
1
Go!
`
	if buf.String() != want {
		t.Fatalf("want %q, got %q", want, buf.String())
	}

	if s.count != 3 {
		t.Fatalf("want 3, got %d", s.count)
	}

	for i, d := range s.args {
		if d != time.Second {
			t.Errorf("expected sleep to be called with 1 second, but got with %v in call number %d", d, i+1)
		}
	}
}

func TestCountDown1(t *testing.T) {

}
