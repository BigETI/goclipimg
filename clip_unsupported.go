// +build !linux,!darwin

package goclipimg

import "errors"

// ErrImagePasteUnsupported means that this system has no implementation for pasting an image.
var ErrImagePasteUnsupported = errors.New("This system doesn't have a paste implementation.")

// GetImageFromClipboard always returns ErrImagePasteUnsupported, since the
// compilation target currently doesn't support pasting images.
func GetImageFromClipboard() ([]byte, error) {
	return nil, ErrImagePasteUnsupported
}
