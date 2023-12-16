package gena

import (
	"bytes"
	"image"
	"image/jpeg"
	"image/png"
	"os"
)

// ToPng saves the image to local with PNG format.
func ToPNG(img image.Image, fpath string) {
	f := must1(os.Create(fpath))
	defer f.Close()

	must(png.Encode(f, img))
}

// ToJpeg saves the image to local with Jpeg format.
func ToJPEG(img image.Image, path string) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	return jpeg.Encode(f, img, nil)
}

// ToBytes returns the image as a jpeg-encoded []byte
func ToBytes(img image.Image) []byte {
	var buffer bytes.Buffer
	must(jpeg.Encode(&buffer, img, nil))

	return buffer.Bytes()
}
