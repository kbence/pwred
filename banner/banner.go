package banner

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/kbence/pwred/utils"
)

func init() {
	rand.Seed(int64(time.Now().Nanosecond()))
}

type Banner struct {
	Width  int
	Height int
	Lines  [][]rune
}

type TransformFunc func(x, y int, r rune) rune

func NewBanner(lines []string) *Banner {
	width := 0
	runes := [][]rune{}

	for _, line := range lines {
		runeLine := []rune{}

		for _, r := range line {
			runeLine = append(runeLine, r)
		}

		if len(runeLine) > width {
			width = len(runeLine)
		}

		runes = append(runes, runeLine)
	}

	return &Banner{Width: width, Height: len(runes), Lines: runes}
}

func (b *Banner) Print() {
	for _, line := range b.Lines {
		fmt.Println(string(line))
	}
}

func (b *Banner) IsInside(x, y int) bool {
	return x >= 0 && y >= 0 && y < len(b.Lines) && x < len(b.Lines[y])
}

func (b *Banner) GetRune(x, y int) rune {
	if b.IsInside(x, y) {
		return rune(b.Lines[y][x])
	}

	return ' '
}

func (b *Banner) Transform(function TransformFunc) *Banner {
	banner := &Banner{Width: b.Width, Height: b.Height}

	for y := 0; y < b.Height; y++ {
		line := make([]rune, b.Width)

		for x := 0; x < b.Width; x++ {
			line[x] = function(x, y, b.GetRune(x, y))
		}

		banner.Lines = append(banner.Lines, line)
	}

	return banner
}

func (b *Banner) Crop(width, height int) *Banner {
	banner := &Banner{Width: utils.Min(b.Width, width), Height: utils.Min(b.Height, height)}

	for y := 0; y < height && y < b.Height; y++ {
		line := b.Lines[y]
		banner.Lines = append(banner.Lines, line[0:utils.Min(len(line), width)])
	}

	return banner
}

type BannerList struct {
	Banners []*Banner
}

func LoadBannerList(filename string) *BannerList {
	bannerList := &BannerList{Banners: []*Banner{}}
	currentBannerLines := []string{}

	f, err := os.Open(filename)
	if err != nil {
		log.Fatalf("error opening banner file: %s", err)
	}

	reader := bufio.NewReader(f)
	separator := ""
	lineNumber := 0

	for {
		line, err := reader.ReadString('\n')

		if len(line) > 0 {
			line = line[0 : len(line)-1]
		}

		if err != nil {
			if len(line) > 0 {
				currentBannerLines = append(currentBannerLines, line)
			}

			if err == io.EOF {
				break
			}

			log.Fatalf("error reading banner file: %s", err)
		}

		if lineNumber == 0 {
			separator = line
		} else {
			if line == separator {
				bannerList.Banners = append(bannerList.Banners, NewBanner(currentBannerLines))
				currentBannerLines = []string{}
			} else {
				currentBannerLines = append(currentBannerLines, line)
			}
		}

		lineNumber++
	}

	if len(currentBannerLines) > 0 {
		bannerList.Banners = append(bannerList.Banners, NewBanner(currentBannerLines))
	}

	return bannerList
}

func (l *BannerList) SelectRandom() *Banner {
	return l.Banners[rand.Int()%len(l.Banners)]
}
