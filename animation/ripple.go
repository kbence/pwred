package animation

import (
	"math"

	"github.com/kbence/pwred/banner"
	"github.com/kbence/pwred/utils"
)

type RippleAnimation struct {
	length int
}

func (a *RippleAnimation) Initialize(settings *AnimationSettings) {
	a.length = settings.Length()
}

func (a *RippleAnimation) Length() int {
	return a.length
}

func (a *RippleAnimation) RuneAtPos(x, y int, frame int, banner *banner.Banner) rune {
	epsilon := 8.0
	rippleDistance := float64(frame) / float64(a.length) * float64(utils.Max(banner.Width, banner.Height))
	floatWidth := float64(x-banner.Width/2) / float64(banner.Width)
	floatHeight := float64(y-banner.Height/2) / float64(banner.Height)
	distance := math.Sqrt(floatWidth*floatWidth+floatHeight*floatHeight) * float64(utils.Max(banner.Width, banner.Height))

	if distance <= rippleDistance+epsilon && distance >= rippleDistance-epsilon {
		strength := 0.99 * math.Cos((distance-rippleDistance)/(2*epsilon)*math.Pi)

		return shadeCharacters[int(strength*float64(len(shadeCharacters)))]
	} else if rippleDistance+epsilon < distance {
		return ' '
	}

	return banner.GetRune(x, y)
}
