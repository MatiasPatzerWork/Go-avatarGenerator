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
	blankAvatar = "blankAvatar.png" // the default name for the output image
	bytesLen    = 20                // this is the len of the array of bytes returned by the encoder
	newAvatar   = "newAvatar.jpg"
)

/*
is the structure that allows to assign colors to the encoded info.
Edits the blankAvatar and generates the identicon.
*/
type imageGenerator struct {
	colorEngine ColorTransformer // it has a []byte to []color.Color converter

}

type ColorTransformer interface {
	//here goes the implementation to transform bytes into colors
	BytesArrayToColorArray(encodedInformations []byte) (colors []color.Color, err error)
}

// drawer tiene una interface que es ColorTransformer o algo asi que transforma de
// byte a color y tiene los parámetros de la imagen width, height y strides
// estas estructuras son privadas. Porque no quiero el que ocupe la libreria
// modifique el colorTransformer ni que formas. Por eso debo darle un puntero a una
// estructura instanciada con los parámetros que el usuario adiciona
type Drawer struct {
	colorEngine ColorTransformer
}

//This is a constructor for the drawer struct. It takes a transformer that implements
// ColorTransformer interface, width and height (int values) and returns a pointer.
func NewDrawer(transformer ColorTransformer) *Drawer {
	return &Drawer{
		transformer,
	}
}

/*
This function is added to the imageGenerator struct.
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

func (g *Drawer) save(filePath string, img *image.RGBA) {
	imgFile, err := os.Create(filePath)
	defer imgFile.Close()
	if err != nil {
		log.Println("Cannot create file:", err)
	}
	png.Encode(imgFile, img.SubImage(img.Rect))
}
