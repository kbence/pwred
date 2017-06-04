package main

import (
	"flag"
	"fmt"
	"os"
	"time"
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

	width, height := GetTerminalDimensions()
	banner := LoadBannerList(args[0]).SelectRandom().Crop(width, height-1)

	settings := &AnimationSettings{
		Width:    width,
		Height:   height - 1,
		Fps:      *fps,
		Duration: time.Duration(*duration) * time.Second,
		Banner:   banner}

	animator := NewAnimator(banner, GetRandomAnimation(settings), settings)
	animator.Animate()
}
