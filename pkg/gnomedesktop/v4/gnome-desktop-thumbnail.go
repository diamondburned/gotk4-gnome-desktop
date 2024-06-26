// Code generated by girgen. DO NOT EDIT.

package gnomedesktop

import (
	"context"
	"fmt"
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gcancel"
	"github.com/diamondburned/gotk4/pkg/core/gerror"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdkpixbuf/v2"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
)

// #include <stdlib.h>
// #include <glib-object.h>
import "C"

// GType values.
var (
	GTypeDesktopThumbnailFactory = coreglib.Type(C.gnome_desktop_thumbnail_factory_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeDesktopThumbnailFactory, F: marshalDesktopThumbnailFactory},
	})
}

type DesktopThumbnailSize C.gint

const (
	DesktopThumbnailSizeNormal DesktopThumbnailSize = iota
	DesktopThumbnailSizeLarge
	DesktopThumbnailSizeXlarge
	DesktopThumbnailSizeXxlarge
)

// String returns the name in string for DesktopThumbnailSize.
func (d DesktopThumbnailSize) String() string {
	switch d {
	case DesktopThumbnailSizeNormal:
		return "Normal"
	case DesktopThumbnailSizeLarge:
		return "Large"
	case DesktopThumbnailSizeXlarge:
		return "Xlarge"
	case DesktopThumbnailSizeXxlarge:
		return "Xxlarge"
	default:
		return fmt.Sprintf("DesktopThumbnailSize(%d)", d)
	}
}

// DesktopThumbnailFactoryOverrides contains methods that are overridable.
type DesktopThumbnailFactoryOverrides struct {
}

func defaultDesktopThumbnailFactoryOverrides(v *DesktopThumbnailFactory) DesktopThumbnailFactoryOverrides {
	return DesktopThumbnailFactoryOverrides{}
}

type DesktopThumbnailFactory struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*DesktopThumbnailFactory)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*DesktopThumbnailFactory, *DesktopThumbnailFactoryClass, DesktopThumbnailFactoryOverrides](
		GTypeDesktopThumbnailFactory,
		initDesktopThumbnailFactoryClass,
		wrapDesktopThumbnailFactory,
		defaultDesktopThumbnailFactoryOverrides,
	)
}

func initDesktopThumbnailFactoryClass(gclass unsafe.Pointer, overrides DesktopThumbnailFactoryOverrides, classInitFunc func(*DesktopThumbnailFactoryClass)) {
	if classInitFunc != nil {
		class := (*DesktopThumbnailFactoryClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapDesktopThumbnailFactory(obj *coreglib.Object) *DesktopThumbnailFactory {
	return &DesktopThumbnailFactory{
		Object: obj,
	}
}

func marshalDesktopThumbnailFactory(p uintptr) (interface{}, error) {
	return wrapDesktopThumbnailFactory(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewDesktopThumbnailFactory creates a new DesktopThumbnailFactory.
//
// This function must be called on the main thread and is non-blocking.
//
// The function takes the following parameters:
//
//   - size: thumbnail size to use.
//
// The function returns the following values:
//
//   - desktopThumbnailFactory: new DesktopThumbnailFactory.
//
func NewDesktopThumbnailFactory(size DesktopThumbnailSize) *DesktopThumbnailFactory {
	var _arg1 C.GnomeDesktopThumbnailSize     // out
	var _cret *C.GnomeDesktopThumbnailFactory // in

	_arg1 = C.GnomeDesktopThumbnailSize(size)

	_cret = C.gnome_desktop_thumbnail_factory_new(_arg1)
	runtime.KeepAlive(size)

	var _desktopThumbnailFactory *DesktopThumbnailFactory // out

	_desktopThumbnailFactory = wrapDesktopThumbnailFactory(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _desktopThumbnailFactory
}

// CanThumbnail returns TRUE if this GnomeDesktopThumbnailFactory can (at least
// try) to thumbnail this file. Thumbnails or files with failed thumbnails won't
// be thumbnailed.
//
// Usage of this function is threadsafe and does blocking I/O.
//
// The function takes the following parameters:
//
//   - uri of a file.
//   - mimeType: mime type of the file.
//   - mtime of the file.
//
// The function returns the following values:
//
//   - ok: TRUE if the file can be thumbnailed.
//
func (factory *DesktopThumbnailFactory) CanThumbnail(uri, mimeType string, mtime int32) bool {
	var _arg0 *C.GnomeDesktopThumbnailFactory // out
	var _arg1 *C.char                         // out
	var _arg2 *C.char                         // out
	var _arg3 C.time_t                        // out
	var _cret C.gboolean                      // in

	_arg0 = (*C.GnomeDesktopThumbnailFactory)(unsafe.Pointer(coreglib.InternObject(factory).Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(uri)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.char)(unsafe.Pointer(C.CString(mimeType)))
	defer C.free(unsafe.Pointer(_arg2))
	_arg3 = C.time_t(mtime)

	_cret = C.gnome_desktop_thumbnail_factory_can_thumbnail(_arg0, _arg1, _arg2, _arg3)
	runtime.KeepAlive(factory)
	runtime.KeepAlive(uri)
	runtime.KeepAlive(mimeType)
	runtime.KeepAlive(mtime)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// CreateFailedThumbnail creates a failed thumbnail for the file so that we
// don't try to re-thumbnail the file later.
//
// Usage of this function is threadsafe and does blocking I/O.
//
// The function takes the following parameters:
//
//   - ctx (optional): GCancellable object, or NULL.
//   - uri of a file.
//   - mtime: modification time of the file.
//
func (factory *DesktopThumbnailFactory) CreateFailedThumbnail(ctx context.Context, uri string, mtime int32) error {
	var _arg0 *C.GnomeDesktopThumbnailFactory // out
	var _arg3 *C.GCancellable                 // out
	var _arg1 *C.char                         // out
	var _arg2 C.time_t                        // out
	var _cerr *C.GError                       // in

	_arg0 = (*C.GnomeDesktopThumbnailFactory)(unsafe.Pointer(coreglib.InternObject(factory).Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg3 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(uri)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.time_t(mtime)

	C.gnome_desktop_thumbnail_factory_create_failed_thumbnail(_arg0, _arg1, _arg2, _arg3, &_cerr)
	runtime.KeepAlive(factory)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(uri)
	runtime.KeepAlive(mtime)

	var _goerr error // out

	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _goerr
}

// The function takes the following parameters:
//
//   - result of the operation.
//
func (factory *DesktopThumbnailFactory) CreateFailedThumbnailFinish(result gio.AsyncResulter) error {
	var _arg0 *C.GnomeDesktopThumbnailFactory // out
	var _arg1 *C.GAsyncResult                 // out
	var _cerr *C.GError                       // in

	_arg0 = (*C.GnomeDesktopThumbnailFactory)(unsafe.Pointer(coreglib.InternObject(factory).Native()))
	_arg1 = (*C.GAsyncResult)(unsafe.Pointer(coreglib.InternObject(result).Native()))

	C.gnome_desktop_thumbnail_factory_create_failed_thumbnail_finish(_arg0, _arg1, &_cerr)
	runtime.KeepAlive(factory)
	runtime.KeepAlive(result)

	var _goerr error // out

	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _goerr
}

// GenerateThumbnail tries to generate a thumbnail for the specified file.
// If it succeeds it returns a pixbuf that can be used as a thumbnail.
//
// Usage of this function is threadsafe and does blocking I/O.
//
// The function takes the following parameters:
//
//   - ctx (optional) object or NULL.
//   - uri of a file.
//   - mimeType: mime type of the file.
//
// The function returns the following values:
//
//   - pixbuf: thumbnail pixbuf if thumbnailing succeeded, NULL otherwise and
//     error will be set.
//
func (factory *DesktopThumbnailFactory) GenerateThumbnail(ctx context.Context, uri, mimeType string) (*gdkpixbuf.Pixbuf, error) {
	var _arg0 *C.GnomeDesktopThumbnailFactory // out
	var _arg3 *C.GCancellable                 // out
	var _arg1 *C.char                         // out
	var _arg2 *C.char                         // out
	var _cret *C.GdkPixbuf                    // in
	var _cerr *C.GError                       // in

	_arg0 = (*C.GnomeDesktopThumbnailFactory)(unsafe.Pointer(coreglib.InternObject(factory).Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg3 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(uri)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.char)(unsafe.Pointer(C.CString(mimeType)))
	defer C.free(unsafe.Pointer(_arg2))

	_cret = C.gnome_desktop_thumbnail_factory_generate_thumbnail(_arg0, _arg1, _arg2, _arg3, &_cerr)
	runtime.KeepAlive(factory)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(uri)
	runtime.KeepAlive(mimeType)

	var _pixbuf *gdkpixbuf.Pixbuf // out
	var _goerr error              // out

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
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _pixbuf, _goerr
}

// The function takes the following parameters:
//
//   - result of the operation.
//
// The function returns the following values:
//
//   - pixbuf: thumbnail pixbuf if thumbnailing succeeded, NULL otherwise.
//
//     Since 43.0.
//
func (factory *DesktopThumbnailFactory) GenerateThumbnailFinish(result gio.AsyncResulter) (*gdkpixbuf.Pixbuf, error) {
	var _arg0 *C.GnomeDesktopThumbnailFactory // out
	var _arg1 *C.GAsyncResult                 // out
	var _cret *C.GdkPixbuf                    // in
	var _cerr *C.GError                       // in

	_arg0 = (*C.GnomeDesktopThumbnailFactory)(unsafe.Pointer(coreglib.InternObject(factory).Native()))
	_arg1 = (*C.GAsyncResult)(unsafe.Pointer(coreglib.InternObject(result).Native()))

	_cret = C.gnome_desktop_thumbnail_factory_generate_thumbnail_finish(_arg0, _arg1, &_cerr)
	runtime.KeepAlive(factory)
	runtime.KeepAlive(result)

	var _pixbuf *gdkpixbuf.Pixbuf // out
	var _goerr error              // out

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
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _pixbuf, _goerr
}

// HasValidFailedThumbnail tries to locate an failed thumbnail for the file
// specified. Writing and looking for failed thumbnails is important to avoid to
// try to thumbnail e.g. broken images several times.
//
// Usage of this function is threadsafe and does blocking I/O.
//
// The function takes the following parameters:
//
//   - uri of a file.
//   - mtime of the file.
//
// The function returns the following values:
//
//   - ok: TRUE if there is a failed thumbnail for the file.
//
func (factory *DesktopThumbnailFactory) HasValidFailedThumbnail(uri string, mtime int32) bool {
	var _arg0 *C.GnomeDesktopThumbnailFactory // out
	var _arg1 *C.char                         // out
	var _arg2 C.time_t                        // out
	var _cret C.gboolean                      // in

	_arg0 = (*C.GnomeDesktopThumbnailFactory)(unsafe.Pointer(coreglib.InternObject(factory).Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(uri)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.time_t(mtime)

	_cret = C.gnome_desktop_thumbnail_factory_has_valid_failed_thumbnail(_arg0, _arg1, _arg2)
	runtime.KeepAlive(factory)
	runtime.KeepAlive(uri)
	runtime.KeepAlive(mtime)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Lookup tries to locate an existing thumbnail for the file specified.
//
// Usage of this function is threadsafe and does blocking I/O.
//
// The function takes the following parameters:
//
//   - uri of a file.
//   - mtime of the file.
//
// The function returns the following values:
//
//   - utf8: absolute path of the thumbnail, or NULL if none exist.
//
func (factory *DesktopThumbnailFactory) Lookup(uri string, mtime int32) string {
	var _arg0 *C.GnomeDesktopThumbnailFactory // out
	var _arg1 *C.char                         // out
	var _arg2 C.time_t                        // out
	var _cret *C.char                         // in

	_arg0 = (*C.GnomeDesktopThumbnailFactory)(unsafe.Pointer(coreglib.InternObject(factory).Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(uri)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.time_t(mtime)

	_cret = C.gnome_desktop_thumbnail_factory_lookup(_arg0, _arg1, _arg2)
	runtime.KeepAlive(factory)
	runtime.KeepAlive(uri)
	runtime.KeepAlive(mtime)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// SaveThumbnail saves thumbnail at the right place. If the save fails a failed
// thumbnail is written.
//
// Usage of this function is threadsafe and does blocking I/O.
//
// The function takes the following parameters:
//
//   - ctx (optional): GCancellable object, or NULL.
//   - thumbnail as a pixbuf.
//   - uri of a file.
//   - originalMtime: modification time of the original file.
//
func (factory *DesktopThumbnailFactory) SaveThumbnail(ctx context.Context, thumbnail *gdkpixbuf.Pixbuf, uri string, originalMtime int32) error {
	var _arg0 *C.GnomeDesktopThumbnailFactory // out
	var _arg4 *C.GCancellable                 // out
	var _arg1 *C.GdkPixbuf                    // out
	var _arg2 *C.char                         // out
	var _arg3 C.time_t                        // out
	var _cerr *C.GError                       // in

	_arg0 = (*C.GnomeDesktopThumbnailFactory)(unsafe.Pointer(coreglib.InternObject(factory).Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg4 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}
	_arg1 = (*C.GdkPixbuf)(unsafe.Pointer(coreglib.InternObject(thumbnail).Native()))
	_arg2 = (*C.char)(unsafe.Pointer(C.CString(uri)))
	defer C.free(unsafe.Pointer(_arg2))
	_arg3 = C.time_t(originalMtime)

	C.gnome_desktop_thumbnail_factory_save_thumbnail(_arg0, _arg1, _arg2, _arg3, _arg4, &_cerr)
	runtime.KeepAlive(factory)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(thumbnail)
	runtime.KeepAlive(uri)
	runtime.KeepAlive(originalMtime)

	var _goerr error // out

	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _goerr
}

// The function takes the following parameters:
//
//   - result of the operation.
//
func (factory *DesktopThumbnailFactory) SaveThumbnailFinish(result gio.AsyncResulter) error {
	var _arg0 *C.GnomeDesktopThumbnailFactory // out
	var _arg1 *C.GAsyncResult                 // out
	var _cerr *C.GError                       // in

	_arg0 = (*C.GnomeDesktopThumbnailFactory)(unsafe.Pointer(coreglib.InternObject(factory).Native()))
	_arg1 = (*C.GAsyncResult)(unsafe.Pointer(coreglib.InternObject(result).Native()))

	C.gnome_desktop_thumbnail_factory_save_thumbnail_finish(_arg0, _arg1, &_cerr)
	runtime.KeepAlive(factory)
	runtime.KeepAlive(result)

	var _goerr error // out

	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _goerr
}

// DesktopThumbnailFactoryClass: instance of this type is always passed by
// reference.
type DesktopThumbnailFactoryClass struct {
	*desktopThumbnailFactoryClass
}

// desktopThumbnailFactoryClass is the struct that's finalized.
type desktopThumbnailFactoryClass struct {
	native *C.GnomeDesktopThumbnailFactoryClass
}
