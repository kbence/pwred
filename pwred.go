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

	banner := LoadBannerList(args[0]).SelectRandom()
	animator := NewAnimator(banner, &RippleAnimation{Width: banner.Width, Height: banner.Height})
	animator.Animate()
}
