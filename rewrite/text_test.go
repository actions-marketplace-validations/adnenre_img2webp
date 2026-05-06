package rewrite

import "testing"

func TestReplaceExtensionsInText(t *testing.T) {
    tests := []struct {
        name     string
        input    string
        oldExts  []string
        newExt   string
        expected string
    }{
        {
            name:     "simple markdown image",
            input:    "![alt](photo.png)",
            oldExts:  []string{".png"},
            newExt:   ".webp",
            expected: "![alt](photo.webp)",
        },
        {
            name:     "HTML img tag in text",
            input:    `<img src="photo.jpg" alt="test">`,
            oldExts:  []string{".jpg"},
            newExt:   ".webp",
            expected: `<img src="photo.webp" alt="test">`,
        },
        {
            name:     "multiple occurrences",
            input:    "a.png b.jpg c.jpeg",
            oldExts:  []string{".png", ".jpg", ".jpeg"},
            newExt:   ".webp",
            expected: "a.webp b.webp c.webp",
        },
        {
            name:     "path with slashes",
            input:    "/images/folder/photo.png",
            oldExts:  []string{".png"},
            newExt:   ".webp",
            expected: "/images/folder/photo.webp",
        },
        {
            name:     "word boundary prevents partial match",
            input:    "photo.png not aphotopng",
            oldExts:  []string{".png"},
            newExt:   ".webp",
            expected: "photo.webp not aphotopng",
        },
        {
            name:     "no match",
            input:    "photo.gif photo",
            oldExts:  []string{".png"},
            newExt:   ".webp",
            expected: "photo.gif photo",
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            out := ReplaceExtensionsInText([]byte(tt.input), tt.oldExts, tt.newExt)
            if string(out) != tt.expected {
                t.Errorf("got %q, want %q", out, tt.expected)
            }
        })
    }
}