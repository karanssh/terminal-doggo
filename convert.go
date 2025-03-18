package main

import (
	"fmt"
	"image"
	"image/color"
	_ "image/jpeg"
	_ "image/png"
	"os"
	"strings"
)

const asciiChars = " .:-=+*#%@"

func convertImage(inputFile string) {
	var width int
	var height int

	width = 50
	height = 50

	// flag.StringVar(&inputFile, "file", "", "Path to the image file")
	// flag.IntVar(&width, "width", 80, "Width of ASCII output")
	// flag.IntVar(&height, "height", 0, "Height of ASCII output (0 for auto)")
	// flag.Parse()

	// if inputFile == "" {
	// 	fmt.Println("Please provide an input file with -file flag")
	// 	flag.Usage()
	// 	os.Exit(1)
	// }

	// Open the image file
	file, err := os.Open(inputFile)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	// Decode the image
	img, _, err := image.Decode(file)
	if err != nil {
		fmt.Printf("Error decoding image: %v\n", err)
		os.Exit(1)
	}

	// fmt.Printf("Image format: %s\n", format)

	// Get original dimensions
	bounds := img.Bounds()
	imgWidth := bounds.Max.X - bounds.Min.X
	imgHeight := bounds.Max.Y - bounds.Min.Y

	// Calculate height if not specified
	if height == 0 {
		// Maintain aspect ratio
		height = int(float64(width) * float64(imgHeight) / float64(imgWidth) / 2)
	}

	// Convert to ASCII
	asciiArt := convertToASCII(img, width, height)
	fmt.Println(asciiArt)
}

func convertToASCII(img image.Image, width, height int) string {
	bounds := img.Bounds()
	origWidth := bounds.Max.X - bounds.Min.X
	origHeight := bounds.Max.Y - bounds.Min.Y

	var result strings.Builder

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			// Map the coordinates
			srcX := bounds.Min.X + (x * origWidth / width)
			srcY := bounds.Min.Y + (y * origHeight / height)

			// Get the color at this pixel
			c := img.At(srcX, srcY)
			// Convert color to grayscale
			grayColor := convertToGrayscale(c)

			// Map grayscale value to ASCII character
			asciiIndex := int(grayColor * float64(len(asciiChars)-1))
			result.WriteString(string(asciiChars[asciiIndex]))
		}
		result.WriteString("\n")
	}

	return result.String()
}

func convertToGrayscale(c color.Color) float64 {
	r, g, b, _ := c.RGBA()
	// Convert to range [0, 1]
	rFloat := float64(r) / 65535.0
	gFloat := float64(g) / 65535.0
	bFloat := float64(b) / 65535.0

	// Weighted conversion to grayscale (standard luminance formula)
	gray := 0.299*rFloat + 0.587*gFloat + 0.114*bFloat

	return gray
}
