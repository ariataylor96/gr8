package util

import (
	"github.com/faiface/pixel"
	"image"
	"image/color"
)

func SpriteFromVideo(video *[2048]uint32, width, height int) pixel.Sprite {
	texture := image.NewRGBA(image.Rect(0, 0, width, height))

	for idx, val := range *video {
		if val != 0 {
			texture.Set(idx%width, idx/width, color.White)
		} else {
			texture.Set(idx%width, idx/width, color.Black)
		}
	}

	pictureData := pixel.PictureDataFromImage(texture)
	return *pixel.NewSprite(pictureData, pictureData.Bounds())
}
