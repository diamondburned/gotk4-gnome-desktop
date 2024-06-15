// Code generated by girgen. DO NOT EDIT.

package gnomedesktop

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v3"
)

// #include <stdlib.h>
import "C"

//export _gotk4_gnomedesktop3_BG_ConnectChanged
func _gotk4_gnomedesktop3_BG_ConnectChanged(arg0 C.gpointer, arg1 C.guintptr) {
	var f func()
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg1))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func())
	}

	f()
}

//export _gotk4_gnomedesktop3_BG_ConnectTransitioned
func _gotk4_gnomedesktop3_BG_ConnectTransitioned(arg0 C.gpointer, arg1 C.guintptr) {
	var f func()
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg1))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func())
	}

	f()
}

//export _gotk4_gnomedesktop3_BGCrossfadeClass_finished
func _gotk4_gnomedesktop3_BGCrossfadeClass_finished(arg0 *C.GnomeBGCrossfade, arg1 *C.GdkWindow) {
	instance0 := coreglib.Take(unsafe.Pointer(arg0))
	overrides := coreglib.OverridesFromObj[BGCrossfadeOverrides](instance0)
	if overrides.Finished == nil {
		panic("gotk4: " + instance0.TypeFromInstance().String() + ": expected BGCrossfadeOverrides.Finished, got none")
	}

	var _window gdk.Windower // out

	{
		objptr := unsafe.Pointer(arg1)
		if objptr == nil {
			panic("object of type gdk.Windower is nil")
		}

		object := coreglib.Take(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
			_, ok := obj.(gdk.Windower)
			return ok
		})
		rv, ok := casted.(gdk.Windower)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gdk.Windower")
		}
		_window = rv
	}

	overrides.Finished(_window)
}

//export _gotk4_gnomedesktop3_RRScreenClass_changed
func _gotk4_gnomedesktop3_RRScreenClass_changed(arg0 *C.GnomeRRScreen) {
	instance0 := coreglib.Take(unsafe.Pointer(arg0))
	overrides := coreglib.OverridesFromObj[RRScreenOverrides](instance0)
	if overrides.Changed == nil {
		panic("gotk4: " + instance0.TypeFromInstance().String() + ": expected RRScreenOverrides.Changed, got none")
	}

	overrides.Changed()
}

//export _gotk4_gnomedesktop3_RRScreenClass_output_connected
func _gotk4_gnomedesktop3_RRScreenClass_output_connected(arg0 *C.GnomeRRScreen, arg1 *C.GnomeRROutput) {
	instance0 := coreglib.Take(unsafe.Pointer(arg0))
	overrides := coreglib.OverridesFromObj[RRScreenOverrides](instance0)
	if overrides.OutputConnected == nil {
		panic("gotk4: " + instance0.TypeFromInstance().String() + ": expected RRScreenOverrides.OutputConnected, got none")
	}

	var _output *RROutput // out

	_output = (*RROutput)(gextras.NewStructNative(unsafe.Pointer(arg1)))

	overrides.OutputConnected(_output)
}

//export _gotk4_gnomedesktop3_RRScreenClass_output_disconnected
func _gotk4_gnomedesktop3_RRScreenClass_output_disconnected(arg0 *C.GnomeRRScreen, arg1 *C.GnomeRROutput) {
	instance0 := coreglib.Take(unsafe.Pointer(arg0))
	overrides := coreglib.OverridesFromObj[RRScreenOverrides](instance0)
	if overrides.OutputDisconnected == nil {
		panic("gotk4: " + instance0.TypeFromInstance().String() + ": expected RRScreenOverrides.OutputDisconnected, got none")
	}

	var _output *RROutput // out

	_output = (*RROutput)(gextras.NewStructNative(unsafe.Pointer(arg1)))

	overrides.OutputDisconnected(_output)
}

//export _gotk4_gnomedesktop3_RRScreen_ConnectChanged
func _gotk4_gnomedesktop3_RRScreen_ConnectChanged(arg0 C.gpointer, arg1 C.guintptr) {
	var f func()
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg1))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func())
	}

	f()
}

//export _gotk4_gnomedesktop3_XkbInfo_ConnectLayoutsChanged
func _gotk4_gnomedesktop3_XkbInfo_ConnectLayoutsChanged(arg0 C.gpointer, arg1 C.guintptr) {
	var f func()
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg1))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func())
	}

	f()
}