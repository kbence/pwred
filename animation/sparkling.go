package animation

import (
	"math/rand"

	"github.com/kbence/pwred/banner"
)

type SparklingAnimation struct {
	width      int
	thresholds []int
	length     int
	shadeChars []rune
}

func (a *SparklingAnimation) Initialize(settings *AnimationSettings) {
	a.width = settings.Banner.Width
	a.thresholds = rand.Perm(settings.Banner.Width * settings.Height)
	a.length = settings.Length()
	a.shadeChars = append(shadeCharacters[0:len(shadeCharacters)], reverseRunes(shadeCharacters)...)
	a.thresholds = randInts(settings.Banner.Width*settings.Height, 1, a.length-len(a.shadeChars)-1)
}

func (a *SparklingAnimation) Length() int {
	return a.length
}

func (a *SparklingAnimation) RuneAtPos(x, y int, frame int, banner *banner.Banner) rune {
	threshold := frame
	startThreshold := (frame - len(a.shadeChars) - 1)
	value := a.thresholds[y*a.width+x]

	if value <= startThreshold {
		return banner.GetRune(x, y)
	}

	if value > startThreshold && value < threshold {
		return a.shadeChars[((value-startThreshold)*len(a.shadeChars))/(threshold-startThreshold)]
	}

	if value > threshold {
		return a.shadeChars[len(a.shadeChars)-1]
	}

	return ' '
}
