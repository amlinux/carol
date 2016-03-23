// Copyright (c) 2013-2016 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package internal

//------------------------------------------------------------------------------

/*
#include "glad.h"

GLuint CompileShader(const GLchar* b, GLenum t);
char* CompileShaderError(GLuint s);
GLuint LinkProgram(GLuint vs, GLuint fs);
char* LinkProgramError(GLuint s);
GLuint SetupVAO();
void ClosePipeline(GLuint p, GLuint vao);

static inline void BindPipeline(GLuint p, GLuint vao) {	glClear (GL_COLOR_BUFFER_BIT |  GL_DEPTH_BUFFER_BIT); glUseProgram(p); glBindVertexArray(vao);};
static inline void DrawArrays(GLenum m, GLuint f, GLuint c) {glDrawArrays(m, f, c);};
*/
import "C"

import (
	"errors"
	"io"
	"io/ioutil"
	"log"
	"unsafe"
)

//------------------------------------------------------------------------------

type Pipeline struct {
	program C.GLuint
	vao     C.GLuint
}

//------------------------------------------------------------------------------

func (p *Pipeline) CompileShaders(
	vertexShader io.Reader,
	fragmentShader io.Reader,
) error {
	vs, err := compileShader(vertexShader, C.GL_VERTEX_SHADER)
	if err != nil {
		log.Print("vertex shader compile error: ", err)
	}
	fs, err := compileShader(fragmentShader, C.GL_FRAGMENT_SHADER)
	if err != nil {
		log.Print("fragment shader compile error: ", err)
	}

	p.program = C.LinkProgram(vs, fs)
	if errm := C.LinkProgramError(p.program); errm != nil {
		defer C.free(unsafe.Pointer(errm))
		err = errors.New(C.GoString(errm))
		log.Print("shader link error: ", err)
		return err
	}
	return err
}

func compileShader(r io.Reader, t C.GLenum) (C.GLuint, error) {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		log.Print("failed to read shader: ", err)
		return 0, err
	}
	cb := C.CString(string(b))
	defer C.free(unsafe.Pointer(cb))

	s := C.CompileShader((*C.GLchar)(unsafe.Pointer(cb)), t)
	if errm := C.CompileShaderError(s); errm != nil {
		defer C.free(unsafe.Pointer(errm))
		return s, errors.New(C.GoString(errm))
	}

	return s, nil
}

func (p *Pipeline) SetupVAO() error {
	p.vao = C.SetupVAO()
	return nil
}

func (p *Pipeline) Bind() {
	C.BindPipeline(p.program, p.vao)
}

func (p *Pipeline) Close() {
	C.ClosePipeline(p.program, p.vao)
}

//------------------------------------------------------------------------------

func Draw(mode uint32, first, count int32) {
	C.DrawArrays(C.GLenum(mode), C.GLuint(first), C.GLuint(count))
}

//------------------------------------------------------------------------------

const (
	GlByteEnum          C.GLenum = C.GL_BYTE
	GlUnsignedByteEnum  C.GLenum = C.GL_UNSIGNED_BYTE
	GlShortEnum         C.GLenum = C.GL_SHORT
	GlUnsignedShortEnum C.GLenum = C.GL_UNSIGNED_SHORT
	GlIntEnum           C.GLenum = C.GL_INT
	GlUnsignedIntEnum   C.GLenum = C.GL_UNSIGNED_INT
	GlFloatEnum         C.GLenum = C.GL_FLOAT
	GlDoubleEnum        C.GLenum = C.GL_DOUBLE
)

//------------------------------------------------------------------------------
