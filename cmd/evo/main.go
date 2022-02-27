package main

import (
	"evo/internal/app/cellmap"
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

	cmap := cellmap.New(cfg.Bounds)

	imd := imdraw.New(nil)

	// bounds of simulation
	simbounds := cfg.Bounds //cfg.Bounds.Resized(cfg.Bounds.Center(), pixel.V(cfg.Bounds.W()-100, cfg.Bounds.H()-100))

	bounds := pixel.R(10, 10, 100, 100)

	last := time.Now()
	for !win.Closed() {
		dt := time.Since(last).Seconds()
		last = time.Now()

		win.Clear(colornames.Black)
		imd.Clear()

		for _, c := range cmap.GetM() {
			c.Move()

			c.CrossBorder(simbounds)

			c.Draw(imd)
		}

		utils.DrawBounds(imd, bounds, colornames.Red)

		imd.Draw(win)

		// fps
		frames++
		select {
		case <-second:
			win.SetTitle(fmt.Sprintf("Cells: %d, FPS: %d, Delta: %f", cmap.Size(), frames, dt))
			frames = 0
			last := cmap.Keys()[cmap.Size()-1]

			cmap.Remove(last)
			fmt.Println(last)

		default:
		}

		win.Update()
	}
}
