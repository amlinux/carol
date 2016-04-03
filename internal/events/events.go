// Copyright (c) 2013-2016 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package events

//------------------------------------------------------------------------------

/*
#cgo windows LDFLAGS: -lSDL2
#cgo linux freebsd darwin pkg-config: sdl2

#include "../sdl.h"

#define PEEP_SIZE 128

SDL_Event Events[PEEP_SIZE];

int PeepEvents();
*/
import "C"

import (
	"time"
	"unsafe"

	"github.com/drakmaniso/glam/geom"
	"github.com/drakmaniso/glam/gfx"
	"github.com/drakmaniso/glam/internal"
	"github.com/drakmaniso/glam/key"
	"github.com/drakmaniso/glam/mouse"
	"github.com/drakmaniso/glam/window"
)

//------------------------------------------------------------------------------

func Process() {
	more := true
	for more && !internal.QuitRequested {
		n := PeepEvents()
		for i := 0; i < n && !internal.QuitRequested; i++ {
			e := EventAt(i)
			dispatch(e)
		}
		more = n >= C.PEEP_SIZE
	}
}

func dispatch(e unsafe.Pointer) {
	ts := time.Duration(((*C.SDL_CommonEvent)(e)).timestamp) * time.Millisecond
	switch ((*C.SDL_CommonEvent)(e))._type {
	case C.SDL_QUIT:
		window.Handler.WindowQuit(ts)
	// Window Events
	case C.SDL_WINDOWEVENT:
		e := (*C.SDL_WindowEvent)(e)
		switch e.event {
		case C.SDL_WINDOWEVENT_NONE:
			// Ignore
		case C.SDL_WINDOWEVENT_SHOWN:
			window.Handler.WindowShown(ts)
		case C.SDL_WINDOWEVENT_HIDDEN:
			window.Handler.WindowHidden(ts)
		case C.SDL_WINDOWEVENT_EXPOSED:
			// Ignore
		case C.SDL_WINDOWEVENT_MOVED:
			// Ignore
		case C.SDL_WINDOWEVENT_RESIZED:
			internal.Window.Width = float32(e.data1)
			internal.Window.Height = float32(e.data2)
			gfx.Viewport(geom.Vec2{X: 0, Y: 0}, geom.Vec2{X: float32(e.data1), Y: float32(e.data2)})
			window.Handler.WindowResized(
				geom.Vec2{X: float32(e.data1), Y: float32(e.data2)},
				ts,
			)
		case C.SDL_WINDOWEVENT_SIZE_CHANGED:
			//TODO
		case C.SDL_WINDOWEVENT_MINIMIZED:
			window.Handler.WindowMinimized(ts)
		case C.SDL_WINDOWEVENT_MAXIMIZED:
			window.Handler.WindowMaximized(ts)
		case C.SDL_WINDOWEVENT_RESTORED:
			window.Handler.WindowRestored(ts)
		case C.SDL_WINDOWEVENT_ENTER:
			internal.HasMouseFocus = true
			window.Handler.WindowMouseEnter(ts)
		case C.SDL_WINDOWEVENT_LEAVE:
			internal.HasMouseFocus = false
			window.Handler.WindowMouseLeave(ts)
		case C.SDL_WINDOWEVENT_FOCUS_GAINED:
			internal.HasFocus = true
			window.Handler.WindowFocusGained(ts)
		case C.SDL_WINDOWEVENT_FOCUS_LOST:
			internal.HasFocus = false
			window.Handler.WindowFocusLost(ts)
		case C.SDL_WINDOWEVENT_CLOSE:
			// Ignore
		default:
			//TODO: log.Print("unkown window event")
		}
	// Keyboard Events
	case C.SDL_KEYDOWN:
		e := (*C.SDL_KeyboardEvent)(e)
		if e.repeat == 0 {
			internal.KeyState[e.keysym.scancode] = true
			key.Handler.KeyDown(
				key.Label(e.keysym.sym),
				key.Position(e.keysym.scancode),
				ts,
			)
		}
	case C.SDL_KEYUP:
		e := (*C.SDL_KeyboardEvent)(e)
		internal.KeyState[e.keysym.scancode] = false
		key.Handler.KeyUp(
			key.Label(e.keysym.sym),
			key.Position(e.keysym.scancode),
			ts,
		)
	// Mouse Events
	case C.SDL_MOUSEMOTION:
		e := (*C.SDL_MouseMotionEvent)(e)
		rel := geom.Vec2{X: float32(e.xrel), Y: float32(e.yrel)}
		internal.MouseDelta = internal.MouseDelta.Plus(rel)
		internal.MousePosition = geom.Vec2{X: float32(e.x), Y: float32(e.y)}
		internal.MouseButtons = uint32(e.state)
		mouse.Handler.MouseMotion(
			rel,
			internal.MousePosition,
			ts,
		)
	case C.SDL_MOUSEBUTTONDOWN:
		e := (*C.SDL_MouseButtonEvent)(e)
		mouse.Handler.MouseButtonDown(
			mouse.Button(e.button),
			int(e.clicks),
			ts,
		)
	case C.SDL_MOUSEBUTTONUP:
		e := (*C.SDL_MouseButtonEvent)(e)
		mouse.Handler.MouseButtonUp(
			mouse.Button(e.button),
			int(e.clicks),
			ts,
		)
	case C.SDL_MOUSEWHEEL:
		e := (*C.SDL_MouseWheelEvent)(e)
		var d float32 = 1
		if e.direction == C.SDL_MOUSEWHEEL_FLIPPED {
			d = -1
		}
		mouse.Handler.MouseWheel(
			geom.Vec2{X: float32(e.x) * d, Y: float32(e.y) * d},
			ts,
		)
	//TODO: Joystick Events
	case C.SDL_JOYAXISMOTION:
	case C.SDL_JOYBALLMOTION:
	case C.SDL_JOYHATMOTION:
	case C.SDL_JOYBUTTONDOWN:
	case C.SDL_JOYBUTTONUP:
	case C.SDL_JOYDEVICEADDED:
	case C.SDL_JOYDEVICEREMOVED:
	//TODO: Controller Events
	case C.SDL_CONTROLLERAXISMOTION:
	case C.SDL_CONTROLLERBUTTONDOWN:
	case C.SDL_CONTROLLERBUTTONUP:
	case C.SDL_CONTROLLERDEVICEADDED:
	case C.SDL_CONTROLLERDEVICEREMOVED:
	case C.SDL_CONTROLLERDEVICEREMAPPED:
	//TODO: Audio Device Events
	case C.SDL_AUDIODEVICEADDED:
	case C.SDL_AUDIODEVICEREMOVED:
	default:
		//TODO: log.Print("unknown SDL event:", ((*C.SDL_CommonEvent)(e))._type)
	}
}

// PeepEvents fill the event buffer and returns the number of events fetched.
func PeepEvents() int {
	return int(C.PeepEvents())
}

// EventAt returns a pointer to an event in the event buffer.
func EventAt(i int) unsafe.Pointer {
	return unsafe.Pointer(&C.Events[i])
}

//------------------------------------------------------------------------------
