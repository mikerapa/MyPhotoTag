package main

import (
	cli "./cli"
	"fmt"
	"image"
	"image/draw"
	"image/jpeg"
	"image/png"
	"log"
	"os"
)

func main() {
	// TODO get values from the command parameters

	// set up the console logger with a default value
	cli.InitLogger("Error")

	photoFilePath, tagFilePath, outputFilePath := cli.ParseCommandParameters()
	tagPhoto(photoFilePath, tagFilePath, outputFilePath)

}

func tagPhoto(photoFilePath string, tagFilePath string, outputFilePath string) {
	cli.ConsoleLogger.Trace("Tagging photo:", photoFilePath)

	image1 := openImageFile(photoFilePath)
	defer image1.Close()

	tagImage := openImageFile(tagFilePath)
	defer tagImage.Close()

	photoJpg, err := jpeg.Decode(image1)
	if err != nil {
		log.Fatalf("failed to decode: %s", err)
	}

	tagPng, err := png.Decode(tagImage)
	if err != nil {
		log.Fatalf("failed to decode: %s", err)
	}

	image3 := createCombinedImage(photoJpg, tagPng)

	third, err := os.Create(outputFilePath)
	if err != nil {
		log.Fatalf("failed to create: %s", err)
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
	offset := calculateTagCoordinate(photoBounds, tagBounds)
	draw.Draw(image3, tagPng.Bounds().Add(offset), tagPng, image.ZP, draw.Over)
	return image3
}

func calculateTagCoordinate(photoBounds image.Rectangle, tagBounds image.Rectangle) (destinationPoint image.Point) {
	if tagBounds.Dx() > photoBounds.Dx() || tagBounds.Dy() > photoBounds.Dy() {
		log.Println("Tag dimensions are greater than photo dimensions.")
		destinationPoint = image.Pt(0, 0)
	} else {
		destinationPoint = image.Pt(photoBounds.Dx()-tagBounds.Dx(), photoBounds.Dy()-tagBounds.Dy())
	}
	return
}

func openImageFile(imageFilePath string) (openedFile *os.File) {
	openedFile, err := os.Open(imageFilePath)
	if err != nil {
		log.Fatalf("failed to open: %s", err)
	}
	return
}
