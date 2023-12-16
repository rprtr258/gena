package gena

import (
	"math/rand"
)

type colorCanva struct {
	seg float64
}

// ColorCanva returns a color canva image.
func ColorCanva(c Canvas, lineWidth float64, seg float64) {
	cc := &colorCanva{seg: seg}
	ctex := NewContextForRGBA(c.Img())
	ctex.SetLineWidth(lineWidth)

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

	ctex.Translate(c.Size() / 2)
	ctex.Scale(complex(0.6, 0.6))
	ctex.Translate(-c.Size() / 2)

	for i := 0; i < len(rects); i++ {
		cc.draw(c, ctex, rects[i])
		cc.draw(c, ctex, rects[i])
	}
}

func (cc *colorCanva) draw(c Canvas, ctex *Context, rect Rect) {
	wh := Mul2(rect.Size/5, Mul(RandomV2(), cc.seg*2)+1)
	switch rand.Intn(4) {
	case 0:
		ctex.DrawRectangle(complex(X(rect.Pos)-X(wh)/2+X(rect.Size)/2, Y(rect.Pos)+Y(wh)/2+Y(rect.Size)/2), wh)
	case 1:
		ctex.DrawRectangle(complex(X(rect.Pos)-X(wh)/2-X(rect.Size)/2, Y(rect.Pos)+Y(wh)/2+Y(rect.Size)/2), wh)
	case 2:
		ctex.DrawRectangle(complex(X(rect.Pos)-X(wh)/2+X(rect.Size)/2, Y(rect.Pos)+Y(wh)/2-Y(rect.Size)/2), wh)
	case 3:
		ctex.DrawRectangle(complex(X(rect.Pos)-X(wh)/2-X(rect.Size)/2, Y(rect.Pos)+Y(wh)/2-Y(rect.Size)/2), wh)
	}
	ctex.SetColor(Black)
	ctex.StrokePreserve()
	ctex.SetColor(c.ColorSchema[rand.Intn(len(c.ColorSchema))])
	ctex.Fill()
}
