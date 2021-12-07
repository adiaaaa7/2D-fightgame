package elm

import (
	_ "image/png"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
)

func (cp *CameraP) InitialPosition() {
	cp.PX = oPxy
	cp.PY = oPxy
	cp.Mv = Right
}

func init() {
	Stuinit()
	PcImage = PcImages[rand.Intn(33)]
}

func (cp *CameraP) Update(realX int, others []*Other, wps []*Wpn) error {
	cp.moveAct(realX)
	cp.othrwonUpd(realX, others, wps)
	cp.jumpAct()
	return nil
}

func (cp *CameraP) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Reset()
	op.GeoM.Scale(1.5, 1.5)
	r := moveAmn(cp.Mv)
	op.GeoM.Translate(float64(cp.PX), float64(cp.PY))
	op.GeoM.Translate(-float64(PcWidth/2), -float64(PcHight/2))
	screen.DrawImage(PcImage.SubImage(r).(*ebiten.Image), op)
	//input.XyP(screen, math.Pow(float64(pjumpH), 2.0), cp.jumpH, cp.jumpA, cp.PY)
}
