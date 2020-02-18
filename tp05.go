/* NIM : 1301194086
   Nama : Mikogizka Satria Kartika
   Hasil PI repeater
*/
package main
import "fmt"
func main () {
	var pi,pmb, pny, pch float32
	var n, i int
	
	
	i = 0
	pi = 0
	pmb = 1
	pny = 1
	fmt.Print("N suku pertama: ")
	fmt.Scanln(&n)
	
	for i != n {
		i = i+1
		if i == 1 {
			pi = pi + 1
		} else if i%2 == 0 {
			pi = pi - pch
		} else {
			pi = pi + pch
		}
		
		pny = pny + 2
		pch = pmb/pny
	}	
		pi = pi*4
		fmt.Println("Hasil PI: ", pi)
	
	fmt.Println("------------------------------------------------------------")
	fmt.Println("Nama: Mikogizka Satria Kartika")
	fmt.Println("NIM: 1301194086")
	
	}		