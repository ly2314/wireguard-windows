// Copyright 2012 The Walk Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package declarative

import (
	"github.com/lxn/walk"
)

type ToolButton struct {
	AssignTo         **walk.ToolButton
	Name             string
	Enabled          Property
	Visible          Property
	Font             Font
	ToolTipText      Property
	MinSize          Size
	MaxSize          Size
	StretchFactor    int
	Row              int
	RowSpan          int
	Column           int
	ColumnSpan       int
	ContextMenuItems []MenuItem
	OnKeyDown        walk.KeyEventHandler
	OnMouseDown      walk.MouseEventHandler
	OnMouseMove      walk.MouseEventHandler
	OnMouseUp        walk.MouseEventHandler
	OnSizeChanged    walk.EventHandler
	Text             Property
	OnClicked        walk.EventHandler
}

func (tb ToolButton) Create(builder *Builder) error {
	w, err := walk.NewToolButton(builder.Parent())
	if err != nil {
		return err
	}

	return builder.InitWidget(tb, w, func() error {
		if tb.OnClicked != nil {
			w.Clicked().Attach(tb.OnClicked)
		}

		if tb.AssignTo != nil {
			*tb.AssignTo = w
		}

		return nil
	})
}

func (w ToolButton) WidgetInfo() (name string, disabled, hidden bool, font *Font, toolTipText string, minSize, maxSize Size, stretchFactor, row, rowSpan, column, columnSpan int, contextMenuItems []MenuItem, OnKeyDown walk.KeyEventHandler, OnMouseDown walk.MouseEventHandler, OnMouseMove walk.MouseEventHandler, OnMouseUp walk.MouseEventHandler, OnSizeChanged walk.EventHandler) {
	return w.Name, false, false, &w.Font, "", w.MinSize, w.MaxSize, w.StretchFactor, w.Row, w.RowSpan, w.Column, w.ColumnSpan, w.ContextMenuItems, w.OnKeyDown, w.OnMouseDown, w.OnMouseMove, w.OnMouseUp, w.OnSizeChanged
}