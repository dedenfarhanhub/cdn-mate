package cdnmate

import (
	"os"
)

// Config holds application configurations
type Config struct {
	CDNImageQuality string
	CDNUrl          string
	UploaderUrl     string
}

// LoadConfig loads configurations from environment variables
func LoadConfig() Config {
	return Config{
		CDNImageQuality: getEnv("CDNImageQuality", "./cdn"),
		CDNUrl:          getEnv("CDNUrl", "https://akcdn.detik.net.id/community/media/detik-live-shopping/"),
		UploaderUrl:     getEnv("UploaderUrl", "http://127.0.0.3:80/media/detik-live-shopping/"),
	}
}

// getEnv retrieves an environment variable with a default value
func getEnv(key, defaultVal string) string {
	if val, exists := os.LookupEnv(key); exists {
		return val
	}
	return defaultVal
}
