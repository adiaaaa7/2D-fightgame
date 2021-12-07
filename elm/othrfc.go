package elm

import (
	"image"
	"math"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
)

//enemy move
func (cpo *Other) moveAct(cp *CameraP) {
	Sjump := rand.Intn(512)
	if Sjump == 30 {
		cpo.walkTime = 60
		Sjump = 0
	}
	if cpo.walkTime > 0 {
		if cpo.OP.PX > cp.PX {
			cpo.OP.PX -= pspeed
			cpo.OP.Mv = Left
		}
		if cpo.OP.PX < cp.PX {
			cpo.OP.PX += pspeed
			cpo.OP.Mv = Right
		}
		if cpo.OP.PY > cp.PY {
			cpo.OP.PY -= pspeed
			cpo.OP.Mv = Up
		}
		if cpo.OP.PY < cp.PY {
			cpo.OP.PY += pspeed
			cpo.OP.Mv = Down
		}
		cpo.walkTime--
	}
}

//enemy jump
func (cpo *Other) jumpAct() {
	Sjump := rand.Intn(256)
	if Sjump == 64 {
		cpo.OP.jumping = true
		cpo.OP.jumpH = 2 * pjumpH
	}
	if cpo.OP.jumpH > 1 && cpo.OP.jumping {
		if cpo.OP.jumpH > pjumpH {
			cpo.OP.jumpA = int(math.Pow(float64(cpo.OP.jumpH-pjumpH), 2.0))
		} else {
			cpo.OP.jumpA = -int(math.Pow(float64(cpo.OP.jumpH-pjumpH), 2.0))
		}
		cpo.OP.PY -= cpo.OP.jumpA
	}
	if cpo.OP.jumping {
		cpo.OP.jumpH -= jumpS
	}
	if cpo.OP.jumpH < 0 {
		cpo.OP.jumping = false
		cpo.OP.jumpA = 0
	}
}

//enemy move animation
func rndmoveAmn(mv Move, walktime int) image.Rectangle {
	per := walktime % 24 / 6
	return image.Rect(per*PcWidth, int(mv)*PcHight, per*PcWidth+PcWidth, (int(mv)+1)*PcHight)
}

//enemy lying down animation
func lyingAmn(cpo *Other, geom *ebiten.GeoM) (image.Rectangle, int) {
	switch {
	case cpo.DeathFace == Left || cpo.DeathFace == Up:
		geom.Translate(-PcWidth, -PcHight)
		geom.Rotate(float64(cpo.DeathTime-120) * math.Pi / 120)
	case cpo.DeathFace == Right || cpo.DeathFace == Down:
		geom.Translate(-PcWidth*0.5, -PcHight)
		geom.Rotate(float64(120-cpo.DeathTime) * math.Pi / 120)
	}
	return image.Rect(0, int(cpo.OP.Mv)*PcHight, PcWidth, (int(cpo.OP.Mv)+1)*PcHight), PcHight / 2
}

//enemy lying down over animation
func lyingoverAmn(cpo *Other, geom *ebiten.GeoM) (image.Rectangle, int) {
	switch {
	case cpo.DeathFace == Left || cpo.DeathFace == Up:
		geom.Translate(-PcWidth, -PcHight)
		geom.Rotate(float64(-60) * math.Pi / 120)
	case cpo.DeathFace == Right || cpo.DeathFace == Down:
		geom.Translate(-PcWidth*0.5, -PcHight)
		geom.Rotate(float64(60) * math.Pi / 120)
	}
	return image.Rect(0, int(cpo.OP.Mv)*PcHight, PcWidth, (int(cpo.OP.Mv)+1)*PcHight), PcHight / 2
}
