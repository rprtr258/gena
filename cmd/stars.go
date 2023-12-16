package main

import (
	"math"
	"math/rand"

	"github.com/rprtr258/gena"
)

func Polygon(n int) []gena.V2 {
	result := make([]gena.V2, n)
	for i := 0; i < n; i++ {
		result[i] = gena.Rotation(math.Pi * (float64(i)*2/float64(n) - 1/2))
	}
	return result
}

func stars() {
	const W = 1200
	const H = 120
	const S = 100
	dc := gena.NewContext(W, H)
	dc.SetHexColor("#FFFFFF")
	dc.Clear()
	n := 5
	points := Polygon(n)
	for x := S / 2; x < W; x += S {
		dc.Stack(func(ctx *gena.Context) {
			s := rand.Float64()*S/4 + S/4
			dc.Translate(complex(float64(x), H/2))
			dc.Rotate(rand.Float64() * 2 * math.Pi)
			dc.Scale(complex(s, s))
			for i := 0; i < n+1; i++ {
				dc.LineToV2(points[(i*2)%n])
			}
			dc.SetLineWidth(10)
			dc.SetHexColor("#FFCC00")
			dc.StrokePreserve()
			dc.SetHexColor("#FFE43A")
			dc.Fill()
		})
	}
	dc.SavePNG("stars.png")
}
