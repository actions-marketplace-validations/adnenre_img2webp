package walk

import (
	"os"
	"path/filepath"
	"strings"
)

type Visitor struct {
	OnFile func(path string) error
}

// --------------------
// IGNORED DIRECTORIES
// --------------------
var skipDirs = map[string]bool{
	"node_modules": true,
	".git":         true,
	"dist":         true,
	"build":        true,
	".next":        true,
	".nuxt":        true,
	".cache":       true,
	"coverage":     true,
	".vscode":      true,
	".idea":        true,
}

func shouldSkipDir(name string) bool {
	return skipDirs[strings.ToLower(name)]
}

// --------------------
// MAIN WALKER
// --------------------
func Walk(root string, v Visitor) error {
	return filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return nil
		}

		// skip directories
		if info.IsDir() {
			if shouldSkipDir(info.Name()) {
				return filepath.SkipDir
			}
			return nil
		}

		// optional: ignore hidden/system files
		if strings.HasPrefix(filepath.Base(path), ".") {
			if !IsImage(path) && !IsSource(path) {
				return nil
			}
		}

		if v.OnFile != nil {
			return v.OnFile(path)
		}

		return nil
	})
}

// --------------------
// IMAGE DETECTION
// --------------------
func IsImage(path string) bool {
	switch strings.ToLower(filepath.Ext(path)) {
	case ".png", ".jpg", ".jpeg", ".webp", ".gif":
		return true
	}
	return false
}

// --------------------
// SOURCE FILE DETECTION
// --------------------
func IsSource(path string) bool {
	switch strings.ToLower(filepath.Ext(path)) {
	case ".html", ".css", ".scss", ".js", ".ts", ".jsx", ".tsx", ".json", ".md", ".vue":
		return true
	}
	return false
}