package animation

import "math/rand"

func reverseRunes(runes []rune) []rune {
	reversed := []rune{}

	for i := len(runes) - 1; i >= 0; i-- {
		reversed = append(reversed, runes[i])
	}

	return reversed
}

func randInts(length, min, max int) []int {
	ints := make([]int, length)

	for i := 0; i < length; i++ {
		ints[i] = rand.Int()%(max-min) + min
	}

	return ints
}

var shadeCharacters = []rune{' ', '░', '▒', '▓', '█'}
