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
			{255, 255, 255, 255}, //background
			{255, 0, 0, 255},  // main
			{204, 0, 0, 255},  // 2dary
			{255, 149, 129, 255},   // 2dary
		},
		"sapphire": {
			{255, 255, 255, 255}, //background
			{0, 128, 255, 255},  // main
			{0, 68, 204, 255},  // 2dary
			{153, 221, 255, 255},   // 2dary
		},
        "diamond": {
			{255, 255, 255, 255}, //background
			{184, 214, 230, 255},  // main
			{122, 177, 204, 255},  // 2dary
			{230, 247, 255, 255},   // 2dary
		},
        "emerald": {
			{255, 255, 255, 255}, //background
			{0, 179, 60, 255},  // main
			{0, 128, 43, 255},  // 2dary
			{138, 230, 168, 255},   // 2dary
		},
        "amber": {
			{255, 255, 255, 255}, //background
			{255, 128, 0, 255},  // main
			{204, 102, 0, 255},  // 2dary
			{255, 187, 51, 255},   // 2dary
		},
        "onyx": {
			{255, 255, 255, 255}, //background
			{26, 26, 26, 255},  // main
			{0, 0, 0, 255},  // 2dary
			{51, 51, 51, 255},   // 2dary
		},
        "amethyst": {
			{255, 255, 255, 255}, //background
			{128, 38, 128, 255},  // main
			{77, 0, 77, 255},  // 2dary
			{204, 143, 204, 255},   // 2dary
		},
        "azurite": {
			{255, 255, 255, 255}, //background
			{25, 25, 255, 255},  // main
			{0, 0, 153, 255},  // 2dary
			{92, 230, 230, 255},   // 2dary
		},
        "garnet": {
			{255, 255, 255, 255}, //background
			{230, 0, 115, 255},  // main
			{153, 0, 77, 255},  // 2dary
			{255, 153, 204, 255},   // 2dary
		},
        "peridot": {
			{255, 255, 255, 255}, //background
			{168, 204, 61, 255},  // main
			{147, 179, 54, 255},  // 2dary
			{207, 230, 138, 255},   // 2dary
		},
        "citrine": {
			{255, 255, 255, 255}, //background
			{255, 204, 51, 255},  // main
			{230, 178, 23, 255},  // 2dary
			{255, 230, 153, 255},   // 2dary
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
