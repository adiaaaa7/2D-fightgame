package input

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const (
	mdelay    = 5
	fdelay    = 15
	mod       = 24
	interval2 = 6
)

func MoveKeyPressing(key ebiten.Key) int {
	d := inpututil.KeyPressDuration(key)
	if d == 1 || (d >= mdelay) {
		// 60%24 2.5/s 24/6 10/s
		return (d % mod / interval2)
	}
	return 0
}

func FlyKeyPressing(key ebiten.Key) int {
	d := inpututil.KeyPressDuration(key)
	if d == 1 {
		return 1
	}
	if d > fdelay {
		return d
	}
	return 0
}

func KeyPressingTime(key ebiten.Key) int {
	return inpututil.KeyPressDuration(key)
}

func KeyPressingStart(key ebiten.Key) bool {
	return inpututil.IsKeyJustPressed(key)
}

func KeyPressingOver(key ebiten.Key) bool {
	return inpututil.IsKeyJustReleased(key)
}

func XyP(screen *ebiten.Image, a ...interface{}) {
	s := fmt.Sprintln(a...)
	ebitenutil.DebugPrint(screen, s)
}
