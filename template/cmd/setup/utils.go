package main

import (
	"os"
	"syscall"
)

const (
	arrowUp   = "\033[A"
	arrowDown = "\033[B"
	enter     = "\n"
)

func getKey() string {
	var buf [3]byte
	n, _ := syscall.Read(int(os.Stdin.Fd()), buf[:])

	if n == 1 {
		return string(buf[0])
	}
	if buf[0] == 27 && buf[1] == 91 {
		return string(buf[:])
	}

	return ""
}
