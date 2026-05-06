// convert/webp.go
package convert

// #cgo pkg-config: libwebp
// #include <webp/encode.h>
// #include <stdlib.h>
import "C"
import (
    "errors"
    "image"
    "image/draw"
    "unsafe"
)



// EncodeWebP encodes an image.Image to WebP format.
func EncodeWebP(img image.Image, opts EncodeOptions) ([]byte, error) {
    bounds := img.Bounds()
    w, h := bounds.Dx(), bounds.Dy()
    if w == 0 || h == 0 {
        return nil, errors.New("image has zero size")
    }

    // Convert to RGBA (libwebp expects this layout)
    rgba := image.NewRGBA(bounds)
    draw.Draw(rgba, bounds, img, bounds.Min, draw.Src)

    data := rgba.Pix
    stride := rgba.Stride

    var webpData *C.uint8_t
    var webpSize C.size_t

    if opts.Lossless {
        webpSize = C.WebPEncodeLosslessRGBA(
            (*C.uint8_t)(&data[0]),
            C.int(w),
            C.int(h),
            C.int(stride),
            &webpData,
        )
    } else {
        quality := C.float(opts.Quality)
        webpSize = C.WebPEncodeRGBA(
            (*C.uint8_t)(&data[0]),
            C.int(w),
            C.int(h),
            C.int(stride),
            quality,
            &webpData,
        )
    }

    if webpSize == 0 {
        return nil, errors.New("libwebp encoding failed")
    }
    defer C.free(unsafe.Pointer(webpData))

    return C.GoBytes(unsafe.Pointer(webpData), C.int(webpSize)), nil
}