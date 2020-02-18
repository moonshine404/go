package main

import "fmt"
// import(
// 	"bufio"
// 	"time"
// )

const N = 999

type anggota struct {
	nama   string
	alamat string
	umur   int
	poin    int
	id     int
}

type buku struct {
	aidi int
	stock int
}

type data1 [N]anggota
type data2 [N]buku

var kartuid data1
var idbuku data2
var volume int
var aydi int

func main() {
	var pilihan int
	var yt string
	
	volume = 0
	fmt.Print("Jalankan aplikasi Taman Baca? ya/tidak? ")
	fmt.Scan(&yt)
	for yt == "ya" && pilihan != 4 {
		fmt.Println("")
		fmt.Println("----------------Selamat Datang-------------------")
		fmt.Println("------di Aplikasi Kelola Anggota Taman Baca------")
		pilihan = menu(pilihan)
		fmt.Printf("\n")
		if pilihan == 1 {
			carianggota()
		} else if pilihan == 2 {
			tambah_anggota()
		} else if pilihan == 3 {
			outputanggota()
		} else if pilihan == 4 {
			fmt.Println("Program ditutup")
		} 
	}
	
	fmt.Println("Terima Kasih")
	fmt.Println("----------------------------")
	fmt.Println("| Mikogizka Satria Kartika |")
	fmt.Println("|            dan           |")
	fmt.Println("|        Aditya Afif       |")
	fmt.Println("----------------------------")
}

func menu(choose int)int {
	fmt.Println("Selamat datang! silahkan pilih menu yang anda tuju")
	fmt.Println("1. Log In")
	fmt.Println("2. Sign Up")
	fmt.Println("3. Tampilkan data Anggota Taman Baca ")
	fmt.Println("4. Keluar")
	fmt.Print("Pilih : ")
	fmt.Scan(&choose)
	return choose
}

func peminjaman(N int, idbuku data2, aydi int) {
	var atas, tengah, bawah, idx, i int
	var found bool
	var tanggal, bulan, tahun int
	var pemisah string
	
	// pinjam := time.now()
	// kembali:= hari.AddDate(0, 0, 7)
	fmt.Println("")
	fmt.Println("Masukkan ID Buku untuk meminjam: ")
	fmt.Scan(&aydi)
	found = false
	atas = 0
	bawah = 10

	idbuku[0].aidi=001
	idbuku[0].stock=2
	idbuku[1].aidi=002
	idbuku[1].stock=1
	idbuku[2].aidi=003
	idbuku[2].stock=4
	idbuku[3].aidi=004
	idbuku[3].stock=0
	idbuku[4].aidi=005
	idbuku[4].stock=1
	idbuku[5].aidi=006
	idbuku[5].stock=5
	idbuku[6].aidi=007
	idbuku[6].stock=1
	idbuku[7].aidi=010
	idbuku[7].stock=3
	idbuku[8].aidi=011
	idbuku[8].stock=1
	idbuku[9].aidi=012
	idbuku[9].stock=6
	for (atas<=bawah) && (!found) {
		tengah = (atas+bawah)/2
		if aydi == idbuku[tengah].aidi {
			found = true
			idx = tengah
		} else if aydi < idbuku[tengah].aidi {
			bawah = tengah -1
		} else {
			atas = tengah + 1
		}
	}
	
	if found && idbuku[i].stock >= 2{
		fmt.Println(idbuku[idx].aidi)
		if idbuku[idx].stock >= 2 {
			fmt.Println("OK")
			fmt.Println("Masukkan tanggal hari ini: ")
			fmt.Scan(&tanggal, &pemisah, &bulan, &pemisah, &tahun)
			
			if bulan==4 || bulan==6 || bulan==9 || bulan==11 {
				if tanggal < 24 {
						tanggal = (tanggal + 7)
						fmt.Println("Tanggal pengembaliannya adalah: ", tanggal, "-", bulan, "-", tahun)
						 
						i++ 
					}else if tanggal >= 24 {
						tanggal = (tanggal + 7) - 30
						bulan = bulan + 1
						fmt.Println("Tanggal pengembaliannya adalah: ", tanggal, "-", bulan, "-", tahun)
						 
						i++
					}


			}else if  bulan==1 || bulan==3 || bulan==5 || bulan==7 || bulan==8 || bulan==10 || bulan==12 {
				if tanggal <= 25 {
					tanggal = (tanggal + 7)
					fmt.Println("Tanggal pengembaliannya adalah: ", tanggal, "-", bulan, "-", tahun)
					 
					i++
				}else if tanggal >= 26 && bulan == 12 {
						tanggal = (tanggal + 7) - 31
						bulan =  1
						tahun = tahun + 1
						fmt.Println("Tanggal pengembaliannya adalah: ", tanggal, "-", bulan, "-", tahun)
						 
						i++
				}else if tanggal >= 26 && bulan < 12 {
						tanggal = (tanggal + 7) - 31
						bulan =  bulan + 1
						fmt.Println("Tanggal pengembaliannya adalah: ", tanggal, "-", bulan, "-", tahun)
						 
						i++
				}
			

			}else if bulan==2 && tahun%4==0 && tanggal <= 29 {
					if tanggal >= 24 {
						tanggal = (tanggal + 7) - 29
						bulan = bulan + 1
						fmt.Println("Tanggal pengembaliannya adalah: ", tanggal, "-", bulan, "-", tahun)
						 
						i++
					}else if tanggal < 24 {
						tanggal = tanggal + 7
						fmt.Println("Tanggal pengembaliannya adalah: ", tanggal, "-", bulan, "-", tahun)
						 
						i++
					}

			}else if bulan==2 && tahun%4!=0 && tanggal <= 28 {
				if tanggal >= 23 {
					tanggal = (tanggal + 7) - 28
					bulan = bulan + 1
					fmt.Println("Tanggal pengembaliannya adalah: ", tanggal, "-", bulan, "-", tahun)
					 
					i++
				}else if tanggal < 23 {
					tanggal = tanggal + 7
					fmt.Println("Tanggal pengembaliannya adalah: ", tanggal, "-", bulan, "-", tahun)
					 
					i++
					
				}
			}
		}
 	}else if idbuku[i].stock < 2 {
		fmt.Println("Buku tidak dapat dipinjam")
	}else if !found{
		fmt.Println("Buku tidak ditemukan")
	}
}


func tambah_anggota() {
	var nama, alamat string
	var ID, umur, i int

	fmt.Println("")
	fmt.Println("Silahkan mendaftar")
	fmt.Print("Masukkan Nama: ")
	fmt.Scan(&nama)
	fmt.Print("Masukkan alamat: ")
	fmt.Scan(&alamat)
	fmt.Print("Masukkan Umur: ")
	fmt.Scan(&umur)
	fmt.Print("Masukkan ID: ")
	fmt.Scan(&ID)
	for i < N && ID != 666 {
		kartuid[i].nama = nama
		kartuid[i].alamat = alamat
		kartuid[i].umur = umur
		kartuid[i].id = ID
		kartuid[i].poin = 0
		fmt.Println("Data anggota baru sudah terinput :)")
		fmt.Println("Untuk menyelesaikan proses pendaftaran anggota baru, masukkan ID 666")
		fmt.Println("") 
		fmt.Print("Masukkan ID: ")
		fmt.Scan(&ID)
		if ID != 666{
			fmt.Print("Masukkan Nama: ")
			fmt.Scan(&nama)
			fmt.Print("Masukkan alamat: ")
			fmt.Scan(&alamat)
			fmt.Print("Masukkan Umur: ")
			fmt.Scan(&umur)
			fmt.Print("Masukkan ID: ")
			fmt.Scan(&ID)
			}
	i++
	volume++
	}
	fmt.Println("Anggota baru telah ditambahkan")
	fmt.Println("Proses tambah anggota telah selesai")
	sortanggota(&kartuid, N)
}

func cariident(id int)int{
	i:=0
	found:=false
	for i < volume && !found{
		if kartuid[i].id == id {
			found=true
			kartuid[i].poin = kartuid[i].poin + 5
			return i
		}
		i++
	}
	return -1
}

func carianggota() {
	var ind, ID int
	fmt.Println("Silahkan Log In")
	fmt.Println("Masukkan ID anggota: ")
	fmt.Scan(&ID)
	ind=cariident(ID)
	if ind == -1 {
		fmt.Println("Anggota tidak ada")
		tambah_anggota()
	} else {
		fmt.Println("ID Anggota: ", kartuid[ind].id)
		fmt.Println("Nama Anggota: ", kartuid[ind].nama)
		fmt.Println("Alamat Anggota: ", kartuid[ind].alamat)
		fmt.Println("Umur Anggota: ", kartuid[ind].umur)
		fmt.Println("Poin Anggota: ", kartuid[ind].poin)
		peminjaman(N, idbuku, aydi)
	}
}

func sortanggota(kartuid *data1, N int){
	
	var temp anggota
	var point, pass, poinmax int

	pass = 0
	for pass <= N-1 {
		poinmax = pass
		point = pass+1
		for point<N {
			if (kartuid[poinmax].poin < kartuid[point].poin) {
				poinmax = point
			}
			point++
		}
		temp = kartuid[poinmax]
		kartuid[poinmax] = kartuid[pass]
		kartuid[pass] = temp
		pass++
	}
}

func outputanggota(){
	var i int
	i = 0
	fmt.Print("Data dan Daftar Anggota Taman Baca: ")
	for i < volume {
		sortanggota(&kartuid, N)
		fmt.Println("")
		fmt.Println("-------------------------")
		fmt.Println("ID: ", kartuid[i].id)
		fmt.Println("Nama: ", kartuid[i].nama)
		fmt.Println("Umur: ", kartuid[i].umur)
		fmt.Println("Alamat: ", kartuid[i].alamat)
		fmt.Println("Poin: ", kartuid[i].poin)
		fmt.Println("--------------------------")
		i++
	}
}