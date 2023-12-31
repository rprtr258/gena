package gena

import (
	"image"
	"image/color"
	"math/cmplx"
	"math/rand"
)

type Rect struct {
	Pos, Size V2
}

// ColorCanva returns a color canva image.
func ColorCanva(c *image.RGBA, colorSchema []color.RGBA, lineWidth float64, seg float64) {
	dc := NewContextForRGBA(c)
	dc.SetLineWidth(lineWidth)

	rects := make([]Rect, 0)
	w := float64(c.Bounds().Dx()) / seg
	for i := 0.; i < seg; i += 1. {
		for j := 0.; j < seg; j += 1. {
			rects = append(rects, Rect{
				Pos:  Plus(Mul(complex(i, j), w), w/2),
				Size: complex(w, w),
			})
		}
	}

	rand.Shuffle(len(rects), func(i, j int) {
		rects[i], rects[j] = rects[j], rects[i]
	})

	dc.Translate(Size(c) / 2)
	dc.Scale(complex(0.6, 0.6))
	dc.Translate(-Size(c) / 2)

	for i := range len(rects) {
		drawColorCanva(dc, seg, colorSchema, rects[i])
		drawColorCanva(dc, seg, colorSchema, rects[i])
	}
}

func drawColorCanva(dc *Context, seg float64, colorSchema []color.RGBA, rect Rect) {
	wh := Mul2(rect.Size/5, Mul(RandomV2(), seg*2)+1)

	var delta V2
	switch rand.Intn(4) {
	case 0:
		delta = complex(-X(rect.Size), -Y(rect.Size))
	case 1:
		delta = complex(X(rect.Size), -Y(rect.Size))
	case 2:
		delta = complex(-X(rect.Size), Y(rect.Size))
	case 3:
		delta = complex(X(rect.Size), Y(rect.Size))
	}
	dc.DrawRectangle(rect.Pos-cmplx.Conj(wh)/2+delta/2, wh)

	dc.SetColor(Black)
	dc.StrokePreserve()
	dc.SetColor(colorSchema[rand.Intn(len(colorSchema))])
	dc.Fill()
}
