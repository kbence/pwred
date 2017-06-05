package animation

import (
	"math/rand"
	"time"

	"github.com/kbence/pwred/banner"
)

type AnimationSettings struct {
	Width    int
	Height   int
	Fps      int
	Duration time.Duration
	Banner   *banner.Banner
}

func (s *AnimationSettings) Length() int {
	return int(int64(s.Duration) / (int64(time.Second) / int64(s.Fps)))
}

type Animation interface {
	Initialize(settings *AnimationSettings)
	Length() int
	RuneAtPos(x, y int, frame int, banner *banner.Banner) rune
}

func GetRandomAnimation(settings *AnimationSettings) Animation {
	var anim Animation

	switch rand.Int() % 3 {
	case 0:
		anim = &RippleAnimation{}
		break
	case 1:
		anim = &SparklingAnimation{}
		break
	case 2:
		anim = &DoomAnimation{}
		break
	}

	anim.Initialize(settings)

	return anim
}
