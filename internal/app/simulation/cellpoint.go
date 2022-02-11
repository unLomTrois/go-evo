package simulation

import (
	"fmt"

	"github.com/faiface/pixel"
)

type CellPosition struct {
	X, Y float64
}

// Dimensions ...
func (p *CellPosition) Dimensions() int {
	return 2
}

// Dimension ...
func (p *CellPosition) Dimension(i int) float64 {
	if i == 0 {
		return p.X
	}
	return p.Y
}

// String ...
func (p *CellPosition) String() string {
	return fmt.Sprintf("{%.2f %.2f}", p.X, p.Y)
}

func (p *CellPosition) ToVec() pixel.Vec {
	return pixel.V(p.X, p.Y)
}

// func (p *CellPosition) Eq(v pixel.Vec) {

// }
