package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}
func rot13(b byte) byte {
	t := rune(b)
	if t >= 'a' && t <= 'm' || t >= 'A' && t <= 'M' {
		b += 13
	}
	if t >= 'n' && t <= 'z' || t >= 'N' && t <= 'Z' {
		b -= 13
	}
	return b
}
func (r rot13Reader) Read(data []byte) (int, error){
	v, s := r.r.Read(data)
	for i := 0; i < len(data); i++ {
		data[i] = rot13(data[i])
	}
	return v, s
}
func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!\n")
	r := rot13Reader{s}
	io.Copy(os.Stdout , &r)

}
