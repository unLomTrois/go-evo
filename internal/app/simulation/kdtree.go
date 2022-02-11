package simulation

import (
	"github.com/faiface/pixel"
	"github.com/kyroy/kdtree"
	"github.com/kyroy/kdtree/points"
)

type Data struct {
	value *Cell
}

func UpdateTree(tree *kdtree.KDTree, cells []*Cell) {
	for _, cell := range cells {
		tree.Insert(points.NewPoint([]float64{12, 4, 6}, Data{value: cell}))

		// tree.Insert(points.NewPoint([]float64{1, 2}, Data{}))
	}
}

func FindCellByPoint(cells []*Cell, point kdtree.Point) *Cell {
	for _, cell := range cells {
		if cell.Position.ToVec().Eq(point.(*CellPosition).ToVec()) {
			return cell
		}
	}
	return nil
}

func FindNearestCell(cells []*Cell, position pixel.Vec, tree *kdtree.KDTree, k int) *Cell {
	point := tree.KNN(
		&points.Point{Coordinates: []float64{position.X, position.Y}},
		1,
	)[0]

	return FindCellByPoint(cells, point)
}
