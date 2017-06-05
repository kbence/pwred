package animation

import (
	"math"
	"math/rand"
	"time"

	"github.com/kbence/pwred/banner"
	"github.com/kbence/pwred/utils"
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
