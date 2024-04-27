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

	/* Logging */
	to_log bool   = false // Let's not log, default.
	log    string = ""    // By default we log to stdin
)

func main() {
	var (
		file string = "splits.txt"
	)
	if len(os.Args) > 1 {
		args := os.Args
		for i := range args {
			if strings.Compare(args[i], "--file") == 0 {
				file = args[i+1]
			}

			if strings.Compare(args[i], "--log") == 0 {
				to_log = true   // Yes, log.
				log = args[i+1] // Log to this file please!
			}
		}
	}
	splits_ = splits.Gen_splits(file)
	/* https://stackoverflow.com/q/19979178#comment29744145_19979829 */
	fmt.Printf("Logging? %s | File: %s\n",
		(map[bool]string{true: "yes", false: "no"})[to_log], log)

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
	if to_log {
		os.Create(log) // Create file
	}

	for i := range splits_ {
		if to_log {
			f, err := os.OpenFile(log, os.O_APPEND|os.O_WRONLY, os.ModeAppend) // Open it for writing
			if err != nil {
				panic(err)
			}
			_, err = fmt.Fprintf(f, "SPLIT: %s | TIME %d.%d\n",
				splits_[i].NAME, splits_[i].TIME_SECONDS,
				splits_[i].TIME_MILLISEC)
			if err != nil {
				panic(err)
			}
			f.Close() // I feel like shit writing this
		} else { // Else just write out stdout
			fmt.Printf("SPLIT: %s | TIME %d.%d\n",
				splits_[i].NAME, splits_[i].TIME_SECONDS,
				splits_[i].TIME_MILLISEC)
		}
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
