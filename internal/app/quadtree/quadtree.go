package quadtree

import (
	sim "evo/internal/app/simulation"
	"evo/internal/app/utils"
	"image/color"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"golang.org/x/image/colornames"
)

type IQuadTree interface {
	Insert() bool
	Subdivide() bool
	Update()
	Collide()
	Query() []*sim.Cell
}

type QuadTree struct {
	is_divided bool
	parent     *QuadTree
	capacity   int
	points     []*sim.Cell
	boundary   pixel.Rect
	// children
	nw *QuadTree
	ne *QuadTree
	sw *QuadTree
	se *QuadTree
}

func NewQuadTree(boundary pixel.Rect) *QuadTree {
	return &QuadTree{
		is_divided: false,
		parent:     nil,
		capacity:   4,
		points:     make([]*sim.Cell, 0),
		boundary:   boundary,
		nw:         nil,
		ne:         nil,
		sw:         nil,
		se:         nil,
	}
}

func (qt *QuadTree) GetBounds() pixel.Rect {
	return qt.boundary
}

func (qt *QuadTree) Insert(cell *sim.Cell) bool {
	if !qt.boundary.Contains(cell.Position) {
		return false
	}

	if !qt.is_divided {
		if len(qt.points) < qt.capacity {
			qt.points = append(qt.points, cell)
			// fmt.Println("insert point", cell)

			if len(qt.points) == qt.capacity {
				qt.Subdivide()
			}

			return true
		}
	}

	// fmt.Println("try to insert point into children")
	return qt.nw.Insert(cell) || qt.ne.Insert(cell) || qt.sw.Insert(cell) || qt.se.Insert(cell)
}

func (qt *QuadTree) Subdivide() bool {
	qt.is_divided = true
	// fmt.Println("subdivide")

	qt.nw = NewQuadTree(pixel.R(qt.boundary.Center().X-qt.boundary.W()/2, qt.boundary.Center().Y, qt.boundary.Center().X, qt.boundary.Max.Y))
	qt.ne = NewQuadTree(pixel.R(qt.boundary.Center().X, qt.boundary.Center().Y, qt.boundary.Max.X, qt.boundary.Max.Y))
	qt.sw = NewQuadTree(pixel.R(qt.boundary.Center().X-qt.boundary.W()/2, qt.boundary.Min.Y, qt.boundary.Center().X, qt.boundary.Center().Y))
	qt.se = NewQuadTree(pixel.R(qt.boundary.Center().X, qt.boundary.Min.Y, qt.boundary.Center().X+qt.boundary.W()/2, qt.boundary.Center().Y))

	// fmt.Println(qt.nw.boundary)
	// fmt.Println(qt.ne.boundary)
	// fmt.Println(qt.sw.boundary)
	// fmt.Println(qt.se.boundary)

	qt.nw.parent = qt
	qt.ne.parent = qt
	qt.sw.parent = qt
	qt.se.parent = qt

	ret := false
	for _, p := range qt.points {
		ret = qt.Insert(p)
	}

	qt.points = nil

	return ret
}

func (qt *QuadTree) Show(imd *imdraw.IMDraw, color color.Color) {
	utils.DrawBounds(imd, qt.GetBounds(), color)

	if qt.is_divided {
		qt.nw.Show(imd, colornames.Red)
		qt.ne.Show(imd, colornames.Blue)
		qt.sw.Show(imd, colornames.Green)
		qt.se.Show(imd, colornames.Yellow)
	}
}
