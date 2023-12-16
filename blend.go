package gena

import (
	"image"
	"image/color"
	"math"
)

type HSV struct {
	H, S, V int
}

// ToRGB converts a HSV color mode to RGB mode
// mh, ms, mv are used to set the maximum number for HSV.
func (hs HSV) ToRGB(mh, ms, mv int) color.RGBA {
	hs.H = min(hs.H, mh)
	hs.S = min(hs.S, ms)
	hs.V = min(hs.V, mv)

	h, s, v := float64(hs.H)/float64(mh), float64(hs.S)/float64(ms), float64(hs.V)/float64(mv)

	if s == 0 { // HSV from 0 to 1
		return color.RGBA{
			R: uint8(v * 255),
			G: uint8(v * 255),
			B: uint8(v * 255),
			A: 0,
		}
	}

	h *= 6
	if h == 6 {
		h = 0
	} // H must be < 1

	i := math.Floor(h)
	v1 := v * (1 - s)
	v2 := v * (1 - s*(h-i))
	v3 := v * (1 - s*(1-(h-i)))

	var r, g, b float64
	switch i {
	case 0:
		r, g, b = v, v3, v1
	case 1:
		r, g, b = v2, v, v1
	case 2:
		r, g, b = v1, v, v3
	case 3:
		r, g, b = v1, v2, v
	case 4:
		r, g, b = v3, v1, v
	default:
		r, g, b = v, v1, v2
	}

	return color.RGBA{
		// RGB results from 0 to 255
		R: uint8(r * 255),
		G: uint8(g * 255),
		B: uint8(b * 255),
		A: 0,
	}
}

type BlendMode int

const (
	Add BlendMode = iota
)

func Blend(src, dest *image.RGBA, mode BlendMode) *image.RGBA {
	img := image.NewRGBA(src.Bounds())
	for i := 0; i < src.Bounds().Max.X; i++ {
		for j := 0; j < src.Bounds().Max.Y; j++ {
			switch mode {
			case Add:
				if src.RGBAAt(i, j) == Black {
					img.Set(i, j, dest.RGBAAt(i, j))
				} else {
					img.SetRGBA(i, j, Mix(src.RGBAAt(i, j), dest.RGBAAt(i, j)))
				}
			}
		}
	}
	return img
}

func LerpV2(a, b V2, t float64) V2 {
	return a + Mul(b-a, t)
}

func Lerp(c1, c2, coeff float64) float64 {
	return c1*(1-coeff) + c2*coeff
}

func lerp8(c1, c2 uint8, coeff float64) uint8 {
	return uint8(float64(c1)*(1-coeff) + float64(c2)*coeff)
}

// LerpColor blends two colors to find a third color somewhere between them.
func LerpColor(c1, c2 color.RGBA, ratio float64) color.RGBA {
	return color.RGBA{
		lerp8(c1.R, c2.R, 1-ratio),
		lerp8(c1.G, c2.G, 1-ratio),
		lerp8(c1.B, c2.B, 1-ratio),
		lerp8(c1.A, c2.A, 1-ratio),
	}
}

func Mix(src, dst color.RGBA) color.RGBA {
	aSrc := 1 - float64(src.A)/255.0
	return color.RGBA{
		R: lerp8(src.R, dst.R, aSrc),
		G: lerp8(src.G, dst.G, aSrc),
		B: lerp8(src.B, dst.B, aSrc),
		A: lerp8(src.A, dst.A, aSrc),
	}
}
