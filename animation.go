package main

import "math"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type Animation interface {
	Length() int
	RuneAtPos(x, y int, frame int, banner *Banner) rune
}

var shadeCharacters = []rune{' ', '░', '▒', '▓', '█'}

type RippleAnimation struct {
	Width  int
	Height int
}

func (a *RippleAnimation) Length() int {
	return int(math.Sqrt(float64(a.Width*a.Width + a.Height*a.Height)))
}

func (a *RippleAnimation) RuneAtPos(x, y int, frame int, banner *Banner) rune {
	epsilon := 8.0
	rippleDistance := float64(frame)
	floatWidth := float64(x-banner.Width/2) / float64(banner.Width)
	floatHeight := float64(y-banner.Height/2) / float64(banner.Height)
	distance := math.Sqrt(floatWidth*floatWidth+floatHeight*floatHeight) * float64(max(banner.Width, banner.Height))

	if distance <= rippleDistance+epsilon && distance >= rippleDistance-epsilon {
		strength := 0.99 * math.Cos((distance-rippleDistance)/(2*epsilon)*math.Pi)

		return shadeCharacters[int(strength*float64(len(shadeCharacters)))]
	} else if rippleDistance+epsilon < distance {
		return ' '
	}

	return banner.GetRune(x, y)
}
