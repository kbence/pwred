package terminal

import (
	"fmt"
	"os"

	"golang.org/x/crypto/ssh/terminal"
)

type direction uint8

const (
	DirUp direction = iota
	DirDown
	DirRight
	DirLeft
)

func GetTerminalDimensions() (int, int) {
	width, height, err := terminal.GetSize(int(os.Stdout.Fd()))

	if err != nil {
		width = 80
		height = 25
	}

	return width, height
}

func GoTo(dir direction, amount int) {
	dirSpec := 'A' + byte(dir)

	fmt.Printf("\033[%d%c", amount, dirSpec)
}
