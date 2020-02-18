/* NIM : 1301194086
   Nama : Mikogizka Satria Kartika
*/
package main

import (
	"fmt"
	"math"
)

var numpoint int64
var jarDek float64

type Point struct {
x float64
y float64
}

const N = 10000

var points [N]Point

func jarak(p1 Point, p2 Point) float64 {
var ttk, ttk1, ttk2 float64
ttk1 = p1.x - p2.x
ttk2 = p1.y - p2.y

ttk = math.Sqrt((ttk1 * ttk1) + (ttk2 * ttk2))
return ttk
}

func bacaTitik() {
i := 0
cek := true
for cek && numpoint < N {
	fmt.Scan(&points[i].x, &points[i].y)
	if points[i].x == 0 && points[i].y == 0 {
		cek = false
} 	else {
		numpoint++
}
i++
}
}

func ambilTitikTerdekat(p1 *Point, p2 *Point) {

var i, a int64
i = 0
var dekat float64
for i+1 <= numpoint {
	a = i + 1
	jarDek = jarak(points[i], points[a])
	dekat = jarak(*p1, *p2)
	if jarDek < dekat && jarDek > 0 {
		dekat = jarDek
		*p1 = points[i]
		*p2 = points[a]
	} else if i == 0 {
		dekat = jarDek
		*p1 = points[i]
		*p2 = points[a]
}
i++
}

}

func main() {
var p1 Point
var p2 Point
bacaTitik()
ambilTitikTerdekat(&p1, &p2)
fmt.Printf("Titik terdekat adalah (%.1f, %.1f) dan (%.1f, %.1f) dengan jarak %.1f ", p1.x, p1.y, p2.x, p2.y, jarak(p1, p2))

fmt.Println("------------------------------------------------------------")
fmt.Println("Nama: Mikogizka Satria Kartika")
fmt.Println("NIM: 1301194086")
}