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

	imd := imdraw.New(nil)
	for _, c := range cells {
		imd.Color = c.Color
		imd.Push(
			c.Position,
		)
		imd.Circle(utils.RandBetween(1, 3), 0)
	}

	// bounds of simulation
	simbounds := cfg.Bounds //cfg.Bounds.Resized(cfg.Bounds.Center(), pixel.V(cfg.Bounds.W()-100, cfg.Bounds.H()-100))

	for !win.Closed() {
		win.Clear(colornames.Black)
		imd.Clear()

		for _, c := range cells {
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
			win.SetTitle(fmt.Sprintf("FPS: %d", frames))
			frames = 0

		default:
		}

		win.Update()

	}
}
