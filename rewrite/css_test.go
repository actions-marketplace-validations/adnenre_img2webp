package rewrite

import "testing"

func TestReplaceExtensionsInCSS(t *testing.T) {
    tests := []struct {
        name     string
        input    string
        oldExts  []string
        newExt   string
        expected string
    }{
        {
            name:     "simple url()",
            input:    "background: url('photo.png');",
            oldExts:  []string{".png"},
            newExt:   ".webp",
            expected: "background: url('photo.webp');",
        },
        {
            name:     "double quotes",
            input:    "background: url(\"photo.jpg\");",
            oldExts:  []string{".jpg"},
            newExt:   ".webp",
            expected: "background: url(\"photo.webp\");",
        },
        {
            name:     "no quotes",
            input:    "background: url(photo.jpeg);",
            oldExts:  []string{".jpeg"},
            newExt:   ".webp",
            expected: "background: url(photo.webp);",
        },
        {
            name:     "multiple extensions",
            input:    "background: url('photo.png'); list: url('icon.jpg');",
            oldExts:  []string{".png", ".jpg"},
            newExt:   ".webp",
            expected: "background: url('photo.webp'); list: url('icon.webp');",
        },
        {
            name:     "no match",
            input:    "background: url('photo.gif');",
            oldExts:  []string{".png", ".jpg"},
            newExt:   ".webp",
            expected: "background: url('photo.gif');",
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            out := ReplaceExtensionsInCSS([]byte(tt.input), tt.oldExts, tt.newExt)
            if string(out) != tt.expected {
                t.Errorf("got %q, want %q", out, tt.expected)
            }
        })
    }
}