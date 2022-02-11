package main

import (
	"evo/internal/app/simulation"
	"evo/internal/app/utils"
	"fmt"
	"math/rand"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"github.com/kyroy/kdtree"
	"github.com/kyroy/kdtree/points"
	"golang.org/x/image/colornames"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	pixelgl.Run(run)
}

type HypePoint struct {
	Coordinates []float64
	Data        interface{}
}

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "evo",
		Bounds: pixel.R(0, 0, 1024, 720),
		// VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	var (
		frames = 0
		second = time.Tick(time.Second)
	)

	cells := simulation.GenerateCells(100, cfg.Bounds)

	tree := kdtree.New([]kdtree.Point{})
	simulation.UpdateTree(tree, cells)

	point := tree.KNN(
		&points.Point{Coordinates: []float64{cfg.Bounds.W() / 2, cfg.Bounds.H() / 2}},
		1,
	)[0]

	fmt.Println(point.(interface{}))
	fmt.Println(point.(interface{}))
	// fmt.Println(point.(interface{}))
	// pdata := reflect.ValueOf(point).Elem()
	// kek := pdata.FieldByName("Data").Elem()
	// fmt.Println(kek.MapRange().Value())

	// simulation.FindCellByPoint(cells, point).Color = colornames.White

	imd := imdraw.New(nil)
	for _, c := range cells {
		imd.Color = c.Color
		imd.Push(
			pixel.Vec(c.Position),
		)
		imd.Circle(utils.RandBetween(1, 3), 0)
	}

	for !win.Closed() {
		// imd.Clear()

		// for _, c := range cells {
		// 	// c.Move()
		// 	// simulation.UpdateTree(tree, cells)

		// 	imd.Color = c.Color
		// 	imd.Push(
		// 		c.Position.ToVec(),
		// 	)
		// 	imd.Circle(c.Radius, 0)
		// }

		win.Clear(colornames.Black)
		imd.Draw(win)
		win.Update()

		// fps
		frames++
		select {
		case <-second:
			win.SetTitle(fmt.Sprintf("%s | FPS: %d", cfg.Title, frames))
			frames = 0
		default:
		}
	}
}
