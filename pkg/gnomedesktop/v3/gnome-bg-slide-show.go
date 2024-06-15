// Code generated by girgen. DO NOT EDIT.

package gnomedesktop

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gerror"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
import "C"

// GType values.
var (
	GTypeBGSlideShow = coreglib.Type(C.gnome_bg_slide_show_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeBGSlideShow, F: marshalBGSlideShow},
	})
}

// BGSlideShowOverrides contains methods that are overridable.
type BGSlideShowOverrides struct {
}

func defaultBGSlideShowOverrides(v *BGSlideShow) BGSlideShowOverrides {
	return BGSlideShowOverrides{}
}

type BGSlideShow struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*BGSlideShow)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*BGSlideShow, *BGSlideShowClass, BGSlideShowOverrides](
		GTypeBGSlideShow,
		initBGSlideShowClass,
		wrapBGSlideShow,
		defaultBGSlideShowOverrides,
	)
}

func initBGSlideShowClass(gclass unsafe.Pointer, overrides BGSlideShowOverrides, classInitFunc func(*BGSlideShowClass)) {
	if classInitFunc != nil {
		class := (*BGSlideShowClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapBGSlideShow(obj *coreglib.Object) *BGSlideShow {
	return &BGSlideShow{
		Object: obj,
	}
}

func marshalBGSlideShow(p uintptr) (interface{}, error) {
	return wrapBGSlideShow(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewBGSlideShow creates a new object to manage a slide show. window background
// between two #cairo_surface_ts.
//
// The function takes the following parameters:
//
//   - filename of the slide show.
//
// The function returns the following values:
//
//   - bgSlideShow: new BGSlideShow.
//
func NewBGSlideShow(filename string) *BGSlideShow {
	var _arg1 *C.char             // out
	var _cret *C.GnomeBGSlideShow // in

	_arg1 = (*C.char)(unsafe.Pointer(C.CString(filename)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gnome_bg_slide_show_new(_arg1)
	runtime.KeepAlive(filename)

	var _bgSlideShow *BGSlideShow // out

	_bgSlideShow = wrapBGSlideShow(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _bgSlideShow
}

// CurrentSlide returns the current slides progress.
//
// The function takes the following parameters:
//
//   - width: monitor width.
//   - height: monitor height.
//
// The function returns the following values:
//
//   - progress (optional): slide progress.
//   - duration (optional): slide duration.
//   - isFixed (optional): if slide is fixed.
//   - file1 (optional): first file in slide.
//   - file2 (optional): second file in slide.
//
func (self *BGSlideShow) CurrentSlide(width, height int) (progress, duration float64, isFixed bool, file1, file2 string) {
	var _arg0 *C.GnomeBGSlideShow // out
	var _arg1 C.int               // out
	var _arg2 C.int               // out
	var _arg3 C.gdouble           // in
	var _arg4 C.double            // in
	var _arg5 C.gboolean          // in
	var _arg6 *C.char             // in
	var _arg7 *C.char             // in

	_arg0 = (*C.GnomeBGSlideShow)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = C.int(width)
	_arg2 = C.int(height)

	C.gnome_bg_slide_show_get_current_slide(_arg0, _arg1, _arg2, &_arg3, &_arg4, &_arg5, &_arg6, &_arg7)
	runtime.KeepAlive(self)
	runtime.KeepAlive(width)
	runtime.KeepAlive(height)

	var _progress float64 // out
	var _duration float64 // out
	var _isFixed bool     // out
	var _file1 string     // out
	var _file2 string     // out

	_progress = float64(_arg3)
	_duration = float64(_arg4)
	if _arg5 != 0 {
		_isFixed = true
	}
	if _arg6 != nil {
		_file1 = C.GoString((*C.gchar)(unsafe.Pointer(_arg6)))
	}
	if _arg7 != nil {
		_file2 = C.GoString((*C.gchar)(unsafe.Pointer(_arg7)))
	}

	return _progress, _duration, _isFixed, _file1, _file2
}

// HasMultipleSizes gets whether or not the slide show has multiple sizes for
// different monitors.
//
// The function returns the following values:
//
//   - ok: TRUE if multiple sizes.
//
func (self *BGSlideShow) HasMultipleSizes() bool {
	var _arg0 *C.GnomeBGSlideShow // out
	var _cret C.gboolean          // in

	_arg0 = (*C.GnomeBGSlideShow)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.gnome_bg_slide_show_get_has_multiple_sizes(_arg0)
	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// NumSlides returns number of slides in slide show.
//
// The function returns the following values:
//
func (self *BGSlideShow) NumSlides() int {
	var _arg0 *C.GnomeBGSlideShow // out
	var _cret C.int               // in

	_arg0 = (*C.GnomeBGSlideShow)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.gnome_bg_slide_show_get_num_slides(_arg0)
	runtime.KeepAlive(self)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// Slide retrieves slide by frame number.
//
// The function takes the following parameters:
//
//   - frameNumber: frame number.
//   - width: monitor width.
//   - height: monitor height.
//
// The function returns the following values:
//
//   - progress (optional): slide progress.
//   - duration (optional): slide duration.
//   - isFixed (optional): if slide is fixed.
//   - file1 (optional): first file in slide.
//   - file2 (optional): second file in slide.
//   - ok: TRUE if successful.
//
func (self *BGSlideShow) Slide(frameNumber, width, height int) (progress, duration float64, isFixed bool, file1, file2 string, ok bool) {
	var _arg0 *C.GnomeBGSlideShow // out
	var _arg1 C.int               // out
	var _arg2 C.int               // out
	var _arg3 C.int               // out
	var _arg4 C.gdouble           // in
	var _arg5 C.double            // in
	var _arg6 C.gboolean          // in
	var _arg7 *C.char             // in
	var _arg8 *C.char             // in
	var _cret C.gboolean          // in

	_arg0 = (*C.GnomeBGSlideShow)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = C.int(frameNumber)
	_arg2 = C.int(width)
	_arg3 = C.int(height)

	_cret = C.gnome_bg_slide_show_get_slide(_arg0, _arg1, _arg2, _arg3, &_arg4, &_arg5, &_arg6, &_arg7, &_arg8)
	runtime.KeepAlive(self)
	runtime.KeepAlive(frameNumber)
	runtime.KeepAlive(width)
	runtime.KeepAlive(height)

	var _progress float64 // out
	var _duration float64 // out
	var _isFixed bool     // out
	var _file1 string     // out
	var _file2 string     // out
	var _ok bool          // out

	_progress = float64(_arg4)
	_duration = float64(_arg5)
	if _arg6 != 0 {
		_isFixed = true
	}
	if _arg7 != nil {
		_file1 = C.GoString((*C.gchar)(unsafe.Pointer(_arg7)))
	}
	if _arg8 != nil {
		_file2 = C.GoString((*C.gchar)(unsafe.Pointer(_arg8)))
	}
	if _cret != 0 {
		_ok = true
	}

	return _progress, _duration, _isFixed, _file1, _file2, _ok
}

// StartTime gets the start time of the slide show.
//
// The function returns the following values:
//
//   - gdouble: timestamp.
//
func (self *BGSlideShow) StartTime() float64 {
	var _arg0 *C.GnomeBGSlideShow // out
	var _cret C.double            // in

	_arg0 = (*C.GnomeBGSlideShow)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.gnome_bg_slide_show_get_start_time(_arg0)
	runtime.KeepAlive(self)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

// TotalDuration gets the total duration of the slide show.
//
// The function returns the following values:
//
//   - gdouble: timestamp.
//
func (self *BGSlideShow) TotalDuration() float64 {
	var _arg0 *C.GnomeBGSlideShow // out
	var _cret C.double            // in

	_arg0 = (*C.GnomeBGSlideShow)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.gnome_bg_slide_show_get_total_duration(_arg0)
	runtime.KeepAlive(self)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

// Load tries to load the slide show.
func (self *BGSlideShow) Load() error {
	var _arg0 *C.GnomeBGSlideShow // out
	var _cerr *C.GError           // in

	_arg0 = (*C.GnomeBGSlideShow)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	C.gnome_bg_slide_show_load(_arg0, &_cerr)
	runtime.KeepAlive(self)

	var _goerr error // out

	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _goerr
}

// BGSlideShowClass: instance of this type is always passed by reference.
type BGSlideShowClass struct {
	*bgSlideShowClass
}

// bgSlideShowClass is the struct that's finalized.
type bgSlideShowClass struct {
	native *C.GnomeBGSlideShowClass
}
