// Code generated by girgen. DO NOT EDIT.

package gnomedesktop

import (
	"fmt"
	_ "runtime/cgo"
	"unsafe"

	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gnome-desktop-3.0 gnome-desktop-4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdlib.h>
// #include <glib-object.h>
import "C"

// GType values.
var (
	GTypeRRDpmsModeType = coreglib.Type(C.gnome_rr_dpms_mode_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeRRDpmsModeType, F: marshalRRDpmsModeType},
	})
}

type RRDpmsModeType C.gint

const (
	RrDpmsOn RRDpmsModeType = iota
	RrDpmsStandby
	RrDpmsSuspend
	RrDpmsOff
	RrDpmsUnknown
)

func marshalRRDpmsModeType(p uintptr) (interface{}, error) {
	return RRDpmsModeType(coreglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for RRDpmsModeType.
func (r RRDpmsModeType) String() string {
	switch r {
	case RrDpmsOn:
		return "On"
	case RrDpmsStandby:
		return "Standby"
	case RrDpmsSuspend:
		return "Suspend"
	case RrDpmsOff:
		return "Off"
	case RrDpmsUnknown:
		return "Unknown"
	default:
		return fmt.Sprintf("RRDpmsModeType(%d)", r)
	}
}
