package main

import (
	"evo/internal/app/cellmap"
	"evo/internal/app/simulation"
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

	// fmt.Println(cells[0])
	// fmt.Println(reflect.TypeOf(cells[0]))

	// for key, value := range chmap {
	// 	fmt.Println(key, value)
	// }
	chmap := cellmap.New()

	imd := imdraw.New(nil)
	for _, c := range cells {
		chmap.Put(c)

		imd.Color = c.Color
		imd.Push(
			c.Position,
		)
		imd.Circle(utils.RandBetween(1, 3), 0)
	}

	// first := &cells[0]
	// for _, value := range chmap.Values() {
	// 	fmt.Println(value)
	// }

	// bounds of simulation
	simbounds := cfg.Bounds //cfg.Bounds.Resized(cfg.Bounds.Center(), pixel.V(cfg.Bounds.W()-100, cfg.Bounds.H()-100))

	for !win.Closed() {
		win.Clear(colornames.Black)
		imd.Clear()

		for _, c := range chmap.GetM() {
			c.Move()

			if !simbounds.Contains(c.Position) {
				lin := pixel.L(c.Position, c.NextPosition())

				intersec := simbounds.IntersectionPoints(lin.Scaled(5))

				if len(intersec) > 0 {
					// var edge pixel.Line
					for i, e := range simbounds.Edges() {
						// fmt.Println(i, e)
						if _, ok := lin.Scaled(5).Intersect(e); ok {
							if i == 0 {
								c.Position = pixel.V(simbounds.Max.X, intersec[0].Y)
							}
							if i == 2 {
								c.Position = pixel.V(simbounds.Min.X, intersec[0].Y)
							}
							if i == 1 {
								c.Position = pixel.V(intersec[0].X, simbounds.Min.Y)
							}
							if i == 3 {
								c.Position = pixel.V(intersec[0].X, simbounds.Max.Y)
							}
						}
					}
				}

				// c.Direction *= -1
				// c.Move()
			}

			imd.Color = c.Color
			imd.Push(
				c.Position,
			)
			imd.Circle(c.Radius, 0)
		}

		imd.Draw(win)

		// fps
		frames++
		select {
		case <-second:
			win.SetTitle(fmt.Sprintf("Cells: %d, FPS: %d", chmap.Size(), frames))
			frames = 0
			last := chmap.Keys()[chmap.Size()-1]

			chmap.Remove(last)
			fmt.Println(last)

		default:
		}

		win.Update()

	}
}
