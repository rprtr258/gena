// package gena provides a simple API for rendering 2D graphics in pure Go.
package gena

import (
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"image/png"
	"io"
	"math"
	"os"
	"strings"

	"github.com/golang/freetype/raster"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/draw"
	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"golang.org/x/image/math/f64"
	"golang.org/x/image/math/fixed"
)

func Size(img *image.RGBA) V2 {
	return complex(float64(img.Bounds().Dx()), float64(img.Bounds().Dy()))
}

// FillBackground fills the background of the Canvas
func FillBackground(img *image.RGBA, bg color.RGBA) {
	draw.Draw(img, img.Bounds(), &image.Uniform{bg}, image.Point{}, draw.Src)
}

type LineCap int

const (
	LineCapRound LineCap = iota
	LineCapButt
	LineCapSquare
)

type LineJoin int

const (
	LineJoinRound LineJoin = iota
	LineJoinBevel
)

type FillRule int

const (
	FillRuleWinding FillRule = iota
	FillRuleEvenOdd
)

type Align int

const (
	AlignLeft Align = iota
	AlignCenter
	AlignRight
)

var (
	defaultFillStyle   = NewSolidPattern(color.White)
	defaultStrokeStyle = NewSolidPattern(color.Black)
)

type Context struct {
	width, height int
	rasterizer    *raster.Rasterizer
	im            *image.RGBA
	mask          *image.Alpha
	color         color.Color
	fillPattern   Pattern
	strokePattern Pattern
	strokePath    raster.Path
	fillPath      raster.Path
	start         V2
	current       V2
	hasCurrent    bool
	dashes        []float64
	dashOffset    float64
	lineWidth     float64
	lineCap       LineCap
	lineJoin      LineJoin
	fillRule      FillRule
	fontFace      font.Face
	fontHeight    float64
	matrix        Matrix
}

// NewContext creates a new image.RGBA with the specified width and height
// and prepares a context for rendering onto that image.
func NewContext(width, height int) *Context {
	return NewContextForRGBA(image.NewRGBA(image.Rect(0, 0, width, height)))
}

// NewContextForRGBA prepares a context for rendering onto the specified image.
// No copy is made.
func NewContextForRGBA(im *image.RGBA) *Context {
	w := im.Bounds().Size().X
	h := im.Bounds().Size().Y
	return &Context{
		width:         w,
		height:        h,
		rasterizer:    raster.NewRasterizer(w, h),
		im:            im,
		color:         color.Transparent,
		fillPattern:   defaultFillStyle,
		strokePattern: defaultStrokeStyle,
		lineWidth:     1,
		fillRule:      FillRuleWinding,
		fontFace:      basicfont.Face7x13,
		fontHeight:    13,
		matrix:        Identity,
	}
}

// GetCurrentPoint will return the current point and if there is a current point.
// The point will have been transformed by the context's transformation matrix.
func (dc *Context) GetCurrentPoint() (V2, bool) {
	return dc.current, dc.hasCurrent
}

// Image returns the image that has been drawn by this context.
func (dc *Context) Image() *image.RGBA {
	return dc.im
}

// Width returns the width of the image in pixels.
func (dc *Context) Width() int {
	return dc.width
}

// Height returns the height of the image in pixels.
func (dc *Context) Height() int {
	return dc.height
}

// EncodePNG encodes the image as a PNG and writes it to the provided io.Writer.
func (dc *Context) EncodePNG(w io.Writer) {
	must(png.Encode(w, dc.im))
}

// EncodeJPG encodes the image as a JPG and writes it to the provided io.Writer
// in JPEG 4:2:0 baseline format with the given options.
// Default parameters are used if a nil *jpeg.Options is passed.
func (dc *Context) EncodeJPG(w io.Writer, o *jpeg.Options) {
	must(jpeg.Encode(w, dc.im, o))
}

// SetDash sets the current dash pattern to use. Call with zero arguments to
// disable dashes. The values specify the lengths of each dash, with
// alternating on and off lengths.
func (dc *Context) SetDash(dashes ...float64) {
	dc.dashes = dashes
}

// SetDashOffset sets the initial offset into the dash pattern to use when
// stroking dashed paths.
func (dc *Context) SetDashOffset(offset float64) {
	dc.dashOffset = offset
}

func (dc *Context) SetLineWidth(lineWidth float64) {
	dc.lineWidth = lineWidth
}

func (dc *Context) SetLineCap(lineCap LineCap) {
	dc.lineCap = lineCap
}

func (dc *Context) SetLineCapRound() {
	dc.lineCap = LineCapRound
}

func (dc *Context) SetLineCapButt() {
	dc.lineCap = LineCapButt
}

func (dc *Context) SetLineCapSquare() {
	dc.lineCap = LineCapSquare
}

func (dc *Context) SetLineJoin(lineJoin LineJoin) {
	dc.lineJoin = lineJoin
}

func (dc *Context) SetLineJoinRound() {
	dc.lineJoin = LineJoinRound
}

func (dc *Context) SetLineJoinBevel() {
	dc.lineJoin = LineJoinBevel
}

func (dc *Context) SetFillRule(fillRule FillRule) {
	dc.fillRule = fillRule
}

func (dc *Context) SetFillRuleWinding() {
	dc.fillRule = FillRuleWinding
}

func (dc *Context) SetFillRuleEvenOdd() {
	dc.fillRule = FillRuleEvenOdd
}

// Color Setters

func (dc *Context) setFillAndStrokeColor(c color.Color) {
	dc.color = c
	dc.fillPattern = NewSolidPattern(c)
	dc.strokePattern = NewSolidPattern(c)
}

// SetFillStyle sets current fill style
func (dc *Context) SetFillStyle(pattern Pattern) {
	// if pattern is SolidPattern, also change dc.color(for dc.Clear, dc.drawString)
	if fillStyle, ok := pattern.(*solidPattern); ok {
		dc.color = fillStyle.color
	}
	dc.fillPattern = pattern
}

// SetStrokeStyle sets current stroke style
func (dc *Context) SetStrokeStyle(pattern Pattern) {
	dc.strokePattern = pattern
}

// SetColor sets the current color(for both fill and stroke).
func (dc *Context) SetColor(c color.Color) {
	dc.setFillAndStrokeColor(c)
}

// SetHexColor sets the current color using a hex string. The leading pound
// sign (#) is optional. Both 3- and 6-digit variations are supported. 8 digits
// may be provided to set the alpha value as well.
// TODO: move out getting color from hex, make single function to set color
func (dc *Context) SetHexColor(x string) {
	x = strings.TrimPrefix(x, "#")
	var r, g, b int
	a := 255
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

	dc.SetRGBA255(color.RGBA{uint8(r), uint8(g), uint8(b), 0}, a)
}

// SetRGBA255 sets the current color. r, g, b, a values should be between 0 and 255, inclusive.
func (dc *Context) SetRGBA255(cl color.RGBA, a int) {
	dc.color = color.NRGBA{cl.R, cl.G, cl.B, uint8(a)}
	dc.setFillAndStrokeColor(dc.color)
}

// SetRGB255 sets the current color. r, g, b values should be between 0 and 255, inclusive.
// Alpha will be set to 255 (fully opaque).
func (dc *Context) SetRGB255(r, g, b int) {
	dc.SetRGBA255(color.RGBA{uint8(r), uint8(g), uint8(b), 0}, 255)
}

// SetRGBA sets the current color. r, g, b, a values should be between 0 and 1,
// inclusive.
func (dc *Context) SetRGBA(r, g, b, a float64) {
	dc.color = color.NRGBA{
		uint8(r * 255),
		uint8(g * 255),
		uint8(b * 255),
		uint8(a * 255),
	}
	dc.setFillAndStrokeColor(dc.color)
}

// SetRGB sets the current color. r, g, b values should be between 0 and 1,
// inclusive. Alpha will be set to 1 (fully opaque).
func (dc *Context) SetRGB(r, g, b float64) {
	dc.SetRGBA(r, g, b, 1)
}

// Path Manipulation

// MoveTo starts a new subpath within the current path starting at the
// specified point.
func (dc *Context) MoveTo(x, y float64) {
	dc.MoveToV2(complex(x, y))
}

// MoveTo starts a new subpath within the current path starting at the
// specified point.
func (dc *Context) MoveToV2(v V2) {
	if dc.hasCurrent {
		dc.fillPath.Add1(Fixed(dc.start))
	}
	p := dc.TransformPoint(v)
	dc.strokePath.Start(Fixed(p))
	dc.fillPath.Start(Fixed(p))
	dc.start = p
	dc.current = p
	dc.hasCurrent = true
}

// LineTo adds a line segment to the current path starting at the current
// point. If there is no current point, it is equivalent to MoveTo(x, y)
func (dc *Context) LineToV2(v V2) {
	if !dc.hasCurrent {
		dc.MoveToV2(v)
	} else {
		p := dc.TransformPoint(v)
		dc.strokePath.Add1(Fixed(p))
		dc.fillPath.Add1(Fixed(p))
		dc.current = p
	}
}

func (dc *Context) LineTo(x, y float64) {
	dc.LineToV2(complex(x, y))
}

// QuadraticTo adds a quadratic bezier curve to the current path starting at
// the current point. If there is no current point, it first performs
// MoveTo(x1, y1)
func (dc *Context) QuadraticTo(a, b V2) {
	if !dc.hasCurrent {
		dc.MoveToV2(a)
	}
	p1 := dc.TransformPoint(a)
	p2 := dc.TransformPoint(b)
	dc.strokePath.Add2(Fixed(p1), Fixed(p2))
	dc.fillPath.Add2(Fixed(p1), Fixed(p2))
	dc.current = p2
}

// CubicTo adds a cubic bezier curve to the current path starting at the
// current point. If there is no current point, it first performs
// MoveTo(x1, y1). Because freetype/raster does not support cubic beziers,
// this is emulated with many small line segments.
func (dc *Context) CubicTo(a, b, c V2) {
	if !dc.hasCurrent {
		dc.MoveToV2(a)
	}
	a0 := dc.current
	a = dc.TransformPoint(a)
	b = dc.TransformPoint(b)
	c = dc.TransformPoint(c)
	points := CubicBezier(a0, a, b, c)
	previous := Fixed(dc.current)
	for _, p := range points[1:] {
		f := Fixed(p)
		if f == previous {
			// TODO: this fixes some rendering issues but not all
			continue
		}
		previous = f
		dc.strokePath.Add1(f)
		dc.fillPath.Add1(f)
		dc.current = p
	}
}

// ClosePath adds a line segment from the current point to the beginning
// of the current subpath. If there is no current point, this is a no-op.
func (dc *Context) ClosePath() {
	if dc.hasCurrent {
		dc.strokePath.Add1(Fixed(dc.start))
		dc.fillPath.Add1(Fixed(dc.start))
		dc.current = dc.start
	}
}

// ClearPath clears the current path. There is no current point after this
// operation.
func (dc *Context) ClearPath() {
	dc.strokePath.Clear()
	dc.fillPath.Clear()
	dc.hasCurrent = false
}

// NewSubPath starts a new subpath within the current path. There is no current
// point after this operation.
func (dc *Context) NewSubPath() {
	if dc.hasCurrent {
		dc.fillPath.Add1(Fixed(dc.start))
	}
	dc.hasCurrent = false
}

// Path Drawing

func (dc *Context) capper() raster.Capper {
	switch dc.lineCap {
	case LineCapButt:
		return raster.ButtCapper
	case LineCapRound:
		return raster.RoundCapper
	case LineCapSquare:
		return raster.SquareCapper
	}
	return nil
}

func (dc *Context) joiner() raster.Joiner {
	switch dc.lineJoin {
	case LineJoinBevel:
		return raster.BevelJoiner
	case LineJoinRound:
		return raster.RoundJoiner
	}
	return nil
}

func rasterPath(paths [][]V2) raster.Path {
	var result raster.Path
	for _, path := range paths {
		var previous fixed.Point26_6
		for i, point := range path {
			f := Fixed(point)
			if i == 0 {
				result.Start(f)
			} else {
				if Abs(f.X-previous.X)+Abs(f.Y-previous.Y) > 8 {
					// TODO: this is a hack for cases where two points are
					// too close - causes rendering issues with joins / caps
					result.Add1(f)
				}
			}
			previous = f
		}
	}
	return result
}

func dashPath(paths [][]V2, dashes []float64, offset float64) [][]V2 {
	if len(dashes) == 0 {
		return paths
	}

	if len(dashes) == 1 {
		dashes = append(dashes, dashes[0])
	}

	var result [][]V2
	for _, path := range paths {
		if len(path) < 2 {
			continue
		}
		previous := path[0]
		pathIndex := 1
		dashIndex := 0
		segmentLength := 0.0

		// offset
		if offset != 0 {
			var totalLength float64
			for _, dashLength := range dashes {
				totalLength += dashLength
			}
			offset = math.Mod(offset, totalLength)
			if offset < 0 {
				offset += totalLength
			}
			for i, dashLength := range dashes {
				offset -= dashLength
				if offset < 0 {
					dashIndex = i
					segmentLength = dashLength + offset
					break
				}
			}
		}

		var segment []V2
		segment = append(segment, previous)
		for pathIndex < len(path) {
			dashLength := dashes[dashIndex]
			point := path[pathIndex]
			d := Dist(previous, point)
			maxd := dashLength - segmentLength
			if d > maxd {
				t := maxd / d
				p := LerpV2(previous, point, t)
				segment = append(segment, p)
				if dashIndex%2 == 0 && len(segment) > 1 {
					result = append(result, segment)
				}
				segment = nil
				segment = append(segment, p)
				segmentLength = 0
				previous = p
				dashIndex = (dashIndex + 1) % len(dashes)
			} else {
				segment = append(segment, point)
				previous = point
				segmentLength += d
				pathIndex++
			}
		}
		if dashIndex%2 == 0 && len(segment) > 1 {
			result = append(result, segment)
		}
	}
	return result
}

func flattenPath(p raster.Path) [][]V2 {
	var result [][]V2
	var path []V2
	var c V2
	for i := 0; i < len(p); {
		switch p[i] {
		case 0:
			if len(path) > 0 {
				result = append(result, path)
				path = nil
			}
			c = complex(unfix(p[i+1]), unfix(p[i+2]))
			path = append(path, c)
			i += 4
		case 1:
			c = complex(unfix(p[i+1]), unfix(p[i+2]))
			path = append(path, c)
			i += 4
		case 2:
			p1 := complex(unfix(p[i+1]), unfix(p[i+2]))
			p2 := complex(unfix(p[i+3]), unfix(p[i+4]))
			points := QuadraticBezier(c, p1, p2)
			path = append(path, points...)
			c = p2
			i += 6
		case 3:
			p1 := complex(unfix(p[i+1]), unfix(p[i+2]))
			p2 := complex(unfix(p[i+3]), unfix(p[i+4]))
			p3 := complex(unfix(p[i+5]), unfix(p[i+6]))
			points := CubicBezier(c, p1, p2, p3)
			path = append(path, points...)
			c = p3
			i += 8
		default:
			panic("bad path")
		}
	}
	if len(path) > 0 {
		result = append(result, path)
	}
	return result
}

func (dc *Context) stroke(painter raster.Painter) {
	path := dc.strokePath
	if len(dc.dashes) > 0 {
		path = rasterPath(dashPath(flattenPath(path), dc.dashes, dc.dashOffset))
	} else {
		// TODO: this is a temporary workaround to remove tiny segments
		// that result in rendering issues
		path = rasterPath(flattenPath(path))
	}
	r := dc.rasterizer
	r.UseNonZeroWinding = true
	r.Clear()
	r.AddStroke(path, fix(dc.lineWidth), dc.capper(), dc.joiner())
	r.Rasterize(painter)
}

func (dc *Context) fill(painter raster.Painter) {
	path := dc.fillPath
	if dc.hasCurrent {
		path = make(raster.Path, len(dc.fillPath))
		copy(path, dc.fillPath)
		path.Add1(Fixed(dc.start))
	}
	r := dc.rasterizer
	r.UseNonZeroWinding = dc.fillRule == FillRuleWinding
	r.Clear()
	r.AddPath(path)
	r.Rasterize(painter)
}

// StrokePreserve strokes the current path with the current color, line width,
// line cap, line join and dash settings. The path is preserved after this
// operation.
func (dc *Context) StrokePreserve() {
	var painter raster.Painter
	if dc.mask == nil {
		if pattern, ok := dc.strokePattern.(*solidPattern); ok {
			// with a nil mask and a solid color pattern, we can be more efficient
			// TODO: refactor so we don't have to do this type assertion stuff?
			p := raster.NewRGBAPainter(dc.im)
			p.SetColor(pattern.color)
			painter = p
		}
	}
	if painter == nil {
		painter = newPatternPainter(dc.im, dc.mask, dc.strokePattern)
	}
	dc.stroke(painter)
}

// Stroke strokes the current path with the current color, line width,
// line cap, line join and dash settings. The path is cleared after this
// operation.
func (dc *Context) Stroke() {
	dc.StrokePreserve()
	dc.ClearPath()
}

// FillPreserve fills the current path with the current color. Open subpaths
// are implicity closed. The path is preserved after this operation.
func (dc *Context) FillPreserve() {
	var painter raster.Painter
	if dc.mask == nil {
		if pattern, ok := dc.fillPattern.(*solidPattern); ok {
			// with a nil mask and a solid color pattern, we can be more efficient
			// TODO: refactor so we don't have to do this type assertion stuff?
			p := raster.NewRGBAPainter(dc.im)
			p.SetColor(pattern.color)
			painter = p
		}
	}
	if painter == nil {
		painter = newPatternPainter(dc.im, dc.mask, dc.fillPattern)
	}
	dc.fill(painter)
}

// Fill fills the current path with the current color. Open subpaths
// are implicity closed. The path is cleared after this operation.
func (dc *Context) Fill() {
	dc.FillPreserve()
	dc.ClearPath()
}

// ClipPreserve updates the clipping region by intersecting the current
// clipping region with the current path as it would be filled by dc.Fill().
// The path is preserved after this operation.
func (dc *Context) ClipPreserve() {
	clip := image.NewAlpha(image.Rect(0, 0, dc.width, dc.height))
	painter := raster.NewAlphaOverPainter(clip)
	dc.fill(painter)
	if dc.mask == nil {
		dc.mask = clip
	} else {
		mask := image.NewAlpha(image.Rect(0, 0, dc.width, dc.height))
		draw.DrawMask(mask, mask.Bounds(), clip, image.Point{}, dc.mask, image.Point{}, draw.Over)
		dc.mask = mask
	}
}

// SetMask allows you to directly set the *image.Alpha to be used as a clipping mask.
// It must be the same size as the context.
func (dc *Context) SetMask(mask *image.Alpha) {
	if mask.Bounds().Size() != dc.im.Bounds().Size() {
		panic("mask size must match context size")
	}

	dc.mask = mask
}

// AsMask returns an *image.Alpha representing the alpha channel of this context.
// This can be useful for advanced clipping operations where you first render the mask geometry and then use it as a mask.
func (dc *Context) AsMask() *image.Alpha {
	mask := image.NewAlpha(dc.im.Bounds())
	draw.Draw(mask, dc.im.Bounds(), dc.im, image.Point{}, draw.Src)
	return mask
}

// InvertMask inverts the alpha values in the current clipping mask such that
// a fully transparent region becomes fully opaque and vice versa.
func (dc *Context) InvertMask() {
	if dc.mask == nil {
		dc.mask = image.NewAlpha(dc.im.Bounds())
	} else {
		for i, a := range dc.mask.Pix {
			dc.mask.Pix[i] = 255 - a
		}
	}
}

// Clip updates the clipping region by intersecting the current
// clipping region with the current path as it would be filled by dc.Fill().
// The path is cleared after this operation.
func (dc *Context) Clip() {
	dc.ClipPreserve()
	dc.ClearPath()
}

// ResetClip clears the clipping region.
func (dc *Context) ResetClip() {
	dc.mask = nil
}

// Convenient Drawing Functions

// Clear fills the entire image with the current color.
func (dc *Context) Clear() {
	src := image.NewUniform(dc.color)
	draw.Draw(dc.im, dc.im.Bounds(), src, image.Point{}, draw.Src)
}

// SetPixel sets the color of the specified pixel using the current color.
func (dc *Context) SetPixel(x, y int) {
	dc.im.Set(x, y, dc.color)
}

// DrawPoint is like DrawCircle but ensures that a circle of the specified
// size is drawn regardless of the current transformation matrix. The position
// is still transformed, but not the shape of the point.
func (dc *Context) DrawPoint(v V2, r float64) {
	dc.Stack(func(dc *Context) {
		t := dc.TransformPoint(v)
		dc.Identity()
		dc.DrawCircleV2(t, r)
	})
}

func (dc *Context) DrawLine(p1, p2 V2) {
	dc.MoveToV2(p1)
	dc.LineToV2(p2)
}

func (dc *Context) DrawRectangle(topLeft, size V2) {
	dc.NewSubPath()
	dc.MoveToV2(topLeft)
	dc.LineToV2(topLeft + complex(X(size), 0))
	dc.LineToV2(topLeft + size)
	dc.LineToV2(topLeft + complex(0, Y(size)))
	dc.ClosePath()
}

func (dc *Context) DrawRoundedRectangle(topLeft, size V2, r float64) {
	x, y, w, h := X(topLeft), Y(topLeft), X(size), Y(size)
	x0, x1, x2, x3 := x, x+r, x+w-r, x+w
	y0, y1, y2, y3 := y, y+r, y+h-r, y+h
	dc.NewSubPath()
	dc.MoveToV2(complex(x1, y0))
	dc.LineToV2(complex(x2, y0))
	dc.DrawArc(complex(x2, y1), r, Radians(270), Radians(360))
	dc.LineToV2(complex(x3, y2))
	dc.DrawArc(complex(x2, y2), r, Radians(0), Radians(90))
	dc.LineToV2(complex(x1, y3))
	dc.DrawArc(complex(x1, y2), r, Radians(90), Radians(180))
	dc.LineToV2(complex(x0, y1))
	dc.DrawArc(complex(x1, y1), r, Radians(180), Radians(270))
	dc.ClosePath()
}

func (dc *Context) DrawEllipticalArc(center, r V2, angle1, angle2 float64) {
	const n = 16
	for i := range Range(n) {
		a1 := Lerp(angle1, angle2, float64(i+0)/n)
		a2 := Lerp(angle1, angle2, float64(i+1)/n)
		v0 := center + Mul2(Rotation(a1), r)
		v1 := center + Mul2(Rotation((a1+a2)/2), r)
		v2 := center + Mul2(Rotation(a2), r)
		if i == 0 {
			if dc.hasCurrent {
				dc.LineToV2(v0)
			} else {
				dc.MoveToV2(v0)
			}
		}
		dc.QuadraticTo(v1*2-v2/2-v0/2, v2)
	}
}

func (dc *Context) DrawEllipse(c, r V2) {
	dc.NewSubPath()
	dc.DrawEllipticalArc(c, r, 0, 2*math.Pi)
	dc.ClosePath()
}

func (dc *Context) DrawArc(v V2, r, angle1, angle2 float64) {
	dc.DrawEllipticalArc(v, complex(r, r), angle1, angle2)
}

func (dc *Context) DrawCircle(x, y, r float64) {
	dc.DrawCircleV2(complex(x, y), r)
}

func (dc *Context) DrawCircleV2(c V2, r float64) {
	dc.NewSubPath()
	dc.DrawEllipticalArc(c, complex(r, r), 0, 2*math.Pi)
	dc.ClosePath()
}

func (dc *Context) DrawRegularPolygon(n int, x, y, r, rotation float64) {
	angle := 2 * math.Pi / float64(n)

	rotation -= math.Pi / 2
	if n%2 == 0 {
		rotation += angle / 2
	}

	dc.NewSubPath()
	for i := range Range(n) {
		a := rotation + angle*float64(i)
		dc.LineToV2(Polar(r, a) + complex(x, y))
	}
	dc.ClosePath()
}

// DrawImage draws the specified image at the specified point.
func (dc *Context) DrawImage(im image.Image, x, y int) {
	dc.DrawImageAnchored(im, x, y, 0, 0)
}

// DrawImageAnchored draws the specified image at the specified anchor point.
// The anchor point is x - w * ax, y - h * ay, where w, h is the size of the
// image. Use ax=0.5, ay=0.5 to center the image at the specified point.
func (dc *Context) DrawImageAnchored(im image.Image, x, y int, ax, ay float64) {
	s := im.Bounds().Size()
	x -= int(ax * float64(s.X))
	y -= int(ay * float64(s.Y))
	transformer := draw.BiLinear
	m := dc.matrix.Translate(complex(float64(x), float64(y)))
	s2d := f64.Aff3{m.XX, m.XY, m.X0, m.YX, m.YY, m.Y0}
	if dc.mask == nil {
		transformer.Transform(dc.im, s2d, im, im.Bounds(), draw.Over, nil)
	} else {
		transformer.Transform(dc.im, s2d, im, im.Bounds(), draw.Over, &draw.Options{
			DstMask:  dc.mask,
			DstMaskP: image.Point{},
		})
	}
}

// Text Functions

func (dc *Context) SetFontFace(fontFace font.Face) {
	dc.fontFace = fontFace
	dc.fontHeight = float64(fontFace.Metrics().Height) / 64
}

// LoadFontFace is a helper function to load the specified font file with
// the specified point size. Note that the returned `font.Face` objects
// are not thread safe and cannot be used in parallel across goroutines.
// You can usually just use the Context.LoadFontFace function instead of
// this package-level function.
func (dc *Context) LoadFontFace(path string, points float64) {
	fontBytes := must1(os.ReadFile(path))
	f := must1(truetype.Parse(fontBytes))

	dc.fontFace = truetype.NewFace(f, &truetype.Options{
		Size: points,
	})
	dc.fontHeight = points * 72 / 96
}

func (dc *Context) FontHeight() float64 {
	return dc.fontHeight
}

func (dc *Context) drawString(im *image.RGBA, s string, x, y float64) {
	d := &font.Drawer{
		Dst:  im,
		Src:  image.NewUniform(dc.color),
		Face: dc.fontFace,
		Dot:  fixp(x, y),
	}
	// based on Drawer.DrawString() in golang.org/x/image/font/font.go
	prevC := rune(-1)
	for _, c := range s {
		if prevC >= 0 {
			d.Dot.X += d.Face.Kern(prevC, c)
		}
		dr, mask, maskp, advance, ok := d.Face.Glyph(d.Dot, c)
		if !ok {
			// TODO: is falling back on the U+FFFD glyph the responsibility of
			// the Drawer or the Face?
			// TODO: set prevC = '\ufffd'?
			continue
		}
		sr := dr.Sub(dr.Min)
		transformer := draw.BiLinear
		m := dc.matrix.Translate(complex(float64(dr.Min.X), float64(dr.Min.Y)))
		s2d := f64.Aff3{m.XX, m.XY, m.X0, m.YX, m.YY, m.Y0}
		transformer.Transform(d.Dst, s2d, d.Src, sr, draw.Over, &draw.Options{
			SrcMask:  mask,
			SrcMaskP: maskp,
		})
		d.Dot.X += advance
		prevC = c
	}
}

// DrawString draws the specified text at the specified point.
func (dc *Context) DrawString(s string, x, y float64) {
	dc.DrawStringAnchored(s, x, y, 0, 0)
}

// DrawStringAnchored draws the specified text at the specified anchor point.
// The anchor point is x - w * ax, y - h * ay, where w, h is the size of the
// text. Use ax=0.5, ay=0.5 to center the text at the specified point.
func (dc *Context) DrawStringAnchored(s string, x, y, ax, ay float64) {
	wh := dc.MeasureString(s)
	x -= ax * X(wh)
	y += ay * Y(wh)
	if dc.mask == nil {
		dc.drawString(dc.im, s, x, y)
	} else {
		im := image.NewRGBA(image.Rect(0, 0, dc.width, dc.height))
		dc.drawString(im, s, x, y)
		draw.DrawMask(dc.im, dc.im.Bounds(), im, image.Point{}, dc.mask, image.Point{}, draw.Over)
	}
}

// DrawStringWrapped word-wraps the specified string to the given max width
// and then draws it at the specified anchor point using the given line
// spacing and text alignment.
func (dc *Context) DrawStringWrapped(s string, x, y, ax, ay, width, lineSpacing float64, align Align) {
	lines := dc.WordWrap(s, width)

	// sync h formula with MeasureMultilineString
	h := float64(len(lines)) * dc.fontHeight * lineSpacing
	h -= (lineSpacing - 1) * dc.fontHeight

	x -= ax * width
	y -= ay * h
	switch align {
	case AlignLeft:
		ax = 0
	case AlignCenter:
		ax = 0.5
		x += width / 2
	case AlignRight:
		ax = 1
		x += width
	}
	ay = 1
	for _, line := range lines {
		dc.DrawStringAnchored(line, x, y, ax, ay)
		y += dc.fontHeight * lineSpacing
	}
}

func (dc *Context) MeasureMultilineString(s string, lineSpacing float64) (width, height float64) {
	lines := strings.Split(s, "\n")

	// sync h formula with DrawStringWrapped
	height = float64(len(lines)) * dc.fontHeight * lineSpacing
	height -= (lineSpacing - 1) * dc.fontHeight

	d := &font.Drawer{
		Face: dc.fontFace,
	}

	// max width from lines
	for _, line := range lines {
		adv := d.MeasureString(line)
		currentWidth := float64(adv >> 6) // from Context.MeasureString
		if currentWidth > width {
			width = currentWidth
		}
	}

	return width, height
}

// MeasureString returns the rendered width and height of the specified text
// given the current font face.
func (dc *Context) MeasureString(s string) V2 {
	d := &font.Drawer{
		Face: dc.fontFace,
	}
	a := d.MeasureString(s)
	return complex(float64(a>>6), dc.fontHeight)
}

// WordWrap wraps the specified string to the given max width and current
// font face.
func (dc *Context) WordWrap(s string, width float64) []string {
	var result []string
	for _, line := range strings.Split(s, "\n") {
		fields := strings.Fields(line)
		if len(fields)%2 == 1 {
			fields = append(fields, "")
		}
		x := ""
		for i := 0; i < len(fields); i += 2 {
			w := X(dc.MeasureString(x + fields[i]))
			if w > width {
				if x == "" {
					result = append(result, fields[i])
					x = ""
					continue
				} else {
					result = append(result, x)
					x = ""
				}
			}
			x += fields[i] + fields[i+1]
		}
		if x != "" {
			result = append(result, x)
		}
	}
	for i, line := range result {
		result[i] = strings.TrimSpace(line)
	}
	return result
}

// Transformation Matrix Operations

// Identity resets the current transformation matrix to the identity matrix.
// This results in no translating, scaling, rotating, or shearing.
func (dc *Context) Identity() {
	dc.matrix = Identity
}

// Translate updates the current matrix with a translation.
func (dc *Context) Translate(v V2) {
	dc.matrix = dc.matrix.Translate(v)
}

// Scale updates the current matrix with a scaling factor.
// Scaling occurs about the origin.
func (dc *Context) Scale(v V2) {
	dc.matrix = dc.matrix.Scale(v)
}

func (dc *Context) RelativeTo(v V2, fn func(*Context)) {
	dc.Translate(v)
	fn(dc)
	dc.Translate(-v)
}

// Rotate updates the current matrix with a anticlockwise rotation.
// Rotation occurs about the origin. Angle is specified in radians.
func (dc *Context) Rotate(angle float64) {
	dc.matrix = dc.matrix.Rotate(angle)
}

// Shear updates the current matrix with a shearing angle.
// Shearing occurs about the origin.
func (dc *Context) Shear(v V2) {
	dc.matrix = dc.matrix.Shear(v)
}

// TransformPoint multiplies the specified point by the current matrix,
// returning a transformed position.
func (dc *Context) TransformPoint(v V2) V2 {
	return dc.matrix.TransformPoint(v)
}

// InvertY flips the Y axis so that Y grows from bottom to top and Y=0 is at
// the bottom of the image.
func (dc *Context) InvertY() {
	dc.Translate(complex(0, float64(dc.height)))
	dc.Scale(1 - 1i)
}

// Stack
func (dc *Context) Stack(fn func(*Context)) {
	// Push saves the current state of the context for later retrieval. These
	// can be nested.
	old := func(dc Context) *Context { return &dc }(*dc) // you cannot just &*x to copy pointer&value

	fn(dc)

	// Pop restores the last saved context state from the stack.
	before := *dc
	*dc = *old
	dc.mask = before.mask
	dc.strokePath = before.strokePath
	dc.fillPath = before.fillPath
	dc.start = before.start
	dc.current = before.current
	dc.hasCurrent = before.hasCurrent
}
