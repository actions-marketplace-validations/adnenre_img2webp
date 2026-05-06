package walk

import (
    "os"
    "path/filepath"
    "testing"
)

func TestWalk(t *testing.T) {
    tmp := t.TempDir()
    // Create a dummy image and source file
    imgFile := filepath.Join(tmp, "test.png")
    srcFile := filepath.Join(tmp, "index.html")
    os.WriteFile(imgFile, []byte("dummy"), 0644)
    os.WriteFile(srcFile, []byte("<html></html>"), 0644)

    var images, sources []string
    err := Walk(tmp, Visitor{
        OnImage:  func(p string) error { images = append(images, p); return nil },
        OnSource: func(p string) error { sources = append(sources, p); return nil },
    })
    if err != nil {
        t.Fatal(err)
    }
    if len(images) != 1 || images[0] != imgFile {
        t.Error("image not found")
    }
    if len(sources) != 1 || sources[0] != srcFile {
        t.Error("source not found")
    }
}