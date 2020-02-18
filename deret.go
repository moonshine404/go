/*NIM : 1301194301
  NAMA : ADIKA AKBAR SAPUTRA
  SOLUSI : pelajari formula leibnix for pi
*/
package main
import "fmt"

func main () {
	var x,pi,a,z float32
	var i,n int
	
	fmt.Println("Masukan suke ke N")
	fmt.Scanln(&n)
	i = 0
	pi = 0
	x = 1
	a = 1
	for i != n {
		i = i + 1
		if i == 1 {
			pi = pi + x
}		else if i % 2 == 0 {
			pi = pi - z
}		else {
			pi = pi + z
}
		a = a + 2
		z = x / a
}
	pi = pi*4 
	fmt.Println("N suku pertama :", n)
	fmt.Println("Hasil pi:", pi)
}