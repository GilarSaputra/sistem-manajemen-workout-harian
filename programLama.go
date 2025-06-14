// package main

// import "fmt"

// // Definisikan struct Latihan
// type latihan struct {
// 	jenis  string
// 	durasi int
// 	kalori int
// }

// // Slice untuk menyimpan semua data latihan
// var daftarLatihan []latihan

// // Fungsi untuk menambahkan data ke slice
// func tambahLatihan(jenis string, durasi, kalori int) {
// 	inputLatihan := latihan{
// 		jenis:  jenis,
// 		durasi: durasi,
// 		kalori: kalori,
// 	}

// 	daftarLatihan = append(daftarLatihan, inputLatihan)
// }

// // Fungsi untuk mengubah data latihan
// func ubahLatihan(nomor int, jenis string, durasi, kalori int) {
// 	if nomor < 0 || nomor >= len(daftarLatihan) {
// 		fmt.Println("Nomor latihan tidak ditemukan.")
// 		return
// 	}

// 	daftarLatihan[nomor] = latihan{
// 		jenis:  jenis,
// 		durasi: durasi,
// 		kalori: kalori,
// 	}

// 	fmt.Println("Data latihan berhasil diubah.")
// }

// // Fungsi untuk menghapus data latihan
// func hapusLatihan() {
// 	if len(daftarLatihan) == 0 {
// 		fmt.Println("Tidak data latihan.")
// 		return
// 	}

// 	fmt.Println("\nDaftar latihan yang ada:")
// 	for i := 0; i < len(daftarLatihan); i++ {
// 		fmt.Println(i+1, ". Jenis:", daftarLatihan[i].jenis,
// 			"| Durasi:", daftarLatihan[i].durasi, "menit",
// 			"| Kalori:", daftarLatihan[i].kalori, "kalori")
// 	}

// 	var nomorLatihan int
// 	fmt.Print("\nMasukkan nomor latihan yang ingin dihapus: ")
// 	fmt.Scan(&nomorLatihan)
// 	if nomorLatihan < 1 || nomorLatihan > len(daftarLatihan) {
// 		fmt.Println("Nomor yang dimasukkan tidak ada.")
// 		return
// 	}

// 	for i := nomorLatihan - 1; i < len(daftarLatihan)-1; i++ {
// 		daftarLatihan[i] = daftarLatihan[i+1]
// 	}
// 	daftarLatihan = daftarLatihan[:len(daftarLatihan)-1]
// 	fmt.Println("\nData latihan berhasil dihapus.")
// }

// // Fungsi untuk menampilkan semua data latihan
// func tampilkanLatihan() {
// 	fmt.Println("DATA LATIHAN:")
// 	for i := 0; i < len(daftarLatihan); i++ {
// 		fmt.Printf("%d. Jenis: %s | Durasi: %d menit | Kalori: %d kalori\n", i+1, daftarLatihan[i].jenis, daftarLatihan[i].durasi, daftarLatihan[i].kalori)
// 	}
// }

// // Fungsi untuk mencari jenis latihan
// func cariJenisLatihan(LatihanYangDicari string) {
// 	Ada := false
// 	fmt.Println("\nHasil yang ditemukan")

// 	for i := 0; i < len(daftarLatihan); i++ {
// 		if daftarLatihan[i].jenis == LatihanYangDicari {
// 			fmt.Print(i+1, ". Jenis: ")
// 			fmt.Print(daftarLatihan[i].jenis)
// 			fmt.Print(" | Durasi: ")
// 			fmt.Print(daftarLatihan[i].durasi)
// 			fmt.Print(" menit | Kalori: ")
// 			fmt.Print(daftarLatihan[i].kalori)
// 			fmt.Println(" kalori")

// 			Ada = true
// 		}
// 	}

// 	if Ada == false {
// 		fmt.Println("Tidak ada latihan yang sesuai")
// 	}
// }

// // Fungsi untuk menampilkan statistik
// func statistik() {
// 	totalLatihan := len(daftarLatihan)
// 	totalDurasi := 0
// 	totalKalori := 0

// 	for i := 0; i < len(daftarLatihan); i++ {
// 		totalDurasi += daftarLatihan[i].durasi
// 		totalKalori += daftarLatihan[i].kalori
// 	}

// 	fmt.Println()
// 	fmt.Println("Statistik Data Latihan Anda ")
// 	fmt.Printf("Latihan yang sudah anda lakukan : %d latihan \n", totalLatihan)
// 	fmt.Printf("Total durasi yang anda lakukan selama latihan : %d menit \n", totalDurasi)
// 	fmt.Printf("Kalori yang sudah terbakar oleh tubuh : %d kalori \n", totalKalori)
// }

// func main() {
// 	var input int

// 	for {
// 		fmt.Println("\nPilih Opsi dibawah ini (nomor 1-9) :")
// 		fmt.Println("1. Tambah Data Latihan")
// 		fmt.Println("2. Ubah Data Latihan")
// 		fmt.Println("3. Hapus Data Latihan")
// 		fmt.Println("4. Tampilkan Data Latihan")
// 		fmt.Println("5. Cari Jenis Latihan")
// 		fmt.Println("6. Urutkan Data Latihan")
// 		fmt.Println("7. Rekomendasi Latihan")
// 		fmt.Println("8. Tampilkan Statistik")
// 		fmt.Println("9. Keluar")

// 		fmt.Print("Pilihan Anda: ")
// 		fmt.Scan(&input)

// 		switch input {
// 		case 1:
// 			var jenis string
// 			var durasi, kalori int
// 			fmt.Print("Masukkan jenis olahraga: ")
// 			fmt.Scan(&jenis)
// 			fmt.Print("Masukkan durasi (menit): ")
// 			fmt.Scan(&durasi)
// 			fmt.Print("Masukkan jumlah kalori terbakar: ")
// 			fmt.Scan(&kalori)

// 			tambahLatihan(jenis, durasi, kalori)

// 		case 2:
// 			if len(daftarLatihan) == 0 {
// 				fmt.Println("Tidak ada data yang diubah.")
// 				break
// 			}
// 			var nomor, durasi, kalori int
// 			var jenis string
// 			fmt.Print("Pilih nomor data latihan yang ingin diubah: ")
// 			fmt.Scan(&nomor)
// 			fmt.Print("Masukkan jenis latihan yang baru: ")
// 			fmt.Scan(&jenis)
// 			fmt.Print("Masukkan durasi: ")
// 			fmt.Scan(&durasi)
// 			fmt.Print("Masukkan kalori: ")
// 			fmt.Scan(&kalori)
// 			ubahLatihan(nomor-1, jenis, durasi, kalori)

// 		case 3:
// 			hapusLatihan()

// 		case 4:
// 			fmt.Println()
// 			if len(daftarLatihan) > 0 {
// 				tampilkanLatihan()
// 			} else {
// 				fmt.Println("DATA LATIHAN:")
// 				fmt.Println("Tidak ada data yang ditemukan!!")
// 			}

// 		case 5:
// 			if len(daftarLatihan) == 0 {
// 				fmt.Println("Belum ada data latihan")
// 				break
// 			}
// 			var LatihanYangDicari string
// 			fmt.Print("Masukkan jenis latihan yang ingin dicari: ")
// 			fmt.Scan(&LatihanYangDicari)
// 			cariJenisLatihan(LatihanYangDicari)

// 		case 6:
// 			fmt.Println("Fitur urutkan belum dibuat.")
// 		case 7:
// 			fmt.Println("Fitur rekomendasi belum dibuat.")
// 		case 8:
// 			statistik()
// 		case 9:
// 			fmt.Println("Terima Kasih.")
// 			return
// 		}
// 	}
// }
