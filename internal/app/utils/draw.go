package utils

import (
	"image/color"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
)

func DrawBounds(imd *imdraw.IMDraw, bounds pixel.Rect, color color.Color) {
	DrawRectangle(imd, bounds, color, 1)
}

func DrawRectangle(imd *imdraw.IMDraw, bounds pixel.Rect, color color.Color, thickness float64) {
	imd.Color = color
	imd.Push(bounds.Min)
	imd.Push(bounds.Max)
	imd.Rectangle(thickness)
}

func DrawCircle(imd *imdraw.IMDraw, pos pixel.Vec, radius float64, color color.Color, thickness float64) {
	imd.Color = color
	imd.Push(pos)
	imd.Circle(radius, thickness)
}
