// Copyright (c) 2013-2017 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package microtext

import (
	"strings"

	"github.com/drakmaniso/glam/color"
	"github.com/drakmaniso/glam/gfx"
	"github.com/drakmaniso/glam/pixel"
	"time"
	"unsafe"
)

//------------------------------------------------------------------------------

var (
	pipeline   gfx.Pipeline
	fontSSBO   gfx.StorageBuffer
	screenSSBO gfx.StorageBuffer
)

//------------------------------------------------------------------------------

func init() {
	screen.nbCols = 1
	screen.nbRows = 1
	screen.pixelSize = 2
	SetColor(color.RGB{1, 1, 1}, color.RGB{0, 0, 0})
	SetBgAlpha(true)
	Text = make([]byte, screen.nbCols*screen.nbRows)
}

//------------------------------------------------------------------------------

const charWidth = 7
const charHeight = 11

func WindowResized(s pixel.Coord, ts time.Duration) {
	screen.nbCols = s.X / (charWidth * screen.pixelSize)
	screen.nbRows = s.Y / (charHeight * screen.pixelSize)
	screen.top = 0
	screen.left = 0

	// Reallocate the SSBO
	Text = make([]byte, screen.nbCols*screen.nbRows)
	screenSSBO.Delete()
	screenSSBO = gfx.NewStorageBuffer(
		unsafe.Sizeof(screen)+uintptr(screen.nbCols*screen.nbRows),
		gfx.DynamicStorage,
	)

	TextUpdated = true

	// Calculate the margins
	l := (s.X - (charWidth * int32(screen.pixelSize) * int32(screen.nbCols))) / 2
	if l > 0 {
		screen.left = int32(l)
	} else {
		screen.left = 0
	}
	if true {
		t := (s.Y - (charHeight * int32(screen.pixelSize) * int32(screen.nbRows))) / 2
		if t > 0 {
			screen.top = int32(t)
		} else {
			screen.top = 0
		}
	} else {
		screen.top = 0
	}

	screenUpdated = true
}

//------------------------------------------------------------------------------

func Setup() {
	pipeline = gfx.NewPipeline(
		gfx.VertexShader(strings.NewReader(vertexShader)),
		gfx.FragmentShader(strings.NewReader(fragmentShader)),
	)

	fontSSBO = gfx.NewStorageBuffer(&Font, gfx.StaticStorage)
}

//------------------------------------------------------------------------------

func Draw() {
	pipeline.Bind()
	gfx.Disable(gfx.DepthTest)
	gfx.CullFace(false, false)
	if screenUpdated {
		screenSSBO.SubData(&screen, 0)
		screenUpdated = false
	}
	if TextUpdated {
		screenSSBO.SubData(Text, unsafe.Sizeof(screen))
		TextUpdated = false
	}
	fontSSBO.Bind(0)
	screenSSBO.Bind(1)
	gfx.Draw(gfx.TriangleStrip, 0, 4)
	pipeline.Unbind()
}

//------------------------------------------------------------------------------

// Data for the SSBO
var (
	screen struct {
		left   int32
		top    int32
		nbCols int32
		nbRows int32
		//
		fgRed     float32
		fgGreen   float32
		fgBlue    float32
		pixelSize int32
		//
		bgRed   float32
		bgGreen float32
		bgBlue  float32
		bgAlpha float32
	}

	Text []byte
)

var (
	screenUpdated bool
	TextUpdated   bool
)

//------------------------------------------------------------------------------

func SetColor(fg, bg color.RGB) {
	screen.fgRed = fg.R
	screen.fgGreen = fg.G
	screen.fgBlue = fg.B

	screen.bgRed = bg.R
	screen.bgGreen = bg.G
	screen.bgBlue = bg.B

	screenUpdated = true
}

func SetBgAlpha(o bool) {
	if o {
		screen.bgAlpha = 1.0
	} else {
		screen.bgAlpha = 0.0
	}
	screenUpdated = true
}

func GetBgAlpha() bool {
	return screen.bgAlpha != 0.0
}

func ToggleBgAlpha() {
	if screen.bgAlpha != 0 {
		screen.bgAlpha = 0
	} else {
		screen.bgAlpha = 1
	}
	screenUpdated = true
}

func Size() (cols, rows int) {
	return int(screen.nbCols), int(screen.nbRows)
}

//------------------------------------------------------------------------------

func Peek(x, y int) byte {
	return Text[x+y*int(screen.nbCols)]
}

func Poke(x, y int, c byte) {
	Text[x+y*int(screen.nbCols)] = c
}

//------------------------------------------------------------------------------
