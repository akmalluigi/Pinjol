package main

import "fmt"

type pengguna struct {
	nama, status, password string
	totalPinjaman          float64
	tenor, umur            int
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
	pilihan = -1
	for pilihan != 0 {
		fmt.Println("=== Menu Pengguna ===")
		fmt.Println("1. Pinjaman Online")
		fmt.Println("2. Ubah Data")
		fmt.Println("3. Bayar Pinjaman")
		fmt.Println("4. Laporan Jumlah Pinjaman")
		fmt.Println("0. Keluar")
		fmt.Println("Pilih : ")
		fmt.Scan(&pilihan)
		switch pilihan {
		case 1:
			painjol(&*p, &*pinjol, &*bPengguna, &idx)
		case 2:
			ubahUser()
		case 3:
			bayar()
		case 4:
			laporan()
		default:
			fmt.Println("Input tidak tersedia")
		}
	}
}

func painjol(p *data, pinjol *bool, bPengguna, idx *int) {
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
	fmt.Println("Pilih :")
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
		for tenor != 6 && tenor != 12 && tenor != 18 && tenor != 24 {
			fmt.Println("Pilih Tenor Baru | 6 | 12 | 18 | 24 |")
			fmt.Scan(&tenor)
			if tenor == 6 || tenor == 12 || tenor == 18 || tenor == 24 {
				(*p)[*idx].tenor += tenor
				*pinjol = true
			}
		}

		pinjem = false

	}
}

func ubahUser() {

}

func bayar() {

}

func laporan() {

}

func menuAdm(p *data, idx int, bPengguna *int, pinjol *bool) {
	var pilihan int

	//	var ad admin
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
				fmt.Println("2.Total Pinjaman")
				fmt.Println("3.Tenor")
				fmt.Println("0.Stop")
				fmt.Scan(&tes)
				switch tes {
				case 1:
					fmt.Print("Masukkan Nama Baru : ")
					fmt.Scan(&(*p)[idx].nama)
				case 2:
					fmt.Print("Masukkan Pinjaman Baru : ")
					fmt.Scan(&(*p)[idx].totalPinjaman)
				case 3:
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
