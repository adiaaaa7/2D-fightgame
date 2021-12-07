package elm

import (
	"game/myds/elm/input"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

//wp takeing
func (wp *Wpn) takingUpd(pc *CameraP) {
	if wp.Stat == Taken {
		wp.WX = pc.PX
		wp.WY = pc.PY
		wp.Face = pc.Mv
	}
}

//wp take
func (wp *Wpn) takeAct(pc *CameraP) {
	if wp.Stat == Ground && !pc.Armed {
		if math.Abs(float64(pc.PX-wp.WX)) < (wpnSize+PcWidth)/2 &&
			math.Abs(float64(pc.PY-wp.WY)) < (wpnSize+PcHight)/2 {
			wp.Stat = Taken
			pc.Armed = true
		}
	}
}

//wp attack
func (wp *Wpn) attackAct(pc *CameraP) {
	if wp.AttackTime > 0 {
		wp.AttackTime -= 2
	}
	if input.KeyPressingStart(ebiten.KeyO) && wp.Stat == Taken {
		wp.AttackTime = 30
		wp.Face = pc.Mv
	}
}

// wpn codnate
func (wp *Wpn) codnateUpd(others []*Other) {
	for _, o := range others {
		if math.Abs(float64(wp.WX-o.OP.PX)) < (wpnSize+PcWidth)/2 &&
			math.Abs(float64(wp.WY-o.OP.PY)) < (wpnSize+PcHight)/2 && o.DeathTime == 0 &&
			(wp.AttackTime > 0 || (wp.Flytime < 300 && wp.Flytime > 0)) {
			o.DeathTime = 120
			o.DeathFace = wp.Face
		}
	}
}

//throw wpn
func (wp *Wpn) throwAct(pc *CameraP) {
	//charge
	if charge := input.FlyKeyPressing(ebiten.KeyP); charge > 0 {
		wp.chargeTime = charge
		return
	}
	if input.KeyPressingOver(ebiten.KeyP) && wp.chargeTime > 0 &&
		wp.Stat == Taken && pc.Armed {
		wp.Stat = Throw
		wp.wspeed = wp.chargeTime / 10
		pc.Armed = false
	}
	if wp.Stat == Throw && wp.Flytime >= 0 {
		wp.Flytime -= 1
		switch wp.Face {
		case Left:
			wp.WX -= wp.wspeed
		case Right:
			wp.WX += wp.wspeed
		case Up:
			wp.WY -= wp.wspeed
		case Down:
			wp.WY += wp.wspeed
		}
	}
}

//wp attack animation
func takingAmn(wp *Wpn, geom *ebiten.GeoM) (int, int) {
	var amx = 0
	var amy = 0
	geom.Translate(-wpnSize, -wpnSize)
	if wp.Stat == Taken {
		switch wp.Face {
		case Left:
			amx, amy = 12, 26
			if wp.AttackTime > 0 {
				geom.Rotate(float64(wp.AttackTime*6-180) * math.Pi / 360) //0->-180
			} else {
				geom.Rotate(float64(0) * math.Pi / 360)
			}
		case Up:
			amx, amy = 15, 12
			if wp.AttackTime > 0 {
				geom.Rotate(float64(wp.AttackTime*6) * math.Pi / 360) //180->0
			} else {
				geom.Rotate(float64(180) * math.Pi / 360)
			}
		case Right:
			amx, amy = 9, 26
			if wp.AttackTime > 0 {
				geom.Rotate(float64(360-wp.AttackTime*6) * math.Pi / 360) //180->360
			} else {
				geom.Rotate(float64(180) * math.Pi / 360)
			}
		case Down:
			amx, amy = 6, 24
			if wp.AttackTime > 0 {
				geom.Rotate(float64(wp.AttackTime*6-360) * math.Pi / 360) //-180->-360
			} else {
				geom.Rotate(float64(-180) * math.Pi / 360)
			}
		}
	}
	return amx, amy
}

//wp attack animation
func throwingAmn(wp *Wpn, geom *ebiten.GeoM) {
	per := wp.Flytime % 10
	geom.Translate(-wpnSize/2, -wpnSize/2)
	if wp.Face == Right {
		geom.Rotate(float64(-per*36) * math.Pi / 360)
	} else {
		geom.Rotate(float64(per*36) * math.Pi / 360)
	}
}
