// Code generated by girgen. DO NOT EDIT.

package gnomedesktop

// #include <stdlib.h>
import "C"

// GetPlatformVersion returns an integer with the major version of GNOME.
// Useful for dynamic languages like Javascript or Python (static languages like
// C should use GNOME_DESKTOP_PLATFORM_VERSION). If this function doesn't exist,
// it can be presumed that the GNOME platform version is 42 or previous.
//
// The function returns the following values:
//
//   - gint: integer with the major version of GNOME.
//
func GetPlatformVersion() int {
	var _cret C.int // in

	_cret = C.gnome_get_platform_version()

	var _gint int // out

	_gint = int(_cret)

	return _gint
}
