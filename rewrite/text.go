package rewrite

import (
    "regexp"
    "strings"
)

// ReplaceExtensionsInText replaces .png, .jpg, .jpeg with newExt
// using word boundaries to avoid partial matches.
func ReplaceExtensionsInText(text []byte, oldExts []string, newExt string) []byte {
    // oldExts includes the dot, e.g., ".png". Remove the dot for grouping.
    extsNoDot := make([]string, len(oldExts))
    for i, ext := range oldExts {
        extsNoDot[i] = strings.TrimPrefix(ext, ".")
    }
    extPattern := strings.Join(extsNoDot, "|")
    // Pattern: \b[\w./-]+\.(png|jpg|jpeg)\b
    pattern := `\b([\w./-]+)\.(` + extPattern + `)\b`
    re := regexp.MustCompile(pattern)
    // newExt already includes the dot (e.g., ".webp")
    return re.ReplaceAll(text, []byte("${1}"+newExt))
}