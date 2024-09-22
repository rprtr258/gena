package main

import (
	"image"
	"image/color"
	"math/cmplx"

	. "github.com/rprtr258/gena"
)

type Rect struct {
	Pos, Size V2
}

// ColorCanvas returns a color canva image.
func ColorCanvas(dc *Context, colorSchema []color.RGBA, lineWidth, seg float64) {
	im := dc.Image()

	rects := make([]Rect, 0)
	w := float64(im.Bounds().Dx()) / seg
	for i := 0.; i < seg; i += 1. {
		for j := 0.; j < seg; j += 1. {
			rects = append(rects, Rect{
				Pos:  P(i, j)*Coeff(w) + Diag(w/2),
				Size: P(w, w),
			})
		}
	}

	Shuffle(rects)

	dc.SetColor(Black)
	dc.Clear()

	dc.SetLineWidth(lineWidth)
	dc.TransformAdd(Translate(Size(im) / 2))
	dc.TransformAdd(Scale(complex(0.6, 0.6)))
	dc.TransformAdd(Translate(-Size(im) / 2))

	for i := range rects {
		drawColorCanva(dc, seg, colorSchema, rects[i])
		drawColorCanva(dc, seg, colorSchema, rects[i])
	}
}

func drawColorCanva(dc *Context, seg float64, colorSchema []color.RGBA, rect Rect) {
	wh := Mul2(rect.Size/5, RandomV2()*Coeff(seg*2)+1)

	var delta V2
	switch RandomInt(4) {
	case 0:
		delta = P(-rect.Size.X(), -rect.Size.Y())
	case 1:
		delta = P(rect.Size.X(), -rect.Size.Y())
	case 2:
		delta = P(-rect.Size.X(), rect.Size.Y())
	case 3:
		delta = P(rect.Size.X(), rect.Size.Y())
	}
	dc.DrawRectangle(rect.Pos-V2(cmplx.Conj(complex128(wh)))/2+delta/2, wh)

	dc.SetColor(Black)
	dc.StrokePreserve()
	dc.SetColor(RandomItem(colorSchema))
	dc.Fill()
}

func colorcanvas() *image.RGBA {
	dc := NewContext(Diag(500))
	ColorCanvas(dc, []color.RGBA{
		{0xF9, 0xC8, 0x0E, 0xFF},
		{0xF8, 0x66, 0x24, 0xFF},
		{0xEA, 0x35, 0x46, 0xFF},
		{0x66, 0x2E, 0x9B, 0xFF},
		{0x43, 0xBC, 0xCD, 0xFF},
	}, 8, 5)
	return dc.Image()
}
