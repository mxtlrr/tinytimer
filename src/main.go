package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/mxtlrr/tinytimer/splits"
)

var (
	seconds      int
	milliseconds int
	running      bool = true

	/* For splits */
	splits_     []splits.Split_t
	split_count int
	iterator    int = 0
)

func main() {
	// Argv[1] is going to be the flag, either being
	// --file, or --log.
	var file string = "splits.txt"
	if len(os.Args) > 1 {
		args := os.Args
		// We're able to change the file, but we're gonna read
		// from this file
		if strings.Compare(args[1], "--file") == 0 {
			file = args[2]
		}
	}
	splits_ = splits.Gen_splits(file)
	fmt.Printf("found splits...\n")
	for i := range splits_ {
		fmt.Printf("%d: %s\n", i, splits_[i].NAME)
	}

	// Set split count
	split_count = len(splits_)

	for running {
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

	// TODO: show output if argv[1] not log
	for i := range splits_ {
		fmt.Printf("SPLIT: %s | TIME %d.%d\n",
			splits_[i].NAME, splits_[i].TIME_SECONDS,
			splits_[i].TIME_MILLISEC)
	}
	exec.Command("stty", "-F", "/dev/tty", "echo").Run()
}

/* Check for input. Set running to false if we get input */
func checkInput() {
	/* Hide enter
	 * TODO: cross-platform */
	exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()
	exec.Command("stty", "-F", "/dev/tty", "-echo").Run()

	if iterator == split_count {
		running = false
		iterator = 0
	}
	var (
		n string
	)
	fmt.Scanf("%s", &n) // Input

	// Set data for split[iterator]
	// splits[iterator].NAME already set.
	splits_[iterator].TIME_MILLISEC = milliseconds
	splits_[iterator].TIME_SECONDS = seconds

	iterator++
}
