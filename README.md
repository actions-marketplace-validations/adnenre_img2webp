# img2webp

A fast, lightweight CLI tool to **convert images to WebP and automatically update references** in HTML, CSS, JS, Markdown, and framework files (React, Vue, Angular).

Designed for developers who care about **performance, compression, and automation**.

---

## 🚀 Features

- Convert JPEG/PNG to WebP (Google’s `libwebp`)
- **Automatically rewrite image references** – `.png` → `.webp` inside `.html`, `.css`, `.js`, `.md`, `.jsx`, `.tsx`, `.vue`
- Batch process entire folders
- Adjustable quality & lossless mode
- Keep or delete original images
- GitHub Action ready – zero‑config CI/CD
- CLI‑first, fast, minimal dependencies

---

## ⚙️ How It Works

1. Scan a directory (e.g., `./public`, `./src/assets`)
2. Convert every `.png`, `.jpg`, `.jpeg` to WebP
3. Find all references to those images in your source files and replace the extension
4. Optionally delete the original heavy images

**Result:** Your site serves WebP automatically – no manual code changes.

---

## 📦 Installation

````bash
go install github.com/adnenre/img2webp@latest
--

**Requirements:**
- Go 1.26.2+
- `libwebp-dev` (Linux: `sudo apt-get install libwebp-dev`, macOS: `brew install libwebp`)
- Windows: encoding not supported natively – use WSL or GitHub Action

---

## 🧪 Quick Start

```bash
img2webp --input ./public --quality 85
--

This converts all images in `./public` and updates every reference inside that folder.

---

## 🖥️ CLI Usage

```bash
img2webp [flags]
--

---

## ⚙️ Flags

| Flag                 | Default | Description                                      |
|----------------------|---------|--------------------------------------------------|
| `--input`            | `.`     | Root directory to scan                           |
| `--quality`          | `75`    | WebP quality (0–100)                             |
| `--lossless`         | `false` | Enable lossless compression                      |
| `--keep-original`    | `false` | Keep original images after conversion            |
| `--update-refs`      | `true`  | Rewrite image references in source files         |
| `--dry-run`          | `false` | Preview changes without writing                  |
| `--verbose`          | `false` | Print detailed logs                              |

---

## 📌 Examples

### Convert a single image (basic)

```bash
img2webp --input photo.jpg
--

### Convert whole folder, keep originals

```bash
img2webp --input ./images --keep-original true
--

### High quality, lossless, dry-run preview

```bash
img2webp --input ./assets --quality 95 --lossless --dry-run
--

### Only rewrite references (no conversion)

```bash
img2webp --input ./public --update-refs true --dry-run
--

---

## 🔁 What “update references” means

Before:
```html
<img src="images/photo.png">
--
```css
.hero { background-image: url('../img/bg.jpg'); }
--

After running `img2webp`:
```html
<img src="images/photo.webp">
--
```css
.hero { background-image: url('../img/bg.webp'); }
--

All extensions are updated automatically.

---

## 📁 Supported Formats

- Input: JPEG (`.jpg`, `.jpeg`), PNG (`.png`)
- Output: WebP (`.webp`)

---

## 🤖 GitHub Action

Add this to `.github/workflows/webp.yml`:

```yaml
name: Optimize images to WebP

on: [push, pull_request]

jobs:
  convert:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - uses: adnenre/img2webp@v0.2.0
        with:
          input-dir: './public'
          quality: '85'
          keep-original: 'false'
      - run: npm run build
--

**Inputs:**

| Input           | Default | Description              |
|-----------------|---------|--------------------------|
| `input-dir`     | `.`     | Directory to scan        |
| `quality`       | `75`    | WebP quality             |
| `lossless`      | `false` | Use lossless compression |
| `keep-original` | `false` | Keep original images     |
| `update-refs`   | `true`  | Rewrite references       |
| `dry-run`       | `false` | Simulate only            |

---

## 🧠 Framework‑specific recommendations

| Framework                  | Recommended `input-dir` |
|----------------------------|--------------------------|
| React (CRA, Vite, Next.js) | `./public`               |
| Vue (Vue CLI, Vite)        | `./public`               |
| Angular                    | `./src/assets`           |
| Plain HTML/CSS             | `./` or `./images`       |

---

## 📊 Performance Notes

- WebP reduces file size by **30–70%** vs JPEG/PNG
- Higher quality = larger file size
- Lossless preserves data but may increase size

---

## 🛠️ Roadmap

- [ ] GIF support
- [ ] Web interface
- [ ] Advanced compression presets
- [ ] Benchmark reports

---

## 🤝 Contributing

Contributions welcome!

1. Fork the repository
2. Create a feature branch
3. Commit your changes
4. Open a pull request

---

## 👨‍💻 Author

**Adnen Rebai**
Software Engineer | Open Source Enthusiast | Performance-focused tools

- 🌐 Website: [adnenre.dev](https://adnenre.dev)
- 🐙 GitHub: [adnenre](https://github.com/adnenre)

---

## 📬 Contact

Open an issue on GitHub for bugs or feature requests.

---

## ⭐ Support

If you find this project useful, give it a star ⭐

---

## 📄 License

MIT License. See [LICENSE](LICENSE) for details.
````
