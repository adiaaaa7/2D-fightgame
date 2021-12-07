package main

import (
	"bytes"
	"fmt"
	"game/myds/elm/png/wpn"

	"image"
	_ "image/png"
	"log"
	"math"

	"golang.org/x/image/math/f64"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	screenWidth  = 320
	screenHeight = 480
)

const (
	pcWidth  = 32
	pcHeight = 48
)

var (
	stu *ebiten.Image
)

func init() {
	//img, _, err := image.Decode(bytes.NewReader(st.Su1_png))
	img, _, err := image.Decode(bytes.NewReader(wpn.Wswd13_png))
	if err != nil {
		log.Fatal(err)
	}
	stu = ebiten.NewImageFromImage(img)
}

type Camera struct {
	ViewPort   f64.Vec2
	Position   f64.Vec2
	Skew       f64.Vec2
	ZoomFactor int
	Rotation   int
}

func (c *Camera) String() string {
	return fmt.Sprintf(
		"T: %.1f,V: %.1f, R: %d, S: %d",
		c.Position, c.ViewPort, c.Rotation, c.ZoomFactor,
	)
}

func (c *Camera) worldMatrix() ebiten.GeoM {
	m := ebiten.GeoM{}
	m.Translate(-float64(34), -float64(34))
	//m.Translate(c.Position[0], c.Position[1])
	//m.Translate(pcWidth/2, pcHeight/2)
	// m.Scale(
	// 	math.Pow(1.01, float64(c.ZoomFactor)),
	// 	math.Pow(1.01, float64(c.ZoomFactor)),
	// )

	//m.Skew(c.Skew[0], c.Skew[1])
	m.Rotate(float64(c.Rotation) * math.Pi / 360)
	//m.Translate(c.ViewPort[0], c.ViewPort[1])
	m.Translate(160.0, 240.0)
	m.Translate(c.Position[0], c.Position[1])
	return m
}

func (c *Camera) Reset() {
	c.Position[0] = 0
	c.Position[1] = 0
	c.ViewPort[0] = 0
	c.ViewPort[1] = 0
	c.Skew[0] = 0
	c.Skew[1] = 0
	c.Rotation = 0
	c.ZoomFactor = 0
}

type Game struct {
	camera Camera
}

func (g *Game) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		g.camera.Position[0] -= 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		g.camera.Position[0] += 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		g.camera.Position[1] -= 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		g.camera.Position[1] += 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		g.camera.ViewPort[0] -= 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		g.camera.ViewPort[0] += 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		g.camera.ViewPort[1] -= 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		g.camera.ViewPort[1] += 1
	}

	if ebiten.IsKeyPressed(ebiten.KeyQ) {
		if g.camera.ZoomFactor > -2400 {
			g.camera.ZoomFactor -= 1
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyE) {
		if g.camera.ZoomFactor < 2400 {
			g.camera.ZoomFactor += 1
		}
	}

	if ebiten.IsKeyPressed(ebiten.KeyJ) {
		g.camera.Skew[0] -= 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyL) {
		g.camera.Skew[0] += 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyI) {
		g.camera.Skew[1] -= 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyK) {
		g.camera.Skew[1] += 1
	}

	if ebiten.IsKeyPressed(ebiten.KeyO) {
		g.camera.Rotation -= 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyP) {
		g.camera.Rotation += 1
	}

	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		g.camera.Reset()
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM = g.camera.worldMatrix()

	screen.DrawImage(stu.SubImage(image.Rect(0, 0, pcWidth, pcHeight)).(*ebiten.Image), op)

	ebitenutil.DebugPrint(
		screen,
		fmt.Sprintf("TPS: %0.2f\nMove (WASD/Arrows)\nZoom (QE)\nRotate (R)\nReset (Space)", ebiten.CurrentTPS()),
	)

	ebitenutil.DebugPrintAt(
		screen,
		fmt.Sprintf("%s\n",
			g.camera.String()),
		0, screenHeight-32,
	)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	g := &Game{}
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("anitool")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
