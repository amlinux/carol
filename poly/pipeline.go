// Copyright (c) 2013-2017 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package poly

//------------------------------------------------------------------------------

import (
	"strings"

	"github.com/drakmaniso/carol/x/gl"
)

//------------------------------------------------------------------------------

func PipelineSetup() gl.PipelineConfig {
	return func(p *gl.Pipeline) {
		gl.VertexShader(strings.NewReader(vertshader))(p)
		gl.Topology(gl.Triangles)(p)
		gl.CullFace(false, true)(p)
		gl.DepthTest(true)(p)
	}
}

//------------------------------------------------------------------------------
