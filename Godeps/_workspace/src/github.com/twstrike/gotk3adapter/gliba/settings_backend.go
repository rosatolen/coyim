package gliba

import (
	"github.com/twstrike/coyim/Godeps/_workspace/src/github.com/gotk3/gotk3/glib"
	"github.com/twstrike/coyim/Godeps/_workspace/src/github.com/twstrike/gotk3adapter/glibi"
)

type settingsBackend struct {
	*Object
	*glib.SettingsBackend
}

func wrapSettingsBackendSimple(v *glib.SettingsBackend) *settingsBackend {
	if v == nil {
		return nil
	}
	return &settingsBackend{WrapObjectSimple(v.Object), v}
}

func unwrapSettingsBackend(v glibi.SettingsBackend) *glib.SettingsBackend {
	if v == nil {
		return nil
	}
	return v.(*settingsBackend).SettingsBackend
}
