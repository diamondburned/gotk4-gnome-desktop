// Code generated by girgen. DO NOT EDIT.

package gnomedesktop

import (
	"unsafe"

	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
import "C"

//export _gotk4_gnomedesktop3_RRScreen_ConnectOutputConnected
func _gotk4_gnomedesktop3_RRScreen_ConnectOutputConnected(arg0 C.gpointer, arg1 C.gpointer, arg2 C.guintptr) {
	var f func(output unsafe.Pointer)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg2))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(output unsafe.Pointer))
	}

	var _output unsafe.Pointer // out

	_output = (unsafe.Pointer)(unsafe.Pointer(arg1))

	f(_output)
}

//export _gotk4_gnomedesktop3_RRScreen_ConnectOutputDisconnected
func _gotk4_gnomedesktop3_RRScreen_ConnectOutputDisconnected(arg0 C.gpointer, arg1 C.gpointer, arg2 C.guintptr) {
	var f func(output unsafe.Pointer)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg2))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(output unsafe.Pointer))
	}

	var _output unsafe.Pointer // out

	_output = (unsafe.Pointer)(unsafe.Pointer(arg1))

	f(_output)
}