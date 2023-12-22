package gena

import (
	"crypto/md5"
	"flag"
	"fmt"
	"image/color"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

var save bool

func init() {
	flag.BoolVar(&save, "save", false, "save PNG output for each test case")
	flag.Parse()
}

func hash(dc *Context) string {
	return fmt.Sprintf("%x", md5.Sum(dc.im.Pix))
}

func assertHash(t *testing.T, dc *Context, expectedHash string) {
	assert.Equal(t, expectedHash, hash(dc))
}

func saveImage(dc *Context, name string) {
	if save {
		SavePNG(name+".png", dc.Image())
	}
}

func TestBlank(t *testing.T) {
	dc := NewContext(100, 100)
	saveImage(dc, "TestBlank")
	assertHash(t, dc, "4e0a293a5b638f0aba2c4fe2c3418d0e")
}

func TestGrid(t *testing.T) {
	dc := NewContext(100, 100)
	dc.SetRGB(1, 1, 1)
	dc.Clear()
	for i := 10; i < 100; i += 10 {
		x := float64(i) + 0.5
		dc.DrawLine(complex(x, 0), complex(x, 100))
		dc.DrawLine(complex(0, x), complex(100, x))
	}
	dc.SetRGB(0, 0, 0)
	dc.Stroke()
	saveImage(dc, "TestGrid")
	assertHash(t, dc, "78606adda71d8abfbd8bb271087e4d69")
}

func TestLines(t *testing.T) {
	dc := NewContext(100, 100)
	dc.SetRGB(0.5, 0.5, 0.5)
	dc.Clear()
	rnd := rand.New(rand.NewSource(99))
	for range 100 {
		p1 := RandomV2() * 100
		p2 := RandomV2() * 100
		dc.DrawLine(p1, p2)
		dc.SetLineWidth(rnd.Float64() * 3)
		dc.SetRGB(rnd.Float64(), rnd.Float64(), rnd.Float64())
		dc.Stroke()
	}
	saveImage(dc, "TestLines")
	assertHash(t, dc, "036bd220e2529955cc48425dd72bb686")
}

func TestCircles(t *testing.T) {
	dc := NewContext(100, 100)
	dc.SetRGB(1, 1, 1)
	dc.Clear()
	rnd := rand.New(rand.NewSource(99))
	for range 10 {
		r := rnd.Float64()*10 + 5
		dc.DrawCircleV2(RandomV2()*100, r)
		dc.SetRGB(rnd.Float64(), rnd.Float64(), rnd.Float64())
		dc.FillPreserve()
		dc.SetRGB(rnd.Float64(), rnd.Float64(), rnd.Float64())
		dc.SetLineWidth(rnd.Float64() * 3)
		dc.Stroke()
	}
	saveImage(dc, "TestCircles")
	assertHash(t, dc, "c52698000df96fabafe7863701afe922")
}

func TestQuadratic(t *testing.T) {
	dc := NewContext(100, 100)
	dc.SetRGB(0.25, 0.25, 0.25)
	dc.Clear()
	rnd := rand.New(rand.NewSource(99))
	for range 100 {
		p1 := RandomV2() * 100
		p2 := RandomV2() * 100
		p3 := RandomV2() * 100
		dc.MoveToV2(p1)
		dc.QuadraticTo(p2, p3)
		dc.SetLineWidth(rnd.Float64() * 3)
		dc.SetRGB(rnd.Float64(), rnd.Float64(), rnd.Float64())
		dc.Stroke()
	}
	saveImage(dc, "TestQuadratic")
	assertHash(t, dc, "56b842d814aee94b52495addae764a77")
}

func TestCubic(t *testing.T) {
	dc := NewContext(100, 100)
	dc.SetRGB(0.75, 0.75, 0.75)
	dc.Clear()
	rnd := rand.New(rand.NewSource(99))
	for range 100 {
		p1 := RandomV2() * 100
		p2 := RandomV2() * 100
		p3 := RandomV2() * 100
		p4 := RandomV2() * 100
		dc.MoveToV2(p1)
		dc.CubicTo(p2, p3, p4)
		dc.SetLineWidth(rnd.Float64() * 3)
		dc.SetRGB(rnd.Float64(), rnd.Float64(), rnd.Float64())
		dc.Stroke()
	}
	saveImage(dc, "TestCubic")
	assertHash(t, dc, "4a7960fc4eaaa33ce74131c5ce0afca8")
}

func TestFill(t *testing.T) {
	dc := NewContext(100, 100)
	dc.SetRGB(1, 1, 1)
	dc.Clear()
	rnd := rand.New(rand.NewSource(99))
	for range 10 {
		dc.NewSubPath()
		for range 10 {
			dc.LineToV2(RandomV2() * 100)
		}
		dc.ClosePath()
		dc.SetRGBA(rnd.Float64(), rnd.Float64(), rnd.Float64(), rnd.Float64())
		dc.Fill()
	}
	saveImage(dc, "TestFill")
	assertHash(t, dc, "7ccb3a2443906a825e57ab94db785467")
}

func TestClip(t *testing.T) {
	dc := NewContext(100, 100)
	dc.SetRGB(1, 1, 1)
	dc.Clear()
	dc.DrawCircle(50, 50, 40)
	dc.Clip()
	rnd := rand.New(rand.NewSource(99))
	for range 1000 {
		r := rnd.Float64()*10 + 5
		dc.DrawCircleV2(RandomV2()*100, r)
		dc.SetRGBA(rnd.Float64(), rnd.Float64(), rnd.Float64(), rnd.Float64())
		dc.Fill()
	}
	saveImage(dc, "TestClip")
	assertHash(t, dc, "762c32374d529fd45ffa038b05be7865")
}

func TestPushPop(t *testing.T) {
	const S = 100
	dc := NewContext(S, S)
	dc.SetRGBA(0, 0, 0, 0.1)
	for i := 0; i < 360; i += 15 {
		dc.Stack(func(dc *Context) {
			dc.RelativeTo(complex(S/2, S/2), func(dc *Context) {
				dc.Rotate(Radians(float64(i)))
			})
			dc.DrawEllipse(complex(S/2, S/2), complex(S*7/16, S/8))
			dc.Fill()
		})
	}
	saveImage(dc, "TestPushPop")
	assertHash(t, dc, "31e908ee1c2ea180da98fd5681a89d05")
}

func TestDrawStringWrapped(t *testing.T) {
	dc := NewContext(100, 100)
	dc.SetRGB(1, 1, 1)
	dc.Clear()
	dc.SetRGB(0, 0, 0)
	dc.DrawStringWrapped("Hello, world! How are you?", 50, 50, 0.5, 0.5, 90, 1.5, AlignCenter)
	saveImage(dc, "TestDrawStringWrapped")
	assertHash(t, dc, "8d92f6aae9e8b38563f171abd00893f8")
}

func TestDrawImage(t *testing.T) {
	src := NewContext(100, 100)
	src.SetRGB(1, 1, 1)
	src.Clear()
	for i := 10; i < 100; i += 10 {
		x := float64(i) + 0.5
		src.DrawLine(complex(x, 0), complex(x, 100))
		src.DrawLine(complex(0, x), complex(100, x))
	}
	src.SetRGB(0, 0, 0)
	src.Stroke()

	dc := NewContext(200, 200)
	dc.SetRGB(0, 0, 0)
	dc.Clear()
	dc.DrawImage(src.Image(), 50, 50)
	saveImage(dc, "TestDrawImage")
	assertHash(t, dc, "282afbc134676722960b6bec21305b15")
}

func TestSetPixel(t *testing.T) {
	dc := NewContext(100, 100)
	dc.SetRGB(0, 0, 0)
	dc.Clear()
	dc.SetRGB(0, 1, 0)
	i := 0
	for y := range 100 {
		for x := range 100 {
			if i%31 == 0 {
				dc.SetPixel(x, y)
			}
			i++
		}
	}
	saveImage(dc, "TestSetPixel")
	assertHash(t, dc, "27dda6b4b1d94f061018825b11982793")
}

func TestDrawPoint(t *testing.T) {
	dc := NewContext(100, 100)
	dc.SetRGB(0, 0, 0)
	dc.Clear()
	dc.SetRGB(0, 1, 0)
	dc.Scale(complex(10, 10))
	for y := 0; y <= 10; y++ {
		for x := 0; x <= 10; x++ {
			dc.DrawPoint(complex(float64(x), float64(y)), 3)
			dc.Fill()
		}
	}
	saveImage(dc, "TestDrawPoint")
	assertHash(t, dc, "55af8874531947ea6eeb62222fb33e0e")
}

func TestLinearGradient(t *testing.T) {
	dc := NewContext(100, 100)
	g := NewLinearGradient(0, 0, 100, 100)
	g.AddColorStop(0, color.RGBA{0, 255, 0, 255})
	g.AddColorStop(1, color.RGBA{0, 0, 255, 255})
	g.AddColorStop(0.5, color.RGBA{255, 0, 0, 255})
	dc.SetFillStyle(g)
	dc.DrawRectangle(0, complex(100, 100))
	dc.Fill()
	saveImage(dc, "TestLinearGradient")
	assertHash(t, dc, "75eb9385c1219b1d5bb6f4c961802c7a")
}

func TestRadialGradient(t *testing.T) {
	dc := NewContext(100, 100)
	g := NewRadialGradient(30+50i, 0, 70+50i, 50)
	g.AddColorStop(0, color.RGBA{0, 255, 0, 255})
	g.AddColorStop(1, color.RGBA{0, 0, 255, 255})
	g.AddColorStop(0.5, color.RGBA{255, 0, 0, 255})
	dc.SetFillStyle(g)
	dc.DrawRectangle(0, complex(100, 100))
	dc.Fill()
	saveImage(dc, "TestRadialGradient")
	assertHash(t, dc, "f170f39c3f35c29de11e00428532489d")
}

func TestDashes(t *testing.T) {
	dc := NewContext(100, 100)
	dc.SetRGB(1, 1, 1)
	dc.Clear()
	rnd := rand.New(rand.NewSource(99))
	for range 100 {
		p1 := RandomV2() * 100
		p2 := RandomV2() * 100
		dc.SetDash(rnd.Float64()*3+1, rnd.Float64()*3+3)
		dc.DrawLine(p1, p2)
		dc.SetLineWidth(rnd.Float64() * 3)
		dc.SetRGB(rnd.Float64(), rnd.Float64(), rnd.Float64())
		dc.Stroke()
	}
	saveImage(dc, "TestDashes")
	assertHash(t, dc, "d188069c69dcc3970edfac80f552b53c")
}

func BenchmarkCircles(b *testing.B) {
	dc := NewContext(1000, 1000)
	dc.SetRGB(1, 1, 1)
	dc.Clear()
	rnd := rand.New(rand.NewSource(99))
	for i := range b.N {
		dc.DrawCircle(rnd.Float64()*1000, rnd.Float64()*1000, 10)
		m := float64(i % 2)
		dc.SetRGB(m, m, m)
		dc.Fill()
	}
}
