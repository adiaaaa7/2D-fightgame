package elm

import (
	"game/myds/elm/input"
	"image"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

//pc move
func (cp *CameraP) moveAct(realX int) {
	if realX < ScreenWidth/2 {
		cp.PX = realX
	}
	if realX > BgWidth-ScreenWidth/2 {
		cp.PX = realX - (BgWidth - ScreenWidth)
	}
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		cp.PY -= pspeed
		cp.Mv = Up
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		cp.PY += pspeed
		cp.Mv = Down
	}
}

//pc pozation to enemy,wpn
func (cp *CameraP) othrwonUpd(realX int, others []*Other, wps []*Wpn) {
	if realX > ScreenWidth/2 && realX < BgWidth-ScreenWidth/2 {
		if per := input.KeyPressingTime(ebiten.KeyD); per > 0 {
			for _, o := range others {
				if o != nil {
					o.OP.PX -= (pspeed + 1)
				}
			}
			for _, w := range wps {
				if w != nil {
					w.WX -= (pspeed + 1)
				}
			}
		}
		if per := input.KeyPressingTime(ebiten.KeyA); per > 0 {
			for _, o := range others {
				if o != nil {
					o.OP.PX += (pspeed + 1)
				}
			}
			for _, w := range wps {
				if w != nil {
					w.WX += (pspeed + 1)
				}
			}
		}
	}
}

//pc jump
func (cp *CameraP) jumpAct() {
	if input.KeyPressingStart(ebiten.KeyI) {
		cp.jumping = true
		cp.jumpH = 2 * pjumpH
	}
	if cp.jumpH >= 0 && cp.jumping {
		if cp.jumpH > pjumpH {
			//up
			cp.jumpA = int(math.Pow(float64(cp.jumpH-pjumpH), 2.0))
		} else {
			//down
			cp.jumpA = -int(math.Pow(float64(cp.jumpH-pjumpH), 2.0))
		}
		cp.PY -= cp.jumpA
	}
	if cp.jumping {
		cp.jumpH -= jumpS
	} else {
		cp.PY = overYBoundrs(cp.PY)
	}
	if cp.jumpH < 0 {
		cp.jumping = false
	}
}

//pc move animation
func moveAmn(mv Move) image.Rectangle {
	var r image.Rectangle
	r = image.Rect(0, int(mv)*PcHight, PcWidth, (int(mv)+1)*PcHight)
	switch mv {
	case Down:
		if per := input.MoveKeyPressing(ebiten.KeyS); per > 0 {
			r = image.Rect(per*PcWidth, 0, per*PcWidth+PcWidth, 1*PcHight)
		}
	case Left:
		if per := input.MoveKeyPressing(ebiten.KeyA); per > 0 {
			r = image.Rect(per*PcWidth, PcHight, per*PcWidth+PcWidth, 2*PcHight)
		}
	case Right:
		if per := input.MoveKeyPressing(ebiten.KeyD); per > 0 {
			r = image.Rect(per*PcWidth, 2*PcHight, per*PcWidth+PcWidth, 3*PcHight)
		}
	case Up:
		if per := input.MoveKeyPressing(ebiten.KeyW); per > 0 {
			r = image.Rect(per*PcWidth, 3*PcHight, per*PcWidth+PcWidth, 4*PcHight)
		}
	}
	return r
}
