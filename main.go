package main

import (
	"game/myds/elm"
	"game/myds/elm/png/rnd"
	"image"
	_ "image/png"
	"log"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	cp    elm.CameraP
	ops   elm.Ops
	wps   elm.Wpns
	count int
	rX    int
}

func NewGame() *Game {
	cp0 := &elm.CameraP{}
	cp0.InitialPosition()
	g := &Game{cp: *cp0}
	g.rX = cp0.PX
	g.init()
	return g
}

func (g *Game) init() {
	rand.Seed(time.Now().UnixNano())
	img1 := rnd.MpRand(rand.Intn(8) + 1)
	elm.BgImage = ebiten.NewImageFromImage(img1)
}

func (g *Game) Update() error {
	g.count++
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		g.rX -= elm.Bgspeed
		g.cp.Mv = elm.Left
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		g.rX += elm.Bgspeed
		g.cp.Mv = elm.Right
	}
	g.rX = elm.BgBoundrs(g.rX)
	g.cp.Update(g.rX, g.ops.Others, g.wps.Wpns)
	g.otherCret()
	g.wpnCret()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Reset()
	lcx := elm.ScreenWidth / 2
	rcx := elm.BgWidth - elm.ScreenWidth/2
	var r image.Rectangle
	//left camera
	if g.rX < lcx {
		r = image.Rect(0, 0, elm.ScreenWidth, elm.ScreenHeight)
	}
	//middle camera
	if g.rX >= lcx && g.rX < rcx {
		r = image.Rect(g.rX-elm.ScreenWidth/2, 0, g.rX+elm.ScreenWidth/2, elm.ScreenHeight)
	}
	//right camera
	if g.rX >= rcx && g.rX <= elm.BgWidth {
		r = image.Rect(elm.BgWidth-elm.ScreenWidth, 0, elm.BgWidth, elm.ScreenHeight)
	}
	screen.DrawImage(elm.BgImage.SubImage(r).(*ebiten.Image), op)

	//drwa pc
	g.cp.Draw(screen)
	//draw others
	for _, o := range g.ops.Others {
		if o != nil {
			o.Draw(screen)
		}
	}
	//drwa wps
	for _, w := range g.wps.Wpns {
		if w != nil {
			w.Draw(screen)
		}
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return elm.ScreenWidth, elm.ScreenHeight
}

func main() {
	ebiten.SetWindowSize(elm.ScreenWidth, elm.ScreenHeight)
	ebiten.SetWindowTitle("my ft ds")
	g := NewGame()
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}

//enmy creat and delete
func (g *Game) otherCret() {
	if g.count%240 == 0 && g.ops.Num < elm.Onum {
		opNew := elm.Other{}
		opNew.InitialPosition()
		g.ops.Num += 1
		g.ops.Others = append(g.ops.Others, &opNew)
	}

	if g.ops.Num >= 1 {
		for i := g.ops.Num; i > 0; i-- {
			g.ops.Others[i-1].Update(&g.cp)
			if g.ops.Others[i-1].DeathTime == 1 {
				g.ops.Others = append(g.ops.Others[:i-1], g.ops.Others[i:]...)
				g.ops.Num--
			}
		}
	}
}

//wpn creat and delete
func (g *Game) wpnCret() {
	if g.count%360 == 0 && g.wps.Num < elm.Wpnum {
		wpNew := elm.Wpn{}
		wpNew.InitialPosition()
		g.wps.Num += 1
		g.wps.Wpns = append(g.wps.Wpns, &wpNew)
	}
	if g.wps.Num >= 1 {
		for i := g.wps.Num; i > 0; i-- {
			g.wps.Wpns[i-1].Update(&g.cp, g.ops.Others)
			if g.wps.Wpns[i-1].Flytime < 0 {
				g.wps.Num -= 1
				g.wps.Wpns = append(g.wps.Wpns[:i-1], g.wps.Wpns[i:]...)
			}
		}
	}
}
