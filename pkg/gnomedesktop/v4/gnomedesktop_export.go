// Code generated by girgen. DO NOT EDIT.

package gnomedesktop

import (
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
import "C"

//export _gotk4_gnomedesktop4_XkbInfo_ConnectLayoutsChanged
func _gotk4_gnomedesktop4_XkbInfo_ConnectLayoutsChanged(arg0 C.gpointer, arg1 C.guintptr) {
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
