// Code generated by girgen. DO NOT EDIT.

package gnomerr

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
	GTypeConfig = coreglib.Type(C.gnome_rr_config_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeConfig, F: marshalConfig},
	})
}

// ConfigOverrides contains methods that are overridable.
type ConfigOverrides struct {
}

func defaultConfigOverrides(v *Config) ConfigOverrides {
	return ConfigOverrides{}
}

type Config struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*Config)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*Config, *ConfigClass, ConfigOverrides](
		GTypeConfig,
		initConfigClass,
		wrapConfig,
		defaultConfigOverrides,
	)
}

func initConfigClass(gclass unsafe.Pointer, overrides ConfigOverrides, classInitFunc func(*ConfigClass)) {
	if classInitFunc != nil {
		class := (*ConfigClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapConfig(obj *coreglib.Object) *Config {
	return &Config{
		Object: obj,
	}
}

func marshalConfig(p uintptr) (interface{}, error) {
	return wrapConfig(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// The function takes the following parameters:
//
// The function returns the following values:
//
func NewConfigCurrent(screen *Screen) (*Config, error) {
	var _arg1 *C.GnomeRRScreen // out
	var _cret *C.GnomeRRConfig // in
	var _cerr *C.GError        // in

	_arg1 = (*C.GnomeRRScreen)(unsafe.Pointer(coreglib.InternObject(screen).Native()))

	_cret = C.gnome_rr_config_new_current(_arg1, &_cerr)
	runtime.KeepAlive(screen)

	var _config *Config // out
	var _goerr error    // out

	_config = wrapConfig(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _config, _goerr
}

// The function takes the following parameters:
//
func (self *Config) Applicable(screen *Screen) error {
	var _arg0 *C.GnomeRRConfig // out
	var _arg1 *C.GnomeRRScreen // out
	var _cerr *C.GError        // in

	_arg0 = (*C.GnomeRRConfig)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = (*C.GnomeRRScreen)(unsafe.Pointer(coreglib.InternObject(screen).Native()))

	C.gnome_rr_config_applicable(_arg0, _arg1, &_cerr)
	runtime.KeepAlive(self)
	runtime.KeepAlive(screen)

	var _goerr error // out

	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _goerr
}

// The function takes the following parameters:
//
func (self *Config) Apply(screen *Screen) error {
	var _arg0 *C.GnomeRRConfig // out
	var _arg1 *C.GnomeRRScreen // out
	var _cerr *C.GError        // in

	_arg0 = (*C.GnomeRRConfig)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = (*C.GnomeRRScreen)(unsafe.Pointer(coreglib.InternObject(screen).Native()))

	C.gnome_rr_config_apply(_arg0, _arg1, &_cerr)
	runtime.KeepAlive(self)
	runtime.KeepAlive(screen)

	var _goerr error // out

	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _goerr
}

// The function takes the following parameters:
//
func (self *Config) ApplyPersistent(screen *Screen) error {
	var _arg0 *C.GnomeRRConfig // out
	var _arg1 *C.GnomeRRScreen // out
	var _cerr *C.GError        // in

	_arg0 = (*C.GnomeRRConfig)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = (*C.GnomeRRScreen)(unsafe.Pointer(coreglib.InternObject(screen).Native()))

	C.gnome_rr_config_apply_persistent(_arg0, _arg1, &_cerr)
	runtime.KeepAlive(self)
	runtime.KeepAlive(screen)

	var _goerr error // out

	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _goerr
}

// The function returns the following values:
//
func (self *Config) EnsurePrimary() bool {
	var _arg0 *C.GnomeRRConfig // out
	var _cret C.gboolean       // in

	_arg0 = (*C.GnomeRRConfig)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.gnome_rr_config_ensure_primary(_arg0)
	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// The function takes the following parameters:
//
// The function returns the following values:
//
func (config1 *Config) Equal(config2 *Config) bool {
	var _arg0 *C.GnomeRRConfig // out
	var _arg1 *C.GnomeRRConfig // out
	var _cret C.gboolean       // in

	_arg0 = (*C.GnomeRRConfig)(unsafe.Pointer(coreglib.InternObject(config1).Native()))
	_arg1 = (*C.GnomeRRConfig)(unsafe.Pointer(coreglib.InternObject(config2).Native()))

	_cret = C.gnome_rr_config_equal(_arg0, _arg1)
	runtime.KeepAlive(config1)
	runtime.KeepAlive(config2)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// The function returns the following values:
//
//   - ok: whether at least two outputs are at (0, 0) offset and they have the
//     same width/height. Those outputs are of course connected and on (i.e.
//     they have a CRTC assigned).
//
func (self *Config) Clone() bool {
	var _arg0 *C.GnomeRRConfig // out
	var _cret C.gboolean       // in

	_arg0 = (*C.GnomeRRConfig)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.gnome_rr_config_get_clone(_arg0)
	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// The function returns the following values:
//
//   - outputInfos: output configuration for this RRConfig.
//
func (self *Config) Outputs() []*OutputInfo {
	var _arg0 *C.GnomeRRConfig      // out
	var _cret **C.GnomeRROutputInfo // in

	_arg0 = (*C.GnomeRRConfig)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.gnome_rr_config_get_outputs(_arg0)
	runtime.KeepAlive(self)

	var _outputInfos []*OutputInfo // out

	{
		var i int
		var z *C.GnomeRROutputInfo
		for p := _cret; *p != z; p = &unsafe.Slice(p, 2)[1] {
			i++
		}

		src := unsafe.Slice(_cret, i)
		_outputInfos = make([]*OutputInfo, i)
		for i := range src {
			_outputInfos[i] = wrapOutputInfo(coreglib.Take(unsafe.Pointer(src[i])))
		}
	}

	return _outputInfos
}

func (self *Config) LoadCurrent() error {
	var _arg0 *C.GnomeRRConfig // out
	var _cerr *C.GError        // in

	_arg0 = (*C.GnomeRRConfig)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	C.gnome_rr_config_load_current(_arg0, &_cerr)
	runtime.KeepAlive(self)

	var _goerr error // out

	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _goerr
}

// The function takes the following parameters:
//
// The function returns the following values:
//
func (config1 *Config) Match(config2 *Config) bool {
	var _arg0 *C.GnomeRRConfig // out
	var _arg1 *C.GnomeRRConfig // out
	var _cret C.gboolean       // in

	_arg0 = (*C.GnomeRRConfig)(unsafe.Pointer(coreglib.InternObject(config1).Native()))
	_arg1 = (*C.GnomeRRConfig)(unsafe.Pointer(coreglib.InternObject(config2).Native()))

	_cret = C.gnome_rr_config_match(_arg0, _arg1)
	runtime.KeepAlive(config1)
	runtime.KeepAlive(config2)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (self *Config) Sanitize() {
	var _arg0 *C.GnomeRRConfig // out

	_arg0 = (*C.GnomeRRConfig)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	C.gnome_rr_config_sanitize(_arg0)
	runtime.KeepAlive(self)
}

// The function takes the following parameters:
//
func (self *Config) SetClone(clone bool) {
	var _arg0 *C.GnomeRRConfig // out
	var _arg1 C.gboolean       // out

	_arg0 = (*C.GnomeRRConfig)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if clone {
		_arg1 = C.TRUE
	}

	C.gnome_rr_config_set_clone(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(clone)
}

// ConfigClass: instance of this type is always passed by reference.
type ConfigClass struct {
	*configClass
}

// configClass is the struct that's finalized.
type configClass struct {
	native *C.GnomeRRConfigClass
}
