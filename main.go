package main

import (
	"gr8/interfaces"
	"gr8/sys"
	"gr8/util"
	"image/color"
	"os"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

const (
	DISPLAY_WIDTH  int = 1024
	DISPLAY_HEIGHT int = 512
)

func run() {
	var (
		system sys.Chip8 = sys.NewChip8()
	)

	// Load our ROM file
	system.LoadROMFromFile(os.Args[1])

	win, err := pixelgl.NewWindow(pixelgl.WindowConfig{
		Title:  "gr8",
		Bounds: pixel.R(0, 0, float64(DISPLAY_WIDTH), float64(DISPLAY_HEIGHT)),
		VSync:  true,
	})

	if err != nil {
		panic(err)
	}

	win.Clear(color.Black)

	for !win.Closed() {
		texture := util.SpriteFromVideo(&system.Video, int(interfaces.VIDEO_WIDTH), int(interfaces.VIDEO_HEIGHT))

		texture.Draw(win, pixel.IM.Scaled(pixel.ZV, 16).Moved(win.Bounds().Center()))
		system.Cycle()

		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}
