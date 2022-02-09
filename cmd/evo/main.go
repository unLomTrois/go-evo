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

	for !win.Closed() {
		imd.Clear()

		for _, c := range cells {
			c.Move()

			imd.Color = c.Color
			imd.Push(
				c.Position,
			)
			imd.Circle(c.Radius, 0)
		}

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
