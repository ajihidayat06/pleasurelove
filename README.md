# boilerplate
## Dokumentasi Menjalankan Proyek
### Persiapan Sebelum Menjalankan
1. Pastikan Anda memiliki akses ke repository proyek ini.
2. Instal perangkat lunak berikut di mesin Anda:
    - **Docker** dan **Docker Compose** (jika menjalankan dengan Docker).
    - **Go** (jika menjalankan secara lokal).
3. Periksa apakah Anda memiliki file `.env` (jika diperlukan) dan sesuaikan konfigurasi sesuai kebutuhan.

### Cara Menjalankan
Ikuti salah satu metode berikut untuk menjalankan proyek:

#### Dengan Docker Compose
- Jalankan perintah berikut di terminal:
  ```bash
  docker-compose up --build
  ```
- Akses aplikasi di `http://localhost:8080`.

#### Secara Lokal
- Navigasikan ke direktori proyek:
  ```bash
  cd /c:/pleasurelove/pleasurelove
  ```
- Instal dependensi:
  ```bash
  go mod tidy
  ```
- Jalankan aplikasi:
  ```bash
  go run main.go
  ```
- Akses aplikasi di `http://localhost:8080`.

### Cara Menghentikan
- **Docker Compose**: Tekan `Ctrl+C` di terminal yang menjalankan Docker Compose, lalu jalankan:
  ```bash
  docker-compose down
  ```
  **Catatan**: Perintah ini akan menghentikan container tanpa menghapus volume. Data di database akan tetap ada.

  Jika Anda ingin menghapus semua container **dan** volume (termasuk data database), gunakan opsi `-v`:
  ```bash
  docker-compose down -v
  ```
  **Peringatan**: Opsi `-v` akan menghapus semua data yang disimpan di volume Docker, termasuk data di database. Gunakan dengan hati-hati.

- **Secara Lokal**: Tekan `Ctrl+C` di terminal yang menjalankan aplikasi Go.

### Troubleshooting
- Periksa log terminal untuk pesan kesalahan.
- Pastikan semua dependensi telah diinstal dan konfigurasi sudah benar.
- Hubungi pengembang jika masalah berlanjut.

### Catatan
- Pastikan untuk membaca file `docker-compose.yml` atau file konfigurasi lainnya untuk detail lebih lanjut.
- Jika mengalami masalah, periksa kembali apakah semua dependensi dan variabel lingkungan telah diatur dengan benar.
- Dokumentasi tambahan dapat ditemukan di file README atau dokumentasi proyek lainnya.

## Menjalankan Aplikasi

#### Mode Normal
Gunakan perintah berikut untuk menjalankan aplikasi dalam mode normal:
```bash
docker-compose up
```

#### Mode Debug
Gunakan salah satu perintah berikut untuk menjalankan aplikasi dalam mode debug:

- **Untuk Shell Berbasis Unix (Linux/MacOS atau Git Bash di Windows)**:
  ```bash
  DEBUG_MODE=true docker-compose up
  ```

- **Untuk PowerShell (Windows)**:
  ```powershell
$env:DEBUG_MODE="true"
  docker-compose up
  ```

#### Penjelasan
- **DEBUG_MODE**: Variabel lingkungan yang menentukan mode aplikasi.
  - `false` (default): Menjalankan aplikasi secara normal.
  - `true`: Menjalankan aplikasi dalam mode debug menggunakan Delve.