package images

import (
	"fmt"
	"image/color"
	"image/color/palette"
)

// ColorCreatorFromBytes has the methods to convert bytes to color.
type ColorCreatorFromBytes struct{}

// NewColorCreatorFromBytes is a constructor for ColorCreatorFromBytes struct.
func NewColorCreatorFromBytes() *ColorCreatorFromBytes {
	return &ColorCreatorFromBytes{}
}

// BytesArrayToColorArray converts a slice of bytes into a slice of color.Color type.
func (c *ColorCreatorFromBytes) BytesArrayToColorArray(encodedInformations []byte) (colors []color.Color, err error) {
	// iterates over the encoded information and converts every byte into a color array
	for _, b := range encodedInformations {
		tempColor, err := c.ByteToColor(b)
		if err != nil {
			fmt.Print(err)
		}
		colors = append(colors, tempColor)
	}

	return colors, err
}

// This function converts a byte to a color array ([R,G,B,A]).
func (c *ColorCreatorFromBytes) ByteToColor(b byte) (colorByte color.Color, err error) {
	pallete := palette.Plan9
	tempColor := pallete[b]
	colorByte = tempColor

	return colorByte, nil
}
