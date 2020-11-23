package main

import (
	"bufio"
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

	displayIndex := *display - 1
	scrAreaSelector := screen.NewAreaSelector(displayIndex)

	fmt.Println("Select the area of the screen to track.")
	selArea := scrAreaSelector.Select()
	fmt.Printf("The area of the screen to track: (%d,%d)-(%d,%d)\n", selArea.X1, selArea.Y1, selArea.X2, selArea.Y2)

	screenTracker := screen.NewTracker(displayIndex, selArea).WithAlert(screen.AlerterFunc(alert.Play))
	screenTrackerMgr := screen.NewTrackerManager(screenTracker, *timeout)
	screenTrackerMgr.Start()

	go func() {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			input := scanner.Text()
			if input == "q" || input == "quit" {
				screenTrackerMgr.Stop()
				break
			} else if input == "s" || input == "start" {
				screenTrackerMgr.Start()

			} else if input == "p" || input == "pause" {
				screenTrackerMgr.Pause()
			}
		}
	}()

	<-screenTrackerMgr.Done()
	if err := screenTrackerMgr.Error(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(2)
	}
}
