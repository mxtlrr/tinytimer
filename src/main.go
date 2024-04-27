package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

var (
	seconds      int
	milliseconds int
	running      bool = true
)

func main() {

	for running != false {
		if milliseconds < 1000 {
			milliseconds++
		} else {
			seconds++
			milliseconds = 0 // One second has elapsed.
		}
		time.Sleep(1 * time.Millisecond)
		fmt.Printf("%d.%d\r", seconds, milliseconds)
		go checkInput()
	}
	fmt.Printf("\n")
}

/* Check for input. Set running to false if we get input */
func checkInput() {
	// Set up terminal to not echo out character we get!
	/* TODO: make it so it's usable on windows */
	exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()

	var b []byte = make([]byte, 1)
	for {
		os.Stdin.Read(b)
		if string(b) != "" {
			running = false // Input!
		}
	}

}
