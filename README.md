# img2webp

Convert images to WebP and automatically update references in HTML, CSS, JS, and Markdown files.

## Features
- Recursively scans a directory
- Converts .png, .jpg, .jpeg to WebP using libwebp (CGO)
- Rewrites image paths in source files
- Optionally removes original images
- Ready for GitHub Actions

## Installation
```
go install github.com/yourusername/img2webp@latest
```
