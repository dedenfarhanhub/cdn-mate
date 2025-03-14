package cdnmate

import (
	"bytes"
	"fmt"
	"image"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
)

// ImageProcessor handles image processing (resizing and storing images)
type ImageProcessor struct {
	Resizer  ImageResizer
	Encoder  ImageEncoder
	Uploader Uploader
	Config   Config
}

// NewImageProcessor initializes an ImageProcessor with dependencies
func NewImageProcessor(config Config) *ImageProcessor {
	return &ImageProcessor{
		Resizer:  DefaultResizer{},
		Encoder:  &DefaultEncoder{},
		Uploader: &DefaultUploader{UploaderUrl: config.UploaderUrl},
		Config:   config,
	}
}

// ProcessImage processes the uploaded image (decode, resize, and save)
func (p *ImageProcessor) ProcessImage(file multipart.File, header *multipart.FileHeader, imageQuality float32) (string, error) {
	// Decode the image from the uploaded file
	img, format, err := DecodeImage(file)
	if err != nil {
		return "", fmt.Errorf("failed to decode image: %w", err)
	}

	// Resize the image to 99% of its original size
	resizedImg := p.Resizer.Resize(img)

	//  Convert the resized image into a byte buffer
	var buf bytes.Buffer
	if err := p.Encoder.EncodeImage(&buf, resizedImg, format, imageQuality); err != nil {
		return "", fmt.Errorf("failed to encode image: %w", err)
	}

	// Generate a unique filename to avoid duplication
	newFileName := GenerateUniqueFilename(header.Filename)

	// Save the processed image to the CDN
	if err := p.Uploader.UploadToCDN(&buf, newFileName); err != nil {
		return "", fmt.Errorf("failed to save image: %w", err)
	}

	savePath := p.Config.CDNUrl + newFileName
	return savePath, nil
}

// DecodeImage reads and decodes an image file from a multipart request
func DecodeImage(file multipart.File) (image.Image, string, error) {
	img, format, err := image.Decode(file)
	if err != nil {
		return nil, "", err
	}
	return img, format, nil
}

// GenerateUniqueFilename generates a unique filename to prevent duplication
func GenerateUniqueFilename(originalName string) string {
	ext := strings.ToLower(filepath.Ext(originalName))
	name := strings.TrimSuffix(originalName, ext)
	return fmt.Sprintf("%s_%d%s", name, os.Getpid(), ext) // Example: image_12345.jpg
}
