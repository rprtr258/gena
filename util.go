package gena

import (
	"fmt"
	"image"
	"image/draw"
	"image/jpeg"
	"image/png"
	"math"
	"os"
	"strings"

	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
)

func LoadImage(path string) image.Image {
	file := must1(os.Open(path))
	defer file.Close()

	im, _, err := image.Decode(file)
	must(err)
	return im
}

func LoadPNG(path string) (image.Image, error) {
	file := must1(os.Open(path))
	defer file.Close()

	return png.Decode(file)
}

func SavePNG(path string, im image.Image) {
	file := must1(os.Create(path))
	defer file.Close()

	must(png.Encode(file, im))
}

func LoadJPG(path string) image.Image {
	file := must1(os.Open(path))
	defer file.Close()

	return must1(jpeg.Decode(file))
}

func SaveJPG(path string, im image.Image, quality int) error {
	file := must1(os.Create(path))
	defer file.Close()

	return jpeg.Encode(file, im, &jpeg.Options{
		Quality: quality,
	})
}

func imageToRGBA(src image.Image) *image.RGBA {
	bounds := src.Bounds()
	dst := image.NewRGBA(bounds)
	draw.Draw(dst, bounds, src, bounds.Min, draw.Src)
	return dst
}

func parseHexColor(x string) (r, g, b, a int) {
	x = strings.TrimPrefix(x, "#")
	a = 255
	switch len(x) {
	case 3:
		format := "%1x%1x%1x"
		fmt.Sscanf(x, format, &r, &g, &b)
		r |= r << 4
		g |= g << 4
		b |= b << 4
	case 6:
		format := "%02x%02x%02x"
		fmt.Sscanf(x, format, &r, &g, &b)
	case 8:
		format := "%02x%02x%02x%02x"
		fmt.Sscanf(x, format, &r, &g, &b, &a)
	}
	return
}

func fixp(x, y float64) fixed.Point26_6 {
	return fixed.Point26_6{fix(x), fix(y)}
}

func Fixed(v V2) fixed.Point26_6 {
	return fixp(X(v), Y(v))
}

func fix(x float64) fixed.Int26_6 {
	return fixed.Int26_6(math.Round(x * 64))
}

func unfix(x fixed.Int26_6) float64 {
	const shift, mask = 6, 1<<6 - 1
	if x >= 0 {
		return float64(x>>shift) + float64(x&mask)/64
	}

	x = -x
	if x >= 0 {
		return -(float64(x>>shift) + float64(x&mask)/64)
	}

	return 0
}

// LoadFontFace is a helper function to load the specified font file with
// the specified point size. Note that the returned `font.Face` objects
// are not thread safe and cannot be used in parallel across goroutines.
// You can usually just use the Context.LoadFontFace function instead of
// this package-level function.
func LoadFontFace(path string, points float64) font.Face {
	fontBytes := must1(os.ReadFile(path))
	f := must1(truetype.Parse(fontBytes))
	return truetype.NewFace(f, &truetype.Options{
		Size: points,
		// Hinting: font.HintingFull,
	})
}
