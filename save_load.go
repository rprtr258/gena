package gena

import (
	"bytes"
	"image"
	"image/color/palette"
	"image/draw"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
)

func Load(path string) image.Image {
	file := must1(os.Open(path))
	defer file.Close()

	im, _, err := image.Decode(file)
	must(err)
	return im
}

// SavePNG saves the image to local with PNG format
func SavePNG(path string, im image.Image) {
	file := must1(os.Create(path))
	defer file.Close()

	must(png.Encode(file, im))
}

// SaveJPG saves the image to local with Jpeg format
//   - quality: ranges from 1 to 100 inclusive, higher is better
func SaveJPG(path string, im image.Image, quality uint8) {
	file := must1(os.Create(path))
	defer file.Close()

	must(jpeg.Encode(file, im, &jpeg.Options{
		Quality: int(quality),
	}))
}

// SaveGIF saves images as gif animation frames
func SaveGIF(path string, frames ...*image.RGBA) {
	file := must1(os.Create(path))
	defer file.Close()

	images := make([]*image.Paletted, len(frames))
	delays := make([]int, len(frames))
	for i, im := range frames {
		pim := image.NewPaletted(image.Rect(0, 0, im.Bounds().Dx(), im.Bounds().Dy()), palette.WebSafe)
		draw.Draw(pim, pim.Bounds(), im, image.Point{}, draw.Src)
		images[i] = pim
		delays[i] = 10
	}

	must(gif.EncodeAll(file, &gif.GIF{
		Image: images,
		Delay: delays,
	}))
}

// ToBytes returns the image as a png-encoded []byte
func ToBytes(img image.Image) []byte {
	var buffer bytes.Buffer
	must(png.Encode(&buffer, img))

	return buffer.Bytes()
}
