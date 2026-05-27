# Tugas Besar Algoritma Pemrograman - Aplikasi Toko Online (Topik 16)

Repository ini berisi implementasi kode program untuk Tugas Besar mata kuliah **CBK1HAB4 Algoritma Pemrograman** pada Semester Genap 2025/2026. Program ini dikembangkan menggunakan bahasa pemrograman **Go (Golang)** dengan mematuhi seluruh spesifikasi akademik dan aturan penulisan algoritma secara ketat.

## 👥 Identitas Mahasiswa
* **Anggota 1 / NIM:** Ali Andriansyah
* **Anggota 2 / NIM:** Raka 
* **Program Studi:** S1 Teknologi Informasi
* **Fakultas:** Informatika, Universitas Telkom

---

## 📝 Deskripsi Aplikasi
**Aplikasi Toko Online** ini dirancang untuk memfasilitasi pemilik toko (admin) dalam mengelola katalog barang dagangan serta memproses transaksi penjualan dari pembeli secara terstruktur. Aplikasi ini berbasis *Command Line Interface* (CLI) dan mengimplementasikan penyimpanan data berbasis memori lokal menggunakan struktur data array statis.

---

## ✨ Fitur Utama Aplikasi

1. **Manajemen Katalog Barang (CRUD)**
   * **Tambah Barang:** Menambahkan produk baru (ID, Nama, Harga, Stok) ke dalam sistem. Setelah ditambahkan, data otomatis diurutkan agar siap dicari kapan saja.
   * **Ubah Barang (Edit):** Memperbarui informasi nama, harga, atau stok barang yang sudah terdaftar berdasarkan ID uniknya.
   * **Hapus Barang (Delete):** Menghapus produk dari katalog dengan mekanisme pergeseran elemen array statis untuk menjaga kontinuitas data.

2. **Sistem Antrean & Persetujuan Transaksi (Approval System)**
   * **Buat Transaksi:** Pembeli dapat melakukan pemesanan dengan memasukkan ID Barang dan jumlah kuantitas. Transaksi yang baru dibuat akan masuk ke dalam antrean dengan status *Menunggu Persetujuan*.
   * **Persetujuan Transaksi (Approve):** Admin dapat menyetujui transaksi yang menggantung. Sistem akan secara otomatis memeriksa ketersediaan stok di gudang. Jika stok mencukupi, status transaksi berubah menjadi *Disetujui* dan stok barang akan berkurang secara otomatis.

3. **Kalkulasi Finansial Otomatis**
   * **Hitung Omzet Toko:** Menghitung total seluruh pendapatan yang diperoleh hanya dari transaksi-transaksi yang telah disetujui (Approved) oleh admin.

4. **Visualisasi Data Terurut**
   * Menampilkan katalog barang dengan dua pilihan mode pengurutan dinamis: berdasarkan ID secara membesar (*Ascending*) atau berdasarkan sisa stok secara mengecil (*Descending*).

---

## 🛠️ Keselarasan Materi Perkuliahan (ALPRO)

Aplikasi ini dibangun murni menggunakan konsep-konsep dasar yang diajarkan dari Pekan 02 hingga Pekan 12:

* **Pekan 02 & 03 (Function & Procedure):** Pemisahan modular yang tegas. Subprogram pencarian dan kalkulasi diimplementasikan sebagai `Function` (mengembalikan nilai), sedangkan operasi CRUD, pengurutan, dan manajemen antrean menggunakan `Procedure` dengan manipulasi data langsung menggunakan *pointer passing* (`*`).
* **Pekan 04 (Rekursif - Nilai Tambah):** Perhitungan total omzet toko diselesaikan menggunakan fungsi rekursif murni (`hitungPendapatanRekursif`) tanpa menggunakan perulangan konvensional, yang memenuhi komponen nilai kreativitas tambahan pada tugas besar.
* **Pekan 05 (Tipe Bentukan dan Array):** Seluruh data diikat menggunakan tipe bentukan `struct` (`Barang` dan `Transaksi`) dan disimpan di dalam Array statis berkapasitas tetap (`[100]`), tanpa memanfaatkan *slice* atau fungsi dinamis bawaan Go.
* **Pekan 06 (Pencarian Nilai Ekstrem):** Logika pencarian nilai maksimum diintegrasikan di dalam algoritma pengurutan untuk menentukan barang dengan stok terbesar pada setiap tahapan iterasi.
* **Pekan 09 (Sequential Search):** Digunakan untuk mencari data transaksi berdasarkan ID secara linear pada fitur persetujuan admin.
* **Pekan 10 (Binary Search):** Digunakan untuk mencari data barang secara efisien dengan kompleksitas $O(\log n)$. Algoritma ini berjalan menggunakan kontrol logika boolean konjungsi murni tanpa melibatkan statemen *break*.
* **Pekan 11 (Selection Sort):** Digunakan untuk menyusun daftar barang secara menurun (*Descending*) berdasarkan jumlah stok terkini.
* **Pekan 12 (Insertion Sort):** Digunakan untuk menyusun katalog barang secara menaik (*Ascending*) berdasarkan ID Barang guna menjamin kevalidan prasyarat algoritma *Binary Search*.

⚠️ **Kepatuhan Aturan Ketat:** Kode program ini **sama sekali tidak menggunakan** statemen keyword `break` maupun `continue` di dalam seluruh struktur perulangan yang ada.

---

## 🚀 Cara Menjalankan Program

1. Pastikan perangkat Anda telah terinstal compiler **Go (Golang)**.
2. Unduh atau clone repository ini, lalu masuk ke direktori proyek.
3. Jalankan perintah berikut pada Terminal / Command Prompt Anda:
   ```bash
   go run main.go