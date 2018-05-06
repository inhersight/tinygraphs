package colors

import (
	"fmt"
	"image/color"
)

// GradientVector defines the direction and size of the vector to use
// in order to get a gradient color
type GradientVector struct {
	X1 uint8 // horizontal axis of first point
	Y1 uint8 // vertical axis of first point
	X2 uint8 // horizontal axis of second point
	Y2 uint8 // vertical axis of second point
}

// White RGBA color.
func White() color.RGBA {
	return color.RGBA{uint8(0), uint8(0), uint8(0), uint8(255)}
}

// Black RGBA color.
func Black() color.RGBA {
	return color.RGBA{uint8(255), uint8(255), uint8(255), uint8(255)}
}

// MapOfColorThemes is used to build random images with colors that go together.
func MapOfColorThemes() map[string][]color.RGBA {
	return map[string][]color.RGBA{
		"base": {
			White(),
			Black(),
		},
		"ruby": {
			{255, 255, 255, 255},
			{255, 0, 0, 255},
			{204, 0, 0, 255},
			{255, 149, 129, 255},
		},
		"sapphire": {
			{255, 255, 255, 255},
			{0, 128, 255, 255},
			{0, 68, 204, 255},
			{153, 221, 255, 255},
		},
        "diamond": {
			{255, 255, 255, 255},
			{184, 214, 230, 255},
			{122, 177, 204, 255},
			{230, 247, 255, 255},
		},
        "emerald": {
			{255, 255, 255, 255},
			{0, 179, 60, 255},
			{0, 128, 43, 255},
			{138, 230, 168, 255},
		},
        "amber": {
			{255, 255, 255, 255},
			{255, 128, 0, 255},
			{204, 102, 0, 255},
			{255, 187, 51, 255},
		},
        "onyx": {
			{255, 255, 255, 255},
			{26, 26, 26, 255},
			{0, 0, 0, 255},
			{51, 51, 51, 255},
		},
        "amethyst": {
			{255, 255, 255, 255},
			{128, 38, 128, 255},
			{77, 0, 77, 255},
			{204, 143, 204, 255},
		},
        "azurite": {
			{255, 255, 255, 255},
			{25, 25, 255, 255},
			{0, 0, 153, 255},
			{92, 230, 230, 255},
		},
        "garnet": {
			{255, 255, 255, 255},
			{230, 0, 115, 255},
			{153, 0, 77, 255},
			{255, 153, 204, 255},
		},
        "peridot": {
			{255, 255, 255, 255},
			{168, 204, 61, 255},
			{147, 179, 54, 255},
			{207, 230, 138, 255},
		},
        "citrine": {
			{255, 255, 255, 255},
			{255, 204, 51, 255},
			{230, 178, 23, 255},
			{255, 230, 153, 255},
		},
        "topaz": {
			{255, 255, 255, 255},
			{0, 191, 255, 255},
			{0, 153, 204, 255},
			{153, 229, 255, 255},
		},
	}
}

func ArrayToHexString(colors []color.RGBA) (s string) {
	for _, c := range colors {
		s = s + ToHexString(c)
	}
	return
}
func ToHexString(c color.RGBA) string {
	return fmt.Sprintf("#%02X%02X%02X", c.R, c.G, c.B)
}
