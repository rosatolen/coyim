package gtki

import "github.com/twstrike/coyim/Godeps/_workspace/src/github.com/twstrike/gotk3adapter/gdki"
import "github.com/twstrike/coyim/Godeps/_workspace/src/github.com/twstrike/gotk3adapter/glibi"

type Widget interface {
	glibi.Object

	Destroy()
	GetWindow() (gdki.Window, error)
	GrabFocus()
	GetAllocatedHeight() int
	GetAllocatedWidth() int
	GetStyleContext() (StyleContext, error)
	GrabDefault()
	HasFocus() bool
	Hide()
	HideOnDelete()
	SetHAlign(Align)
	SetHExpand(bool)
	SetMarginBottom(int)
	SetMarginTop(int)
	SetName(string)
	SetSensitive(bool)
	SetSizeRequest(int, int)
	SetVisible(bool)
	Show()
	ShowAll()
}

func AssertWidget(_ Widget) {}
