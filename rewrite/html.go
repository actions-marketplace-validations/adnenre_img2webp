package rewrite

import (
    "bytes"
    "strings"
    "golang.org/x/net/html"
)

func ReplaceExtensionsInHTML(htmlContent []byte, oldExts []string, newExt string) ([]byte, error) {
    doc, err := html.Parse(bytes.NewReader(htmlContent))
    if err != nil {
        return nil, err
    }
    var traverse func(*html.Node)
    traverse = func(n *html.Node) {
        if n.Type == html.ElementNode {
            for i, attr := range n.Attr {
                if attr.Key == "src" || attr.Key == "href" {
                    newVal := replaceExtension(attr.Val, oldExts, newExt)
                    if newVal != attr.Val {
                        n.Attr[i].Val = newVal
                    }
                }
            }
        }
        for c := n.FirstChild; c != nil; c = c.NextSibling {
            traverse(c)
        }
    }
    traverse(doc)

    var buf bytes.Buffer
    if err := html.Render(&buf, doc); err != nil {
        return nil, err
    }
    return buf.Bytes(), nil
}

func replaceExtension(path string, oldExts []string, newExt string) string {
    lower := strings.ToLower(path)
    for _, ext := range oldExts {
        if strings.HasSuffix(lower, ext) {
            return path[:len(path)-len(ext)] + newExt
        }
    }
    return path
}