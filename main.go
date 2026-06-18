// TUGAS BESAR ALGORITMA PEMROGRAMAN
// TOPIK 16: APLIKASI TOKO ONLINE
// Dikerjakan oleh:
// Ali Andriansyah
// RAKA SATRIA HAFIZ

package main

import "fmt"

// array
const MAX = 100

type Barang struct {
	ID    int
	Nama  string
	Harga int
	Stok  int
}

type Transaksi struct {
	ID        int
	IDBarang  int
	Pembeli   string
	Jumlah    int
	Total     int
	Disetujui bool
}

type ArrBarang [MAX]Barang
type ArrTransaksi [MAX]Transaksi

var dataBarang ArrBarang
var nBarang int

var dataTransaksi ArrTransaksi
var nTransaksi int

// function binary search untuk mencari indeks data barang dalam array secara cepat menggunakan metode pencarian biner.
func binarySearchBarang(T ArrBarang, n int, idCari int) int {
	var kiri, kanan, tengah, ketemu int
	kiri = 0
	kanan = n - 1
	ketemu = -1

	for kiri <= kanan && ketemu == -1 {
		tengah = (kiri + kanan) / 2
		if T[tengah].ID == idCari {
			ketemu = tengah
		} else if idCari < T[tengah].ID {
			kanan = tengah - 1
		} else {
			kiri = tengah + 1
		}
	}
	return ketemu
}

// function sequential search untuk mencari indeks transaksi dalam array secara berurutan atau sekuensial dari awal sampai akhir.
func seqSearchTransaksi(T ArrTransaksi, n int, idCari int) int {
	var i, ketemu int
	ketemu = -1
	i = 0

	for i < n && ketemu == -1 {
		if T[i].ID == idCari {
			ketemu = i
		}
		i++
	}
	return ketemu
}

// function rekursif untuk menghitung total pendapatan atau omzet penjualan dari transaksi yang sudah berhasil disetujui oleh admin.
func hitungPendapatanRekursif(T ArrTransaksi, n int, i int) int {
	if i >= n {
		return 0
	}
	if T[i].Disetujui {
		return T[i].Total + hitungPendapatanRekursif(T, n, i+1)
	}
	return hitungPendapatanRekursif(T, n, i+1)
}

// procedure insertion sort untuk mengurutkan daftar barang secara ascending atau menaik berdasarkan nomor ID barang secara teratur.
func insertionSortBarangByID(T *ArrBarang, n int) {
	var i, j int
	var temp Barang

	for i = 1; i < n; i++ {
		temp = T[i]
		j = i - 1
		for j >= 0 && T[j].ID > temp.ID {
			T[j+1] = T[j]
			j = j - 1
		}
		T[j+1] = temp
	}
}

// procedure selection sort & extreme value search untuk mengurutkan barang secara descending berdasarkan jumlah stok terbanyak.
func selectionSortBarangByStokDesc(T *ArrBarang, n int) {
	var i, j, idxMax int
	var temp Barang

	for i = 0; i < n-1; i++ {
		idxMax = i
		for j = i + 1; j < n; j++ {
			if T[j].Stok > T[idxMax].Stok {
				idxMax = j
			}
		}
		temp = T[i]
		T[i] = T[idxMax]
		T[idxMax] = temp
	}
}

// procedure buat nambah barang baru ke dalam daftar array barang toko online jika kapasitas array belum penuh.
func tambahBarang(T *ArrBarang, n *int) {
	var b Barang
	if *n < MAX {
		fmt.Print("Masukkan ID Barang      : ")
		fmt.Scan(&b.ID)
		fmt.Print("Masukkan Nama Barang    : ")
		fmt.Scan(&b.Nama)
		fmt.Print("Masukkan Harga Barang   : ")
		fmt.Scan(&b.Harga)
		fmt.Print("Masukkan Stok Barang    : ")
		fmt.Scan(&b.Stok)

		T[*n] = b
		*n = *n + 1

		insertionSortBarangByID(T, *n)
		fmt.Println(">> Barang berhasil ditambahkan!")
	} else {
		fmt.Println(">> Kapasitas array barang penuh!")
	}
}

// procedure edit data barang yang sudah ada untuk memperbarui nama, harga, dan stok barang berdasarkan ID tertentu.
func editBarang(T *ArrBarang, n int) {
	var idCari, idx int
	fmt.Print("Masukkan ID Barang yang akan diubah: ")
	fmt.Scan(&idCari)

	idx = binarySearchBarang(*T, n, idCari)
	if idx != -1 {
		fmt.Print("Masukkan Nama Baru  : ")
		fmt.Scan(&T[idx].Nama)
		fmt.Print("Masukkan Harga Baru : ")
		fmt.Scan(&T[idx].Harga)
		fmt.Print("Masukkan Stok Baru  : ")
		fmt.Scan(&T[idx].Stok)
		fmt.Println(">> Data barang berhasil diperbarui!")
	} else {
		fmt.Println(">> Barang tidak ditemukan.")
	}
}

// procedure buat hapus barang dari daftar array dengan menggeser posisi elemen setelah indeks barang yang dihapus tersebut.
func hapusBarang(T *ArrBarang, n *int) {
	var idCari, idx, i int
	fmt.Print("Masukkan ID Barang yang akan dihapus: ")
	fmt.Scan(&idCari)

	idx = binarySearchBarang(*T, *n, idCari)
	if idx != -1 {
		for i = idx; i < *n-1; i++ {
			T[i] = T[i+1]
		}
		*n = *n - 1
		fmt.Println(">> Barang berhasil dihapus!")
	} else {
		fmt.Println(">> Barang tidak ditemukan.")
	}
}

// procedure buat nampilin barang ke layar dengan pilihan pengurutan berdasarkan ID menaik atau jumlah stok menurun.
func tampilBarang(T *ArrBarang, n int) {
	var pil, i int
	fmt.Println("\nPilih Kriteria Penampilan Data:")
	fmt.Println("1. Ascending berdasarkan ID (Insertion Sort)")
	fmt.Println("2. Descending berdasarkan Stok (Selection Sort)")
	fmt.Print("Pilihan Anda: ")
	fmt.Scan(&pil)

	if pil == 1 {
		insertionSortBarangByID(T, n)
	} else if pil == 2 {
		selectionSortBarangByStokDesc(T, n)
	}

	fmt.Println("\n------------------- DAFTAR BARANG -------------------")
	for i = 0; i < n; i++ {
		fmt.Printf("ID: %d | Nama: %s | Harga: Rp%d | Stok: %d\n", T[i].ID, T[i].Nama, T[i].Harga, T[i].Stok)
	}
	fmt.Println("-----------------------------------------------------")

	if pil == 2 {
		insertionSortBarangByID(T, n)
	}
}

// procedure tambah transaksi baru untuk mencatat pemesanan barang oleh pembeli dan menambahkannya ke dalam antrean transaksi yang tersedia di sistem toko.
func tambahTransaksi(TTrans *ArrTransaksi, nTrans *int, TBar ArrBarang, nBar int) {
	var t Transaksi
	var idxBarang int

	if *nTrans < MAX {
		fmt.Print("Masukkan ID Transaksi : ")
		fmt.Scan(&t.ID)
		fmt.Print("Masukkan Nama Pembeli : ")
		fmt.Scan(&t.Pembeli)
		fmt.Print("Masukkan ID Barang    : ")
		fmt.Scan(&t.IDBarang)

		idxBarang = binarySearchBarang(TBar, nBar, t.IDBarang)
		if idxBarang != -1 {
			fmt.Print("Masukkan Kuantitas    : ")
			fmt.Scan(&t.Jumlah)

			t.Total = t.Jumlah * TBar[idxBarang].Harga
			t.Disetujui = false

			TTrans[*nTrans] = t
			*nTrans = *nTrans + 1
			fmt.Println(">> Transaksi masuk antrean. Menunggu persetujuan (Approve)!")
		} else {
			fmt.Println(">> ID Barang tidak terdaftar. Transaksi dibatalkan.")
		}
	} else {
		fmt.Println(">> Kapasitas array transaksi penuh!")
	}
}

// procedure approve transaksi untuk menyetujui pembelian barang sekaligus mengurangi jumlah stok barang di toko secara otomatis.
func setujuiTransaksi(TTrans *ArrTransaksi, nTrans int, TBar *ArrBarang, nBar int) {
	var idCari, idxTrans, idxBarang int
	fmt.Print("Masukkan ID Transaksi untuk di-Approve: ")
	fmt.Scan(&idCari)

	idxTrans = seqSearchTransaksi(*TTrans, nTrans, idCari)
	if idxTrans != -1 {
		if !TTrans[idxTrans].Disetujui {
			idxBarang = binarySearchBarang(*TBar, nBar, TTrans[idxTrans].IDBarang)
			if idxBarang != -1 {
				if TBar[idxBarang].Stok >= TTrans[idxTrans].Jumlah {
					TBar[idxBarang].Stok = TBar[idxBarang].Stok - TTrans[idxTrans].Jumlah
					TTrans[idxTrans].Disetujui = true
					fmt.Println(">> Transaksi Berhasil Disetujui! Stok barang otomatis dikurangi.")
				} else {
					fmt.Println(">> Stok barang tidak mencukupi. Persetujuan ditolak.")
				}
			}
		} else {
			fmt.Println(">> Transaksi ini telah disetujui sebelumnya.")
		}
	} else {
		fmt.Println(">> ID Transaksi tidak ditemukan.")
	}
}

// procedure nampilin transaksi ke layar untuk melihat status persetujuan dan rincian lengkap dari riwayat seluruh transaksi.
func tampilTransaksi(TTrans ArrTransaksi, nTrans int) {
	var i int
	var status string

	fmt.Println("\n------------------- DAFTAR TRANSAKSI -------------------")
	for i = 0; i < nTrans; i++ {
		if TTrans[i].Disetujui {
			status = "Telah Disetujui"
		} else {
			status = "Menunggu Persetujuan"
		}
		fmt.Printf("ID: %d | Brg: %d | Pembeli: %s | Qty: %d | Total: Rp%d | Status: %s\n",
			TTrans[i].ID, TTrans[i].IDBarang, TTrans[i].Pembeli, TTrans[i].Jumlah, TTrans[i].Total, status)
	}
	fmt.Println("--------------------------------------------------------")
}

// procedure cek total pendapatan dengan memanggil fungsi rekursif untuk menjumlahkan semua transaksi yang berstatus telah disetujui.
func cetakTotalPendapatan(TTrans ArrTransaksi, nTrans int) {
	var total int
	// Pemanggilan Fungsi Rekursif (Materi Pekan 04)
	total = hitungPendapatanRekursif(TTrans, nTrans, 0)
	fmt.Printf("\n>> Total Pendapatan dari Transaksi yang Disetujui: Rp%d\n", total)
}

// program utama untuk menginisialisasi data dummy awal dan menjalankan menu interaktif aplikasi toko online secara berulang.
func main() {
	var berjalan bool
	var pilihan int

	// data dummy

	// data barang terurut by ID
	dataBarang[0] = Barang{ID: 101, Nama: "Kemeja_Flanel", Harga: 150000, Stok: 20}
	dataBarang[1] = Barang{ID: 102, Nama: "Celana_Chino", Harga: 180000, Stok: 15}
	dataBarang[2] = Barang{ID: 103, Nama: "Kaos_Polos", Harga: 50000, Stok: 50}
	dataBarang[3] = Barang{ID: 104, Nama: "Jaket_Parka", Harga: 250000, Stok: 5}
	nBarang = 4

	// data transaksi
	dataTransaksi[0] = Transaksi{ID: 1, IDBarang: 101, Pembeli: "Ali", Jumlah: 2, Total: 300000, Disetujui: true}
	dataTransaksi[1] = Transaksi{ID: 2, IDBarang: 103, Pembeli: "Budi", Jumlah: 3, Total: 150000, Disetujui: false}
	dataTransaksi[2] = Transaksi{ID: 3, IDBarang: 104, Pembeli: "Citra", Jumlah: 1, Total: 250000, Disetujui: false}
	nTransaksi = 3

	berjalan = true

	// perulangan
	for berjalan {
		fmt.Println("\n=============================================")
		fmt.Println("      APLIKASI TOKO ONLINE (E-COMMERCE)      ")
		fmt.Println("=============================================")
		fmt.Println("1. Tambah Data Barang Baru")
		fmt.Println("2. Edit Data Barang (Update)")
		fmt.Println("3. Hapus Data Barang (Delete)")
		fmt.Println("4. Tampilkan Daftar Barang")
		fmt.Println("5. Buat Transaksi Penjualan")
		fmt.Println("6. Tampilkan Riwayat Transaksi")
		fmt.Println("7. Persetujuan Transaksi Penjualan")
		fmt.Println("8. Hitung Total Omzet (Fungsi Rekursif)")
		fmt.Println("0. Keluar Aplikasi")
		fmt.Print("Pilih menu operasi (0-8): ")
		fmt.Scan(&pilihan)

		// if else biar gak pake break
		if pilihan == 1 {
			tambahBarang(&dataBarang, &nBarang)
		} else if pilihan == 2 {
			editBarang(&dataBarang, nBarang)
		} else if pilihan == 3 {
			hapusBarang(&dataBarang, &nBarang)
		} else if pilihan == 4 {
			tampilBarang(&dataBarang, nBarang)
		} else if pilihan == 5 {
			tambahTransaksi(&dataTransaksi, &nTransaksi, dataBarang, nBarang)
		} else if pilihan == 6 {
			tampilTransaksi(dataTransaksi, nTransaksi)
		} else if pilihan == 7 {
			setujuiTransaksi(&dataTransaksi, nTransaksi, &dataBarang, nBarang)
		} else if pilihan == 8 {
			cetakTotalPendapatan(dataTransaksi, nTransaksi)
		} else if pilihan == 0 {
			berjalan = false
			fmt.Println("\nKeluar dari program. Terima kasih!")
		} else {
			fmt.Println("\n>> Pilihan Anda tidak terdaftar dalam menu.")
		}
	}
}
