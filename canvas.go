package gena

import (
	"bytes"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"image/png"
	"os"
)

func must(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func must1[T any](res T, err error) T {
	if err != nil {
		panic(err.Error())
	}

	return res
}

type Canvas struct {
	height, width int
	img           *image.RGBA
	LineColor     color.RGBA
	LineWidth     float64
	ColorSchema   []color.RGBA
	Alpha         int
}

// NewCanvas returns a Canvas
func NewCanvas(w, h int) Canvas {
	return Canvas{
		height: h,
		width:  w,
		img:    image.NewRGBA(image.Rect(0, 0, w, h)),
		// Set some defaults value
		LineColor:   Tomato,
		LineWidth:   3,
		ColorSchema: Youthful,
		Alpha:       255,
	}
}

func (c Canvas) Img() *image.RGBA {
	return c.img
}

func (c Canvas) Width() int {
	return c.width
}

func (c Canvas) Height() int {
	return c.height
}

func (c Canvas) Size() V2 {
	return complex(float64(c.Width()), float64(c.Height()))
}

func (c Canvas) SetColorSchema(rgbas []color.RGBA) Canvas {
	c.ColorSchema = rgbas
	return c
}

func (c Canvas) SetLineColor(rgba color.RGBA) Canvas {
	c.LineColor = rgba
	return c
}

func (c Canvas) SetLineWidth(lw float64) Canvas {
	c.LineWidth = lw
	return c
}

func (c Canvas) SetAlpha(alpha int) Canvas {
	c.Alpha = alpha
	return c
}

// FillBackground fills the background of the Canvas
func (c Canvas) FillBackground(bg color.RGBA) {
	draw.Draw(c.Img(), c.Img().Bounds(), &image.Uniform{bg}, image.Point{}, draw.Src)
}

// ToPng saves the image to local with PNG format.
func (c Canvas) ToPNG(fpath string) {
	f := must1(os.Create(fpath))
	defer f.Close()

	must(png.Encode(f, c.Img()))
}

// ToJpeg saves the image to local with Jpeg format.
func (c Canvas) ToJPEG(path string) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	return jpeg.Encode(f, c.Img(), nil)
}

// ToBytes returns the image as a jpeg-encoded []byte
func (c Canvas) ToBytes() []byte {
	var buffer bytes.Buffer
	must(jpeg.Encode(&buffer, c.Img(), nil))

	return buffer.Bytes()
}
