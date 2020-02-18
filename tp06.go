/* NIM : 1301194086
   Nama : Mikogizka Satria Kartika
*/
package main

import "fmt"

func fungsiF(x float64) float64 {
	return x * x
	
}
func fungsiG(x float64) float64 {
	return x - 2

}
func fungsiH(x float64) float64 {
	return x + 1
}
func fungsiFoGoH(x float64) float64 {
	return x * x - (x - 2 + x + 1)
}
func main() {
	var x float64
	var f string
	fmt.Println("masukan nilai x : ")
	fmt.Scanln(&x)
	fmt.Println(f)
	fmt.Println(fungsiF(x))
	fmt.Println(fungsiG(x))
	fmt.Println(fungsiH(x))
	fmt.Println(fungsiFoGoH(x))
	
	fmt.Println("------------------------------------------------------------")
	fmt.Println("Nama: Mikogizka Satria Kartika")
	fmt.Println("NIM: 1301194086")
}