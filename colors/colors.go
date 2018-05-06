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
			{204, 204, 204, 255},
			{102, 102, 102, 255},
			{51, 51, 51, 255},
			{153, 153, 153, 255},
		},
		"ruby": {
			{255, 255, 255, 255},
			{255, 0, 0, 255},
			{230, 0, 0, 255},
			{255, 128, 128, 255},
		},
		"sapphire": {
			{255, 255, 255, 255},
			{0, 128, 255, 255},
			{0, 115, 230, 255},
			{128, 191, 255, 255},
		},
        "diamond": {
			{255, 255, 255, 255},
			{161, 207, 230, 255},
			{122, 177, 204, 255},
			{217, 242, 255, 255},
		},
        "emerald": {
			{255, 255, 255, 255},
			{0, 204, 34, 255},
			{0, 179, 30, 255},
			{115, 230, 134, 255},
		},
        "amber": {
			{255, 255, 255, 255},
			{255, 149, 0, 255},
			{230, 134, 0, 255},
			{255, 202, 128, 255},
		},
        "onyx": {
			{255, 255, 255, 255},
			{26, 26, 26, 255},
			{0, 0, 0, 255},
			{51, 51, 51, 255},
		},
        "amethyst": {
			{255, 255, 255, 255},
			{153, 61, 153, 255},
			{128, 51, 128, 255},
			{204, 163, 204, 255},
		},
        "garnet": {
			{255, 255, 255, 255},
			{230, 0, 115, 255},
			{204, 0, 102, 255},
			{255, 128, 191, 255},
		},
        "peridot": {
			{255, 255, 255, 255},
			{177, 204, 41, 255},
			{155, 179, 36, 255},
			{218, 230, 161, 255},
		},
        "citrine": {
			{255, 255, 255, 255},
			{255, 198, 25, 255},
			{230, 178, 23, 255},
			{255, 230, 153, 255},
		},
        "topaz": {
			{255, 255, 255, 255},
			{0, 191, 255, 255},
			{0, 172, 230, 255},
			{128, 223, 255, 255},
		},
        "pearl": {
			{255, 255, 255, 255},
			{255, 204, 179, 255},
			{242, 178, 145, 255},
			{255, 221, 204, 255},
		},
        "quartz": {
			{255, 255, 255, 255},
			{204, 204, 204, 255},
			{179, 179, 179, 255},
			{230, 230, 230, 255},
		},
        "beryl": {
			{255, 255, 255, 255},
			{0, 179, 164, 255},
			{0, 153, 140, 255},
			{115, 230, 220, 255},
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
