package rewrite

import "testing"

func TestReplaceExtensionsInHTML(t *testing.T) {
    input := `<img src="photo.png">`
    // The html package will render a full document with html/head/body.
    expected := `<html><head></head><body><img src="photo.webp"/></body></html>`
    oldExts := []string{".png"}
    out, err := ReplaceExtensionsInHTML([]byte(input), oldExts, ".webp")
    if err != nil {
        t.Fatal(err)
    }
    if string(out) != expected {
        t.Errorf("got %q, want %q", out, expected)
    }
}