package main

import "fmt"

type direction uint8

const (
	DirUp direction = iota
	DirDown
	DirRight
	DirLeft
)

func GoTo(dir direction, amount int) {
	dirSpec := 'A' + byte(dir)

	fmt.Printf("\033[%d%c", amount, dirSpec)
}

func Invisible() {
	fmt.Printf("\033[8m")
}

func Reset() {
	fmt.Printf("\033(B\033[m")
}
