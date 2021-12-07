package elm

import (
	"image"
	_ "image/png"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
)

func (cpo *Other) InitialPosition() {
	cpo.OP.PX = rand.Intn(PxmaxB-PxminB) + PxminB
	cpo.OP.PY = rand.Intn(PymaxB-PyminB) + PyminB
	cpo.OP.Mv = RndMov(rand.Intn(4) + 1)
	cpo.OpcImage = PcImages[rand.Intn(32)]
	cpo.DeathTime = 0
}

func (cpo *Other) Update(cp *CameraP) error {
	if cpo.DeathTime > 0 {
		cpo.DeathTime -= 1
	} else {
		cpo.moveAct(cp)
		cpo.jumpAct()
	}
	cpo.OP.PY = overYBoundrs(cpo.OP.PY)
	return nil
}

func (cpo *Other) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Reset()
	op.GeoM.Scale(1.5, 1.5)
	//move animation
	var r image.Rectangle
	var amh = 0
	if cpo.DeathTime == 0 {
		op.GeoM.Translate(-float64(PcWidth/2), -float64(PcHight/2))
		r = rndmoveAmn(cpo.OP.Mv, cpo.walkTime)
	} else if cpo.DeathTime > 60 {
		r, amh = lyingAmn(cpo, &op.GeoM)
	} else if cpo.DeathTime > 0 && cpo.DeathTime <= 60 {
		r, amh = lyingoverAmn(cpo, &op.GeoM)
	}
	op.GeoM.Translate(float64(cpo.OP.PX), float64(cpo.OP.PY+amh))
	screen.DrawImage(cpo.OpcImage.SubImage(r).(*ebiten.Image), op)
	//input.XyP(screen, cpo.DeathTime)
}
