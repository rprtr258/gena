package main

import (
	"fmt"
	"image"
	"image/color"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/schollz/progressbar/v3"

	. "github.com/rprtr258/gena"
)

func cmap(r, m1, m2 float64) color.RGBA {
	return color.RGBA{
		R: uint8(Clamp(m1*200*r, 0, 255)),
		G: uint8(Clamp(r*200, 0, 255)),
		B: uint8(Clamp(m2*255*r, 70, 255)),
		A: 255,
	}
}

func newImage(w, h int) *image.RGBA {
	return image.NewRGBA(image.Rect(0, 0, w, h))
}

func dotline() {
	im := newImage(2080, 2080)
	FillBackground(im, color.RGBA{230, 230, 230, 255})
	DotLine(im, DarkPink, 10, 100, 20, 50, false, 15000)
	SavePNG("dotline.png", im)
}

func dotswave() {
	im := newImage(500, 500)
	FillBackground(im, Black)
	DotsWave(im, []color.RGBA{
		{0xFF, 0xBE, 0x0B, 0xFF},
		{0xFB, 0x56, 0x07, 0xFF},
		{0xFF, 0x00, 0x6E, 0xFF},
		{0x83, 0x38, 0xEC, 0xFF},
		{0x3A, 0x86, 0xFF, 0xFF},
	}, 300)
	SavePNG("dotswave.png", im)
}

func gridsquare() {
	im := newImage(600, 600)
	GirdSquares(im, DarkPink, 24, 10, 0.2, 20)
	SavePNG("gsquare.png", im)
}

func _http() {
	// Initialize handlers
	http.HandleFunc("/",
		// drawHandler is writes a piece of generative art
		// as a response to an http request
		func(w http.ResponseWriter, r *http.Request) {
			// Log Requests
			log.Printf("method=%s path=%s ", r.Method, r.RequestURI)

			// Draw the image to bytes
			b := func() []byte {
				// Generate a new image
				im := newImage(500, 500)
				FillBackground(im, Black)
				Janus(im, DarkRed, LightPink, 0.2)

				// Return the image as []byte
				return ToBytes(im)
			}()

			// Set content headers
			w.Header().Set("Content-Type", "image/jpeg")
			w.Header().Set("Content-Length", strconv.Itoa(len(b)))
			_, _ = w.Write(b) // Write image to response
		})

	const address = ":8090"
	log.Println("Server started at", address)
	log.Fatal(http.ListenAndServe(address, nil))
}

func janus() {
	im := newImage(500, 500)
	FillBackground(im, Black)
	Janus(im, DarkRed, LightPink, 0.2)
	SavePNG("janus.png", im)
}

func julia() {
	im := newImage(500, 500)
	FillBackground(im, Azure)
	Julia(im, Viridis, func(z complex128) complex128 {
		return z*z + complex(-0.1, 0.651)
	}, 40, 1.5+1.5i, 800)
	SavePNG("julia.png", im)
}

func maze() {
	im := newImage(600, 600)
	FillBackground(im, Azure)
	Maze(im, 3, Orange, 20)
	SavePNG("maze.png", im)
}

func noiseline() {
	im := newImage(1000, 1000)
	FillBackground(im, color.RGBA{0xF0, 0xFE, 0xFF, 0xFF})
	NoiseLine(im, []color.RGBA{
		{0x06, 0x7B, 0xC2, 0xFF},
		{0x84, 0xBC, 0xDA, 0xFF},
		{0xEC, 0xC3, 0x0B, 0xFF},
		{0xF3, 0x77, 0x48, 0xFF},
		{0xD5, 0x60, 0x62, 0xFF},
	}, 1000)
	SavePNG("noiseline.png", im)
}

func oceanfish() {
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
	im := newImage(500, 500)
	OceanFish(im, colors, 100, 8)
	SavePNG("oceanfish.png", im)
}

func perlinpearls() {
	im := newImage(500, 500)
	FillBackground(im, White)
	PerlinPearls(im, 0.3, 120, 10, 200, 40, 80, 200)
	SavePNG("perlinperls.png", im)
}

func pixelhole() {
	im := newImage(800, 800)
	FillBackground(im, Black)
	PixelHole(im, []color.RGBA{
		{0xF9, 0xC8, 0x0E, 0xFF},
		{0xF8, 0x66, 0x24, 0xFF},
		{0xEA, 0x35, 0x46, 0xFF},
		{0x66, 0x2E, 0x9B, 0xFF},
		{0x43, 0xBC, 0xCD, 0xFF},
	}, 60, 1200)
	SavePNG("pixelhole.png", im)
}

func pointribbon() {
	im := newImage(500, 500)
	FillBackground(im, Lavender)
	PointRibbon(im, 2, 50, 150000)
	SavePNG("pointribbon.png", im)
}

func randcircle() {
	im := newImage(500, 500)
	FillBackground(im, MistyRose)
	RandCircle(im, Plasma, 1, color.RGBA{122, 122, 122, 30}, 30, 80, 0.2, 2, 10, 30, true, 4)
	SavePNG("randcircle.png", im)
}

func randomshape() {
	im := newImage(500, 500)
	FillBackground(im, White)
	RandomShape(im, []color.RGBA{
		{0xCF, 0x2B, 0x34, 0xFF},
		{0xF0, 0x8F, 0x46, 0xFF},
		{0xF0, 0xC1, 0x29, 0xFF},
		{0x19, 0x6E, 0x94, 0xFF},
		{0x35, 0x3A, 0x57, 0xFF},
	}, 150)
	SavePNG("randomshape.png", im)
}

func silksky() {
	im := newImage(600, 600)
	SilkSky(im, 10, 15, 5)
	SavePNG("silksky.png", im)
}

func silksmoke() {
	im := newImage(500, 500)
	FillBackground(im, Black)
	SilkSmoke(im, Plasma, 1, MediumAquamarine, 30, 400, 20, 0.2, 2, 10, 30, false)
	SavePNG("silksmoke.png", im)
}

func solarflare() {
	im := newImage(500, 500)
	FillBackground(im, Black)
	SolarFlare(im, color.RGBA{255, 64, 8, 128})
	SavePNG("solarflare.png", im)
}

func spiralsquare() {
	im := newImage(500, 500)
	FillBackground(im, MistyRose)
	SpiralSquare(im, Plasma, 10, Orange, 40, 400, 0.05, Tomato, true)
	SavePNG("spiralsquare.png", im)
}

func swirl() {
	bg := Azure
	im := newImage(1600, 1600)
	FillBackground(im, bg)
	Swirl(
		im,
		color.RGBA{113, 3, 0, 140}, bg,
		0.970, -1.899, 1.381, -1.506,
		2.4+2.4i,
		4000000,
	)
	SavePNG("swirl.png", im)
}

func yarn() {
	im := newImage(500, 500)
	FillBackground(im, Orange)
	Yarn(im, 0.3, color.RGBA{A: 60}, 2000)
	SavePNG("yarn.png", im)
}

func blackhole() {
	im := newImage(500, 500)
	FillBackground(im, color.RGBA{30, 30, 30, 255})
	BlackHole(im, 1, Tomato, 200, 400, 0.03, NewPerlinNoiseDeprecated(), RandomF64(0, 1))
	SavePNG("blackhole.png", im)
}

func circlegrid() {
	im := newImage(500, 500)
	FillBackground(im, color.RGBA{0xDF, 0xEB, 0xF5, 0xFF})
	CircleGrid(im, []color.RGBA{
		{0xED, 0x34, 0x41, 0xFF},
		{0xFF, 0xD6, 0x30, 0xFF},
		{0x32, 0x9F, 0xE3, 0xFF},
		{0x15, 0x42, 0x96, 0xFF},
		{0x00, 0x00, 0x00, 0xFF},
		{0xFF, 0xFF, 0xFF, 0xFF},
	}, 2, RandomIntN(4, 6))
	SavePNG("circlegrid.png", im)
}

func circleline() {
	im := newImage(600, 600)
	FillBackground(im, Tan)
	CircleLine(im, 1, LightPink, 0.02, 600, 1.5, complex(2, 2))
	SavePNG("circleline.png", im)
}

func circleloop() {
	im := newImage(500, 500)
	CircleLoop(im, 1, Orange, 30, 100, 1000)
	SavePNG("circleloop.png", im)
}

func circleloop2() {
	im := newImage(500, 500)
	FillBackground(im, color.RGBA{8, 10, 20, 255})
	CircleLoop2(im, []color.RGBA{
		{0xF9, 0xC8, 0x0E, 0xFF},
		{0xF8, 0x66, 0x24, 0xFF},
		{0xEA, 0x35, 0x46, 0xFF},
		{0x66, 0x2E, 0x9B, 0xFF},
		{0x43, 0xBC, 0xCD, 0xFF},
	}, 7)
	SavePNG("colorloop2.png", im)
}

func circlemove() {
	im := newImage(1200, 500)
	FillBackground(im, White)
	CircleMove(im, 1000)
	SavePNG("circlemove.png", im)
}

func circlenoise() {
	const _dots = 2000
	const _iters = 1 // 000
	noise := NewPerlinNoiseDeprecated()
	bar := progressbar.Default(_iters)
	for i, ralpha := range RangeF64(0, 2*PI, _iters) {
		_ = bar.Add(1)

		angles := make([]float64, _dots)
		for j := range angles {
			angles[j] = noise.NoiseV2_1(complex(float64(j), 0)+Polar(1, ralpha)) + ralpha
		}

		peepo := [4096]float64{}
		for j := range peepo {
			peepo[j] = noise.NoiseV2_1(complex(float64(j), 1) + Polar(1, ralpha))
		}
		noise2 := NewPerlinNoise(peepo)

		im := newImage(500, 500)
		FillBackground(im, White)
		CircleNoise(im, 0.3, 80, angles, 60, 80, 400, noise2)
		SavePNG(fmt.Sprintf("circlenoise%04d.png", i), im)
	}
}

func colorcanvas() {
	im := newImage(500, 500)
	FillBackground(im, Black)
	ColorCanva(im, []color.RGBA{
		{0xF9, 0xC8, 0x0E, 0xFF},
		{0xF8, 0x66, 0x24, 0xFF},
		{0xEA, 0x35, 0x46, 0xFF},
		{0x66, 0x2E, 0x9B, 0xFF},
		{0x43, 0xBC, 0xCD, 0xFF},
	}, 8, 5)
	SavePNG("colorcanva.png", im)
}

func colorcircle() {
	im := newImage(1000, 1000)
	FillBackground(im, White)
	ColorCircle(im, []color.RGBA{
		{0xFF, 0xC6, 0x18, 0xFF},
		{0xF4, 0x25, 0x39, 0xFF},
		{0x41, 0x78, 0xF4, 0xFF},
		{0xFE, 0x84, 0xFE, 0xFF},
		{0xFF, 0x81, 0x19, 0xFF},
		{0x56, 0xAC, 0x51, 0xFF},
		{0x98, 0x19, 0xFA, 0xFF},
		{0xFF, 0xFF, 0xFF, 0xFF},
	}, 500)
	SavePNG("colorcircle.png", im)
}

func colorcircle2() {
	im := newImage(800, 800)
	FillBackground(im, White)
	ColorCircle2(im, []color.RGBA{
		{0x11, 0x60, 0xC6, 0xFF},
		{0xFD, 0xD9, 0x00, 0xFF},
		{0xF5, 0xB4, 0xF8, 0xFF},
		{0xEF, 0x13, 0x55, 0xFF},
		{0xF4, 0x9F, 0x0A, 0xFF},
	}, 30)
	SavePNG("colorcircle2.png", im)
}

func contourline() {
	im := newImage(1600, 1600)
	FillBackground(im, color.RGBA{0x1a, 0x06, 0x33, 0xFF})
	ContourLine(im, []color.RGBA{
		{0x58, 0x18, 0x45, 0xFF},
		{0x90, 0x0C, 0x3F, 0xFF},
		{0xC7, 0x00, 0x39, 0xFF},
		{0xFF, 0x57, 0x33, 0xFF},
		{0xFF, 0xC3, 0x0F, 0xFF},
	}, 500)
	SavePNG("contourline.png", im)
}

func domainwrap() {
	im := newImage(500, 500)
	FillBackground(im, Black)
	DomainWarp(im, 0.01, 4, complex(4, 20), cmap, 0, 0, "")
	SavePNG("domainwarp.png", im)
}

func domainwrapFrames() {
	im := newImage(500, 500)
	FillBackground(im, Black)
	DomainWarp(im, 0.01, 4, complex(4, 20), cmap, 0.005, 100, "./temp")
}

func test() {
	dest := image.NewRGBA(image.Rect(0, 0, 500, 500))

	dc := NewContextFromRGBA(dest)
	dc.Stack(func(dc *Context) {
		dc.TransformAdd(Translate(complex(500/2, 500/2)))
		dc.TransformAdd(Rotate(40))
		dc.SetColor(color.RGBA{0xFF, 0x00, 0x00, 255})
		for i := range Range(361) {
			theta := Radians(float64(i))
			p := complex(
				Cos(theta)-Pow(Sin(theta), 2)/Sqrt(2.0),
				Cos(theta)*Sin(theta),
			)

			alpha := Radians(float64(i + 1))

			p1 := complex(
				Cos(alpha)-Pow(Sin(alpha), 2)/Sqrt(2.0),
				Cos(alpha)*Sin(alpha),
			)

			dc.DrawLine(p*100, p1*100)
			dc.Stroke()
		}
	})
	SavePNG("test.png", dc.Image())
}

func out() {
	dc := NewContext(complex(1000, 1000))
	dc.DrawCircle(complex(350, 500), 300)
	dc.Clip()
	dc.DrawCircle(complex(650, 500), 300)
	dc.Clip()
	dc.DrawRectangle(0, complex(1000, 1000))
	dc.SetColor(ColorRGB(0, 0, 0))
	dc.Fill()
	SavePNG("out.png", dc.Image())
}

func main() {
	rand.Seed(time.Now().Unix())

	beziers()
	blackhole()
	circlegrid()
	circle()
	circleline()
	circleloop()
	circleloop2()
	circlemove()
	circlenoise()
	clip()
	colorcanvas()
	colorcircle()
	colorcircle2()
	concat()
	contourline()
	crisp()
	cubic()
	dotswave()
	domainwrap()
	if false { // TODO: generates lots of images, change to gif
		domainwrapFrames()
	}
	dotline()
	ellipse()
	gofont()
	gradientLinear()
	gradientRadial()
	gradientConic()
	if false { // TODO: fix font loading
		gradientText()
	}
	gridsquare()
	invertMask()
	janus()
	julia()
	lines()
	linewidth()
	lorem()
	mask()
	maze()
	if false { // TODO: fix font loading
		meme()
	}
	mystar()
	noiseline()
	oceanfish()
	openfill()
	out()
	perlinpearls()
	pixelhole()
	patternFill()
	pointribbon()
	if false { // TODO: fix font loading
		quadratic()
	}
	randcircle()
	randomshape()
	rotatedImage()
	rotatedText()
	scatter()
	silksky()
	silksmoke()
	sine()
	solarflare()
	spiral()
	spiralsquare()
	star(5)
	stars()
	swirl()
	test()
	if false { // TODO: fix font loading
		text()
	}
	tiling()
	if false { // TODO: fix font loading
		unicode()
		wrap()
	}
	yarn()

	if false {
		_http()
	}
}
