package main

import (
    "bytes"
    "flag"
    "fmt"
    "image"
    _ "image/jpeg"
    _ "image/png"
    "os"
    "path/filepath"
    "strings"

    "github.com/adnenre/img2webp/convert"
    "github.com/adnenre/img2webp/rewrite"
    "github.com/adnenre/img2webp/walk"
)

var (
    inputDir     = flag.String("input", ".", "Root directory to scan")
    quality      = flag.Float64("quality", 75, "WebP quality (0-100)")
    lossless     = flag.Bool("lossless", false, "Use lossless compression")
    keepOriginal = flag.Bool("keep-original", false, "Do not delete original images")
    updateRefs   = flag.Bool("update-refs", true, "Rewrite references in source files")
    dryRun       = flag.Bool("dry-run", false, "Only simulate changes")
    verbose      = flag.Bool("verbose", false, "Print detailed logs")
)

var imageExts = []string{".png", ".jpg", ".jpeg"}

func main() {
    flag.Parse()

    // Clean and validate input directory
    dir := filepath.Clean(*inputDir)
    info, err := os.Stat(dir)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error: input directory '%s' does not exist\n", dir)
        os.Exit(1)
    }
    if !info.IsDir() {
        fmt.Fprintf(os.Stderr, "Error: '%s' is not a directory\n", dir)
        os.Exit(1)
    }

    if *verbose {
        fmt.Println("img2webp starting...")
        fmt.Printf("Scanning directory: %s\n", dir)
    }

    stats := struct {
        converted, deleted, rewritten int
    }{}

    visitor := walk.Visitor{
        OnImage: func(imgPath string) error {
            webpPath := strings.TrimSuffix(imgPath, filepath.Ext(imgPath)) + ".webp"
            if *verbose {
                fmt.Printf("Converting: %s -> %s\n", imgPath, webpPath)
            }
            if !*dryRun {
                if err := convertImage(imgPath, webpPath); err != nil {
                    // Log error but continue processing other files
                    fmt.Fprintf(os.Stderr, "Warning: %s\n", err)
                    return nil // continue walking
                }
                stats.converted++
                if !*keepOriginal {
                    if err := os.Remove(imgPath); err != nil {
                        fmt.Fprintf(os.Stderr, "Warning: failed to remove %s: %v\n", imgPath, err)
                        return nil
                    }
                    stats.deleted++
                }
            }
            return nil
        },
        OnSource: func(srcPath string) error {
            if !*updateRefs {
                return nil
            }
            if *verbose {
                fmt.Printf("Rewriting: %s\n", srcPath)
            }
            if !*dryRun {
                if err := rewriteFile(srcPath); err != nil {
                    fmt.Fprintf(os.Stderr, "Warning: failed to rewrite %s: %v\n", srcPath, err)
                    return nil
                }
                stats.rewritten++
            }
            return nil
        },
    }

    if err := walk.Walk(dir, visitor); err != nil {
        fmt.Fprintf(os.Stderr, "Error: %v\n", err)
        os.Exit(1)
    }

    fmt.Printf("Done. Converted: %d, Deleted originals: %d, Rewritten source files: %d\n",
        stats.converted, stats.deleted, stats.rewritten)
}

func convertImage(src, dst string) error {
    f, err := os.Open(src)
    if err != nil {
        return err
    }
    defer f.Close()
    img, _, err := image.Decode(f)
    if err != nil {
        return err
    }
    opts := convert.EncodeOptions{
        Quality:  float32(*quality),
        Lossless: *lossless,
    }
    data, err := convert.EncodeWebP(img, opts)
    if err != nil {
        return err
    }
    return os.WriteFile(dst, data, 0644)
}

func rewriteFile(path string) error {
    data, err := os.ReadFile(path)
    if err != nil {
        return err
    }
    ext := strings.ToLower(filepath.Ext(path))
    var newData []byte
    var rewriteErr error

    switch ext {
    case ".html", ".htm":
        newData, rewriteErr = rewrite.ReplaceExtensionsInHTML(data, imageExts, ".webp")
    case ".css":
        newData = rewrite.ReplaceExtensionsInCSS(data, imageExts, ".webp")
    default:
        newData = rewrite.ReplaceExtensionsInText(data, imageExts, ".webp")
    }
    if rewriteErr != nil {
        return rewriteErr
    }
    if !bytes.Equal(data, newData) {
        return os.WriteFile(path, newData, 0644)
    }
    return nil
}