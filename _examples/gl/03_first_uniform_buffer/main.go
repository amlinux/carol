// Copyright (c) 2013-2016 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package main

//------------------------------------------------------------------------------

import (
	"github.com/drakmaniso/carol"
	"github.com/drakmaniso/carol/colour"
	"github.com/drakmaniso/carol/x/gl"
	"github.com/drakmaniso/carol/plane"
)

//------------------------------------------------------------------------------

func main() {
	err := carol.Run(loop{})
	if err != nil {
		carol.ShowError(err)
		return
	}
}

//------------------------------------------------------------------------------

// OpenGL objects
var (
	pipeline    *gl.Pipeline
	perFrameUBO gl.UniformBuffer
)

// Uniform buffer
var perFrame struct {
	transform plane.GPUMatrix
}

// Vertex buffer
type mesh []struct {
	position plane.Coord `layout:"0"`
	color    colour.RGB  `layout:"1"`
}

// Animation state
var (
	angle float64
)

//------------------------------------------------------------------------------

type loop struct {
	carol.Handlers
}

//------------------------------------------------------------------------------

func (loop) Setup() error {
	var triangle mesh

	// Create and configure the pipeline
	pipeline = gl.NewPipeline(
		gl.Shader(carol.Path()+"shader.vert"),
		gl.Shader(carol.Path()+"shader.frag"),
		gl.VertexFormat(0, triangle),
		gl.Topology(gl.Triangles),
	)
	gl.Enable(gl.FramebufferSRGB)

	// Create the uniform buffer
	perFrameUBO = gl.NewUniformBuffer(&perFrame, gl.DynamicStorage)

	// Fill and create the vertex buffer
	triangle = mesh{
		{plane.Coord{0, 0.75}, colour.RGB{R: 0.3, G: 0, B: 0.8}},
		{plane.Coord{-0.65, -0.465}, colour.RGB{R: 0.8, G: 0.3, B: 0}},
		{plane.Coord{0.65, -0.465}, colour.RGB{R: 0, G: 0.6, B: 0.2}},
	}
	vbo := gl.NewVertexBuffer(triangle, gl.StaticStorage)

	// Bind the vertex buffer to the pipeline
	pipeline.Bind()
	vbo.Bind(0, 0)
	pipeline.Unbind()

	return carol.Error("gl", gl.Err())
}

//------------------------------------------------------------------------------

func (loop) Update() error {
	return nil
}

//------------------------------------------------------------------------------

func (loop) Draw(dt, _ float64) error {
	angle -= 1.0 * dt

	pipeline.Bind()
	gl.ClearColorBuffer(colour.RGBA{0.9, 0.9, 0.9, 1.0})

	perFrame.transform = plane.Rotation(float32(angle)).GPU()
	perFrameUBO.SubData(&perFrame, 0)
	perFrameUBO.Bind(0)

	gl.Draw(0, 3)
	pipeline.Unbind()

	return gl.Err()
}

//------------------------------------------------------------------------------
