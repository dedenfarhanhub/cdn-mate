package cdnmate

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

// Uploader defines the interface for uploading images
type Uploader interface {
	UploadToCDN(imageBuffer *bytes.Buffer, filename string) error
}

// DefaultUploader uploads images to a CDN
type DefaultUploader struct {
	UploaderUrl string
}

// UploadToCDN uploads an image directly to CDN from memory (without saving to disk)
func (u *DefaultUploader) UploadToCDN(imageBuffer *bytes.Buffer, filename string) error {
	url := fmt.Sprintf("%s/%s", u.UploaderUrl, filename)

	req, err := http.NewRequest("PUT", url, imageBuffer)
	if err != nil {
		return fmt.Errorf("failed to create upload request: %w", err)
	}

	req.Header.Set("User-Agent", "cdn-uploader")
	req.Header.Set("Content-Type", "image/jpeg")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to upload to CDN: %w", err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("CDN upload failed with status: %d", resp.StatusCode)
	}

	return nil
}
