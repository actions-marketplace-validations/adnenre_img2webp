package rewrite

import (
    "regexp"
    "strings"
)

func ReplaceExtensionsInCSS(cssContent []byte, oldExts []string, newExt string) []byte {
    re := regexp.MustCompile(`url\(([^)]+)\)`)
    return re.ReplaceAllFunc(cssContent, func(match []byte) []byte {
        // Extract content inside parentheses
        inner := match[4 : len(match)-1] // remove "url(" and ")"
        // Trim spaces
        str := strings.TrimSpace(string(inner))
        // Detect and preserve quotes
        quote := ""
        if strings.HasPrefix(str, "'") && strings.HasSuffix(str, "'") {
            quote = "'"
            str = strings.Trim(str, "'")
        } else if strings.HasPrefix(str, "\"") && strings.HasSuffix(str, "\"") {
            quote = "\""
            str = strings.Trim(str, "\"")
        }
        newStr := replaceExtension(str, oldExts, newExt)
        if newStr == str {
            return match
        }
        if quote != "" {
            return []byte("url(" + quote + newStr + quote + ")")
        }
        return []byte("url(" + newStr + ")")
    })
}