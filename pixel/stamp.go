// Copyright (c) 2013-2017 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package pixel

//------------------------------------------------------------------------------

import (
	"github.com/drakmaniso/carol/core/gl"
)

//------------------------------------------------------------------------------

type stamp struct {
	//  word
	mode, bin int16

	//  word
	x, y int16

	//  word
	w, h int16

	//  word
	u, v int16
}

var stamps []stamp

//------------------------------------------------------------------------------

var stampPipeline *gl.Pipeline

var stampSSBO gl.StorageBuffer

var screenUBO gl.UniformBuffer

var screenUniforms struct {
	PixelSize struct{ X, Y float32 }
}

//------------------------------------------------------------------------------