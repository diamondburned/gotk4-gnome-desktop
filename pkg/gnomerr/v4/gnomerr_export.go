// Code generated by girgen. DO NOT EDIT.

package gnomerr

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
import "C"

//export _gotk4_gnomerr4_ScreenClass_changed
func _gotk4_gnomerr4_ScreenClass_changed(arg0 *C.GnomeRRScreen) {
	instance0 := coreglib.Take(unsafe.Pointer(arg0))
	overrides := coreglib.OverridesFromObj[ScreenOverrides](instance0)
	if overrides.Changed == nil {
		panic("gotk4: " + instance0.TypeFromInstance().String() + ": expected ScreenOverrides.Changed, got none")
	}

	overrides.Changed()
}

//export _gotk4_gnomerr4_ScreenClass_output_connected
func _gotk4_gnomerr4_ScreenClass_output_connected(arg0 *C.GnomeRRScreen, arg1 *C.GnomeRROutput) {
	instance0 := coreglib.Take(unsafe.Pointer(arg0))
	overrides := coreglib.OverridesFromObj[ScreenOverrides](instance0)
	if overrides.OutputConnected == nil {
		panic("gotk4: " + instance0.TypeFromInstance().String() + ": expected ScreenOverrides.OutputConnected, got none")
	}

	var _output *Output // out

	_output = (*Output)(gextras.NewStructNative(unsafe.Pointer(arg1)))

	overrides.OutputConnected(_output)
}

//export _gotk4_gnomerr4_ScreenClass_output_disconnected
func _gotk4_gnomerr4_ScreenClass_output_disconnected(arg0 *C.GnomeRRScreen, arg1 *C.GnomeRROutput) {
	instance0 := coreglib.Take(unsafe.Pointer(arg0))
	overrides := coreglib.OverridesFromObj[ScreenOverrides](instance0)
	if overrides.OutputDisconnected == nil {
		panic("gotk4: " + instance0.TypeFromInstance().String() + ": expected ScreenOverrides.OutputDisconnected, got none")
	}

	var _output *Output // out

	_output = (*Output)(gextras.NewStructNative(unsafe.Pointer(arg1)))

	overrides.OutputDisconnected(_output)
}

//export _gotk4_gnomerr4_Screen_ConnectChanged
func _gotk4_gnomerr4_Screen_ConnectChanged(arg0 C.gpointer, arg1 C.guintptr) {
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
