// Code generated by girgen. DO NOT EDIT.

package gnomedesktop

import (
	"unsafe"

	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
import "C"

//export _gotk4_gnomedesktop3_BGCrossfade_ConnectFinished
func _gotk4_gnomedesktop3_BGCrossfade_ConnectFinished(arg0 C.gpointer, arg1 *C.GObject, arg2 C.guintptr) {
	var f func(window *coreglib.Object)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg2))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(window *coreglib.Object))
	}

	var _window *coreglib.Object // out

	_window = coreglib.Take(unsafe.Pointer(arg1))

	f(_window)
}