package main

import (
    //"code.google.com/p/go-tour/pic"
    "golang.org/x/tour/pic"
    "image"
    "image/color"
)
type Image struct{
    w, h int
}

func (img *Image) ColorModel() color.Model {
    return color.RGBAModel
}

func (img *Image) Bounds() image.Rectangle {
    return image.Rect(0, 0, img.w , img.h)
}

func (img *Image) At(x, y int) color.Color {
    return color.RGBA{uint8(x), uint8(y), 255, 255}
}

func main() {
    m := &Image{256, 256}
    pic.ShowImage(m)
}

