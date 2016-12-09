package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	flag.Parse()
	args := flag.Args()

	if len(args) != 1 {
		fmt.Printf("Usage: pwred <banner_file>\n")
		os.Exit(1)
	}

	width, height := GetTerminalDimensions()
	banner := LoadBannerList(args[0]).SelectRandom().Crop(width, height-1)
	animator := NewAnimator(banner, &RippleAnimation{Width: banner.Width, Height: banner.Height})
	animator.Animate()
}
