package animation

import (
	"github.com/kbence/pwred/banner"
	"github.com/kbence/pwred/utils"
)

type DoomAnimation struct {
	length         int
	height         int
	startPositions []int
}

func (a *DoomAnimation) Initialize(settings *AnimationSettings) {
	a.length = settings.Length() * 2
	a.height = settings.Banner.Height
	a.startPositions = randInts(settings.Banner.Width, 2*a.height, 3*a.height)
}

func (a *DoomAnimation) Length() int {
	return a.length
}

func (a *DoomAnimation) RuneAtPos(x, y int, frame int, banner *banner.Banner) rune {
	offset := a.startPositions[x] * (a.length - frame + 1) / a.length

	if y+offset > a.height {
		return shadeCharacters[utils.Max(0, utils.Min(len(shadeCharacters)-1, y+offset-a.height))]
	}

	return banner.GetRune(x, y+offset)
}
