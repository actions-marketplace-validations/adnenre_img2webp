package convert

import (
    "image"
    "image/color"
    "testing"
)

func TestEncodeWebP(t *testing.T) {
    img := image.NewRGBA(image.Rect(0, 0, 2, 2))
    img.Set(0, 0, color.RGBA{255, 0, 0, 255})
    data, err := EncodeWebP(img, EncodeOptions{Quality: 75})
    if err != nil {
        t.Fatal(err)
    }
    // Check minimal WebP signature (RIFF...WEBP)
    if len(data) < 12 || string(data[0:4]) != "RIFF" || string(data[8:12]) != "WEBP" {
        t.Error("encoded data is not a valid WebP file")
    }
}