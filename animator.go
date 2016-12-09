package main

import "time"

type Animator struct {
	Banner    *Banner
	Animation Animation
	Fps       int
}

func NewAnimator(banner *Banner, animation Animation) *Animator {
	return &Animator{Banner: banner, Animation: animation, Fps: 100}
}

func (a *Animator) Animate() {
	animLength := a.Animation.Length()

	Invisible()

	for frame := 0; frame < animLength; frame++ {
		a.Banner.Transform(func(x, y int, r rune) rune {
			return a.Animation.RuneAtPos(x, y, frame, a.Banner)
		}).Print()

		time.Sleep(time.Duration(int(time.Second) / a.Fps))

		if frame < animLength-1 {
			GoTo(DirUp, a.Banner.Height)
		}
	}

	Reset()
}
