package ansi

import (
	"math"
)

// Color defines an interface for all possible forms of
// color used by the application.
type Color interface {
	// RGB returns the red, green and blue components
	// of the color, each in the range 0-255.
	RGB() (R, G, B uint8)
}

// RGBFrom takes a concrete [Color] and returns its RGB
// represetation. 
// 
// This serves to avoid recalculations for more complex
// [Color] types.
func RGBFromColor(c Color) RGB {
	r, g, b := c.RGB()
	return RGB{r, g, b}
}

// RGB is a color defined by its red, green and blue
// components.
type RGB struct{ R, G, B uint8 }

// RGB returns the red, green and blue components,
// satisfying the [Color] interface.
func (c RGB) RGB() (uint8, uint8, uint8) { return c.R, c.G, c.B }

// HSL is a color defined by its hue, saturation and
// lightness components.
type HSL struct{ H, S, L float32 }

// RGB returns the red, green and blue components,
// satisfying the [Color] interface.
func (c HSL) RGB() (uint8, uint8, uint8) {
	c.H = _FMod(c.H, 360)
	c.S = _FClamp(0, c.S, 1)
	c.L = _FClamp(0, c.L, 1)

	C := (1 - _FAbs(2*c.L-1)) * c.S
	H := c.H / 60
	X := C * (1 - _FAbs(_FMod(H, 2)-1))

	var r, g, b float32
	switch int(H) {
	case 0: r, g = C, X
	case 1: r, g = X, C
	case 2: g, b = C, X
	case 3: g, b = X, C
	case 4: r, b = X, C
	case 5: r, b = C, X
	}

	m := c.L - C/2
	return uint8((r + m) * 255), uint8((g + m) * 255), uint8((b + m) * 255)
}

func _FClamp(mn, x, mx float32) float32 {
	return max(mn, min(x, mx))
}

func _FAbs(x float32) float32 {
	return math.Float32frombits(math.Float32bits(x) &^ (1 << 31))
}

func _FFloor(x float32) float32 {
	return float32(int(x))
}

func _FMod(x, mod float32) float32 {
	res := x - _FFloor(x/mod)*mod
	if res < 0 {
		return res + mod
	}

	return res
}
