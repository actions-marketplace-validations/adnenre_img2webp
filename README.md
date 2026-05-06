# img2webp

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

--codebash
go install github.com/adnenre/img2webp@latest
--code

**Requirements:**

- Go 1.26.2+
- `cwebp` (WebP tools)
  - Linux: `sudo apt-get install webp`
  - macOS: `brew install webp`
  - Windows: install WebP tools or use WSL

---

## 🧪 Quick Start

--codebash
img2webp --input ./public --quality 85 --replace
--code

This converts all images in `./public` and updates every reference inside that folder.

---

## 🖥️ CLI Usage

--codebash
img2webp [flags]
--code

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

--codebash
img2webp --input ./testdata --replace
--code

---

### High quality conversion

--codebash
img2webp --input ./assets --q 90 --replace
--code

---

### Dry run (preview changes only)

--codebash
img2webp --input ./public --dry-run
--code

---

## 🔁 What “replace” means

Before:

--codehtml
<img src="images/photo.png" />
--code

--codecss
.hero {
background-image: url("../img/bg.jpg");
}
--code

After running `img2webp`:

--codehtml
<img src="images/photo.webp" />
--code

--codecss
.hero {
background-image: url("../img/bg.webp");
}
--code

All extensions are updated automatically.

---

## 📁 Supported Formats

- Input: `.jpg`, `.jpeg`, `.png`
- Output: `.webp`

---

## 🤖 GitHub Action

Add this to `.github/workflows/webp.yml`:

--codeyaml
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

--code

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
