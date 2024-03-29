// Copyright (c) 2013-2017 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package internal

//------------------------------------------------------------------------------

/*
#include "sdl.h"
*/
import "C"

//------------------------------------------------------------------------------

// A KeyLabel designate a key by its label in the current layout of the keyboard.
// For printable characters, the value is the rune that would be generated by
// pressing the key without any modifiers.
type KeyLabel rune

// A KeyPosition designate a key by its physical position on the keyboard.
// It is not affected by the layout or any other language settings.
type KeyPosition uint32

//------------------------------------------------------------------------------

// KeyLabelOf returns the key label at the specified position in the current
// layout.
func KeyLabelOf(pos KeyPosition) KeyLabel {
	return KeyLabel(C.SDL_GetKeyFromScancode(C.SDL_Scancode(pos)))
}

// KeySearchPositionOf searches the current position of label in the current
// layout.
func KeySearchPositionOf(l KeyLabel) KeyPosition {
	return KeyPosition(C.SDL_GetScancodeFromKey(C.SDL_Keycode(l)))
}

//------------------------------------------------------------------------------
