package main

import (
	"fmt"
	"net"
	"test/mocking"
)

type CustomerWriter struct{}

func (w *CustomerWriter) Write(b []byte) (int, error) {
	fmt.Println("--------------------")
	fmt.Println(string(b))

	return len(b), nil
}

func main() {
	lis, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		panic(err)
	}
	for {
		conn, err := lis.Accept()
		if err != nil {
			panic(err)
		}
		mocking.Countdown(conn, mocking.RealSleeper{})
	}
}
