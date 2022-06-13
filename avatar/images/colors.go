package images

import (
	"fmt"
	"image/color"
	"image/color/palette"
)

type ColorCreatorFromBytes struct{}

func NewColorCreatorFromBytes() *ColorCreatorFromBytes {
	return &ColorCreatorFromBytes{}
}

func (c *ColorCreatorFromBytes) BytesArrayToColorArray(encodedInformations []byte) (colors []color.Color, err error) {
	// the first rusty implementation is to asign some range of byte values to a defined color
	// first I defined that I will use just 4 colours
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
