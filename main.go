package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/adnenre/img2webp/walk"
	"flag"
)

// -------------------------
// FLAG NORMALIZER (IMPORTANT)
// -------------------------
func normalizeArgs() {
	for i, arg := range os.Args {
		if strings.HasPrefix(arg, "--") {
			os.Args[i] = "-" + strings.TrimPrefix(arg, "--")
		}
	}
}

var (
	inputDir = flag.String("input", ".", "project root")
	quality  = flag.Int("q", 75, "webp quality")
	replace  = flag.String("replace", "true", "replace references in files")
	dryRun   = flag.String("dry-run", "false", "preview only")
	verbose  = flag.String("v", "false", "verbose")
)

type ImageMap struct {
	Old string
	New string
}

func parseBool(v string) bool {
	b, err := strconv.ParseBool(v)
	if err != nil {
		return false
	}
	return b
}

func main() {

	// 🔥 FIX: allow --flags support
	normalizeArgs()

	flag.Parse()

	root := filepath.Clean(*inputDir)

	doReplace := parseBool(*replace)
	isDryRun := parseBool(*dryRun)
	isVerbose := parseBool(*verbose)

	var images []ImageMap

	// -------------------------
	// 1. CONVERT IMAGES
	// -------------------------
	walk.Walk(root, walk.Visitor{
		OnFile: func(path string) error {

			if !walk.IsImage(path) {
				return nil
			}

			ext := filepath.Ext(path)
			output := strings.TrimSuffix(path, ext) + ".webp"

			if isVerbose {
				fmt.Println("convert:", path, "->", output)
			}

			_ = os.MkdirAll(filepath.Dir(output), 0755)

			if !isDryRun {
				if err := convert(path, output); err != nil {
					fmt.Println("conversion failed:", err)
					return nil
				}

				if _, err := os.Stat(output); err != nil {
					fmt.Println("missing output, skip delete:", path)
					return nil
				}

				_ = os.Remove(path)
			}

			images = append(images, ImageMap{
				Old: filepath.Base(path),
				New: strings.TrimSuffix(filepath.Base(path), ext) + ".webp",
			})

			return nil
		},
	})

	if len(images) == 0 {
		fmt.Println("no images found")
		return
	}

	// -------------------------
	// 2. REPLACE REFERENCES
	// -------------------------
	if doReplace {

		walk.Walk(root, walk.Visitor{
			OnFile: func(path string) error {

				if walk.IsImage(path) {
					return nil
				}

				data, err := os.ReadFile(path)
				if err != nil {
					return nil
				}

				content := string(data)
				original := content

				for _, img := range images {
					content = strings.ReplaceAll(content, img.Old, img.New)
				}

				if content == original {
					return nil
				}

				if isDryRun {
					fmt.Println("[dry-run] would update:", path)
					return nil
				}

				if isVerbose {
					fmt.Println("updated:", path)
				}

				return os.WriteFile(path, []byte(content), 0644)
			},
		})
	}

	fmt.Println("done ✔")
}

// -------------------------
// CWEBP CONVERTER
// -------------------------
func convert(src, dst string) error {

	args := []string{
		src,
		"-o", dst,
		"-q", fmt.Sprintf("%d", *quality),
	}

	cmd := exec.Command("cwebp", args...)

	var out bytes.Buffer
	cmd.Stderr = &out

	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("%v: %s", err, out.String())
	}

	return nil
}