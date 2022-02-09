package simulation

import (
	"evo/internal/app/utils"
	"image/color"
	"math"

	"github.com/faiface/pixel"
)

type Cell struct {
	Position pixel.Vec
	Color    color.Color
	// angle
	Radius    float64
	Direction float64
}

func NewCell(position pixel.Vec, color color.Color, radius float64) *Cell {
	return &Cell{position, color, radius, 0}
}

func GenerateCells(count int, bounds pixel.Rect) []*Cell {
	var cells []*Cell
	for i := 0; i < count; i++ {
		cells = append(
			cells,
			NewCell(utils.RandPosition(bounds), utils.RandColor(), utils.RandBetween(1, 3)),
		)
	}
	return cells
}

func (c *Cell) Move() {
	unitVec := pixel.Unit(c.Direction)

	c.Position.X += unitVec.X
	c.Position.Y += unitVec.Y

	c.Direction += 0.1 + math.Sin(c.Position.X) + math.Cos(c.Position.Y)

	if c.Direction >= 1.57 {
		c.Direction = 0
	}
}
