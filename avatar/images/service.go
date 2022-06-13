package images

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
)

const (
	blankAvatar = "blankAvatar.png" // blankAvatar is the default name for the base image.
	bytesLen    = 20                // bytesLen is the len of the array of bytes returned by the encoder.
	newAvatar   = "newAvatar.jpg"   // newAvatar is the name of the output file.
)

/*
imageGenerator is a struct that allows to assign colors to the encoded info.
Edits the blankAvatar and generates the identicon.
*/
type imageGenerator struct {
	colorEngine ColorTransformer // it has a []byte to []color.Color converter

}

// colorTransformer is an interface designed to isolate the BytesArrayToColorArray conversion process.
type ColorTransformer interface {
	//here goes the implementation to transform bytes into colors
	BytesArrayToColorArray(encodedInformations []byte) (colors []color.Color, err error)
}

// Drawer is a struct that has the methods to build and save the avatar. Color engine
// is the implementation of the color transformer interface. Is the way that
// bytes are converted to a color array.
type Drawer struct {
	colorEngine ColorTransformer
}

//NewDrawer is a constructor for the drawer struct. It takes a transformer that implements
// ColorTransformer interface and returns a pointer to the struct.
func NewDrawer(transformer ColorTransformer) *Drawer {
	return &Drawer{
		transformer,
	}
}

/*
BuildAndSaveImage is added to the drawer struct.
It receives an slice of bytes and builds and saves an identicon for the array of bytes given.
It only returns an error. It returns nil when execution was succesfull.
*/
func (drawer *Drawer) BuildAndSaveImage(encodedInformation []byte) error {
	// transforms the encodedInformation []bytes to colorArray

	colorsArray, err := drawer.colorEngine.BytesArrayToColorArray(encodedInformation)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
	// if nothing bad happens
	img := drawer.load(blankAvatar) //opens the blankAvatar
	// color editing logic
	indexColor := 0
	for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y++ {
		for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
			r, g, b, a := img.At(x, y).RGBA()
			// normalize the values
			r /= 257
			g /= 257
			b /= 257
			a /= 257
			// check if isn't black (it considers up to a gray scale and at least 1 for opacity)
			if r > 50 && g > 50 && b > 50 && a > 1 {
				// to avoid the out of bounds error we reset the indexColor
				if indexColor == (len(colorsArray) - 1) {
					indexColor = 0 // resets the index of the colorArrays
				}
				img.Set(x, y, colorsArray[indexColor]) // changes the color of the pixel
				indexColor += 1
			}
		}
	} //ends the iterations over the image
	fmt.Println("Image editting finished")
	// saves the newAvatar
	drawer.save(newAvatar, (*image.RGBA)(img))
	return nil
}

// load loads an image from the filePath string given. It returns a *image.RGBA
func (g *Drawer) load(filePath string) *image.RGBA {
	imgFile, err := os.Open(filePath)
	defer imgFile.Close()
	if err != nil {
		log.Println("Cannot read file:", err)
	}
	img, _, err := image.Decode(imgFile)
	if err != nil {
		log.Println("Cannot decode file:", err)
	}
	return img.(*image.RGBA)
}

// save function saves the img (*image.RGBA type) given in to the filePath specified.
func (g *Drawer) save(filePath string, img *image.RGBA) {
	imgFile, err := os.Create(filePath)
	defer imgFile.Close()
	if err != nil {
		log.Println("Cannot create file:", err)
	}
	png.Encode(imgFile, img.SubImage(img.Rect))
}
