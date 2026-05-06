package walk

import (
	"os"
	"path/filepath"
	"testing"
)

func TestWalk(t *testing.T) {
	tmp := t.TempDir()

	// --------------------
	// create test files
	// --------------------
	imgFile := filepath.Join(tmp, "test.png")
	srcFile := filepath.Join(tmp, "index.html")

	err := os.WriteFile(imgFile, []byte("img"), 0644)
	if err != nil {
		t.Fatal(err)
	}

	err = os.WriteFile(srcFile, []byte("<html></html>"), 0644)
	if err != nil {
		t.Fatal(err)
	}

	// --------------------
	// results collector
	// --------------------
	var images []string
	var sources []string

	err = Walk(tmp, Visitor{
		OnFile: func(p string) error {

			if IsImage(p) {
				images = append(images, p)
			}

			if IsSource(p) {
				sources = append(sources, p)
			}

			return nil
		},
	})

	if err != nil {
		t.Fatal(err)
	}

	// --------------------
	// assertions
	// --------------------
	if len(images) != 1 {
		t.Fatalf("expected 1 image, got %d", len(images))
	}

	if images[0] != imgFile {
		t.Errorf("expected image %s, got %s", imgFile, images[0])
	}

	if len(sources) != 1 {
		t.Fatalf("expected 1 source, got %d", len(sources))
	}

	if sources[0] != srcFile {
		t.Errorf("expected source %s, got %s", srcFile, sources[0])
	}
}