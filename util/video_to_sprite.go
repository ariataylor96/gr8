package util

import (
	"github.com/faiface/pixel"
	"gr8/interfaces"
	"image"
	"image/color"
)

var (
	width   = int(interfaces.VIDEO_WIDTH)
	height  = int(interfaces.VIDEO_HEIGHT)
	texture = image.NewRGBA(image.Rect(0, 0, width, height))
)

func SpriteFromVideo(video *[2048]byte, output *pixel.Sprite) {
	for idx, val := range *video {
		if val > 0 {
			texture.Set(idx%width, idx/width, color.White)
		} else {
			texture.Set(idx%width, idx/width, color.Black)
		}
	}

	pictureData := pixel.PictureDataFromImage(texture)
	output.Set(pictureData, pictureData.Rect)
}

func BlankSprite() pixel.Sprite {
	texture := image.NewRGBA(image.Rect(0, 0, width, height))
	pictureData := pixel.PictureDataFromImage(texture)

	return *pixel.NewSprite(pictureData, pictureData.Rect)
}
