package main

import (
	"image"
	"image/color"

	. "github.com/rprtr258/gena"
)

type colorCircle2 struct {
	circleNum int
}

func (cc *colorCircle2) circle(dc *Context, colorSchema []color.RGBA, v V2, d, dx float64) {
	col := RandomItem(colorSchema)
	for j := 0.0; j < dx; j += 1.0 {
		dd := d + j*2.0
		alpha := min(int((dx/j)*255.0), 255)
		col.A = uint8(alpha)
		dc.SetColor(col)
		for range Range(150) {
			theta := RandomF64(0, PI*2)
			dc.DrawPoint(v+Polar(dd/2, theta), 0.51)
			dc.Fill()
		}
	}
}

// ColorCircle2 is version 2 of ColorCircle.
// It still draws the circle and point cloud.
//   - n: number of circle
func ColorCircle2(dc *Context, colorSchema []color.RGBA, n int) {
	cc := &colorCircle2{
		circleNum: n,
	}

	dc.SetColor(White)
	dc.Clear()

	im := dc.Image()
	for range Range(cc.circleNum) {
		v := Mul2(RandomV2(), Size(im))

		r1 := RandomF64(50.0, float64(im.Bounds().Dx())/4)
		r2 := RandomF64(10.0, float64(im.Bounds().Dx())/3)

		cc.circle(dc, colorSchema, v, r1, r2)
		if Random() < 0.3 {
			col := RandomItem(colorSchema)
			dc.SetColor(col)
			dc.DrawCircle(v, r1/2.0)
			dc.Fill()
		}
	}
}

func colorcircle2() *image.RGBA {
	dc := NewContext(Diag(800))
	ColorCircle2(dc, []color.RGBA{
		{0x11, 0x60, 0xC6, 0xFF},
		{0xFD, 0xD9, 0x00, 0xFF},
		{0xF5, 0xB4, 0xF8, 0xFF},
		{0xEF, 0x13, 0x55, 0xFF},
		{0xF4, 0x9F, 0x0A, 0xFF},
	}, 30)
	return dc.Image()
}
