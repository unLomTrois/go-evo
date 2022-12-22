package main

import (
	"evo/internal/app/cellmap"
	"evo/internal/app/quadtree"
	sim "evo/internal/app/simulation"
	"evo/internal/app/utils"
	"fmt"
	"math/rand"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	pixelgl.Run(run)
}

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "evo",
		Bounds: pixel.R(0, 0, 1024, 720),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	var (
		frames = 0
		second = time.Tick(time.Second)
	)

	cmap := cellmap.New(win.Bounds())

	imd := imdraw.New(nil)

	// bounds of simulation
	simbounds := win.Bounds()

	// bounds := pixel.R(utils.RandBetween(0, 400), utils.RandBetween(0, 400), utils.RandBetween(400, 800), utils.RandBetween(400, 800))

	qt := quadtree.NewQuadTree(simbounds)
	qt.InsertMap(cmap.GetM())

	last := time.Now()
	for !win.Closed() {
		dt := time.Since(last).Seconds()
		last = time.Now()

		// управление
		if win.JustPressed(pixelgl.MouseButtonLeft) {
			newcell := sim.NewCell(win.MousePosition(), utils.RandColor(), utils.RandBetween(0, 5))

			cmap.Put(newcell)
			qt.Insert(newcell)
		}

		// отрисовка
		win.Clear(colornames.Black)
		imd.Clear()

		// update cells
		for _, c := range cmap.GetM() {
			c.Move()
			c.CrossBorder(simbounds)
			c.Draw(imd)

			// find nearest neighbors
			searchzone := pixel.R(c.Position.X-40, c.Position.Y-40, c.Position.X+40, c.Position.Y+40)
			neighbors := qt.Query(searchzone)
			// utils.DrawCircle(imd, c.Position, 40, c.Color, 1)
			if len(neighbors) > 0 {
				for _, n := range neighbors {
					utils.DrawLine(imd, c.Position, n.Position, c.Color)
				}
			}
		}

		// utils.DrawBounds(imd, pixel.R(200, 400, 400, 600), colornames.Red)

		// utils.DrawBounds(imd, bounds, colornames.Red)
		// utils.DrawBounds(imd, qt.GetBounds(), colornames.Blue)
		qt.Show(imd, colornames.Orange)
		imd.Draw(win)

		// fps
		frames++
		select {
		case <-second:
			win.SetTitle(fmt.Sprintf("Cells: %d, FPS: %d, Delta: %f", cmap.Size(), frames, dt))
			qt.Update(cmap.GetM())
			frames = 0

			// last := cmap.Keys()[cmap.Size()-1]
			// cmap.Remove(last)
			// fmt.Println(last)

		default:
		}

		win.Update()
	}
}
