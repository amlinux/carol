// Copyright (c) 2013-2017 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package colour

//------------------------------------------------------------------------------

// SRGBnA represents a color in *non* alpha-premultiplied sRGB color space. Each
// value ranges within [0, 1].
//
// An alpha-premultiplied color component c has been scaled by alpha (a), so has
// valid values 0 <= c <= a.
//
// Note that additive blending can also be achieved when alpha is set to 0 while
// the color components are non-null.
type SRGBnA struct {
	R float32
	G float32
	B float32
	A float32
}

//------------------------------------------------------------------------------

// SRGBnAOf converts any color to sRGB non alpha-premultiplied sRGB color space.
func SRGBnAOf(c Colour) SRGBnA {
	r, g, b, a := c.Linear()
	r = linearToSrgb(r) / a
	g = linearToSrgb(g) / a
	b = linearToSrgb(b) / a
	return SRGBnA{r, g, b, a}
}

//------------------------------------------------------------------------------

// Linear implements the Colour interface: it returns the color converted to
// alpha-premultipled linear color space.
func (c SRGBnA) Linear() (r, g, b, a float32) {
	a = c.A
	r = srgbToLinear(a * c.R)
	g = srgbToLinear(a * c.G)
	b = srgbToLinear(a * c.B)
	return r, g, b, a
}

//------------------------------------------------------------------------------

// RGBA implements the image.Color interface: it returns the four components
// scaled by 0xFFFF.
func (c SRGBnA) RGBA() (r, g, b, a uint32) {
	return uint32(c.A * c.R * 0xFFFF), uint32(c.A * c.G * 0xFFFF), uint32(c.A * c.B * 0xFFFF), uint32(c.A * 0xFFFF)
}

//------------------------------------------------------------------------------
