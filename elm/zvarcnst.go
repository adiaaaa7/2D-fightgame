package elm

import (
	"bytes"
	"game/myds/elm/png/pc"
	"game/myds/elm/png/wpn"
	"image"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

var (
	BgImage  *ebiten.Image
	PcImages []*ebiten.Image
	WpImages []*ebiten.Image
	PcImage  *ebiten.Image
)

type Move int
type Shoot int
type Where int

const (
	ScreenWidth  = 640
	ScreenHeight = 480
	Bgspeed      = 3
	PcWidth      = 32
	PcHight      = 48
	wpnSize      = 34
	jumpS        = 1
	Onum         = 2
	Wpnum        = 2
	oPxy         = ScreenHeight / 2
	PymaxB       = ScreenHeight - PcHight
	PyminB       = ScreenHeight - 5*PcHight
	PrxmaxB      = BgWidth - PcWidth/2
	PxmaxB       = ScreenWidth - PcWidth/2
	PxminB       = PcWidth / 2
	BgWidth      = ScreenWidth * 4
	pspeed       = Bgspeed
	pjumpH       = PcHight / 8
)
const (
	Down Move = iota
	Left
	Right
	Up
)

const (
	Ground Where = iota
	Taken
	Throw
)

type CameraP struct {
	PX      int
	PY      int
	Mv      Move
	jumping bool
	jumpH   int
	jumpA   int
	Armed   bool
}

type Other struct {
	OP        CameraP
	OpcImage  *ebiten.Image
	DeathTime int
	DeathFace Move
	walkTime  int
}

type Wpn struct {
	WX         int
	WY         int
	Face       Move
	WpnImage   *ebiten.Image
	Stat       Where
	AttackTime int
	wspeed     int
	chargeTime int
	Flytime    int
	WpnFace    Move
}

type Wpns struct {
	Wpns []*Wpn
	Num  int
}

type Ops struct {
	Others []*Other
	Num    int
}

func RndMov(rnd int) Move {
	switch rnd {
	case 0:
		return Left
	case 1:
		return Down
	case 2:
		return Right
	case 3:
		return Up
	}
	return Left
}

func Stuinit() {
	PcImages = make([]*ebiten.Image, 33)
	PcImages[0] = addImage(pc.Su1_png)
	PcImages[1] = addImage(pc.Su2_png)
	PcImages[2] = addImage(pc.Su3_png)
	PcImages[3] = addImage(pc.Su4_png)
	PcImages[4] = addImage(pc.Su5_png)
	PcImages[5] = addImage(pc.Su6_png)
	PcImages[6] = addImage(pc.Su7_png)
	PcImages[7] = addImage(pc.Su8_png)
	PcImages[8] = addImage(pc.Su9_png)
	PcImages[9] = addImage(pc.Su10_png)
	PcImages[10] = addImage(pc.Su11_png)
	PcImages[11] = addImage(pc.Su12_png)
	PcImages[12] = addImage(pc.Su13_png)
	PcImages[13] = addImage(pc.Su14_png)
	PcImages[14] = addImage(pc.Su15_png)
	PcImages[15] = addImage(pc.Su16_png)
	PcImages[16] = addImage(pc.Su17_png)
	PcImages[17] = addImage(pc.Su18_png)
	PcImages[18] = addImage(pc.Su19_png)
	PcImages[19] = addImage(pc.Su20_png)
	PcImages[20] = addImage(pc.Su21_png)
	PcImages[21] = addImage(pc.Su22_png)
	PcImages[22] = addImage(pc.Su23_png)
	PcImages[23] = addImage(pc.Su24_png)
	PcImages[24] = addImage(pc.Su25_png)
	PcImages[25] = addImage(pc.Su26_png)
	PcImages[26] = addImage(pc.Su27_png)
	PcImages[27] = addImage(pc.Su28_png)
	PcImages[28] = addImage(pc.Su29_png)
	PcImages[29] = addImage(pc.Su30_png)
	PcImages[30] = addImage(pc.Su31_png)
	PcImages[31] = addImage(pc.Su32_png)
}

func Wswdinit() {
	WpImages = make([]*ebiten.Image, 22)
	WpImages[0] = addImage(wpn.Wswd1_png)
	WpImages[1] = addImage(wpn.Wswd2_png)
	WpImages[2] = addImage(wpn.Wswd3_png)
	WpImages[3] = addImage(wpn.Wswd4_png)
	WpImages[4] = addImage(wpn.Wswd5_png)
	WpImages[5] = addImage(wpn.Wswd6_png)
	WpImages[6] = addImage(wpn.Wswd7_png)
	WpImages[7] = addImage(wpn.Wswd8_png)
	WpImages[8] = addImage(wpn.Wswd9_png)
	WpImages[9] = addImage(wpn.Wswd10_png)
	WpImages[10] = addImage(wpn.Wswd11_png)
	WpImages[11] = addImage(wpn.Wswd12_png)
	WpImages[12] = addImage(wpn.Wswd13_png)
	WpImages[13] = addImage(wpn.Wswd14_png)
	WpImages[14] = addImage(wpn.Wswd15_png)
	WpImages[15] = addImage(wpn.Wswd16_png)
	WpImages[16] = addImage(wpn.Wswd17_png)
	WpImages[17] = addImage(wpn.Wswd18_png)
	WpImages[18] = addImage(wpn.Wswd19_png)
	WpImages[19] = addImage(wpn.Wswd20_png)
	WpImages[20] = addImage(wpn.Wswd21_png)
}

func addImage(b []byte) *ebiten.Image {
	img, _, err := image.Decode(bytes.NewReader(b))
	errLog(err)
	return ebiten.NewImageFromImage(img)
}

func errLog(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

//left right bounds
func BgBoundrs(x int) int {
	if x < PxminB {
		x = PxminB
	}
	if x >= PrxmaxB {
		x = PrxmaxB
	}
	return x
}

//pc up down bounds
func overYBoundrs(y int) int {
	if y < PyminB {
		y = PyminB
	}
	if y > PymaxB {
		y = PymaxB
	}
	return y
}
