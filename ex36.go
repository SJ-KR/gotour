package main

import (
	"code.google.com/p/go-tour/pic"
)

func Pic(dx, dy int) [][]uint8 {
	p := make([][]uint8, dx)
	for i := 0;  i < dx; i++ {
		p[i] = make([]uint8, dy)
		for j := 0; j < dy; j++ {
			p[i][j] = uint8(i^j)
		}
	}
	return p

}
func main() {
	pic.Show(Pic)


}
