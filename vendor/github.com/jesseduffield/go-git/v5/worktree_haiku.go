//go:build haiku
package git

//import "os"

// isSymlinkWindowsNonAdmin is a no-op on Haiku.
func isSymlinkWindowsNonAdmin(err error) bool {
	return false
}
