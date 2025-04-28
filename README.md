# pleasurelove

Aplikasi web pleasure love.

## 📦 Dokumentasi Menjalankan Proyek

### 🛠️ Persiapan Sebelum Menjalankan

1. Pastikan Anda memiliki akses ke repository proyek ini.
2. Instal perangkat lunak berikut:
   - **Go** (jika menjalankan secara lokal)
   - **Docker** dan **Docker Compose** (jika menjalankan menggunakan Docker)
3. Siapkan file `.env` jika diperlukan, dan sesuaikan konfigurasinya.

---

### 🚀 Cara Menjalankan Aplikasi

#### 💡 Dengan Docker Compose

```bash
docker-compose up --build
```

Akses aplikasi di: [http://localhost:8080](http://localhost:8080)

#### 💡 Secara Lokal (Tanpa Docker)

```bash
cd /c:/pleasurelove/pleasurelove
go mod tidy
go run cmd/main.go
```

Akses aplikasi di: [http://localhost:8080](http://localhost:8080)

---

### 📛 Cara Menghentikan Aplikasi

#### Jika menggunakan Docker:

- Tekan `Ctrl+C` di terminal yang menjalankan Docker Compose.
- Untuk menghentikan container:
  ```bash
  docker-compose down
  ```
- Untuk menghentikan dan menghapus volume (termasuk data database):
  ```bash
  docker-compose down -v
  ```

#### Jika secara lokal:

- Tekan `Ctrl+C` di terminal yang menjalankan aplikasi Go.

---

### 🦞 Troubleshooting

- Periksa error di terminal saat menjalankan aplikasi.
- Pastikan semua dependensi telah di-install dan konfigurasi `.env` sudah benar.
- Jika error berlanjut, hubungi pengembang terkait.

---

### 🥪 Menjalankan Mode Debug

#### Untuk Unix / Mac / Git Bash:

```bash
DEBUG_MODE=true docker-compose up
```

#### Untuk PowerShell (Windows):

```powershell
$env:DEBUG_MODE="true"
docker-compose up
```

- `DEBUG_MODE=true`: Aplikasi akan berjalan dalam mode debug menggunakan Delve.

---

## 🔀 Git Branch Workflow

Proyek ini menggunakan struktur branch yang jelas untuk memudahkan kolaborasi dan deployment.

### 🧹 Struktur Branch

```
main               → branch utama untuk produksi
develop            → branch staging (testing sebelum ke production)

feature/*          → fitur baru
bugfix/*           → perbaikan bug minor
hotfix/*           → perbaikan cepat langsung ke production
```

Contoh branch:

- `feature/login-page`
- `bugfix/cart-total-wrong`
- `hotfix/fix-payment-crash`

---

### 🔧 Cara Membuat Branch Baru

Semua branch fitur dan bugfix **dibuat dari `main`**.

```bash
git checkout main
git checkout -b feature/nama-fitur
git push -u origin feature/nama-fitur
```

---

### 🔀 Alur Pengembangan

1. Buat branch dari `main` → `feature/*` atau `bugfix/*`
2. Lakukan coding dan commit
3. Merge langsung ke `main` jika sudah siap produksi
4. Setelah lulus uji QA/test → merge `feature/*` atau `bugfix/*` ke `main`
5. Tag versi jika perlu:

```bash
git tag v1.0.0
git push origin v1.0.0
```

---

### 📄 Catatan Tambahan

- Simpan semua penamaan branch dalam format lowercase + dash (`-`)
- Hindari membuat branch langsung dari `develop` kecuali memang dibutuhkan
- Dokumentasikan fitur-fitur besar di Notion atau `CHANGELOG.md`

---

📌 *Dokumentasi ini akan terus diperbarui sesuai perkembangan proyek.*

