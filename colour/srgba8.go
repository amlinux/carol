// Copyright (c) 2013-2017 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package colour

//------------------------------------------------------------------------------

// SRGBA8 represents a 32-bit color in alpha-premultiplied sRGB color space.
// There is 8 bits for each components.
//
// An alpha-premultiplied color component c has been scaled by alpha (a), so has
// valid values 0 <= c <= a.
//
// Note that additive blending can also be achieved when alpha is set to 0 while
// the color components are non-null.
type SRGBA8 struct {
	R uint8
	G uint8
	B uint8
	A uint8
}

//------------------------------------------------------------------------------

// SRGBA8Of converts any color to alpha-premultiplied sRGB color space.
func SRGBA8Of(c Colour) SRGBA8 {
	r, g, b, a := c.Linear()
	r = linearToSrgb(r)
	g = linearToSrgb(g)
	b = linearToSrgb(b)
	return SRGBA8{uint8(r * 0xFF), uint8(g * 0xFF), uint8(b * 0xFF), uint8(a * 0xFF)}
}

//------------------------------------------------------------------------------

// Linear implements the Colour interface.
func (c SRGBA8) Linear() (r, g, b, a float32) {
	r = srgbToLinear(float32(c.R) / float32(0xFF))
	g = srgbToLinear(float32(c.G) / float32(0xFF))
	b = srgbToLinear(float32(c.B) / float32(0xFF))
	a = srgbToLinear(float32(c.A) / float32(0xFF))
	return r, g, b, a
}

//------------------------------------------------------------------------------

// RGBA implements the image.Color interface.
func (c SRGBA8) RGBA() (r, g, b, a uint32) {
	r = uint32((float64(c.R) / float64(0xFF)) * float64(0xFFFF))
	g = uint32((float64(c.G) / float64(0xFF)) * float64(0xFFFF))
	b = uint32((float64(c.B) / float64(0xFF)) * float64(0xFFFF))
	a = uint32((float64(c.A) / float64(0xFF)) * float64(0xFFFF))
	return r, g, b, a
}

//------------------------------------------------------------------------------
