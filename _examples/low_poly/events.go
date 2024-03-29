// Copyright (c) 2013-2016 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package main

//------------------------------------------------------------------------------

import (
	"github.com/drakmaniso/carol/key"
	"github.com/drakmaniso/carol/mouse"
	"github.com/drakmaniso/carol/plane"
	"github.com/drakmaniso/carol/space"
)

//------------------------------------------------------------------------------

func (l loop) WindowResized(_, _ int32) {
	camera.WindowResized()
}

//------------------------------------------------------------------------------

func (l loop) MouseWheel(x, y int32) {
	camera.ChangeDistance(float32(-y))
}

func (l loop) MouseButtonDown(b mouse.Button, _ int) {
	switch b {
	case mouse.Left:
		dragStart = misc.worldFromObject
		current.dragDelta = plane.Coord{0, 0}
		mouse.SetRelativeMode(true)
	case mouse.Extra1:
		camera.SetFocus(space.Coord{0, 0, 0})
		camera.SetDistance(4)
		camera.SetOrientation(0, 0, 0)
	case mouse.Extra2:
		misc.worldFromObject = space.Identity()
	default:
		mouse.SetRelativeMode(true)
	}
}

func (l loop) MouseButtonUp(b mouse.Button, _ int) {
	mouse.SetRelativeMode(false)
}

//------------------------------------------------------------------------------

func (l loop) KeyDown(lb key.Label, p key.Position) {
	const s = 2.0
	switch p {
	case key.PositionW:
		forward = -s
	case key.PositionS:
		forward = s
	case key.PositionA:
		lateral = -s
	case key.PositionD:
		lateral = s
	case key.PositionSpace:
		vertical = s
	case key.PositionLShift:
		vertical = -s
	case key.PositionQ:
		rolling = -1.0
	case key.PositionE:
		rolling = 1.0
	default:
		l.Handlers.KeyDown(lb, p)
	}
}

func (l loop) KeyUp(_ key.Label, p key.Position) {
	const s = 5.0
	switch p {
	case key.PositionW, key.PositionS:
		forward = 0.0
	case key.PositionA, key.PositionD:
		lateral = 0.0
	case key.PositionSpace, key.PositionLShift:
		vertical = 0.0
	case key.PositionQ, key.PositionE:
		rolling = 0.0
	}
}

//------------------------------------------------------------------------------
