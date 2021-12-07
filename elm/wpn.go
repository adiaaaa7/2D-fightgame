package elm

import (
	"image"
	_ "image/png"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
)

func (wp *Wpn) InitialPosition() {
	wp.WX = rand.Intn(PxmaxB-wpnSize-PxminB-wpnSize) + PxminB + wpnSize
	wp.WY = rand.Intn(PymaxB-wpnSize-PyminB-wpnSize) + PyminB + wpnSize
	wp.Stat = Ground
	wp.Face = Right
	wp.AttackTime = 0
	wp.Flytime = 300
	wp.WpnImage = WpImages[rand.Intn(22)]
}

func init() {
	Wswdinit()
}

func (wp *Wpn) Update(pc *CameraP, others []*Other) error {
	wp.takeAct(pc)
	wp.takingUpd(pc)
	wp.throwAct(pc)
	wp.attackAct(pc)
	wp.codnateUpd(others)
	return nil
}

func (wp *Wpn) Draw(screen *ebiten.Image) {
	var wr image.Rectangle
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Reset()
	if wp.Stat == Taken {
		amx, amy := takingAmn(wp, &op.GeoM)
		op.GeoM.Translate(float64(wp.WX+amx), float64(wp.WY+amy))
	} else if wp.Stat == Throw {
		throwingAmn(wp, &op.GeoM)
		op.GeoM.Translate(float64(wp.WX), float64(wp.WY))
	} else {
		op.GeoM.Translate(float64(wp.WX), float64(wp.WY))
	}
	wr = image.Rect(0, 0, wpnSize, wpnSize)
	screen.DrawImage(wp.WpnImage.SubImage(wr).(*ebiten.Image), op)
}
