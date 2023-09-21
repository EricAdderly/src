package main

import (
	"bytes"
	"image"
	"image/jpeg"
	"io"
	"log"
	"os"

	"github.com/nfnt/resize"
)

func main() {
	file, err := os.Open("2219339.jpg")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, file); err != nil {
		return
	}

	image, _, err := image.Decode(bytes.NewReader(buf.Bytes()))
	if err != nil {
		return
	}

	width := image.Bounds().Dx()
	height := image.Bounds().Dy()

	if width > 1024 || height > 1024 {
		return
	}

	if width > 1024 && height > 1024 {
		newImage := resize.Resize(1024, 1024, image, resize.Lanczos3)
		changePermission(newImage)
		return
	}

	if height < 1024 {
		newHeight := uint(height)
		newImage := resize.Resize(1024, newHeight, image, resize.Lanczos3)
		changePermission(newImage)
		return
	}

	if height > 1024 {
		newWidth := uint(width)
		newImage := resize.Resize(newWidth, 1024, image, resize.Lanczos3)
		changePermission(newImage)
		return
	}
}

func changePermission(newImage image.Image) {
	result := bytes.NewBuffer(nil)

	if err := jpeg.Encode(result, newImage, nil); err != nil {
		return
	}

	if err := os.WriteFile("1.jpg", result.Bytes(), 0644); err != nil {
		return
	}
}
