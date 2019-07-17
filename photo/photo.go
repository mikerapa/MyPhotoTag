package photo

import (
	"../cli"
	"fmt"
	"image"
	"image/draw"
	"image/jpeg"
	"image/png"
	"os"
)

func TagPhoto(photoFilePath string, tagFilePath string, outputFilePath string) {
	cli.ConsoleLogger.Trace("Tagging photo:", photoFilePath)

	image1 := openImageFile(photoFilePath)
	defer image1.Close()

	tagImage := openImageFile(tagFilePath)
	defer tagImage.Close()

	photoJpg, err := jpeg.Decode(image1)
	if err != nil {
		cli.ConsoleLogger.Fatalf("failed to decode jpeg file: %s", err)
	}

	tagPng, err := png.Decode(tagImage)
	if err != nil {
		cli.ConsoleLogger.Fatalf("Failed to decode png file: %s", err)
	}

	image3 := createCombinedImage(photoJpg, tagPng)

	third, err := os.Create(outputFilePath)
	if err != nil {
		cli.ConsoleLogger.Fatalf("failed to create output file: %s", err)
	}
	jpeg.Encode(third, image3, &jpeg.Options{jpeg.DefaultQuality})
	defer third.Close()

	fmt.Println("Outputting file:", outputFilePath)
}

func createCombinedImage(photoJpg image.Image, tagPng image.Image) *image.RGBA {
	photoBounds := photoJpg.Bounds()
	tagBounds := tagPng.Bounds()
	image3 := image.NewRGBA(photoBounds)
	draw.Draw(image3, photoBounds, photoJpg, image.ZP, draw.Src)
	offset := CalculateTagCoordinate(photoBounds, tagBounds)
	draw.Draw(image3, tagPng.Bounds().Add(offset), tagPng, image.ZP, draw.Over)
	return image3
}

func CalculateTagCoordinate(photoBounds image.Rectangle, tagBounds image.Rectangle) (destinationPoint image.Point) {
	cli.ConsoleLogger.Trace("Calculating tag location")
	if tagBounds.Dx() > photoBounds.Dx() || tagBounds.Dy() > photoBounds.Dy() {
		cli.ConsoleLogger.Error("Tag dimensions are greater than photo dimensions. The tag will be located in the upper left position.")
		destinationPoint = image.Pt(0, 0)
	} else {
		destinationPoint = image.Pt(photoBounds.Dx()-tagBounds.Dx(), photoBounds.Dy()-tagBounds.Dy())
	}
	cli.ConsoleLogger.Infof("Tag location: %v", destinationPoint)
	return
}

func openImageFile(imageFilePath string) (openedFile *os.File) {
	openedFile, err := os.Open(imageFilePath)
	if err != nil {
		cli.ConsoleLogger.Fatalf("failed to open: %s", err)
	}
	return
}
