package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/kbence/pwred/animation"
	"github.com/kbence/pwred/banner"
	"github.com/kbence/pwred/terminal"
)

var fps = flag.Int("fps", 30, "Printed frames per second")
var duration = flag.Int("duration", 2, "Duration in seconds")

func main() {
	flag.Parse()
	args := flag.Args()

	if len(args) != 1 {
		fmt.Printf("Usage: pwred <banner_file>\n")
		os.Exit(1)
	}

	width, height := terminal.GetTerminalDimensions()
	banner := banner.LoadBannerList(args[0]).SelectRandom().Crop(width, height-1)

	settings := &animation.AnimationSettings{
		Width:    width,
		Height:   height - 1,
		Fps:      *fps,
		Duration: time.Duration(*duration) * time.Second,
		Banner:   banner}

	animator := animation.NewAnimator(banner, animation.GetRandomAnimation(settings), settings)
	animator.Animate()
}
