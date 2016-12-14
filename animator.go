package main

import "time"

type Animator struct {
	Banner    *Banner
	Animation Animation
	Settings  *AnimationSettings
}

func NewAnimator(banner *Banner, animation Animation, settings *AnimationSettings) *Animator {
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
			GoTo(DirUp, a.Banner.Height)
		}
	}
}
