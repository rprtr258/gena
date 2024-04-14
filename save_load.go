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

func LoadPNG(path string) image.Image {
	file := must1(os.Open(path))
	defer file.Close()

	return must1(png.Decode(file))
}

// SavePNG saves the image to local with PNG format.
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

// SaveJPG saves the image to local with Jpeg format.
func SaveJPG(path string, im image.Image, quality int) {
	file := must1(os.Create(path))
	defer file.Close()

	must(jpeg.Encode(file, im, &jpeg.Options{
		Quality: quality,
	}))
}

// ToBytes returns the image as a jpeg-encoded []byte
func ToBytes(img image.Image) []byte {
	var buffer bytes.Buffer
	must(jpeg.Encode(&buffer, img, nil))

	return buffer.Bytes()
}
