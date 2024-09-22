package main

import (
	"image"
	"image/color"

	. "github.com/rprtr258/gena"
)

func drawOceanFishLines(dc *Context, size V2, colorSchema []color.RGBA, n int) {
	for range Range(n) {
		dc.SetColor(RandomItem(colorSchema))
		dc.SetLineWidth(RandomF64(9, 20))
		y := RandomF64(0, size.Y())
		dc.DrawLine(
			P(0, y+RandomF64(-50, 50)),
			P(size.X(), y+RandomF64(-50, 50)),
		)
		dc.Stroke()
	}
}

func fishPt(r, theta float64) V2 {
	return P(
		r*Cos(theta)-r*Pow(Sin(theta), 2)/Sqrt(2.0),
		r*Cos(theta)*Sin(theta),
	)
}

func drawfish(dc *Context, r float64) {
	dc.TransformAdd(Rotate(PI * 3 / 2))

	for _, a := range RangeF64(0, PI*2, 360) {
		dc.LineTo(fishPt(r, a))
	}
	dc.ClosePath()
}

// OceanFish draws an ocean and some fishes in the center.
//   - lineNum: The number of the line used to simulate the ocean wave.
//   - fishNum: The number of fish.
//
// TODO: make similar to original
func OceanFish(dc *Context, colorSchema []color.RGBA, lineNum, fishNum int) {
	im := dc.Image()
	W := Size(im).X()

	dc.SetColor(Black)
	dc.Clear()

	drawOceanFishLines(dc, Size(im), colorSchema, lineNum)

	dc.TransformAdd(Translate(Size(im) / 2))
	for i := range Range(fishNum) {
		dc.Stack(func(ctx *Context) {
			dc.TransformAdd(Translate(P(W/4, 0)))
			drawfish(dc, W/10)
			dc.Clip()

			dc.TransformAdd(Rotate(-float64(i) * PI * 2 / float64(fishNum)))
			dc.TransformAdd(Rotate(PI / 2))
			dc.TransformAdd(Translate(-P(W/4, 0)))
			drawOceanFishLines(dc, Size(im), colorSchema, lineNum)
			dc.TransformAdd(Rotate(float64(i) * PI * 2 / float64(fishNum)))
		})
		dc.ClearPath()
		dc.ResetClip()
		dc.TransformAdd(Rotate(PI * 2 / float64(fishNum)))
	}
}

func oceanfish() *image.RGBA {
	// colors := []color.RGBA{
	// 	{0x05, 0x1F, 0x34, 0xFF},
	// 	{0x02, 0x74, 0x95, 0xFF},
	// 	{0x01, 0xA9, 0xC1, 0xFF},
	// 	{0xBA, 0xD6, 0xDB, 0xFF},
	// 	{0xF4, 0xF5, 0xF5, 0xFF},
	// }
	colors := []color.RGBA{
		{0xCF, 0x2B, 0x34, 0xFF},
		{0xF0, 0x8F, 0x46, 0xFF},
		{0xF0, 0xC1, 0x29, 0xFF},
		{0x19, 0x6E, 0x94, 0xFF},
		{0x35, 0x3A, 0x57, 0xFF},
	}
	dc := NewContext(Diag(500))
	OceanFish(dc, colors, 100, 8)
	return dc.Image()
}
