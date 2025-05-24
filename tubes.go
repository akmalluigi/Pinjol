package main

import (
	"fmt"
	"os"
	"os/exec"
)

type riwayat struct {
	pinjamID int
	jumlah   float64
	tenor    int
}
type pengguna struct {
	nama, password       string
	totalPinjaman, bayar float64
	tenor, id, bulan     int
	status               bool
	riwayatpinjam        [maxpin]riwayat
	jumlahriwayat        int
}

type admin struct {
	nama, password string
	id             int
}

const max = 100
const maxpin = 10

type data [max]pengguna
type adm [max]admin

func main() {
	menu1()
}

func menu1() {
	var pilihan, banyakPengguna, banyakAdmin int
	var bayar float64
	var p data
	var ad adm
	loaduser(&p, &banyakPengguna)
	loadadm(&ad, &banyakAdmin)
	for i := 0; i < banyakPengguna; i++ {
		bayar = bunga(p, i)
		p[i].bayar = bayar
		p[i].totalPinjaman = p[i].bayar
	}
	banyakAdmin++
	pilihan = -1
	for pilihan != 0 {
		fmt.Println("╔════════════════════════════════════════════════╗")
		fmt.Printf("║%27s %-19s ║\n", "PINJAMAN", "")
		fmt.Println("╠════════════════════════════════════════════════╣")
		fmt.Println("║〔1〕 Login                                     ║")
		fmt.Println("║〔2〕 Register                                  ║")
		fmt.Println("╚════════════════════════════════════════════════╝")
		fmt.Print(" Pilih opsi〔0 untuk keluar〕 : ")
		fmt.Scan(&pilihan)
		switch pilihan {
		case 1:
			login(&p, &ad, &banyakPengguna, &banyakAdmin)
		case 2:
			register(&p, &ad, &banyakPengguna, &banyakAdmin)
		case 0:
			return
		default:
			fmt.Println("Input Tidak Sesuai")

		}
		pilihan = -1
	}
}

func register(p *data, ad *adm, bPengguna, bAdmin *int) {
	clear()
	var adm, id int
	id = 10
	fmt.Println("╔════════════════════════════════════════════════╗")
	fmt.Printf("║%27s %-19s ║\n", "REGISTER", "")
	fmt.Println("╠════════════════════════════════════════════════╣")
	fmt.Printf("║%-27s %-19s ║\n", "〔1〕 Admin", "")
	fmt.Printf("║%-27s %-19s ║\n", "〔2〕 Pengguna", "")
	fmt.Println("╚════════════════════════════════════════════════╝")
	fmt.Print("Pilihan 〔0 untuk kembali〕: ")
	fmt.Scan(&adm)
	if adm == 1 {
		fmt.Print("Nama :")
		fmt.Scan(&ad[*bAdmin].nama)
		fmt.Print("Password : ")
		fmt.Scan(&ad[*bAdmin].password)
		id++
		ad[*bAdmin].id = id
		*bAdmin += 1

	} else if adm == 2 {
		fmt.Println("Nama : ")
		fmt.Scan(&(*p)[*bPengguna].nama)
		fmt.Println("Password : ")
		fmt.Scan(&(*p)[*bPengguna].password)
		id++
		(*p)[*bPengguna].id = id
		*bPengguna += 1

	} else if adm == 0 {
		return
	} else {
		fmt.Println("Tidak Valid")

	}

}

func login(p *data, ad *adm, bPengguna, bAdmin *int) {
	clear()
	var idx, i int
	var nama, password string
	var user, pinjol bool
	var role int
	role = -1

	for role != 0 {
		clear()
		fmt.Println("╔════════════════════════════════════════════════╗")
		fmt.Printf("║%27s %-19s ║\n", "LOGIN", "")
		fmt.Println("╠════════════════════════════════════════════════╣")
		fmt.Printf("║%-27s %-17s ║\n", "〔1〕Login sebagai Admin", "")
		fmt.Printf("║%-27s %-17s ║\n", "〔2〕Login sebagai Pengguna", "")
		fmt.Println("╚════════════════════════════════════════════════╝")
		fmt.Print("Pilihan 〔0 untuk kembali〕: ")
		fmt.Scan(&role)

		idx = -1
		user = false

		if role == 1 {
			clear()
			fmt.Println("╔════════════════════════════════════════════════╗")
			fmt.Printf("║%27s %-19s ║\n", "LOGIN", "")
			fmt.Println("╠════════════════════════════════════════════════╣")
			fmt.Printf("║%-27s %-8s ║\n", " Silahkan Masukan Data Yang Dibutuhkan", "")
			fmt.Println("╚════════════════════════════════════════════════╝")
			fmt.Print("Masukkan username admin: ")
			fmt.Scan(&nama)
			fmt.Print("Masukkan password admin: ")
			fmt.Scan(&password)

			for i = 0; i < *bAdmin; i++ {
				if nama == (*ad)[i].nama {
					idx = i
					user = true
					i = *bAdmin
				}
			}

			if user {
				if password == (*ad)[idx].password {
					menuAdm(&*p, &*ad, idx, &(*bPengguna), &pinjol)
				} else {
					fmt.Println("Password salah.")
				}
			} else {
				fmt.Println("Username admin tidak ditemukan.")
			}

		} else if role == 2 {
			fmt.Print("Masukkan username pengguna: ")
			fmt.Scan(&nama)
			fmt.Print("Masukkan password pengguna: ")
			fmt.Scan(&password)

			for i = 0; i < *bPengguna; i++ {
				if nama == (*p)[i].nama {
					idx = i
					user = true
					i = *bPengguna
				}
			}

			if user {
				if password == (*p)[idx].password {
					menuUser(&*p, idx, &(*bPengguna), &pinjol)
				} else {
					fmt.Println("Password salah.")
				}
			} else {
				fmt.Println("Username pengguna tidak ditemukan.")
			}

		} else if role == 0 {
			fmt.Println("Keluar dari menu login.")
			return
		} else {
			fmt.Println("Input tidak sesuai.")
		}
	}
}

func menuUser(p *data, idx int, bPengguna *int, pinjol *bool) {
	clear()
	var pilihan int
	var bayar float64
	pilihan = -1
	for pilihan != 0 {
		fmt.Println("╔════════════════════════════════════════════════════════════════════════════════════╗")
		fmt.Printf("║%50s %-33s║\n", "MENU PENGGUNA", "")
		fmt.Println("╠════════════════════════════════════════════════════════════════════════════════════╣")
		fmt.Printf("║%-48s %-33s║\n", "〔1〕 Ajukan Pinjaman", "")
		fmt.Printf("║%-48s %-33s║\n", "〔2〕 Pengaturan Akun", "")
		fmt.Printf("║%-48s %-33s║\n", "〔3〕 Bayar Pinjaman", "")
		fmt.Printf("║%-48s %-33s║\n", "〔4〕 Lihat Riwayat Pinjaman", "")
		fmt.Printf("║%-48s %-33s║\n", "〔0〕 Logout", "")
		fmt.Println("╚════════════════════════════════════════════════════════════════════════════════════╝")
		fmt.Print(" Pilih opsi (1-4): ")

		fmt.Scan(&pilihan)
		switch pilihan {
		case 1:
			painjol(p, pinjol, bPengguna, &idx, &bayar)
		case 2:
			ubahUser(p, idx)
		case 3:
			membayar(p, idx, &bayar)
		case 4:
			Laporan(p, idx)
		case 0:
			return
		default:
			fmt.Println("Input tidak tersedia")
		}
	}
}

func painjol(p *data, pinjol *bool, bPengguna, idx *int, bayar *float64) {
	clear()
	//	var duid, bunga float64
	var pilih, tenor int
	var pinjem bool
	var jumlah float64
	fmt.Println("Halo ", p[*idx].nama)
	fmt.Println("Butuh Pinjaman Berapa")
	fmt.Println("1. RP.500.000")
	fmt.Println("2. RP.1.000.000")
	fmt.Println("3. RP.2.000.000")
	fmt.Println("4. RP.3.000.000")
	fmt.Println("5. RP.5.000.000")
	fmt.Println("6. RP.10.000.000")
	fmt.Println("0. Keluar")
	fmt.Print("Pilih :")
	fmt.Scan(&pilih)
	switch pilih {
	case 1:
		jumlah = 500000
		pinjem = true
	case 2:
		jumlah = 1000000
		pinjem = true
	case 3:
		jumlah = 2000000
		pinjem = true
	case 4:
		jumlah = 3000000
		pinjem = true
	case 5:
		jumlah = 5000000
		pinjem = true
	case 6:
		jumlah = 10000000
		pinjem = true
	case 0:
		return
	default:
		fmt.Println("Tidak ada pilihan tersebut")
	}
	if pinjem {
		if p[*idx].jumlahriwayat >= maxpin {
			fmt.Println("Kapasitas Peminjaman Penuh!!!")
			return
		}
		fmt.Println("Batas Tenor adalah 36 bulan")
		for tenor != 6 && tenor != 12 && tenor != 18 && tenor != 24 && tenor != 36 {
			fmt.Println("Pilih Tenor Baru | 6 | 12 | 18 | 24 | 36 |")
			fmt.Scan(&tenor)
			if tenor == 6 || tenor == 12 || tenor == 18 || tenor == 24 || tenor == 36 {
				p[*idx].riwayatpinjam[p[*idx].jumlahriwayat] = riwayat{
					pinjamID: p[*idx].jumlahriwayat + 1,
					jumlah:   jumlah,
					tenor:    tenor,
				}
				p[*idx].jumlahriwayat++
				p[*idx].totalPinjaman += jumlah
				p[*idx].tenor = tenor
				p[*idx].bayar = bunga(*p, *idx)
				p[*idx].totalPinjaman = p[*idx].bayar
				if p[*idx].tenor > 36 {
					p[*idx].tenor = 36
				}
				*pinjol = true
			}
		}
		fmt.Printf("%s, Anda telah meminjam sebesar %v dengan tenor selama %d bulan\n", p[*idx].nama, (*p)[*idx].totalPinjaman, (*p)[*idx].tenor)
		pinjem = false

	}
}

func ubahUser(p *data, idx int) {
	clear()
	var pilih int
	pilih = -1
	for pilih != 0 {
		fmt.Println("Halo", p[idx].nama, "Ingin Mengubah Data Apa?")
		fmt.Println("1.Nama")
		fmt.Println("2.Password")
		fmt.Println("0.Keluar")
		fmt.Print("Pilihan : ")
		fmt.Scan(&pilih)
		switch pilih {
		case 1:
			fmt.Println("Nama Baru :")
			fmt.Scan(&p[idx].nama)
		case 2:
			fmt.Println("Password Baru :")
			fmt.Scan(&p[idx].password)
		default:
			fmt.Println("Input Tidak Sesuai")
		}
	}
}

func membayar(p *data, idx int, bayar *float64) {
	clear()
	var pilih, tenor int

	tenor = p[idx].tenor
	for pilih != 2 {
		if p[idx].totalPinjaman <= 0 {
			p[idx].totalPinjaman = 0
			p[idx].tenor = 0
			p[idx].bulan = 0
			*bayar = 0
			p[idx].status = true
		}
		clear()
		fmt.Printf("Total Yang harus Dibayar adalah RP. %.2f Dengan tenor selama %d bulan\n", p[idx].totalPinjaman, (*p)[idx].tenor)
		if p[idx].totalPinjaman <= 0 || p[idx].bulan == p[idx].tenor {
			fmt.Println("Pembayaran Sudah Lunas")
		} else {
			fmt.Printf("Total Yang harus dibayarkan bulan ini adalah Rp. %.2f\n", p[idx].bayar/float64(tenor))
		}
		if p[idx].totalPinjaman > 0 && p[idx].status != true {
			fmt.Println("Sudah Bayar ", p[idx].bulan, " Bulan")
			fmt.Println("Bayar Sekarang?")
			fmt.Println("1. Ya")
			fmt.Println("2. Nanti Saja (kembali)")
			fmt.Print("Pilihan : ")
			fmt.Scan(&pilih)
		} else {
			clear()
			return
		}
		switch pilih {
		case 1:
			if p[idx].totalPinjaman > 0 && p[idx].status != true {
				p[idx].totalPinjaman -= *&(p)[idx].bayar / float64(tenor)

				p[idx].bulan++

			} else {
				fmt.Println("Sudah Lunas")
			}
		case 2:
			clear()
			return
		default:
			fmt.Println("Tidak Ada Pilihan Tersebut")
		}
	}

}

func Laporan(p *data, idx int) {
	clear()
	var x int
	fmt.Printf("Riwayat Pinjaman %s (ID: %d):\n", p[idx].nama, p[idx].id)
	if p[idx].jumlahriwayat == 0 {
		fmt.Println("Belum ada pinjaman.")
		fmt.Println("Tekan Enter untuk kembali...")
		fmt.Scanln()
		return
	}
	fmt.Printf("%-10s %-20s %-15s\n", "Loan ID", "Jumlah (Rp)", "Tenor (bulan)")
	for i := 0; i < p[idx].jumlahriwayat; i++ {
		loan := p[idx].riwayatpinjam[i]
		fmt.Printf("%-10d %-20.2f %-15d\n", loan.pinjamID, loan.jumlah, loan.tenor)
	}
	fmt.Printf("Total Pinjaman Sekarang (ditambah bunga): %.2f\n", p[idx].totalPinjaman)
	fmt.Println("ketik 0 untuk kembali...")
	fmt.Scan(&x)
	switch x {
	case 0:
		clear()
		return
	}
}

func bunga(p data, idx int) float64 {
	var bunga, total float64
	bunga = 9.0 / 100 / 12
	for i := 0; i < p[idx].jumlahriwayat; i++ {
		loan := p[idx].riwayatpinjam[i]
		total = loan.jumlah * bunga * float64(loan.tenor)
	}
	return total + p[idx].totalPinjaman
}

func menuAdm(p *data, ad *adm, idx int, bPengguna *int, pinjol *bool) {
	clear()
	var pilihan int

	pilihan = -1
	for pilihan != 0 {
		cetaktotal(*p, *bPengguna)
		fmt.Println("╔════════════════════════════════════════════════════════════════════════════════════╗")
		fmt.Printf("║%50s %-33s║\n", "MENU ADMINISTRATOR", "")
		fmt.Println("╠════════════════════════════════════════════════════════════════════════════════════╣")
		fmt.Printf("║%-48s %-33s║\n", "〔1〕 Ubah Data Peminjam", "")
		fmt.Printf("║%-48s %-33s║\n", "〔2〕 Ubah Pengaturan Akun", "")
		fmt.Printf("║%-48s %-33s║\n", "〔3〕 Hapus Data", "")
		fmt.Printf("║%-48s %-33s║\n", "〔4〕 Cari Data", "")
		fmt.Printf("║%-48s %-33s║\n", "〔5〕 Urut Data", "")
		fmt.Printf("║%-48s %-33s║\n", "〔0〕 Logout", "")
		fmt.Println("╚════════════════════════════════════════════════════════════════════════════════════╝")
		fmt.Print("Pilihan 〔0 untuk kembali〕: ")
		fmt.Scan(&pilihan)
		switch pilihan {
		case 1:
			ubah(&*p, &(*bPengguna))
		case 2:
			ubahAdm(&*ad, idx)
		case 3:
			hapus(&*p, &*bPengguna)
		case 4:
			caridata(*p, *bPengguna)
		case 5:
			clear()
			Urut(p, *bPengguna)
		}
	}
}

func ubah(p *data, bPengguna *int) {
	var idx, i, tes, tenor, id int
	var nama string
	if *bPengguna >= 1 {
		tes = -1
		idx = -1
		clear()
		for tes != 0 {
			cetaktotal(*p, *bPengguna)
			fmt.Print("Input Nama : ")
			fmt.Scan(&nama)
			fmt.Print("Input Id : ")
			fmt.Scan(&id)

			for i = 0; i < *bPengguna; i++ {
				if nama == (*p)[i].nama && id == (*p)[i].id {
					idx = i
				}
			}
			if idx == -1 {
				clear()
				fmt.Println("Data Tidak Ditemukan")
			} else {
				for tes != 0 {
					clear()
					fmt.Println("╔════════════════════════════════════════════════════════════════════════════════════╗")
					fmt.Printf("║%48s %-35s║\n", "DATA DTEMUKAN", "")
					fmt.Println("╠════════════════════════════════════════════════════════════════════════════════════╣")
					fmt.Printf("║ Nama %-12s: %-48s %-15s║\n", "", (*p)[idx].nama, "")
					fmt.Printf("║ Total Pinjaman %-2s: %-48.2f %-15s║\n", "", (*p)[idx].totalPinjaman, "")
					fmt.Printf("║ Tenor %-11s: %-48d %-15s║\n", "", (*p)[idx].tenor, "")
					fmt.Println("╠════════════════════════════════════════════════════════════════════════════════════╣")
					fmt.Printf("║ %-48s %-34s║\n", "Data mana yang ingin anda ubah", "")
					fmt.Println("╠════════════════════════════════════════════════════════════════════════════════════╣")
					fmt.Printf("║%-48s %-33s║\n", "〔1〕Nama", "")
					fmt.Printf("║%-48s %-33s║\n", "〔2〕Password", "")
					fmt.Printf("║%-48s %-33s║\n", "〔3〕Total Pinjaman", "")
					fmt.Printf("║%-48s %-33s║\n", "〔4〕Tenor", "")
					fmt.Printf("║%-48s %-33s║\n", "〔0〕Kembali", "")
					fmt.Println("╚════════════════════════════════════════════════════════════════════════════════════╝")
					fmt.Print("Pilihan 〔0 untuk kembali〕: ")
					fmt.Scan(&tes)
					switch tes {
					case 1:
						fmt.Print("Masukkan Nama Baru : ")
						fmt.Scan(&(*p)[idx].nama)
					case 2:
						fmt.Print("Masukkan Password Baru : ")
						fmt.Scan(&(*p)[idx].password)
					case 3:
						fmt.Print("Masukkan Pinjaman Baru : ")
						fmt.Scan(&(*p)[idx].totalPinjaman)
					case 4:
						for tenor != 6 && tenor != 12 && tenor != 18 && tenor != 24 {
							clear()
							fmt.Println("╔════════════════════════════════════════════════════════════════════════════════════╗")
							fmt.Printf("║%48s %-35s║\n", "DATA DTEMUKAN", "")
							fmt.Println("╠════════════════════════════════════════════════════════════════════════════════════╣")
							fmt.Printf("║ Nama %-12s: %-48s %-15s║\n", "", (*p)[idx].nama, "")
							fmt.Printf("║ Total Pinjaman %-2s: %-48.2f %-15s║\n", "", (*p)[idx].totalPinjaman, "")
							fmt.Printf("║ Tenor %-11s: %-48d %-15s║\n", "", (*p)[idx].tenor, "")
							fmt.Println("╠════════════════════════════════════════════════════════════════════════════════════╣")
							fmt.Printf("║ %-48s %-34s║\n", "Data mana yang ingin anda ubah", "")
							fmt.Println("╠════════════════════════════════════════════════════════════════════════════════════╣")
							fmt.Printf("║%-48s %-33s║\n", "〔1〕Nama", "")
							fmt.Printf("║%-48s %-33s║\n", "〔2〕Password", "")
							fmt.Printf("║%-48s %-33s║\n", "〔3〕Total Pinjaman", "")
							fmt.Printf("║%-48s %-33s║\n", "〔4〕Tenor", "")
							fmt.Printf("║%-48s %-33s║\n", "〔0〕Kembali", "")
							fmt.Println("╠════════════════════════════════════════════════════════════════════════════════════╣")
							fmt.Println("╠════════════════════════════════════════════════════════════════════════════════════╣")
							fmt.Printf("║ %-48s %-34s║\n", "Pilih Tenor Baru | 6 | 12 | 18 | 24 |", "")
							fmt.Println("╚════════════════════════════════════════════════════════════════════════════════════╝")
							fmt.Print("Pilihan 〔0 untuk kembali〕: ")
							fmt.Scan(&tenor)
							if tenor == 6 || tenor == 12 || tenor == 18 || tenor == 24 {
								(*p)[idx].tenor = tenor
							}
						}
					case 0:
						clear()
						return
					}
				}
			}
		}
	} else {
		fmt.Println("Belum ada Data Pengguna")
	}
}

func ubahAdm(ad *adm, idx int) {
	var pilih int
	pilih = -1
	for pilih != 0 {
		clear()
		fmt.Println("╔════════════════════════════════════════════════════════════════════════════════════╗")
		fmt.Printf("║ %-5s %-15s %-61s║\n", "Halo", (*ad)[idx].nama, "Ingin Mengubah Data Apa?")
		fmt.Println("╠════════════════════════════════════════════════════════════════════════════════════╣")
		fmt.Printf("║%-48s %-33s║\n", "〔1〕Nama", "")
		fmt.Printf("║%-48s %-33s║\n", "〔2〕Password", "")
		fmt.Printf("║%-48s %-33s║\n", "〔0〕Kembali", "")
		fmt.Println("╠════════════════════════════════════════════════════════════════════════════════════╣")

		fmt.Print("Pilihan : ")
		fmt.Scan(&pilih)
		switch pilih {
		case 1:
			fmt.Print("Nama Baru :")
			fmt.Scan(&(*ad)[idx].nama)
		case 2:
			fmt.Print("Password Baru :")
			fmt.Scan(&(*ad)[idx].password)
		default:
			fmt.Println("Input Tidak Sesuai")
		}
	}
}

func hapus(p *data, bPengguna *int) {
	clear()
	var nama string
	var idx, i, id int
	idx = -1
	cetaktotal(*p, *bPengguna)
	fmt.Println("╔════════════════════════════════════════════════════════════════════════════════════╗")
	fmt.Printf("║%48s %-35s║\n", "Masukkan Nama dan Id pengguna yang ingin dihapus", "")
	fmt.Println("╚════════════════════════════════════════════════════════════════════════════════════╝")

	fmt.Print("Masukkan Nama : ")
	fmt.Scan(&nama)
	fmt.Print("Masukkan ID : ")
	fmt.Scan(&id)

	for i = 0; i < *bPengguna; i++ {
		if (*p)[i].nama == nama && (*p)[i].id == id && idx == -1 {
			idx = i
		}
	}

	if idx == -1 {
		fmt.Println("Pengguna tidak ditemukan.")
		return
	}

	for i = idx; i < *bPengguna-1; i++ {
		(*p)[i] = (*p)[i+1]
	}

	*bPengguna--
	clear()
	fmt.Println("Data berhasil dihapus.")
}
func Urut(p *data, bPengguna int) {
	var pilih int
	cetaktotal(*p, bPengguna)
	fmt.Println("1. Urutkan Berdasarkan Jumlah Peminjaman")
	fmt.Println("2. Urutkan Berdasarkan Jumlah Tenor")
	fmt.Print("Data yang ingin diurutkan : ")
	fmt.Scan(&pilih)
	switch pilih {
	case 1:
		urutpinjam(p, bPengguna)
		clear()
	case 2:
		uruttenor(p, bPengguna)
		clear()
	}

}
func urutpinjam(p *data, bPengguna int) {
	var pass, pil int
	var temp pengguna
	pass = 1
	fmt.Print(" (1) untuk menaik, (2) untuk menurun : ")
	fmt.Scan(&pil)
	if pil == 1 {
		for pass < bPengguna {
			i := pass
			temp = p[pass]
			for i > 0 && temp.totalPinjaman < p[i-1].totalPinjaman {
				p[i] = p[i-1]
				i--
			}
			p[i] = temp
			pass++
		}
	} else if pil == 2 {
		for pass < bPengguna {
			i := pass
			temp = p[pass]
			for i > 0 && temp.totalPinjaman > p[i-1].totalPinjaman {
				p[i] = p[i-1]
				i--
			}
			p[i] = temp
			pass++
		}
	} else {
		fmt.Println("Invalid")
	}
}
func uruttenor(p *data, bPengguna int) {
	var pass, pil int
	var temp pengguna
	pass = 1
	fmt.Print(" (1) untuk menaik, (2) untuk menurun : ")
	fmt.Scan(&pil)
	if pil == 1 {
		for pass < bPengguna {
			i := pass
			temp = p[pass]
			for i > 0 && temp.tenor < p[i-1].tenor {
				p[i] = p[i-1]
				i--
			}
			p[i] = temp
			pass++
		}
	} else if pil == 2 {
		for pass < bPengguna {
			i := pass
			temp = p[pass]
			for i > 0 && temp.tenor > p[i-1].tenor {
				p[i] = p[i-1]
				i--
			}
			p[i] = temp
			pass++
		}
	} else {
		fmt.Println("Invalid")
	}
}
func clear() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
func caridata(p data, bPengguna int) {
	var x string
	var cek int
	cetaktotal(p, bPengguna)
	fmt.Print("Ketikan nama pengguna yang ingin anda cari: ")
	fmt.Scan(&x)
	carid(p, bPengguna, x, &cek)
	clear()
	if cek == -1 {
		fmt.Println("╔════════════════════════════════════════════════════════════════════════════════════╗")
		fmt.Printf("║%48s %-35s║\n", "Nama pengguna tidak ditemukan.", "")
		fmt.Println("╚════════════════════════════════════════════════════════════════════════════════════╝")
		fmt.Println("Tekan 0 untuk kembali...")
		var input int
		for {
			fmt.Scan(&input)
			if input == 0 {
				clear()
				return
			}
			fmt.Println("Input tidak valid, tekan 0 untuk kembali...")
		}
	} else {
		cetakcari(p, bPengguna, cek)
		fmt.Println("Tekan 0 untuk kembali...")
		var input int
		for {
			fmt.Scan(&input)
			if input == 0 {
				clear()
				return
			}
			fmt.Println("Input tidak valid, tekan 0 untuk kembali...")
		}
	}
}

func cetakcari(p data, bPengguna, cek int) {
	var status string
	fmt.Println("╔════════════════════════════════════════════════════════════════════════════════════╗")
	fmt.Printf("╠ %-5s║ %-20s║ %-20s║ %-15s║ %-15s╣\n", "ID", "Nama", "Total Peminjaman", "Tenor(bulan)", "Status")
	fmt.Println("╠════════════════════════════════════════════════════════════════════════════════════╣")
	if cek >= 0 && cek < bPengguna {
		if p[cek].status == false {
			status = "Belum Lunas"
		} else {
			status = "Sudah Lunas"
		}
		fmt.Printf("║ %-5d║ %-20s║ %-20.2f║ %-15d║ %-15s║\n", p[cek].id, p[cek].nama, p[cek].totalPinjaman, p[cek].tenor, status)
	}
	fmt.Println("╚════════════════════════════════════════════════════════════════════════════════════╝")
}

// Fungsi lain tetap sama seperti kode asli
func carid(p data, bPengguna int, x string, idx *int) {
	var left, right, mid int
	*idx = -1
	left = 0
	right = bPengguna - 1
	for left <= right {
		mid = (left + right) / 2
		if x < p[mid].nama {
			right = mid - 1
		} else if x > p[mid].nama {
			left = mid + 1
		} else {
			*idx = mid
			return
		}
	}
}

func cetaktotal(p data, bPengguna int) {
	fmt.Println("╔════════════════════════════════════════════════════════════════════════════════════╗")
	fmt.Printf("║%48s %-35s║\n", "Data Pinjaman", "")
	fmt.Println("╠════════════════════════════════════════════════════════════════════════════════════╣")

	fmt.Printf("╠ %-5s║ %-20s║ %-20s║ %-15s║ %-15s╣\n", "ID", "Nama", "Total Peminjaman", "Tenor(bulan)", "Status")
	for i := 0; i < bPengguna; i++ {
		status := "Belum Lunas"
		if p[i].status {
			status = "Sudah Lunas"
		}
		fmt.Printf("║ %-5d║ %-20s║ %-20.2f║ %-15d║ %-15s║\n", p[i].id, p[i].nama, p[i].totalPinjaman, p[i].tenor, status)
	}
	fmt.Println("╚════════════════════════════════════════════════════════════════════════════════════╝")
}
func dummyuser(nama, password string, totalpinjaman float64, tenor, id int, status bool) pengguna {
	return pengguna{nama: nama, password: password, totalPinjaman: totalpinjaman, tenor: tenor, id: id, status: status, riwayatpinjam: [maxpin]riwayat{
		{pinjamID: 1, jumlah: totalpinjaman, tenor: tenor}}, jumlahriwayat: 1}
}
func loaduser(A *data, bPengguna *int) {
	A[*bPengguna] = dummyuser("akmal", "1", 500000, 6, 1, false)
	*bPengguna++
	A[*bPengguna] = dummyuser("budi", "2", 750000, 12, 2, true)
	*bPengguna++
	A[*bPengguna] = dummyuser("citra", "3", 600000, 6, 3, false)
	*bPengguna++
	A[*bPengguna] = dummyuser("dian", "4", 850000, 12, 4, true)
	*bPengguna++
	A[*bPengguna] = dummyuser("eka", "5", 700000, 18, 5, false)
	*bPengguna++
}

func dummyadm(nama, password string, id int) admin {
	return admin{nama: nama, password: password, id: id}
}
func loadadm(A *adm, bAdmin *int) {
	A[*bAdmin] = dummyadm("adm", "1", 1)
	*bAdmin += 1
}
