// Code generated by girgen. DO NOT EDIT.

package gnomedesktop

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// extern void _gotk4_gnomedesktop3_XkbInfo_ConnectLayoutsChanged(gpointer, guintptr);
import "C"

// GType values.
var (
	GTypeXkbInfo = coreglib.Type(C.gnome_xkb_info_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeXkbInfo, F: marshalXkbInfo},
	})
}

// XkbInfoOverrides contains methods that are overridable.
type XkbInfoOverrides struct {
}

func defaultXkbInfoOverrides(v *XkbInfo) XkbInfoOverrides {
	return XkbInfoOverrides{}
}

type XkbInfo struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*XkbInfo)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*XkbInfo, *XkbInfoClass, XkbInfoOverrides](
		GTypeXkbInfo,
		initXkbInfoClass,
		wrapXkbInfo,
		defaultXkbInfoOverrides,
	)
}

func initXkbInfoClass(gclass unsafe.Pointer, overrides XkbInfoOverrides, classInitFunc func(*XkbInfoClass)) {
	if classInitFunc != nil {
		class := (*XkbInfoClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapXkbInfo(obj *coreglib.Object) *XkbInfo {
	return &XkbInfo{
		Object: obj,
	}
}

func marshalXkbInfo(p uintptr) (interface{}, error) {
	return wrapXkbInfo(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

func (self *XkbInfo) ConnectLayoutsChanged(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(self, "layouts-changed", false, unsafe.Pointer(C._gotk4_gnomedesktop3_XkbInfo_ConnectLayoutsChanged), f)
}

// The function returns the following values:
//
//   - xkbInfo: new XkbInfo instance.
//
func NewXkbInfo() *XkbInfo {
	var _cret *C.GnomeXkbInfo // in

	_cret = C.gnome_xkb_info_new()

	var _xkbInfo *XkbInfo // out

	_xkbInfo = wrapXkbInfo(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _xkbInfo
}

// The function takes the following parameters:
//
//   - groupId: identifier for group.
//
// The function returns the following values:
//
//   - utf8: translated description for the group group_id.
//
func (self *XkbInfo) DescriptionForGroup(groupId string) string {
	var _arg0 *C.GnomeXkbInfo // out
	var _arg1 *C.gchar        // out
	var _cret *C.gchar        // in

	_arg0 = (*C.GnomeXkbInfo)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(groupId)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gnome_xkb_info_description_for_group(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(groupId)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// The function takes the following parameters:
//
//   - groupId: identifier for group containing the option.
//   - id: option identifier.
//
// The function returns the following values:
//
//   - utf8: translated description for the option id.
//
func (self *XkbInfo) DescriptionForOption(groupId, id string) string {
	var _arg0 *C.GnomeXkbInfo // out
	var _arg1 *C.gchar        // out
	var _arg2 *C.gchar        // out
	var _cret *C.gchar        // in

	_arg0 = (*C.GnomeXkbInfo)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(groupId)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(id)))
	defer C.free(unsafe.Pointer(_arg2))

	_cret = C.gnome_xkb_info_description_for_option(_arg0, _arg1, _arg2)
	runtime.KeepAlive(self)
	runtime.KeepAlive(groupId)
	runtime.KeepAlive(id)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// AllLayouts returns a list of all layout identifiers we know about.
//
// The function returns the following values:
//
//   - list: list of layout names. The caller takes ownership of the #GList but
//     not of the strings themselves, those are internally allocated and must
//     not be modified.
//
func (self *XkbInfo) AllLayouts() []string {
	var _arg0 *C.GnomeXkbInfo // out
	var _cret *C.GList        // in

	_arg0 = (*C.GnomeXkbInfo)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.gnome_xkb_info_get_all_layouts(_arg0)
	runtime.KeepAlive(self)

	var _list []string // out

	_list = make([]string, 0, gextras.ListSize(unsafe.Pointer(_cret)))
	gextras.MoveList(unsafe.Pointer(_cret), true, func(v unsafe.Pointer) {
		src := (*C.gchar)(v)
		var dst string // out
		dst = C.GoString((*C.gchar)(unsafe.Pointer(src)))
		_list = append(_list, dst)
	})

	return _list
}

// AllOptionGroups returns a list of all option group identifiers we know about.
//
// The function returns the following values:
//
//   - list: list of option group ids. The caller takes ownership of the #GList
//     but not of the strings themselves, those are internally allocated and
//     must not be modified.
//
func (self *XkbInfo) AllOptionGroups() []string {
	var _arg0 *C.GnomeXkbInfo // out
	var _cret *C.GList        // in

	_arg0 = (*C.GnomeXkbInfo)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.gnome_xkb_info_get_all_option_groups(_arg0)
	runtime.KeepAlive(self)

	var _list []string // out

	_list = make([]string, 0, gextras.ListSize(unsafe.Pointer(_cret)))
	gextras.MoveList(unsafe.Pointer(_cret), true, func(v unsafe.Pointer) {
		src := (*C.gchar)(v)
		var dst string // out
		dst = C.GoString((*C.gchar)(unsafe.Pointer(src)))
		_list = append(_list, dst)
	})

	return _list
}

// LanguagesForLayout returns a list of all languages supported by a layout,
// given by layout_id.
//
// The function takes the following parameters:
//
//   - layoutId: layout identifier.
//
// The function returns the following values:
//
//   - list of ISO 639 code strings. The caller takes ownership of the #GList
//     but not of the strings themselves, those are internally allocated and
//     must not be modified.
//
func (self *XkbInfo) LanguagesForLayout(layoutId string) []string {
	var _arg0 *C.GnomeXkbInfo // out
	var _arg1 *C.gchar        // out
	var _cret *C.GList        // in

	_arg0 = (*C.GnomeXkbInfo)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(layoutId)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gnome_xkb_info_get_languages_for_layout(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(layoutId)

	var _list []string // out

	_list = make([]string, 0, gextras.ListSize(unsafe.Pointer(_cret)))
	gextras.MoveList(unsafe.Pointer(_cret), true, func(v unsafe.Pointer) {
		src := (*C.gchar)(v)
		var dst string // out
		dst = C.GoString((*C.gchar)(unsafe.Pointer(src)))
		_list = append(_list, dst)
	})

	return _list
}

// LayoutInfo retrieves information about a layout. Both display_name and
// short_name are suitable to show in UIs and might be localized if translations
// are available.
//
// Some layouts don't provide a short name (2 or 3 letters) or don't specify a
// XKB variant, in those cases short_name or xkb_variant are empty strings, i.e.
// "".
//
// If the given layout doesn't exist the return value is FALSE and all the (out)
// parameters are set to NULL.
//
// The function takes the following parameters:
//
//   - id layout's identifier about which to retrieve the info.
//
// The function returns the following values:
//
//   - displayName (optional): location to store the layout's display name,
//     or NULL.
//   - shortName (optional): location to store the layout's short name, or NULL.
//   - xkbLayout (optional): location to store the layout's XKB name, or NULL.
//   - xkbVariant (optional): location to store the layout's XKB variant,
//     or NULL.
//   - ok: TRUE if the layout exists or FALSE otherwise.
//
func (self *XkbInfo) LayoutInfo(id string) (displayName, shortName, xkbLayout, xkbVariant string, ok bool) {
	var _arg0 *C.GnomeXkbInfo // out
	var _arg1 *C.gchar        // out
	var _arg2 *C.gchar        // in
	var _arg3 *C.gchar        // in
	var _arg4 *C.gchar        // in
	var _arg5 *C.gchar        // in
	var _cret C.gboolean      // in

	_arg0 = (*C.GnomeXkbInfo)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(id)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gnome_xkb_info_get_layout_info(_arg0, _arg1, &_arg2, &_arg3, &_arg4, &_arg5)
	runtime.KeepAlive(self)
	runtime.KeepAlive(id)

	var _displayName string // out
	var _shortName string   // out
	var _xkbLayout string   // out
	var _xkbVariant string  // out
	var _ok bool            // out

	if _arg2 != nil {
		_displayName = C.GoString((*C.gchar)(unsafe.Pointer(_arg2)))
	}
	if _arg3 != nil {
		_shortName = C.GoString((*C.gchar)(unsafe.Pointer(_arg3)))
	}
	if _arg4 != nil {
		_xkbLayout = C.GoString((*C.gchar)(unsafe.Pointer(_arg4)))
	}
	if _arg5 != nil {
		_xkbVariant = C.GoString((*C.gchar)(unsafe.Pointer(_arg5)))
	}
	if _cret != 0 {
		_ok = true
	}

	return _displayName, _shortName, _xkbLayout, _xkbVariant, _ok
}

// LayoutsForCountry returns a list of all layout identifiers we know about for
// country_code.
//
// The function takes the following parameters:
//
//   - countryCode: ISO 3166 code string.
//
// The function returns the following values:
//
//   - list: list of layout ids. The caller takes ownership of the #GList but
//     not of the strings themselves, those are internally allocated and must
//     not be modified.
//
func (self *XkbInfo) LayoutsForCountry(countryCode string) []string {
	var _arg0 *C.GnomeXkbInfo // out
	var _arg1 *C.gchar        // out
	var _cret *C.GList        // in

	_arg0 = (*C.GnomeXkbInfo)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(countryCode)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gnome_xkb_info_get_layouts_for_country(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(countryCode)

	var _list []string // out

	_list = make([]string, 0, gextras.ListSize(unsafe.Pointer(_cret)))
	gextras.MoveList(unsafe.Pointer(_cret), true, func(v unsafe.Pointer) {
		src := (*C.gchar)(v)
		var dst string // out
		dst = C.GoString((*C.gchar)(unsafe.Pointer(src)))
		_list = append(_list, dst)
	})

	return _list
}

// LayoutsForLanguage returns a list of all layout identifiers we know about for
// language_code.
//
// The function takes the following parameters:
//
//   - languageCode: ISO 639 code string.
//
// The function returns the following values:
//
//   - list: list of layout ids. The caller takes ownership of the #GList but
//     not of the strings themselves, those are internally allocated and must
//     not be modified.
//
func (self *XkbInfo) LayoutsForLanguage(languageCode string) []string {
	var _arg0 *C.GnomeXkbInfo // out
	var _arg1 *C.gchar        // out
	var _cret *C.GList        // in

	_arg0 = (*C.GnomeXkbInfo)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(languageCode)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gnome_xkb_info_get_layouts_for_language(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(languageCode)

	var _list []string // out

	_list = make([]string, 0, gextras.ListSize(unsafe.Pointer(_cret)))
	gextras.MoveList(unsafe.Pointer(_cret), true, func(v unsafe.Pointer) {
		src := (*C.gchar)(v)
		var dst string // out
		dst = C.GoString((*C.gchar)(unsafe.Pointer(src)))
		_list = append(_list, dst)
	})

	return _list
}

// OptionsForGroup returns a list of all option identifiers we know about for
// group group_id.
//
// The function takes the following parameters:
//
//   - groupId group's identifier about which to retrieve the options.
//
// The function returns the following values:
//
//   - list: list of option ids. The caller takes ownership of the #GList but
//     not of the strings themselves, those are internally allocated and must
//     not be modified.
//
func (self *XkbInfo) OptionsForGroup(groupId string) []string {
	var _arg0 *C.GnomeXkbInfo // out
	var _arg1 *C.gchar        // out
	var _cret *C.GList        // in

	_arg0 = (*C.GnomeXkbInfo)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(groupId)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gnome_xkb_info_get_options_for_group(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(groupId)

	var _list []string // out

	_list = make([]string, 0, gextras.ListSize(unsafe.Pointer(_cret)))
	gextras.MoveList(unsafe.Pointer(_cret), true, func(v unsafe.Pointer) {
		src := (*C.gchar)(v)
		var dst string // out
		dst = C.GoString((*C.gchar)(unsafe.Pointer(src)))
		_list = append(_list, dst)
	})

	return _list
}

// XkbInfoClass: instance of this type is always passed by reference.
type XkbInfoClass struct {
	*xkbInfoClass
}

// xkbInfoClass is the struct that's finalized.
type xkbInfoClass struct {
	native *C.GnomeXkbInfoClass
}
