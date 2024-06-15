// Code generated by girgen. DO NOT EDIT.

package gnomerr

import (
	"fmt"
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gerror"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
)

// #include <stdlib.h>
// #include <glib-object.h>
// extern void _gotk4_gnomerr4_Screen_ConnectOutputDisconnected(gpointer, GnomeRROutput*, guintptr);
// extern void _gotk4_gnomerr4_Screen_ConnectOutputConnected(gpointer, GnomeRROutput*, guintptr);
// extern void _gotk4_gnomerr4_Screen_ConnectChanged(gpointer, guintptr);
// extern void _gotk4_gnomerr4_ScreenClass_output_disconnected(GnomeRRScreen*, GnomeRROutput*);
// extern void _gotk4_gnomerr4_ScreenClass_output_connected(GnomeRRScreen*, GnomeRROutput*);
// extern void _gotk4_gnomerr4_ScreenClass_changed(GnomeRRScreen*);
// void _gotk4_gnomerr4_Screen_virtual_changed(void* fnptr, GnomeRRScreen* arg0) {
//   ((void (*)(GnomeRRScreen*))(fnptr))(arg0);
// };
// void _gotk4_gnomerr4_Screen_virtual_output_connected(void* fnptr, GnomeRRScreen* arg0, GnomeRROutput* arg1) {
//   ((void (*)(GnomeRRScreen*, GnomeRROutput*))(fnptr))(arg0, arg1);
// };
// void _gotk4_gnomerr4_Screen_virtual_output_disconnected(void* fnptr, GnomeRRScreen* arg0, GnomeRROutput* arg1) {
//   ((void (*)(GnomeRRScreen*, GnomeRROutput*))(fnptr))(arg0, arg1);
// };
import "C"

// GType values.
var (
	GTypeScreen = coreglib.Type(C.gnome_rr_screen_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeScreen, F: marshalScreen},
	})
}

const CONNECTOR_TYPE_PANEL = "Panel"

type Error C.gint

const (
	RrErrorUnknown Error = iota
	RrErrorNoRandrExtension
	RrErrorRandrError
	RrErrorBoundsError
	RrErrorCrtcAssignment
	RrErrorNoMatchingConfig
	RrErrorNoDpmsExtension
)

// String returns the name in string for Error.
func (e Error) String() string {
	switch e {
	case RrErrorUnknown:
		return "Unknown"
	case RrErrorNoRandrExtension:
		return "NoRandrExtension"
	case RrErrorRandrError:
		return "RandrError"
	case RrErrorBoundsError:
		return "BoundsError"
	case RrErrorCrtcAssignment:
		return "CrtcAssignment"
	case RrErrorNoMatchingConfig:
		return "NoMatchingConfig"
	case RrErrorNoDpmsExtension:
		return "NoDpmsExtension"
	default:
		return fmt.Sprintf("Error(%d)", e)
	}
}

// ErrorQuark returns the error domain used by the GnomeRR API.
//
// The function returns the following values:
//
//   - quark: gnomeRR error domain.
//
func ErrorQuark() glib.Quark {
	var _cret C.GQuark // in

	_cret = C.gnome_rr_error_quark()

	var _quark glib.Quark // out

	_quark = uint32(_cret)
	type _ = glib.Quark
	type _ = uint32

	return _quark
}

// ScreenOverrides contains methods that are overridable.
type ScreenOverrides struct {
	Changed func()
	// The function takes the following parameters:
	//
	OutputConnected func(output *Output)
	// The function takes the following parameters:
	//
	OutputDisconnected func(output *Output)
}

func defaultScreenOverrides(v *Screen) ScreenOverrides {
	return ScreenOverrides{
		Changed:            v.changed,
		OutputConnected:    v.outputConnected,
		OutputDisconnected: v.outputDisconnected,
	}
}

type Screen struct {
	_ [0]func() // equal guard
	*coreglib.Object

	gio.AsyncInitable
	gio.Initable
}

var (
	_ coreglib.Objector = (*Screen)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*Screen, *ScreenClass, ScreenOverrides](
		GTypeScreen,
		initScreenClass,
		wrapScreen,
		defaultScreenOverrides,
	)
}

func initScreenClass(gclass unsafe.Pointer, overrides ScreenOverrides, classInitFunc func(*ScreenClass)) {
	pclass := (*C.GnomeRRScreenClass)(unsafe.Pointer(C.g_type_check_class_cast((*C.GTypeClass)(gclass), C.GType(GTypeScreen))))

	if overrides.Changed != nil {
		pclass.changed = (*[0]byte)(C._gotk4_gnomerr4_ScreenClass_changed)
	}

	if overrides.OutputConnected != nil {
		pclass.output_connected = (*[0]byte)(C._gotk4_gnomerr4_ScreenClass_output_connected)
	}

	if overrides.OutputDisconnected != nil {
		pclass.output_disconnected = (*[0]byte)(C._gotk4_gnomerr4_ScreenClass_output_disconnected)
	}

	if classInitFunc != nil {
		class := (*ScreenClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapScreen(obj *coreglib.Object) *Screen {
	return &Screen{
		Object: obj,
		AsyncInitable: gio.AsyncInitable{
			Object: obj,
		},
		Initable: gio.Initable{
			Object: obj,
		},
	}
}

func marshalScreen(p uintptr) (interface{}, error) {
	return wrapScreen(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

func (screen *Screen) ConnectChanged(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(screen, "changed", false, unsafe.Pointer(C._gotk4_gnomerr4_Screen_ConnectChanged), f)
}

// ConnectOutputConnected: this signal is emitted when a display device is
// connected to a port, or a port is hotplugged with an active output. The
// latter can happen if a laptop is docked, and the dock provides a new active
// output.
//
// The output value is not a #GObject. The returned output value can only
// assume to be valid during the emission of the signal (i.e. within your signal
// handler only), as it may change later when the screen is modified due to an
// event from the X server, or due to another place in the application modifying
// the screen and the output. Therefore, deal with changes to the output right
// in your signal handler, instead of keeping the output reference for an async
// or idle function.
func (screen *Screen) ConnectOutputConnected(f func(output *Output)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(screen, "output-connected", false, unsafe.Pointer(C._gotk4_gnomerr4_Screen_ConnectOutputConnected), f)
}

// ConnectOutputDisconnected: this signal is emitted when a display device is
// disconnected from a port, or a port output is hot-unplugged. The latter can
// happen if a laptop is undocked, and the dock provided the output.
//
// The output value is not a #GObject. The returned output value can only
// assume to be valid during the emission of the signal (i.e. within your signal
// handler only), as it may change later when the screen is modified due to an
// event from the X server, or due to another place in the application modifying
// the screen and the output. Therefore, deal with changes to the output right
// in your signal handler, instead of keeping the output reference for an async
// or idle function.
func (screen *Screen) ConnectOutputDisconnected(f func(output *Output)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(screen, "output-disconnected", false, unsafe.Pointer(C._gotk4_gnomerr4_Screen_ConnectOutputDisconnected), f)
}

// NewScreen creates a unique RRScreen instance for the specified display.
//
// The function takes the following parameters:
//
//   - display: windowing system connection used to query the display data.
//
// The function returns the following values:
//
//   - screen: unique RRScreen instance, specific to the screen, or NULL if this
//     could not be created, for instance if the driver does not support Xrandr
//     1.2. Each Display thus has a single instance of RRScreen.
//
func NewScreen(display *gdk.Display) (*Screen, error) {
	var _arg1 *C.GdkDisplay    // out
	var _cret *C.GnomeRRScreen // in
	var _cerr *C.GError        // in

	_arg1 = (*C.GdkDisplay)(unsafe.Pointer(coreglib.InternObject(display).Native()))

	_cret = C.gnome_rr_screen_new(_arg1, &_cerr)
	runtime.KeepAlive(display)

	var _screen *Screen // out
	var _goerr error    // out

	_screen = wrapScreen(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _screen, _goerr
}

// NewScreenFinish finishes the asynchronous creation of a new RRScreen
// instance.
//
// The function takes the following parameters:
//
//   - result of the asynchronous operation.
//
// The function returns the following values:
//
//   - screen: newly created instance; on error, this function will return NULL
//     and set the given #GError.
//
func NewScreenFinish(result gio.AsyncResulter) (*Screen, error) {
	var _arg1 *C.GAsyncResult  // out
	var _cret *C.GnomeRRScreen // in
	var _cerr *C.GError        // in

	_arg1 = (*C.GAsyncResult)(unsafe.Pointer(coreglib.InternObject(result).Native()))

	_cret = C.gnome_rr_screen_new_finish(_arg1, &_cerr)
	runtime.KeepAlive(result)

	var _screen *Screen // out
	var _goerr error    // out

	_screen = wrapScreen(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _screen, _goerr
}

// CrtcByID retrieves the CRTC of the screen using the given identifier.
//
// The function takes the following parameters:
//
//   - id: identifier of a CRTC.
//
// The function returns the following values:
//
//   - crtc: CRTC identified by id.
//
func (screen *Screen) CrtcByID(id uint32) *Crtc {
	var _arg0 *C.GnomeRRScreen // out
	var _arg1 C.guint32        // out
	var _cret *C.GnomeRRCrtc   // in

	_arg0 = (*C.GnomeRRScreen)(unsafe.Pointer(coreglib.InternObject(screen).Native()))
	_arg1 = C.guint32(id)

	_cret = C.gnome_rr_screen_get_crtc_by_id(_arg0, _arg1)
	runtime.KeepAlive(screen)
	runtime.KeepAlive(id)

	var _crtc *Crtc // out

	_crtc = (*Crtc)(gextras.NewStructNative(unsafe.Pointer(_cret)))

	return _crtc
}

// The function returns the following values:
//
//   - mode: current RRDpmsMode of this screen.
//
func (screen *Screen) DpmsMode() (DpmsMode, error) {
	var _arg0 *C.GnomeRRScreen  // out
	var _arg1 C.GnomeRRDpmsMode // in
	var _cerr *C.GError         // in

	_arg0 = (*C.GnomeRRScreen)(unsafe.Pointer(coreglib.InternObject(screen).Native()))

	C.gnome_rr_screen_get_dpms_mode(_arg0, &_arg1, &_cerr)
	runtime.KeepAlive(screen)

	var _mode DpmsMode // out
	var _goerr error   // out

	_mode = DpmsMode(_arg1)
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _mode, _goerr
}

// OutputByID retrieves the output of a screen using the given identifier.
//
// The function takes the following parameters:
//
//   - id: identifier of an output.
//
// The function returns the following values:
//
//   - output identified by id.
//
func (screen *Screen) OutputByID(id uint32) *Output {
	var _arg0 *C.GnomeRRScreen // out
	var _arg1 C.guint32        // out
	var _cret *C.GnomeRROutput // in

	_arg0 = (*C.GnomeRRScreen)(unsafe.Pointer(coreglib.InternObject(screen).Native()))
	_arg1 = C.guint32(id)

	_cret = C.gnome_rr_screen_get_output_by_id(_arg0, _arg1)
	runtime.KeepAlive(screen)
	runtime.KeepAlive(id)

	var _output *Output // out

	_output = (*Output)(gextras.NewStructNative(unsafe.Pointer(_cret)))

	return _output
}

// OutputByName retrieves the output for the given name.
//
// The function takes the following parameters:
//
// The function returns the following values:
//
//   - output identified by name.
//
func (screen *Screen) OutputByName(name string) *Output {
	var _arg0 *C.GnomeRRScreen // out
	var _arg1 *C.char          // out
	var _cret *C.GnomeRROutput // in

	_arg0 = (*C.GnomeRRScreen)(unsafe.Pointer(coreglib.InternObject(screen).Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(name)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gnome_rr_screen_get_output_by_name(_arg0, _arg1)
	runtime.KeepAlive(screen)
	runtime.KeepAlive(name)

	var _output *Output // out

	_output = (*Output)(gextras.NewStructNative(unsafe.Pointer(_cret)))

	return _output
}

// Ranges: get the ranges of the screen.
//
// The function returns the following values:
//
//   - minWidth: minimum width.
//   - maxWidth: maximum width.
//   - minHeight: minimum height.
//   - maxHeight: maximum height.
//
func (screen *Screen) Ranges() (minWidth, maxWidth, minHeight, maxHeight int) {
	var _arg0 *C.GnomeRRScreen // out
	var _arg1 C.int            // in
	var _arg2 C.int            // in
	var _arg3 C.int            // in
	var _arg4 C.int            // in

	_arg0 = (*C.GnomeRRScreen)(unsafe.Pointer(coreglib.InternObject(screen).Native()))

	C.gnome_rr_screen_get_ranges(_arg0, &_arg1, &_arg2, &_arg3, &_arg4)
	runtime.KeepAlive(screen)

	var _minWidth int  // out
	var _maxWidth int  // out
	var _minHeight int // out
	var _maxHeight int // out

	_minWidth = int(_arg1)
	_maxWidth = int(_arg2)
	_minHeight = int(_arg3)
	_maxHeight = int(_arg4)

	return _minWidth, _maxWidth, _minHeight, _maxHeight
}

// ListCloneModes lists all available XRandR clone modes.
//
// The function returns the following values:
//
//   - modes: available XRandR clone modes.
//
func (screen *Screen) ListCloneModes() []*Mode {
	var _arg0 *C.GnomeRRScreen // out
	var _cret **C.GnomeRRMode  // in

	_arg0 = (*C.GnomeRRScreen)(unsafe.Pointer(coreglib.InternObject(screen).Native()))

	_cret = C.gnome_rr_screen_list_clone_modes(_arg0)
	runtime.KeepAlive(screen)

	var _modes []*Mode // out

	{
		var i int
		var z *C.GnomeRRMode
		for p := _cret; *p != z; p = &unsafe.Slice(p, 2)[1] {
			i++
		}

		src := unsafe.Slice(_cret, i)
		_modes = make([]*Mode, i)
		for i := range src {
			_modes[i] = (*Mode)(gextras.NewStructNative(unsafe.Pointer(src[i])))
		}
	}

	return _modes
}

// ListCrtcs: list all CRTCs of the given screen.
//
// The function returns the following values:
//
//   - crtcs: available CRTCs.
//
func (screen *Screen) ListCrtcs() []*Crtc {
	var _arg0 *C.GnomeRRScreen // out
	var _cret **C.GnomeRRCrtc  // in

	_arg0 = (*C.GnomeRRScreen)(unsafe.Pointer(coreglib.InternObject(screen).Native()))

	_cret = C.gnome_rr_screen_list_crtcs(_arg0)
	runtime.KeepAlive(screen)

	var _crtcs []*Crtc // out

	{
		var i int
		var z *C.GnomeRRCrtc
		for p := _cret; *p != z; p = &unsafe.Slice(p, 2)[1] {
			i++
		}

		src := unsafe.Slice(_cret, i)
		_crtcs = make([]*Crtc, i)
		for i := range src {
			_crtcs[i] = (*Crtc)(gextras.NewStructNative(unsafe.Pointer(src[i])))
		}
	}

	return _crtcs
}

// ListModes lists all available XRandR modes.
//
// The function returns the following values:
//
//   - modes: available XRandR modes.
//
func (screen *Screen) ListModes() []*Mode {
	var _arg0 *C.GnomeRRScreen // out
	var _cret **C.GnomeRRMode  // in

	_arg0 = (*C.GnomeRRScreen)(unsafe.Pointer(coreglib.InternObject(screen).Native()))

	_cret = C.gnome_rr_screen_list_modes(_arg0)
	runtime.KeepAlive(screen)

	var _modes []*Mode // out

	{
		var i int
		var z *C.GnomeRRMode
		for p := _cret; *p != z; p = &unsafe.Slice(p, 2)[1] {
			i++
		}

		src := unsafe.Slice(_cret, i)
		_modes = make([]*Mode, i)
		for i := range src {
			_modes[i] = (*Mode)(gextras.NewStructNative(unsafe.Pointer(src[i])))
		}
	}

	return _modes
}

// ListOutputs: list all outputs of the given screen.
//
// The function returns the following values:
//
//   - outputs: available outputs.
//
func (screen *Screen) ListOutputs() []*Output {
	var _arg0 *C.GnomeRRScreen  // out
	var _cret **C.GnomeRROutput // in

	_arg0 = (*C.GnomeRRScreen)(unsafe.Pointer(coreglib.InternObject(screen).Native()))

	_cret = C.gnome_rr_screen_list_outputs(_arg0)
	runtime.KeepAlive(screen)

	var _outputs []*Output // out

	{
		var i int
		var z *C.GnomeRROutput
		for p := _cret; *p != z; p = &unsafe.Slice(p, 2)[1] {
			i++
		}

		src := unsafe.Slice(_cret, i)
		_outputs = make([]*Output, i)
		for i := range src {
			_outputs[i] = (*Output)(gextras.NewStructNative(unsafe.Pointer(src[i])))
		}
	}

	return _outputs
}

// Refresh refreshes the screen configuration, and calls the screen's callback
// if it exists and if the screen's configuration changed.
func (screen *Screen) Refresh() error {
	var _arg0 *C.GnomeRRScreen // out
	var _cerr *C.GError        // in

	_arg0 = (*C.GnomeRRScreen)(unsafe.Pointer(coreglib.InternObject(screen).Native()))

	C.gnome_rr_screen_refresh(_arg0, &_cerr)
	runtime.KeepAlive(screen)

	var _goerr error // out

	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _goerr
}

// SetDpmsMode: this method also disables the DPMS timeouts.
//
// The function takes the following parameters:
//
func (screen *Screen) SetDpmsMode(mode DpmsMode) error {
	var _arg0 *C.GnomeRRScreen  // out
	var _arg1 C.GnomeRRDpmsMode // out
	var _cerr *C.GError         // in

	_arg0 = (*C.GnomeRRScreen)(unsafe.Pointer(coreglib.InternObject(screen).Native()))
	_arg1 = C.GnomeRRDpmsMode(mode)

	C.gnome_rr_screen_set_dpms_mode(_arg0, _arg1, &_cerr)
	runtime.KeepAlive(screen)
	runtime.KeepAlive(mode)

	var _goerr error // out

	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _goerr
}

func (screen *Screen) changed() {
	gclass := (*C.GnomeRRScreenClass)(coreglib.PeekParentClass(screen))
	fnarg := gclass.changed

	var _arg0 *C.GnomeRRScreen // out

	_arg0 = (*C.GnomeRRScreen)(unsafe.Pointer(coreglib.InternObject(screen).Native()))

	C._gotk4_gnomerr4_Screen_virtual_changed(unsafe.Pointer(fnarg), _arg0)
	runtime.KeepAlive(screen)
}

// The function takes the following parameters:
//
func (screen *Screen) outputConnected(output *Output) {
	gclass := (*C.GnomeRRScreenClass)(coreglib.PeekParentClass(screen))
	fnarg := gclass.output_connected

	var _arg0 *C.GnomeRRScreen // out
	var _arg1 *C.GnomeRROutput // out

	_arg0 = (*C.GnomeRRScreen)(unsafe.Pointer(coreglib.InternObject(screen).Native()))
	_arg1 = (*C.GnomeRROutput)(gextras.StructNative(unsafe.Pointer(output)))

	C._gotk4_gnomerr4_Screen_virtual_output_connected(unsafe.Pointer(fnarg), _arg0, _arg1)
	runtime.KeepAlive(screen)
	runtime.KeepAlive(output)
}

// The function takes the following parameters:
//
func (screen *Screen) outputDisconnected(output *Output) {
	gclass := (*C.GnomeRRScreenClass)(coreglib.PeekParentClass(screen))
	fnarg := gclass.output_disconnected

	var _arg0 *C.GnomeRRScreen // out
	var _arg1 *C.GnomeRROutput // out

	_arg0 = (*C.GnomeRRScreen)(unsafe.Pointer(coreglib.InternObject(screen).Native()))
	_arg1 = (*C.GnomeRROutput)(gextras.StructNative(unsafe.Pointer(output)))

	C._gotk4_gnomerr4_Screen_virtual_output_disconnected(unsafe.Pointer(fnarg), _arg0, _arg1)
	runtime.KeepAlive(screen)
	runtime.KeepAlive(output)
}

// ScreenClass: instance of this type is always passed by reference.
type ScreenClass struct {
	*screenClass
}

// screenClass is the struct that's finalized.
type screenClass struct {
	native *C.GnomeRRScreenClass
}
