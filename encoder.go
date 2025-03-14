package cdnmate

import (
	"bytes"
	"fmt"
	"github.com/chai2010/webp"
	"image"
	"image/jpeg"
	"image/png"
)

// ImageEncoder handles encoding and decoding images
type ImageEncoder interface {
	EncodeImage(buf *bytes.Buffer, img image.Image, format string, imageQuality float32) error
}

// DefaultEncoder implements ImageEncoder
type DefaultEncoder struct{}

// EncodeImage converts an image to a buffer in the given format
func (s *DefaultEncoder) EncodeImage(buf *bytes.Buffer, img image.Image, format string, imageQuality float32) error {
	switch format {
	case "jpeg", "jpg":
		return jpeg.Encode(buf, img, &jpeg.Options{Quality: int(imageQuality)})
	case "png":
		return png.Encode(buf, img)
	case "webp":
		return webp.Encode(buf, img, &webp.Options{Quality: imageQuality})
	default:
		return fmt.Errorf("unsupported format: %s", format)
	}
}
