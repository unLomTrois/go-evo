package simulation

import (
	"evo/internal/app/utils"
	"image/color"
	"math"

	"github.com/faiface/pixel"
)

type Cell struct {
	Position  pixel.Vec
	Color     color.Color
	Radius    float64
	Direction float64
	Speed     float64
}

func NewCell(position pixel.Vec, color color.Color, radius float64) *Cell {
	return &Cell{position, color, radius, utils.RandBetween(-2*math.Pi, 2*math.Pi), utils.RandBetween(0, 0.1)}
	// utils.RandBetween(-2*math.Pi, 2*math.Pi)
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

func (c *Cell) NextPosition() pixel.Vec {
	unitVec := pixel.Unit(c.Direction).Scaled(c.Speed)

	// fmt.Println("unitvec", unitVec)

	return pixel.V(c.Position.X+unitVec.X, c.Position.Y+unitVec.Y)
}

func (c *Cell) Move() {
	nextpos := c.NextPosition()
	// fmt.Println(c.Direction, math.Sin(c.Position.X), math.Cos(c.Position.Y))
	// fmt.Println(nextpos)

	c.Position.X = nextpos.X
	c.Position.Y = nextpos.Y

	// c.Direction -= math.Sin(c.Position.X) + math.Cos(c.Direction)
	c.Direction += utils.RandBetween(-0.01, 0.01)

	if c.Direction <= -math.Pi || c.Direction >= math.Pi {
		c.Direction += utils.RandBetween(-0.01, 0.01)
	}
}
