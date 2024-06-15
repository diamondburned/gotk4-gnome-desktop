// Code generated by girgen. DO NOT EDIT.

package gnomedesktop

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/cairo"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v3"
	"github.com/diamondburned/gotk4/pkg/gdkpixbuf/v2"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
)

// #include <stdlib.h>
// #include <glib-object.h>
// extern void _gotk4_gnomedesktop3_BG_ConnectTransitioned(gpointer, guintptr);
// extern void _gotk4_gnomedesktop3_BG_ConnectChanged(gpointer, guintptr);
import "C"

// GType values.
var (
	GTypeBG = coreglib.Type(C.gnome_bg_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeBG, F: marshalBG},
	})
}

type BG struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*BG)(nil)
)

func wrapBG(obj *coreglib.Object) *BG {
	return &BG{
		Object: obj,
	}
}

func marshalBG(p uintptr) (interface{}, error) {
	return wrapBG(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

func (bg *BG) ConnectChanged(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(bg, "changed", false, unsafe.Pointer(C._gotk4_gnomedesktop3_BG_ConnectChanged), f)
}

func (bg *BG) ConnectTransitioned(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(bg, "transitioned", false, unsafe.Pointer(C._gotk4_gnomedesktop3_BG_ConnectTransitioned), f)
}

// The function returns the following values:
//
func NewBG() *BG {
	var _cret *C.GnomeBG // in

	_cret = C.gnome_bg_new()

	var _bG *BG // out

	_bG = wrapBG(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _bG
}

// The function returns the following values:
//
func (bg *BG) ChangesWithTime() bool {
	var _arg0 *C.GnomeBG // out
	var _cret C.gboolean // in

	_arg0 = (*C.GnomeBG)(unsafe.Pointer(coreglib.InternObject(bg).Native()))

	_cret = C.gnome_bg_changes_with_time(_arg0)
	runtime.KeepAlive(bg)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// CreateFrameThumbnail creates a thumbnail for a certain frame, where 'frame'
// is somewhat vaguely defined as 'suitable point to show while single-stepping
// through the slideshow'.
//
// The function takes the following parameters:
//
//   - factory
//   - screen
//   - destWidth
//   - destHeight
//   - frameNum
//
// The function returns the following values:
//
//   - pixbuf: newly created thumbnail or or NULL if frame_num is out of bounds.
//
func (bg *BG) CreateFrameThumbnail(factory *DesktopThumbnailFactory, screen *gdk.Screen, destWidth, destHeight, frameNum int) *gdkpixbuf.Pixbuf {
	var _arg0 *C.GnomeBG                      // out
	var _arg1 *C.GnomeDesktopThumbnailFactory // out
	var _arg2 *C.GdkScreen                    // out
	var _arg3 C.int                           // out
	var _arg4 C.int                           // out
	var _arg5 C.int                           // out
	var _cret *C.GdkPixbuf                    // in

	_arg0 = (*C.GnomeBG)(unsafe.Pointer(coreglib.InternObject(bg).Native()))
	_arg1 = (*C.GnomeDesktopThumbnailFactory)(unsafe.Pointer(coreglib.InternObject(factory).Native()))
	_arg2 = (*C.GdkScreen)(unsafe.Pointer(coreglib.InternObject(screen).Native()))
	_arg3 = C.int(destWidth)
	_arg4 = C.int(destHeight)
	_arg5 = C.int(frameNum)

	_cret = C.gnome_bg_create_frame_thumbnail(_arg0, _arg1, _arg2, _arg3, _arg4, _arg5)
	runtime.KeepAlive(bg)
	runtime.KeepAlive(factory)
	runtime.KeepAlive(screen)
	runtime.KeepAlive(destWidth)
	runtime.KeepAlive(destHeight)
	runtime.KeepAlive(frameNum)

	var _pixbuf *gdkpixbuf.Pixbuf // out

	{
		obj := coreglib.AssumeOwnership(unsafe.Pointer(_cret))
		_pixbuf = &gdkpixbuf.Pixbuf{
			Object: obj,
			LoadableIcon: gio.LoadableIcon{
				Icon: gio.Icon{
					Object: obj,
				},
			},
		}
	}

	return _pixbuf
}

// CreateSurface: create a surface that can be set as background for window.
//
// The function takes the following parameters:
//
//   - window
//   - width
//   - height
//
// The function returns the following values:
//
//   - surface: NULL on error (e.g. out of X connections).
//
func (bg *BG) CreateSurface(window gdk.Windower, width, height int) *cairo.Surface {
	var _arg0 *C.GnomeBG         // out
	var _arg1 *C.GdkWindow       // out
	var _arg2 C.int              // out
	var _arg3 C.int              // out
	var _cret *C.cairo_surface_t // in

	_arg0 = (*C.GnomeBG)(unsafe.Pointer(coreglib.InternObject(bg).Native()))
	_arg1 = (*C.GdkWindow)(unsafe.Pointer(coreglib.InternObject(window).Native()))
	_arg2 = C.int(width)
	_arg3 = C.int(height)

	_cret = C.gnome_bg_create_surface(_arg0, _arg1, _arg2, _arg3)
	runtime.KeepAlive(bg)
	runtime.KeepAlive(window)
	runtime.KeepAlive(width)
	runtime.KeepAlive(height)

	var _surface *cairo.Surface // out

	_surface = cairo.WrapSurface(uintptr(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(_surface, func(v *cairo.Surface) {
		C.cairo_surface_destroy((*C.cairo_surface_t)(unsafe.Pointer(v.Native())))
	})

	return _surface
}

// The function takes the following parameters:
//
//   - factory
//   - screen
//   - destWidth
//   - destHeight
//
// The function returns the following values:
//
//   - pixbuf showing the background as a thumbnail.
//
func (bg *BG) CreateThumbnail(factory *DesktopThumbnailFactory, screen *gdk.Screen, destWidth, destHeight int) *gdkpixbuf.Pixbuf {
	var _arg0 *C.GnomeBG                      // out
	var _arg1 *C.GnomeDesktopThumbnailFactory // out
	var _arg2 *C.GdkScreen                    // out
	var _arg3 C.int                           // out
	var _arg4 C.int                           // out
	var _cret *C.GdkPixbuf                    // in

	_arg0 = (*C.GnomeBG)(unsafe.Pointer(coreglib.InternObject(bg).Native()))
	_arg1 = (*C.GnomeDesktopThumbnailFactory)(unsafe.Pointer(coreglib.InternObject(factory).Native()))
	_arg2 = (*C.GdkScreen)(unsafe.Pointer(coreglib.InternObject(screen).Native()))
	_arg3 = C.int(destWidth)
	_arg4 = C.int(destHeight)

	_cret = C.gnome_bg_create_thumbnail(_arg0, _arg1, _arg2, _arg3, _arg4)
	runtime.KeepAlive(bg)
	runtime.KeepAlive(factory)
	runtime.KeepAlive(screen)
	runtime.KeepAlive(destWidth)
	runtime.KeepAlive(destHeight)

	var _pixbuf *gdkpixbuf.Pixbuf // out

	{
		obj := coreglib.AssumeOwnership(unsafe.Pointer(_cret))
		_pixbuf = &gdkpixbuf.Pixbuf{
			Object: obj,
			LoadableIcon: gio.LoadableIcon{
				Icon: gio.Icon{
					Object: obj,
				},
			},
		}
	}

	return _pixbuf
}

// The function takes the following parameters:
//
func (bg *BG) Draw(dest *gdkpixbuf.Pixbuf) {
	var _arg0 *C.GnomeBG   // out
	var _arg1 *C.GdkPixbuf // out

	_arg0 = (*C.GnomeBG)(unsafe.Pointer(coreglib.InternObject(bg).Native()))
	_arg1 = (*C.GdkPixbuf)(unsafe.Pointer(coreglib.InternObject(dest).Native()))

	C.gnome_bg_draw(_arg0, _arg1)
	runtime.KeepAlive(bg)
	runtime.KeepAlive(dest)
}

// The function returns the following values:
//
func (bg *BG) Filename() string {
	var _arg0 *C.GnomeBG // out
	var _cret *C.gchar   // in

	_arg0 = (*C.GnomeBG)(unsafe.Pointer(coreglib.InternObject(bg).Native()))

	_cret = C.gnome_bg_get_filename(_arg0)
	runtime.KeepAlive(bg)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// The function takes the following parameters:
//
//   - factory
//   - bestWidth
//   - bestHeight
//   - width
//   - height
//
// The function returns the following values:
//
func (bg *BG) ImageSize(factory *DesktopThumbnailFactory, bestWidth, bestHeight int, width, height *int) bool {
	var _arg0 *C.GnomeBG                      // out
	var _arg1 *C.GnomeDesktopThumbnailFactory // out
	var _arg2 C.int                           // out
	var _arg3 C.int                           // out
	var _arg4 *C.int                          // out
	var _arg5 *C.int                          // out
	var _cret C.gboolean                      // in

	_arg0 = (*C.GnomeBG)(unsafe.Pointer(coreglib.InternObject(bg).Native()))
	_arg1 = (*C.GnomeDesktopThumbnailFactory)(unsafe.Pointer(coreglib.InternObject(factory).Native()))
	_arg2 = C.int(bestWidth)
	_arg3 = C.int(bestHeight)
	_arg4 = (*C.int)(unsafe.Pointer(width))
	_arg5 = (*C.int)(unsafe.Pointer(height))

	_cret = C.gnome_bg_get_image_size(_arg0, _arg1, _arg2, _arg3, _arg4, _arg5)
	runtime.KeepAlive(bg)
	runtime.KeepAlive(factory)
	runtime.KeepAlive(bestWidth)
	runtime.KeepAlive(bestHeight)
	runtime.KeepAlive(width)
	runtime.KeepAlive(height)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// The function returns the following values:
//
func (bg *BG) HasMultipleSizes() bool {
	var _arg0 *C.GnomeBG // out
	var _cret C.gboolean // in

	_arg0 = (*C.GnomeBG)(unsafe.Pointer(coreglib.InternObject(bg).Native()))

	_cret = C.gnome_bg_has_multiple_sizes(_arg0)
	runtime.KeepAlive(bg)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// The function takes the following parameters:
//
//   - destWidth
//   - destHeight
//
// The function returns the following values:
//
func (bg *BG) IsDark(destWidth, destHeight int) bool {
	var _arg0 *C.GnomeBG // out
	var _arg1 C.int      // out
	var _arg2 C.int      // out
	var _cret C.gboolean // in

	_arg0 = (*C.GnomeBG)(unsafe.Pointer(coreglib.InternObject(bg).Native()))
	_arg1 = C.int(destWidth)
	_arg2 = C.int(destHeight)

	_cret = C.gnome_bg_is_dark(_arg0, _arg1, _arg2)
	runtime.KeepAlive(bg)
	runtime.KeepAlive(destWidth)
	runtime.KeepAlive(destHeight)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// The function takes the following parameters:
//
func (bg *BG) LoadFromPreferences(settings *gio.Settings) {
	var _arg0 *C.GnomeBG   // out
	var _arg1 *C.GSettings // out

	_arg0 = (*C.GnomeBG)(unsafe.Pointer(coreglib.InternObject(bg).Native()))
	_arg1 = (*C.GSettings)(unsafe.Pointer(coreglib.InternObject(settings).Native()))

	C.gnome_bg_load_from_preferences(_arg0, _arg1)
	runtime.KeepAlive(bg)
	runtime.KeepAlive(settings)
}

// The function takes the following parameters:
//
func (bg *BG) SaveToPreferences(settings *gio.Settings) {
	var _arg0 *C.GnomeBG   // out
	var _arg1 *C.GSettings // out

	_arg0 = (*C.GnomeBG)(unsafe.Pointer(coreglib.InternObject(bg).Native()))
	_arg1 = (*C.GSettings)(unsafe.Pointer(coreglib.InternObject(settings).Native()))

	C.gnome_bg_save_to_preferences(_arg0, _arg1)
	runtime.KeepAlive(bg)
	runtime.KeepAlive(settings)
}

// The function takes the following parameters:
//
func (bg *BG) SetFilename(filename string) {
	var _arg0 *C.GnomeBG // out
	var _arg1 *C.char    // out

	_arg0 = (*C.GnomeBG)(unsafe.Pointer(coreglib.InternObject(bg).Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(filename)))
	defer C.free(unsafe.Pointer(_arg1))

	C.gnome_bg_set_filename(_arg0, _arg1)
	runtime.KeepAlive(bg)
	runtime.KeepAlive(filename)
}
