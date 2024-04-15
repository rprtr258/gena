package gena

import (
	"bytes"
	"image"
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

// SavePNG saves the image to local with PNG format.
func SavePNG(path string, im image.Image) {
	file := must1(os.Create(path))
	defer file.Close()

	must(png.Encode(file, im))
}

// SaveJPG saves the image to local with Jpeg format.
//   - quality ranges from 1 to 100 inclusive, higher is better.
func SaveJPG(path string, im image.Image, quality uint8) {
	file := must1(os.Create(path))
	defer file.Close()

	must(jpeg.Encode(file, im, &jpeg.Options{
		Quality: int(quality),
	}))
}

// ToBytes returns the image as a jpeg-encoded []byte
func ToBytes(img image.Image) []byte {
	var buffer bytes.Buffer
	must(jpeg.Encode(&buffer, img, nil))

	return buffer.Bytes()
}
