/*NIM : 1301194301
  NAMA : ADIKA AKBAR SAPUTRA
  SOLUSI : pelajari formula leibnix for pi
*/
package main
import "fmt"

func main () {
	var x,pi,a,z,y float32
	var i int
	y = 0
	i = 0
	pi = 0
	x = 1
	a = 1
	z = 0
	for x == 1 {
		i = i + 1
		y = pi
		if i == 1 {
			pi = pi + x
}		else if i % 2 == 0 {
			pi = pi - z
}		else {
			pi = pi + z
}		
		a = a + 2
		z = 1 / a
		
		if (pi*4) - 0.00001 == (y*4) {
			x = 0
}		else {
			x = 1
}
}
		y = y*4
		pi = pi*4
	fmt.Println("Hasil pi:", y)
	fmt.Println("Hasil pi:", pi)
	fmt.Println("Pada i ke:", i)
}