package main

import "fmt"

type pengguna struct {
	nama, password string
	totalPinjaman  float64
	tenor, umur    int
	status         bool
}

type admin struct {
	nama, password string
	Nomor          int
}

const max = 100

type data [max]pengguna
type adm [max]admin

func main() {
	menu1()
}

func menu1() {
	var pilihan, banyakPengguna, banyakAdmin int
	var p data
	var ad adm
	pilihan = -1
	for pilihan != 0 {
		fmt.Println("Menu 1")
		fmt.Println("1. Login")
		fmt.Println("2. Register")
		fmt.Println("0. Keluar")
		fmt.Print("Pilihan : ")
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
	var adm int
	fmt.Println("Menu Register")
	fmt.Println("1. Admin")
	fmt.Println("2. Pengguna")
	fmt.Print("Pilihan : ")
	fmt.Scan(&adm)
	if adm == 1 {
		fmt.Println("Nama :")
		fmt.Scan(&(*ad)[*bAdmin].nama)
		fmt.Println("Password : ")
		fmt.Scan(&(*ad)[*bAdmin].password)
		fmt.Println("Nomor HP : ")
		fmt.Scan(&(*ad)[*bAdmin].Nomor)
		*bAdmin += 1
	} else if adm == 2 {
		fmt.Println("Nama : ")
		fmt.Scan(&(*p)[*bPengguna].nama)
		fmt.Println("Password : ")
		fmt.Scan(&(*p)[*bPengguna].password)

		*bPengguna += 1

	} else {
		fmt.Print("Tidak Valid")
	}

}

func login(p *data, ad *adm, bPengguna, bAdmin *int) {
	var idx, i int
	var nama, password string
	var user, pinjol bool
	var role int
	role = -1

	for role != 0 {
		fmt.Println("=== Menu Login ===")
		fmt.Println("1. Login sebagai Admin")
		fmt.Println("2. Login sebagai Pengguna")
		fmt.Println("0. Keluar")
		fmt.Print("Pilih: ")
		fmt.Scan(&role)

		idx = -1
		user = false

		if role == 1 {
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
					menuAdm(&*p, idx, &(*bPengguna), &pinjol)
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
	var pilihan int
	var bayar float64
	pilihan = -1
	for pilihan != 0 {
		fmt.Println("=== Menu Pengguna ===")
		fmt.Println("1. Pinjaman Online")
		fmt.Println("2. Ubah Data")
		fmt.Println("3. Bayar Pinjaman")
		fmt.Println("4. Laporan Jumlah Pinjaman")
		fmt.Println("0. Keluar")
		fmt.Print("Pilih : ")
		fmt.Scan(&pilihan)
		switch pilihan {
		case 1:
			painjol(&*p, &*pinjol, &*bPengguna, &idx, &bayar)
		case 2:
			ubahUser(&*p, idx)
		case 3:
			membayar(&*p, idx, &bayar)
		case 4:
			laporan()
		default:
			fmt.Println("Input tidak tersedia")
		}
	}
}

func painjol(p *data, pinjol *bool, bPengguna, idx *int, bayar *float64) {
	//	var duid, bunga float64
	var pilih, tenor int
	var pinjem bool
	fmt.Println("Halo ", (*p)[*idx].nama)
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
		(*p)[*idx].totalPinjaman += 500000
		pinjem = true
	case 2:
		(*p)[*idx].totalPinjaman += 1000000
		pinjem = true
	case 3:
		(*p)[*idx].totalPinjaman += 2000000
		pinjem = true
	case 4:
		(*p)[*idx].totalPinjaman += 3000000
		pinjem = true
	case 5:
		(*p)[*idx].totalPinjaman += 5000000
		pinjem = true
	case 6:
		(*p)[*idx].totalPinjaman += 10000000
		pinjem = true
	case 0:
		return
	default:
		fmt.Println("Tidak ada pilihan tersebut")
	}
	if pinjem {
		fmt.Println("Batas Tenor adalah 36 bulan")
		for tenor != 6 && tenor != 12 && tenor != 18 && tenor != 24 && tenor != 36 {
			fmt.Println("Pilih Tenor Baru | 6 | 12 | 18 | 24 | 36 |")
			fmt.Scan(&tenor)
			if tenor == 6 || tenor == 12 || tenor == 18 || tenor == 24 || tenor == 36 {
				(*p)[*idx].tenor += tenor
				*bayar = bunga(*p, *idx)
				(*p)[*idx].totalPinjaman = *bayar
				if (*p)[*idx].tenor > 36 {
					(*p)[*idx].tenor = 36
				}
				*pinjol = true
			}
		}

		pinjem = false

	}
}

func ubahUser(p *data, idx int) {
	var pilih int
	pilih = -1
	for pilih != 0 {
		fmt.Println("Halo", (*p)[idx].nama, "Ingin Mengubah Data Apa?")
		fmt.Println("1.Nama")
		fmt.Println("2.Password")
		fmt.Println("0.Keluar")
		fmt.Print("Pilihan : ")
		fmt.Scan(&pilih)
		switch pilih {
		case 1:
			fmt.Println("Nama Baru :")
			fmt.Scan(&(*p)[idx].nama)
		case 2:
			fmt.Println("Password Baru :")
			fmt.Scan(&(*p)[idx].password)
		default:
			fmt.Println("Input Tidak Sesuai")
		}
	}
}

func membayar(p *data, idx int, bayar *float64) {
	var pilih, bulan, tenor int

	tenor = (*p)[idx].tenor
	for pilih != 2 {
		if (*p)[idx].totalPinjaman <= 0 {
			(*p)[idx].totalPinjaman = 0
			(*p)[idx].tenor = 0
		}
		fmt.Printf("Total Yang harus Dibayar adalah RP. %.2f Dengan tenor selama %d\n", bunga(*p, idx), (*p)[idx].tenor)
		fmt.Printf("Total Yang harus dibayarkan bulan ini adalah %.2f\n", *bayar/float64((*p)[idx].tenor))
		fmt.Println("Sudah Bayar ", bulan, " Bulan")
		fmt.Println("Bayar Sekarang?")
		fmt.Println("1. Ya")
		fmt.Println("2. Nanti Saja")
		fmt.Print("Pilihan : ")
		fmt.Scan(&pilih)
		switch pilih {
		case 1:
			if (*p)[idx].totalPinjaman > 0 {
				(*p)[idx].totalPinjaman -= *bayar / float64(tenor)
				(p)[idx].status = true
				bulan++
			} else {
				fmt.Println("Sudah Lunas")
			}
		case 2:
			return
		default:
			fmt.Println("Tidak Ada Pilihan Tersebut")
		}
	}

}

func laporan() {

}

func bunga(p data, idx int) float64 {
	var tenor int
	var bunga, total float64
	tenor = p[idx].tenor
	bunga = 9.0 / 100 / 12
	total = float64(p[idx].totalPinjaman) * bunga * float64(tenor)
	return float64(p[idx].totalPinjaman) + total
}

func menuAdm(p *data, idx int, bPengguna *int, pinjol *bool) {
	var pilihan int

	pilihan = -1
	for pilihan != 0 {
		fmt.Println("=== Menu Admin ===")
		fmt.Println("1. Ubah Data")
		fmt.Println("2. Hapus Data")
		fmt.Println("3. Cari Data")
		fmt.Println("4. Urutkan Pinjaman")
		fmt.Println("5. Urutkan Tenor")
		fmt.Println("0. Keluar")
		fmt.Println("Pilih : ")
		fmt.Scan(&pilihan)
		switch pilihan {
		case 1:
			ubah(&*p, &(*bPengguna))
		case 2:
			hapus(&*p, &*bPengguna)
		}
	}
}

func ubah(p *data, bPengguna *int) {
	var idx, i, tes, tenor int
	var nama string
	tes = -1
	idx = -1
	for tes != 0 {
		fmt.Println("Input Nama yang datanya ingin diubah")
		fmt.Scan(&nama)
		for i = 0; i < *bPengguna; i++ {
			if nama == (*p)[i].nama {
				idx = i
			}
		}
		if idx == -1 {
			fmt.Println("Data Tidak Ditemukan")
		} else {
			for tes != 0 {
				fmt.Println("Data Ditemukan ")
				fmt.Println("Nama : ", (*p)[idx].nama)
				fmt.Println("Total Pinjaman : ", (*p)[idx].totalPinjaman)
				fmt.Println("Tenor : ", (*p)[idx].tenor)
				fmt.Println("Data mana yang ingin diubah")
				fmt.Println("1.Nama")
				fmt.Println("2.Password")
				fmt.Println("3.Total Pinjaman")
				fmt.Println("4.Tenor")
				fmt.Println("0.Stop")
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
						fmt.Println("Pilih Tenor Baru | 6 | 12 | 18 | 24 |")
						fmt.Scan(&tenor)
						if tenor == 6 || tenor == 12 || tenor == 18 || tenor == 24 {
							(*p)[idx].tenor = tenor
						}
					}
				case 0:
					return
				}
			}
		}
	}
}

func hapus(p *data, bPengguna *int) {
	var nama string
	var idx, i int
	idx = -1

	fmt.Print("Masukkan nama pengguna yang ingin dihapus: ")
	fmt.Scan(&nama)

	for i = 0; i < *bPengguna; i++ {
		if (*p)[i].nama == nama && idx == -1 {
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
	fmt.Println("Data berhasil dihapus.")
}
