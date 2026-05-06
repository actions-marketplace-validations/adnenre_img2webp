package walk

import (
    "os"
    "path/filepath"
    "strings"
)

type Visitor struct {
    OnImage  func(path string) error
    OnSource func(path string) error
}

func Walk(root string, v Visitor) error {
    return filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
        if err != nil {
            return err
        }
        if info.IsDir() {
            return nil
        }
        ext := strings.ToLower(filepath.Ext(path))
        if isImageExt(ext) {
            if v.OnImage != nil {
                return v.OnImage(path)
            }
        } else if isSourceExt(ext) {
            if v.OnSource != nil {
                return v.OnSource(path)
            }
        }
        return nil
    })
}

func isImageExt(ext string) bool {
    return ext == ".png" || ext == ".jpg" || ext == ".jpeg"
}

func isSourceExt(ext string) bool {
    exts := []string{".html", ".htm", ".css", ".js", ".md", ".jsx", ".tsx", ".vue", ".php"}
    for _, e := range exts {
        if ext == e {
            return true
        }
    }
    return false
}