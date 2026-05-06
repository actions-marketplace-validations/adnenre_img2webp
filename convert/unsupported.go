//go:build !cgo || windows

package convert

import (
    "errors"
    "image"
)

func EncodeWebP(img image.Image, opts EncodeOptions) ([]byte, error) {
    return nil, errors.New("WebP encoding requires CGO and libwebp (not available on this platform)")
}