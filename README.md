# Tubes-Alpro-TokoOnline

# Tugas Besar Algoritma Pemrograman - Aplikasi Toko Online

[cite_start]Repository ini berisi implementasi *source code* untuk Tugas Besar mata kuliah Algoritma Pemrograman (CBK1HAB4)[cite: 1]. [cite_start]Proyek ini disusun sebagai bagian dari pemenuhan tugas di Program Studi Teknologi Informasi, Fakultas Informatika[cite: 2, 3]. 

[cite_start]Aplikasi yang dikembangkan adalah **Aplikasi Toko Online** (Topik 16) yang ditulis sepenuhnya menggunakan bahasa pemrograman Go (Golang)[cite: 13, 192].

## 👥 Anggota Kelompok
[cite_start]Tugas ini dikerjakan secara berkelompok[cite: 10]:
* **Nama:** Ali Andriansyah
* **NIM:** [Isi NIM Anda di sini]
* *(Jika Anda memiliki rekan satu kelompok, tambahkan Nama dan NIM-nya di sini)* ## 📝 Deskripsi Aplikasi
[cite_start]Aplikasi ini dirancang untuk membantu pemilik toko online dalam mengelola data barang dan pembeli secara efisien[cite: 193]. [cite_start]Program berjalan berbasis *Command Line Interface* (CLI) dan dirancang secara modular menggunakan konsep *Subprogram* (Fungsi dan Prosedur)[cite: 35]. Data disimpan secara statis dalam memori selama program berjalan.

## ✨ Fitur Utama
Sesuai dengan spesifikasi tugas, aplikasi ini mencakup fitur-fitur fungsional berikut:
1. [cite_start]**Manajemen Barang:** Pemilik toko dapat melakukan penambahan, pengubahan (edit), dan penghapusan data barang yang dijual[cite: 196].
2. [cite_start]**Pengurutan Data:** Menampilkan data barang secara terurut berdasarkan jumlah stok atau kriteria lain (seperti ID)[cite: 197]. [cite_start]Pengurutan dapat dilakukan secara *ascending* maupun *descending*[cite: 40].
3. [cite_start]**Persetujuan Transaksi:** Fasilitas bagi pemilik toko untuk melakukan persetujuan ( *approve*) terhadap transaksi penjualan yang masuk[cite: 199].
4. [cite_start]**Kalkulasi Pendapatan:** Sistem dapat menghitung total uang yang diperoleh dari hasil transaksi[cite: 198].

## 🛠️ Algoritma & Struktur Data
* [cite_start]**Struktur Data:** Menggunakan implementasi Array statis dan Tipe Bentukan Struktur (`struct`) murni, tanpa *slice* atau *array* dinamis[cite: 37]. 
* [cite_start]**Pencarian Data (Searching)[cite: 38]:**
  * *Binary Search*: Diterapkan untuk mencari ID Barang, memanfaatkan efisiensi tinggi pada *array* yang sudah terurut.
  * *Sequential Search*: Diterapkan untuk mencari ID Transaksi.
* [cite_start]**Pengurutan Data (Sorting)[cite: 39]:**
  * *Insertion Sort*: Menjaga *array* Barang terurut secara *Ascending* berdasarkan ID.
  * *Selection Sort*: Digunakan saat user ingin menampilkan daftar barang secara *Descending* berdasarkan jumlah stok terbanyak.
* **Standar Kode Tambahan:**
  * [cite_start]Algoritma **Rekursif** diimplementasikan pada kalkulasi total pendapatan sebagai nilai tambahan kreativitas aplikasi[cite: 17, 41].
  * [cite_start]Tidak menggunakan *statement* `break` atau `continue` dalam perulangan apa pun[cite: 42].
  * [cite_start]Penggunaan variabel global dibatasi secara ketat hanya untuk *array* data utama[cite: 43].

## 🚀 Cara Menjalankan Program
1. Pastikan Anda sudah menginstal **Go (Golang)** di komputer Anda.
2. *Clone* repository ini ke lokal mesin Anda, atau unduh langsung file `main.go`.
3. Buka *Terminal* atau *Command Prompt*, dan arahkan direktori ke lokasi file `main.go`.
4. Jalankan perintah kompilasi berikut:
   ```bash
   go run main.go
