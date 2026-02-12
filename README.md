# Auto-Commit Tool

Tool otomatis untuk membuat random commit ke multiple GitHub repositories.

## Fitur

- 📦 Konfigurasi multiple repositories dari `.env`
- 🔄 Auto-clone repositories (sekali saja)
- 🎲 Random selection dari list repositories
- 📝 **Random sentences** - Generate konten random dengan natural language
- 💬 **Random commit messages** - Setiap commit punya message yang unik dan random
- 🚀 Auto add, commit, dan push ke GitHub

## Setup

1. **Copy file .env.example ke .env:**

   ```bash
   cp .env.example .env
   ```

2. **Edit .env dengan credentials Anda:**

   ```env
   GIT_USERNAME=your-github-username
   GIT_ACCESS_TOKEN=your-github-personal-access-token
   REPOS=username/repo1,username/repo2,username/repo3
   ```

3. **Buat GitHub Personal Access Token:**
   - Buka https://github.com/settings/tokens
   - Generate new token (classic)
   - Pilih scope: `repo` (Full control of private repositories)
   - Copy token dan paste ke `.env`

## Cara Pakai

Jalankan program:

```bash
go run main.go
```

Atau build dulu:

```bash
go build -o auto-commit
./auto-commit
```

## Cara Kerja

1. Program membaca konfigurasi dari file `.env`
2. Membuat folder `repos/` untuk menyimpan repositories
3. Clone semua repositories yang belum ada (hanya sekali)
4. Pilih 1 repository secara random
5. Generate **random sentence** dengan natural language (5-14 kata)
6. Tambahkan random sentence ke `README.md` repo tersebut
7. Buat commit dengan **random tech phrase** sebagai commit message
8. Add, commit, dan push perubahan ke GitHub

## Contoh Output

```
📦 Loaded 3 repositories

🔄 Checking repositories...
📥 Cloning username/repo1...
✓ Cloned username/repo1
✓ username/repo2 already exists
✓ username/repo3 already exists

🎲 Selected random repo: username/repo2

📝 Committing changes...
  📌 Staging changes...
  💾 Creating commit: calculating the bandwidth won't do anything
  🚀 Pushing to remote...

✅ Successfully committed and pushed changes!
```

## Contoh Random Content di README.md

Setiap kali dijalankan, program akan menambahkan konten seperti:

```markdown
## Update 2026-02-12 10:30:45

The blue cat quickly jumps over the lazy dog while eating pizza.

## Update 2026-02-12 14:20:15

Seven mysterious dragons fly above ancient mountains during sunset.
```

## Struktur Folder

```
auto-commit/
├── main.go           # Program utama
├── go.mod            # Go module
├── .env              # Konfigurasi (jangan commit!)
├── .env.example      # Template konfigurasi
└── repos/            # Folder repositories (auto-generated)
    ├── repo1/
    ├── repo2/
    └── repo3/
```

## Tips

- Jalankan dengan cron job untuk auto-commit berkala
- Pastikan semua repos yang di-list memiliki akses write
- Token harus memiliki permission `repo`
- File `.env` tidak akan di-commit (sudah ada di `.gitignore`)

## Dependencies

- [gofakeit/v6](https://github.com/brianvoe/gofakeit) - Generate random fake data (sentences, phrases, words)
- [godotenv](https://github.com/joho/godotenv) - Load environment variables from .env file

## License

MIT

GOOS=linux GOARCH=amd64 go build -o auto-commit .
