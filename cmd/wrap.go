package main

import (
	"image"

	. "github.com/rprtr258/gena"
)

const TEXT = "Call me Ishmael. Some years ago—never mind how long precisely—having little or no money in my purse, and nothing particular to interest me on shore, I thought I would sail about a little and see the watery part of the world. It is a way I have of driving off the spleen and regulating the circulation. Whenever I find myself growing grim about the mouth; whenever it is a damp, drizzly November in my soul; whenever I find myself involuntarily pausing before coffin warehouses, and bringing up the rear of every funeral I meet; and especially whenever my hypos get such an upper hand of me, that it requires a strong moral principle to prevent me from deliberately stepping into the street, and methodically knocking people's hats off—then, I account it high time to get to sea as soon as I can. This is my substitute for pistol and ball. With a philosophical flourish Cato throws himself upon his sword; I quietly take to the ship. There is nothing surprising in this. If they but knew it, almost all men in their degree, some time or other, cherish very nearly the same feelings towards the ocean with me."

func wrap() *image.RGBA {
	const W = 1024
	const H = 1024
	const p = 16
	dc := NewContext(P(W, H))
	dc.SetColor(ColorRGB(1, 1, 1))
	dc.Clear()
	dc.DrawLine(P(W/2, 0), P(W/2, H))
	dc.DrawLine(P(0, H/2), P(W, H/2))
	dc.DrawRectangle(P(p, p), P(W, H)-Diag(p*2))
	dc.SetColor(ColorRGBA(0, 0, 1, 0.25))
	dc.SetLineWidth(3)
	dc.Stroke()
	dc.SetColor(ColorRGB(0, 0, 0))
	dc.LoadFontFace("/Library/Fonts/Arial Bold.ttf", 18)
	dc.DrawStringWrapped("UPPER LEFT", P(p, p), 0, 0, 1.5, AlignLeft)
	dc.DrawStringWrapped("UPPER RIGHT", P(W-p, p), P(1, 0), 0, 1.5, AlignRight)
	dc.DrawStringWrapped("BOTTOM LEFT", P(p, H-p), P(0, 1), 0, 1.5, AlignLeft)
	dc.DrawStringWrapped("BOTTOM RIGHT", P(W-p, H-p), P(1, 1), 0, 1.5, AlignRight)
	dc.DrawStringWrapped("UPPER MIDDLE", P(W/2, p), P(0.5, 0), 0, 1.5, AlignCenter)
	dc.DrawStringWrapped("LOWER MIDDLE", P(W/2, H-p), P(0.5, 1), 0, 1.5, AlignCenter)
	dc.DrawStringWrapped("LEFT MIDDLE", P(p, H/2), P(0, 0.5), 0, 1.5, AlignLeft)
	dc.DrawStringWrapped("RIGHT MIDDLE", P(W-p, H/2), P(1, 0.5), 0, 1.5, AlignRight)
	dc.LoadFontFace("/Library/Fonts/Arial.ttf", 12)
	SZ := P(W, H)
	dc.DrawStringWrapped(TEXT, SZ/Coeff(2)-P(p, p), P(1, 1), W/3, 1.75, AlignLeft)
	dc.DrawStringWrapped(TEXT, SZ/Coeff(2)-P(-p, p), P(0, 1), W/3, 2, AlignLeft)
	dc.DrawStringWrapped(TEXT, SZ/Coeff(2)-P(p, -p), P(1, 0), W/3, 2.25, AlignLeft)
	dc.DrawStringWrapped(TEXT, SZ/Coeff(2)-P(-p, -p), P(0, 0), W/3, 2.5, AlignLeft)
	return dc.Image()
}
