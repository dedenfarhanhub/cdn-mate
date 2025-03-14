package cdnmate

import (
	"bytes"
	"fmt"
	"github.com/kolesa-team/go-webp/webp"
	"image"
	"image/jpeg"
	"image/png"

	"github.com/kolesa-team/go-webp/encoder"
)

// ImageEncoder handles encoding images into different formats
type ImageEncoder interface {
	EncodeImage(buf *bytes.Buffer, img image.Image, format string, quality int) error
}

// DefaultEncoder implements ImageEncoder
type DefaultEncoder struct{}

// EncodeImage converts an image to a buffer in the specified format with quality settings
func (e *DefaultEncoder) EncodeImage(buf *bytes.Buffer, img image.Image, format string, quality int) error {
	switch format {
	case "jpeg", "jpg":
		opts := jpeg.Options{Quality: quality}
		return jpeg.Encode(buf, img, &opts)
	case "png":
		return png.Encode(buf, img)
	case "webp":
		options, err := encoder.NewLossyEncoderOptions(encoder.PresetDefault, float32(quality))
		if err != nil {
			return fmt.Errorf("failed to create WebP encoder options: %w", err)
		}
		// Encode image to WebP
		if err = webp.Encode(buf, img, options); err != nil {
			return fmt.Errorf("failed to encode WebP image: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unsupported format: %s", format)
	}
}
