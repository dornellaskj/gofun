package main

import (
    "fmt"
    "io"
    "os"
	"time"
)

func sleep() {
	time.Sleep(1 * time.Second)
}
func Countdown(out io.Writer, num int) {
	for i := num; i > 0; i-- {
		sleep()
		fmt.Fprintln(out, i)
	}

	sleep()
    fmt.Fprint(out, "Go!")
}

func main() {
	Countdown(os.Stdout, 3)
}