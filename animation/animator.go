package animation

import (
	"time"

	"github.com/kbence/pwred/banner"
	"github.com/kbence/pwred/terminal"
)

type Animator struct {
	Banner    *banner.Banner
	Animation Animation
	Settings  *AnimationSettings
}

func NewAnimator(banner *banner.Banner, animation Animation, settings *AnimationSettings) *Animator {
	return &Animator{Banner: banner, Animation: animation, Settings: settings}
}

func (a *Animator) Animate() {
	animLength := a.Animation.Length()

	for frame := 0; frame < animLength; frame++ {
		a.Banner.Transform(func(x, y int, r rune) rune {
			return a.Animation.RuneAtPos(x, y, frame, a.Banner)
		}).Print()

		time.Sleep(time.Duration(int(time.Second) / a.Settings.Fps))

		if frame < animLength-1 {
			terminal.GoTo(terminal.DirUp, a.Banner.Height)
		}
	}
}
