package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

func main() {
	flag.Parse()
	args := flag.Args()

	if len(args) != 1 {
		fmt.Printf("Usage: pwred <banner_file>\n")
		os.Exit(1)
	}

	width, height := GetTerminalDimensions()
	settings := &AnimationSettings{Width: width, Height: height - 1, Fps: 30, Duration: 1 * time.Second}

	banner := LoadBannerList(args[0]).SelectRandom().Crop(settings.Width, settings.Height-1)
	animator := NewAnimator(banner, GetRandomAnimation(settings), settings)
	animator.Animate()
}
