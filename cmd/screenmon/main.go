package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"time"

	"github.com/vps2/screenmon/internal/alert"
	"github.com/vps2/screenmon/internal/screen"

	flag "github.com/spf13/pflag"
)

func main() {
	timeout := flag.DurationP("timeout", "t", 15*time.Second, "timeout between screen capture")
	display := flag.IntP("display", "d", 1, "number of the display to track changes on. It may not match the number in the system settings.")
	help := flag.BoolP("help", "h", false, "show help.")

	_ = flag.CommandLine.MarkHidden("help")
	flag.Parse()

	if *help {
		flag.Usage()
		os.Exit(0)
	}

	if *display <= 0 || *display > screen.NumActiveScreens() {
		fmt.Fprintln(os.Stderr, "wrong display number:", *display)
		os.Exit(1)
	}

	ctx, cancel := context.WithCancel(context.Background())

	go func(cancel func()) {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			input := scanner.Text()
			if input == "q" || input == "quit" || input == "exit" {
				cancel()
				break
			}
		}

	}(cancel)

	displayIndex := *display - 1
	screenTracker := screen.NewTracker(displayIndex).WithAlert(screen.AlerterFunc(alert.Play))
	if err := screenTracker.TrackChanges(ctx, *timeout); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(2)
	}
}
