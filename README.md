# img2webp

[![Test img2webp Action](https://github.com/adnenre/img2webp/actions/workflows/test-action.yml/badge.svg)](https://github.com/adnenre/img2webp/actions/workflows/test-action.yml)
[![GitHub Action](https://img.shields.io/badge/GitHub%20Action-Convert%20to%20WebP-blue)](https://github.com/marketplace/actions/img2webp)
[![Go Reference](https://pkg.go.dev/badge/github.com/adnenre/img2webp.svg)](https://pkg.go.dev/github.com/adnenre/img2webp)
[![Marketplace](https://img.shields.io/badge/Marketplace-Install-blue?logo=github)](https://github.com/marketplace/actions/img2webp)

A fast, lightweight CLI tool to **convert images to WebP and automatically update references** in HTML, CSS, JS, Markdown, and framework files (React, Vue, Angular).

Designed for developers who care about **performance, compression, and automation**.

---

## 🚀 Features

- Convert JPEG/PNG to WebP (Google’s `cwebp`)
- **Automatically rewrite image references** – `.png` → `.webp` inside `.html`, `.css`, `.js`, `.md`, `.jsx`, `.tsx`, `.vue`
- Batch process entire folders
- Adjustable quality & lossless mode
- Keep or delete original images
- GitHub Action ready – zero-config CI/CD
- CLI-first, fast, minimal dependencies

---

## ⚙️ How It Works

1. Scan a directory (e.g., `./public`, `./src/assets`)
2. Convert every `.png`, `.jpg`, `.jpeg` to WebP
3. Find all references to those images in your source files and replace the extension
4. Optionally delete the original heavy images

**Result:** Your site serves WebP automatically – no manual code changes.

---

## 📦 Installation

```bash
go install github.com/adnenre/img2webp@latest
```

**Requirements:**

- Go 1.26.2+
- `cwebp` (WebP tools)
  - Linux: `sudo apt-get install webp`
  - macOS: `brew install webp`
  - Windows: install WebP tools or use WSL

---

## 🧪 Quick Start

```bash
img2webp --input ./public --quality 85 --replace
```

This converts all images in `./public` and updates every reference inside that folder.

---

## 🖥️ CLI Usage

```bash
img2webp [flags]
```

---

## ⚙️ Flags

| Flag        | Default | Description                              |
| ----------- | ------- | ---------------------------------------- |
| `--input`   | `.`     | Root directory to scan                   |
| `--q`       | `75`    | WebP quality (0–100)                     |
| `--replace` | `true`  | Rewrite image references in source files |
| `--dry-run` | `false` | Preview changes without writing          |
| `--v`       | `false` | Print detailed logs                      |

---

## 📌 Examples

### Convert a folder (default)

```bash
img2webp --input ./testdata --replace
```

---

### High quality conversion

```bash
img2webp --input ./assets --q 90 --replace
```

---

### Dry run (preview changes only)

```bash
img2webp --input ./public --dry-run
```

---

## 🔁 What “replace” means

Before:

```html
<img src="images/photo.png" />
```

```css
.hero {
  background-image: url("../img/bg.jpg");
}
```

After running `img2webp`:

```html
<img src="images/photo.webp" />
```

```css
.hero {
  background-image: url("../img/bg.webp");
}
```

All extensions are updated automatically.

---

## 📁 Supported Formats

- Input: `.jpg`, `.jpeg`, `.png`
- Output: `.webp`

---

## 🤖 GitHub Action

Add this to `.github/workflows/webp.yml`:

```yaml
name: Optimize images to WebP

on: [push, pull_request]

jobs:
convert:
runs-on: ubuntu-latest
steps: - uses: actions/checkout@v4

      - name: Install webp tools
        run: sudo apt-get install webp -y

      - name: Run img2webp
        run: |
          go install github.com/adnenre/img2webp@latest
          img2webp --input ./public --q 85 --replace

      - name: Build project
        run: npm run build

```

---

## 🧠 Framework Recommendations

| Framework          | Recommended `--input` |
| ------------------ | --------------------- |
| React (Vite / CRA) | `./public`            |
| Vue                | `./public`            |
| Angular            | `./src/assets`        |
| Plain HTML/CSS     | `./` or `./images`    |

---

## 📊 Performance Notes

- WebP reduces file size by **30–70%**
- Higher quality = larger file size
- `--q 75` is recommended default balance

---

## 🛠️ Roadmap

- [ ] CSS url() advanced parser
- [ ] React/Angular AST support
- [ ] Backup & rollback mode
- [ ] Parallel conversion
- [ ] Plugin system

---

## 🤝 Contributing

1. Fork repo
2. Create feature branch
3. Commit changes
4. Open PR

---

## 👨‍💻 Author

**Adnen Rebai**

- 🌐 Website: https://adnenre.dev
- 🐙 GitHub: https://github.com/adnenre

---

## 📬 Contact

Open an issue on GitHub for bugs or feature requests.

---

## ⭐ Support

If you find this project useful, give it a star ⭐

---

## 📄 License

MIT License. See `LICENSE` for details.
