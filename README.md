# img2webp

[![GitHub Action](https://img.shields.io/badge/GitHub%20Action-Convert%20to%20WebP-blue)](https://github.com/adnenre/img2webp)
[![Go Reference](https://pkg.go.dev/badge/github.com/adnenre/img2webp.svg)](https://pkg.go.dev/github.com/adnenre/img2webp)

Convert images to WebP and automatically update references in HTML, CSS, JS, and Markdown files.  
Perfect for static sites, React/Vue/Angular apps, or any project that wants to serve modern WebP images.

## Features

- 🖼️ Converts `.png`, `.jpg`, `.jpeg` → `.webp` (using Google's `libwebp` via CGO)
- 🔁 Rewrites references in `.html`, `.css`, `.js`, `.md`, `.jsx`, `.tsx`, `.vue` – automatically
- 🧹 Optionally deletes original images
- 🚀 Ready for GitHub Actions (no extra configuration)
- 📦 Can be used as a Go library, CLI tool, or GitHub Action

## Requirements

- **Go 1.26.2+** (if building from source)
- **libwebp-dev** (system library)  
  Install on Ubuntu/Debian: `sudo apt-get install libwebp-dev`  
  Install on macOS: `brew install libwebp`  
  Not required when using the GitHub Action – it installs automatically on Ubuntu runners.

## Usage

### 1. As a CLI tool

Install globally:

--code
go install github.com/adnenre/img2webp@latest
--code

Then run:

--code
img2webp --input ./public --quality 85 --keep-original false
--code

All flags:

| Flag              | Default | Description                              |
| ----------------- | ------- | ---------------------------------------- |
| `--input`         | `.`     | Root directory to scan                   |
| `--quality`       | `75`    | WebP quality (0–100)                     |
| `--lossless`      | `false` | Use lossless compression                 |
| `--keep-original` | `false` | Keep original images after conversion    |
| `--update-refs`   | `true`  | Rewrite image references in source files |
| `--dry-run`       | `false` | Preview changes without writing          |
| `--verbose`       | `false` | Print detailed logs                      |

### 2. As a GitHub Action

Add this to your `.github/workflows/webp.yml`:

--code
name: Optimize images to WebP

on: [push, pull_request]

jobs:
convert:
runs-on: ubuntu-latest
steps: - uses: actions/checkout@v4 - uses: adnenre/img2webp@v1
with:
input-dir: './public' # change to your image folder
quality: '85'
keep-original: 'false' # Now build/deploy your site – it will use the new .webp files - run: npm run build
--code

#### Inputs

| Input           | Default | Description              |
| --------------- | ------- | ------------------------ |
| `input-dir`     | `.`     | Directory to scan        |
| `quality`       | `75`    | WebP quality             |
| `lossless`      | `false` | Use lossless compression |
| `keep-original` | `false` | Keep original images     |
| `update-refs`   | `true`  | Rewrite references       |
| `dry-run`       | `false` | Simulate only            |

### 3. As a Go library

--code
import "github.com/adnenre/img2webp/convert"

func main() {
img, _ := os.Open("photo.png")
defer img.Close()
src, _, \_ := image.Decode(img)

    data, _ := convert.EncodeWebP(src, convert.EncodeOptions{
        Quality:  85,
        Lossless: false,
    })
    os.WriteFile("photo.webp", data, 0644)

}
--code

## Framework‑specific recommendations

| Framework                  | Recommended `input-dir` | Image location example       |
| -------------------------- | ----------------------- | ---------------------------- |
| React (CRA, Vite, Next.js) | `./public`              | `public/images/logo.png`     |
| Vue (Vue CLI, Vite)        | `./public`              | `public/images/logo.png`     |
| Angular                    | `./src/assets`          | `src/assets/images/logo.png` |
| Plain HTML/CSS             | `./` or `./images`      | `images/photo.png`           |

**Important:** Place your images in these folders, reference them using **relative paths** (e.g., `logo.png`). After conversion, references are updated to `logo.webp` automatically.

## Example: full CI pipeline

--code
name: Build and deploy with WebP

on:
push:
branches: [main]

jobs:
build:
runs-on: ubuntu-latest
steps: - uses: actions/checkout@v4 - uses: actions/setup-node@v4
with:
node-version: 20 - name: Convert images to WebP
uses: adnenre/img2webp@v1
with:
input-dir: './public'
quality: '85' - run: npm ci - run: npm run build - name: Deploy
uses: peaceiris/actions-gh-pages@v3
with:
github_token: ${{ secrets.GITHUB_TOKEN }}
publish_dir: ./build
--code

## How it works

1. The tool recursively walks the input directory.
2. For every `.png`, `.jpg`, `.jpeg` file, it encodes a WebP version using `libwebp`.
3. It scans all `.html`, `.css`, `.js`, `.md`, `.jsx`, `.tsx`, `.vue` files and replaces occurrences `image.png` → `image.webp`.
4. Original images are kept or deleted according to the `keep-original` flag.

Errors during conversion are logged but do **not** stop the entire process – other images and files continue.

## Development

If you want to build from source:

--code
git clone https://github.com/adnenre/img2webp.git
cd img2webp

# Install libwebp-dev (see Requirements)

go build -o img2webp .
./img2webp --help
--code

## License

MIT

## Author

Adnen Rebai
website : https://adnenre.dev
