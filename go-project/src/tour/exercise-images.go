package tour

import "image"
import "image/color"

type Image struct{
	x int
	y int
}

func (im Image) Bounds() (image.Rectangle) {
	return image.Rect(0, 0, im.x, im.y)
}

func (im Image) ColorModel() (color.Model) {
	return color.RGBAModel
}

func (im Image) At(x, y int) (color.Color) {
    v := uint8((x + y) / 2)
    return color.RGBA{v, v, 255, 255}
}
