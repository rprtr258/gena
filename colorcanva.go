package gena

import (
	"image/color"
	"math/rand"
)

type colorCanva struct {
	seg float64
}

// ColorCanva returns a color canva image.
func ColorCanva(c Canvas, colorSchema []color.RGBA, lineWidth float64, seg float64) {
	cc := &colorCanva{seg: seg}
	dc := NewContextForRGBA(c.Img())
	dc.SetLineWidth(lineWidth)

	rects := make([]Rect, 0)
	w := float64(c.Width) / cc.seg
	for i := 0.; i < cc.seg; i += 1. {
		for j := 0.; j < cc.seg; j += 1. {
			rects = append(rects, Rect{
				Pos:  Plus(Mul(complex(i, j), w), w/2),
				Size: complex(w, w),
			})
		}
	}

	rand.Shuffle(len(rects), func(i, j int) {
		rects[i], rects[j] = rects[j], rects[i]
	})

	dc.Translate(c.Size() / 2)
	dc.Scale(complex(0.6, 0.6))
	dc.Translate(-c.Size() / 2)

	for i := 0; i < len(rects); i++ {
		cc.draw(c, dc, colorSchema, rects[i])
		cc.draw(c, dc, colorSchema, rects[i])
	}
}

func (cc *colorCanva) draw(c Canvas, dc *Context, colorSchema []color.RGBA, rect Rect) {
	wh := Mul2(rect.Size/5, Mul(RandomV2(), cc.seg*2)+1)
	switch rand.Intn(4) {
	case 0:
		dc.DrawRectangle(complex(X(rect.Pos)-X(wh)/2+X(rect.Size)/2, Y(rect.Pos)+Y(wh)/2+Y(rect.Size)/2), wh)
	case 1:
		dc.DrawRectangle(complex(X(rect.Pos)-X(wh)/2-X(rect.Size)/2, Y(rect.Pos)+Y(wh)/2+Y(rect.Size)/2), wh)
	case 2:
		dc.DrawRectangle(complex(X(rect.Pos)-X(wh)/2+X(rect.Size)/2, Y(rect.Pos)+Y(wh)/2-Y(rect.Size)/2), wh)
	case 3:
		dc.DrawRectangle(complex(X(rect.Pos)-X(wh)/2-X(rect.Size)/2, Y(rect.Pos)+Y(wh)/2-Y(rect.Size)/2), wh)
	}
	dc.SetColor(Black)
	dc.StrokePreserve()
	dc.SetColor(colorSchema[rand.Intn(len(colorSchema))])
	dc.Fill()
}
