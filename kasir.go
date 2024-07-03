package main

import "fmt"

const NMAX int = 1000

type barang struct {
	nama  string
	harga int
}

type tabBarang [NMAX]barang

var count int

func main() {
	var x, y, n int
	var p tabBarang
	var hari, member, status string

	x = menuUtama()

	for x != 1 && x != 3 {
		fmt.Println()
		fmt.Println("Tidak ada riwayat transaksi")
		fmt.Println()
		x = menuUtama()
	}
	for x != 3 {
		if x == 1 {
			y = menuTransaksi()
			if y == 1 {
				fmt.Print("Masukkan jumlah barang yang ingin ditambahkan: ")
				fmt.Scan(&n)
				for i := 0; i < n; i++ {
					masukanP(&p, &count)
					tampilkanRiwayat(&p, count)
				}
			} else if y == 2 {
				editP(&p, &count)
			} else if y == 3 {
				hapusP(&p, &count)
			} else if y == 4 {
				fmt.Println("Masukkan hari, status member, dan status pelajar/mahasiswa (contoh: Senin Ya Pelajar):")
				fmt.Scan(&hari, &member, &status)
				terapkanDiskon(&p, count, hari, member, status)
			} else if y == 5 {
				bayar(&p, count)
				x = 3 // Selesai belanja, keluar dari loop
			} else if y == 6 {
				x = menuUtama()
			}
		} else if x == 2 {
			fmt.Println("Riwayat Transaksi Dzaky Mart")
			tampilkanRiwayat(&p, count)
			x = menuUtama()
		}
	}
	if x == 3 {
		fmt.Println("Terima Kasih")
	}
}

func menuUtama() int {
	var pilih int

	fmt.Println("||====================================||")
	fmt.Println("||             DZAKY MART             ||")
	fmt.Println("||====================================||")
	fmt.Println("||              >> Menu               ||")
	fmt.Println("||------------------------------------||")
	fmt.Println("||  1.   Transaksi Baru               ||")
	fmt.Println("||  2.   Riwayat Transaksi            ||")
	fmt.Println("||  3.   Exit                         ||")
	fmt.Println("||------------------------------------||")
	fmt.Print("Pilih opsi (1/2/3): ")
	fmt.Scan(&pilih)
	return pilih
}

func menuTransaksi() int {
	var pilih int

	fmt.Println("||====================================||")
	fmt.Println("||             DZAKY MART             ||")
	fmt.Println("||====================================||")
	fmt.Println("||              >> Menu               ||")
	fmt.Println("||------------------------------------||")
	fmt.Println("||  1.   Masukan Produk               ||")
	fmt.Println("||  2.   Edit Produk                  ||")
	fmt.Println("||  3.   Hapus Produk                 ||")
	fmt.Println("||  4.   Cek Diskon                   ||")
	fmt.Println("||  5.   Bayar                        ||")
	fmt.Println("||  6.   Menu Utama                   ||")
	fmt.Println("||------------------------------------||")
	fmt.Print("Pilih opsi (1/2/3/4/5/6): ")
	fmt.Scan(&pilih)
	return pilih
}

func masukanP(A *tabBarang, n *int) {
	fmt.Println("Masukan nama barang:")
	fmt.Scan(&A[*n].nama)
	fmt.Println("Masukan harga barang:")
	fmt.Scan(&A[*n].harga)
	*n++
}

func editP(A *tabBarang, n *int) {
	var newNama, targetNama, edit string
	var newHarga, targetHarga int
	fmt.Println("Masukan nama barang yang ingin diedit:")
	fmt.Scan(&targetNama)
	fmt.Println("Masukan harga barang yang ingin diedit:")
	fmt.Scan(&targetHarga)

	found := false
	for i := 0; i < *n; i++ {
		if A[i].nama == targetNama && A[i].harga == targetHarga {
			found = true
			fmt.Println("Barang ditemukan. Ingin mengubah apa?")
			fmt.Print("Nama / Harga: ")
			fmt.Scan(&edit)
			if edit == "Nama" {
				fmt.Print("Nama baru: ")
				fmt.Scan(&newNama)
				A[i].nama = newNama
			} else if edit == "Harga" {
				fmt.Print("Harga baru: ")
				fmt.Scan(&newHarga)
				A[i].harga = newHarga
			}
			fmt.Println("Barang berhasil diubah")
			break
		}
	}
	if !found {
		fmt.Println("Barang tidak ditemukan")
	}
}

func hapusP(A *tabBarang, n *int) {
	var barang string
	found := false
	fmt.Print("Masukan nama barang yang ingin dihapus: ")
	fmt.Scan(&barang)
	for i := 0; i < *n; i++ {
		if A[i].nama == barang {
			found = true
			for j := i; j < *n-1; j++ {
				A[j] = A[j+1]
			}
			*n--
			fmt.Println("Barang berhasil dihapus")
			break
		}
	}
	if !found {
		fmt.Println("Barang tidak ditemukan")
	}
}

func tampilkanRiwayat(A *tabBarang, n int) {
	fmt.Println("Riwayat Barang yang Sudah Dimasukkan:")
	for i := 0; i < n; i++ {
		fmt.Printf("Nama Barang: %s, Harga: %d\n", A[i].nama, A[i].harga)
	}
}

func terapkanDiskon(A *tabBarang, n int, hari, member, status string) {
	for i := 0; i < n; i++ {
		harga := A[i].harga
		if (hari == "Jumat" || hari == "Sabtu" || hari == "Minggu") && member == "Ya" {
			harga = harga * 80 / 100 // diskon 20%
		} else if (hari == "Senin" || hari == "Selasa" || hari == "Rabu" || hari == "Kamis") && member == "Ya" {
			harga = harga * 90 / 100 // diskon 10%
		} else if (hari == "Senin" || hari == "Selasa" || hari == "Rabu" || hari == "Kamis") && (status == "Pelajar" || status == "Mahasiswa") {
			harga = harga * 85 / 100 // diskon 15% untuk pelajar dan mahasiswa
		}
		A[i].harga = harga
		fmt.Printf("Nama Barang: %s, Harga Setelah Diskon: %d\n", A[i].nama, harga)
	}
}

func bayar(A *tabBarang, n int) {
	var total, uang, kembalian int

	for i := 0; i < n; i++ {
		total += A[i].harga
	}

	fmt.Printf("Total belanja: %d\n", total)
	fmt.Print("Masukkan jumlah uang: ")
	fmt.Scan(&uang)

	if uang >= total {
		kembalian = uang - total
		fmt.Printf("Kembalian: %d\n", kembalian)
		fmt.Println("Terima kasih sudah berbelanja. Semoga Anda puas berbelanja di Dzaky Mart!")
	} else {
		fmt.Println("Uang tidak cukup untuk membayar total belanja.")
	}
}

