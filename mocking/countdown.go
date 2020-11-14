package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

var finalWord = "Go!"
var countdownStart = 3

func Countdown(out io.Writer) {

	for i := countdownStart; i > 0; i-- {
		time.Sleep(1 * time.Second)
		fmt.Fprintln(out, i)
	}
	time.Sleep(1 * time.Second)
	fmt.Fprint(out, finalWord)
}

func main() {
	Countdown(os.Stdout)
}
