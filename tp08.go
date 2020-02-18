/* NIM : 1301194086
   Nama : Mikogizka Satria Kartika
*/
package main

import "fmt"

const N=2019
type Rectype struct {
				f1 string
				f2 int
				f3 float64
			}
type ArrType [N]RecType
var arr ArrType

func rmax(data ArrType) int {
	var i int
	var big int
	i = 0
	var N = 2019
	big = 0
	for i <= N {
		if data[i].f3 != f3 {
			i++
		} else if data[i].f3 == f3 {
			if data[i] > big {
				big = data[i]
			}
		}
	}
	return big
}

func imin(data ArrType) int {
	var i int
	var lit int
	var N = 2019
	i = 0
	lit = N
	for i <= N {
		if data[i].f2 != f2 {
			i++
		} else if data[i].f2 == f2 {
			if data[i] < lit {
				lit = data[i]
			}
		}
	}
	return lit
}

func found(data ArrType, key string) bool {
	var i int
	var find int
	var N = 2019
	i = 0
	for i <= N {
		if data[i].f1 != key.f1 {
			i++
		} else if data[i].f1 == key.f1 {
			find = data[i]
		}
	}
	return find
}

func pos(data ArrType, key string) int {
	var idx int
	found = false
	atas = 0
	bawah = 2019
	for (atas <= bawah) && (!found) {
		tengah = (atas + bawah) / 2
		if key == data[tengah].key {
			found = true
			idx = tengah
		} else if kode < data[tengah] {
			bawah = tengah - 1
		} else if kode > data[tengah] {
			atas = tengah + 1
		} else {
			found = true
			idx = -1
		}
	}
	return idx
}

func main () {
isi := ""
	tambahArr()
	fmt.Println(rmax(arr))
	fmt.Println(imin(arr))
	fmt.Print("Cari f1 : ")
	fmt.Scanln(&isi)
	fmt.Println("ada ditemukan :", found(arr, isi))
	if pos(arr, isi) >= 0 {
		fmt.Println("Pada index ke", pos(arr, isi))
	} else {
		fmt.Println("Tidak ditemukan!")
	}
	
}