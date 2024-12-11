package utils

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"os"
	"strings"

	"github.com/disintegration/imaging"
	"github.com/labstack/echo/v4"
)

func GetBasePath() string {
	basePath := os.Getenv("ASSETS_BASE_PATH")
	if basePath == "" {
		basePath = "../assets/rentals"
	}
	return basePath
}

// compressOrResizeImage compresses or resizes an image to standard web sizes.
func ResizeImage(src io.Reader, dstPath string) error {
	// Decode the image
	img, format, err := image.Decode(src)
	if err != nil {
		return fmt.Errorf("failed to decode image: %w", err)
	}

	// Resize the image to 1200x800 while maintaining aspect ratio
	resizedImg := imaging.Resize(img, 1200, 0, imaging.Lanczos)

	// Create the destination file
	dstFile, err := os.Create(dstPath)
	if err != nil {
		return fmt.Errorf("failed to create destination file: %w", err)
	}
	defer dstFile.Close()

	// Encode the resized image back to the destination file
	switch format {
	case "jpeg":
		err = jpeg.Encode(dstFile, resizedImg, &jpeg.Options{Quality: 80}) // Compress with 80% quality
	case "png":
		err = png.Encode(dstFile, resizedImg)
	default:
		return fmt.Errorf("unsupported image format: %s", format)
	}

	if err != nil {
		return fmt.Errorf("failed to encode image: %w", err)
	}

	return nil
}

// mapImagePathsToURLs converts file paths to public URLs.
func MapImagePathsToURLs(c echo.Context, imagePaths []string) []string {
	baseURL := fmt.Sprintf("http://%s/", c.Request().Host) // Base URL for assets
	var urls []string
	for _, imagePath := range imagePaths {
		url := fmt.Sprintf("%s%s", baseURL, strings.TrimPrefix(imagePath, "../"))
		urls = append(urls, url)
	}
	return urls
}
