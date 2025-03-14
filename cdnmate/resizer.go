package cdnmate

import (
	xdraw "golang.org/x/image/draw"
	"image"
	"image/draw"
)

// ImageResizer handles image resizing
type ImageResizer interface {
	Resize(img image.Image) image.Image
}

// DefaultResizer implements ImageResizer with a basic resizing method
type DefaultResizer struct{}

// Resize resizes an image to 99% of its original size
func (r DefaultResizer) Resize(img image.Image) image.Image {
	width := img.Bounds().Dx() * 99 / 100
	height := img.Bounds().Dy() * 99 / 100
	dst := image.NewRGBA(image.Rect(0, 0, width, height))
	xdraw.CatmullRom.Scale(dst, dst.Bounds(), img, img.Bounds(), draw.Over, nil)

	return dst
}
