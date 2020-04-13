package main

import (
	"fmt"
	"github.com/faiface/pixel"
	"image"
	"os"
)

// getRectInGrid: Get a rectangle in the game's grid
func getRectInGrid(width float64, height float64, totalX int, totalY int, x int, y int) pixel.Rect {
	gridWidth := width / float64(totalX)
	gridHeight := height / float64(totalY)
	return pixel.R(float64(x)*gridWidth, float64(y)*gridHeight, float64(x+1)*gridWidth, float64(y+1)*gridHeight)
}

// getFrame: Get frame coordinates from the X and y values
func getFrame(frameWidth float64, frameHeight float64,
	xGrid int, yGrid int) pixel.Rect {
	return pixel.R(
		float64(xGrid)*frameWidth,
		float64(yGrid)*frameHeight,
		float64(xGrid+1)*frameWidth,
		float64(yGrid+1)*frameHeight,
	)
}

// getSheet: Load spritemap into memory
func getSheet(filePath string) (pixel.Picture, error) {
	imgFile, err := os.Open(filePath)
	defer imgFile.Close()
	if err != nil {
		fmt.Println("Cannot read file:", err)
		return nil, err
	}

	img, _, err := image.Decode(imgFile)
	if err != nil {
		fmt.Println("Cannot decode file:", err)
		return nil, err
	}
	sheet := pixel.PictureDataFromImage(img)
	return sheet, nil
}
