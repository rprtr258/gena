# gena

[![Go](https://github.com/rprtr258/gena/actions/workflows/go.yml/badge.svg)](https://github.com/rprtr258/gena/actions/workflows/go.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/rprtr258/gena)](https://goreportcard.com/report/github.com/rprtr258/gena)

`gena` is a `Go` package to create generative art.

## How to use

### Hello, Circle!
```go
package main

import . "github.com/rprtr258/gena"

func main() {
    dc := NewContext(complex(1000, 1000))
    dc.DrawCircle(complex(500, 500), 400)
    dc.SetColor(ColorRGB(0, 0, 0))
    dc.Fill()
    SavePNG("out.png", dc.Image())
}
```

### Creating Contexts

There are a few ways of creating a context.

```go
NewContext(size V2) *Context
NewContextFromRGBA(im *image.RGBA) *Context
```

### Drawing Functions
```go
DrawPoint(pos V2, r float64)
DrawLine(from, to V2)
DrawRectangle(topLeft, size V2)
DrawRoundedRectangle(topLeft, size V2, r float64)
DrawCircle(center V2, r float64)
DrawArc(center V2, r, angle1, angle2 float64)
DrawEllipse(center, r V2)
DrawEllipticalArc(center, r V2, angle1, angle2 float64)
DrawRegularPolygon(n int, center V2, r, rotation float64)
DrawImage(im image.Image, pos V2)
DrawImageAnchored(im image.Image, pos, a V2)
SetPixel(pos V2)

MoveTo(pos V2)
LineTo(pos V2)
QuadraticTo(v1, v2 V2)
CubicTo(v1, v2, v3 V2)
ClearPath()
NewSubPath()
ClosePath()

Clear()
Stroke()
Fill()
StrokePreserve()
FillPreserve()
```

It is often desired to center an image at a point. Use `DrawImageAnchored` with `a` set to `0.5` to do this. Use `0` to left or top align. Use `1` to right or bottom align. `DrawStringAnchored` does the same for text, so you don't need to call `MeasureString` yourself.

### Text Functions
It will even do word wrap for you!

```go
DrawString(s string, v V2)
DrawStringAnchored(s string, v, a V2)
DrawStringWrapped(s string, v, a V2, width, lineSpacing float64, align Align)
MeasureString(s string) V2
MeasureMultilineString(s string, lineSpacing float64) V2
WordWrap(s string, w float64) []string
SetFontFace(fontFace font.Face)
LoadFontFace(path string, points float64) error
```

### Color Functions
```go
ColorRGB(r, g, b float64) color.Color
ColorRGBA(r, g, b, a float64) color.Color
ColorRGB255(r, g, b int) color.Color
ColorRGBA255(r, g, b, a int) color.Color
ColorHex(x string) color.Color
```

### Stroke & Fill Options
```go
SetLineWidth(lineWidth float64)
SetLineCap(lineCap LineCap)
SetLineJoin(lineJoin LineJoin)
SetDash(dashes ...float64)
SetDashOffset(offset float64)
SetFillRule(fillRule FillRule)
```

### Gradients & Patterns
Linear, radial and conic gradients are supported. You can also use surface patterns or implement your own.

```go
// apply patterns
SetFillStyle(pattern Pattern)
SetStrokeStyle(pattern Pattern)

// builtin patterns
PatternSolid(color color.Color) Pattern
PatternGradientLinear(v0, v1 V2) Pattern
PatternGradientRadial(v0 V2, r0 float64, v1 V2, r1 float64) Pattern
PatternGradientConic(c V2, deg float64) Pattern
PatternPatternSurface(im image.Image, op RepeatOp) Pattern
```

### Transformation Functions
```go
Identity()
Translate(v V2)
Scale(v V2)
Rotate(angle float64)
Shear(v V2)
ScaleAbout(s, v V2)
RotateAbout(angle float64, v V2)
ShearAbout(s, v V2)
TransformPoint(v V2) V2
InvertY()
```

It is often desired to rotate or scale about a point that is not the origin. The functions `RotateAbout`, `ScaleAbout`, `ShearAbout` are provided as a convenience.

`InvertY` is provided in case Y should increase from bottom to top vs. The default is top to bottom.

### Clipping Functions
Use clipping regions to restrict drawing operations to an area that you defined using paths.

```go
Clip()
ClipPreserve()
ResetClip()
AsMask() *image.Alpha
SetMask(mask *image.Alpha)
InvertMask()
```

### Helper Functions
```go
Radians(degrees float64) float64
Degrees(radians float64) float64
LoadImage(path string) image.Image
SavePNG(path string, im image.Image)
SaveJPG(path string, im image.Image, quality uint8)
```


## Examples
![](./images/examples.png)

### Black Hole
![](images/blackhole.png)

### Circle Composes Circle
![](images/colorloop2.png)

### Circle Grid
![](images/circlegrid.png)

### Circle Line
![](images/circleline.png)

### Circle Loop
![](images/circleloop.png)

### Circle Loop2
![](images/circleloop2.png)

### Circle Move
![](images/circlemove.png)

### Circle Noise
![](images/circlenoise.png)

### Color Canva
![](images/colorcanva.png)

### Color Circle
![](images/colorcircle.png)

### Color Circle2
![](images/colorcircle2.png)

### Contour Line
![](images/contourline.png)

### Domain Warp
![](images/domainwarp.png)

### Dot Line
![](images/dotline.png)

### Dots Wave
![](images/dotswave.png)

### Noise Line
![](images/noiseline.png)

### Janus
![](images/janus.png)

### Julia Set
![](images/julia.png)

### Maze
![](images/maze.png)

### Ocean Fish
![](images/oceanfish.png)

### Perlin Perls
![](images/perlinperls.png)

### Pixel Hole
![](images/pixelhole.png)

### Point Ribbon
![](images/pointribbon.png)

### Random Circle
![](images/randcircle.png)

### Random Circle Trails
![](images/random_circle_trails.png)

### Random Shapes
![](images/randomshape.png)

### Silk Sky
![](images/silksky.png)

### Silk Smoke
![](images/silksmoke.png)

### Spiral Square
![](images/spiralsquare.png)

### Square Grid
![](images/squaregrid.png)

### Stars
![](images/stars.png)

### Swirl
![](images/swirl.png)

### Yarn
![](images/yarn.png)
