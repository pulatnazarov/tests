package mocking

import (
	"fmt"
	"io"
	"time"
)

type Sleeper interface {
	Sleep(d time.Duration)
}

type RealSleeper struct{}

func (s RealSleeper) Sleep(d time.Duration) {
	time.Sleep(d)
}

func Countdown(w io.Writer, s Sleeper) {
	for i := 3; i > 0; i-- {
		fmt.Fprintln(w, i)
		s.Sleep(time.Second)
	}
	fmt.Fprintln(w, "Go!")
}

func CountDown1() {
	for i := 3; i > 0; i-- {
		fmt.Println(i)
		time.Sleep(time.Second)
	}
	fmt.Println("Go!")
}
