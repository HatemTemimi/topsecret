package utils

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	types "server/internal/rental/types"
	"strings"
	"sync"

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
		if !strings.Contains(imagePath, "https://") {
			url := fmt.Sprintf("%s%s", baseURL, strings.TrimPrefix(imagePath, "../"))
			urls = append(urls, url)
		} else {
			urls = append(urls, imagePath)
		}
	}
	return urls
}

// processImages processes multiple images concurrently. Each image is resized or compressed, and the result is stored in the rental images array.
func ProcessImages(files []*multipart.FileHeader, rentalFolder string, rental *types.Rental, wg *sync.WaitGroup, mu *sync.Mutex) error {
	var processingErr error

	// Create goroutines for processing each image
	for _, file := range files {
		wg.Add(1)
		go func(file *multipart.FileHeader) {
			defer wg.Done()

			src, err := file.Open()
			if err != nil {
				processingErr = err
				return
			}
			defer src.Close()

			// Destination path for saving the processed image
			dstPath := filepath.Join(rentalFolder, file.Filename)

			// Resize or compress the image
			if err := ResizeImage(src, dstPath); err != nil {
				processingErr = err
				return
			}

			// Append processed image path to rental images
			mu.Lock()
			rental.Images = append(rental.Images, strings.TrimPrefix(dstPath, "../"))
			mu.Unlock()
		}(file)
	}

	// Wait for all goroutines to finish
	wg.Wait()

	// Check if any error occurred during image processing
	if processingErr != nil {
		return fmt.Errorf("failed to process images: %w", processingErr)
	}

	return nil
}
