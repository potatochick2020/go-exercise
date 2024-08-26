package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rot *rot13Reader) Read(p []byte) (n int, err error) {
	n, err = rot.r.Read(p)
	//copy logic from https://github.com/jreisinger/gokatas/blob/main/rot13/rot13.go
	for i, b := range p {
		var a, z byte
		switch {
		case 'a' <= b && b <= 'z':
			a, z = 'a', 'z'
		case 'A' <= b && b <= 'Z':
			a, z = 'A', 'Z'
		default:
			p[i] = b
		}
		// return (b-a+13)%26 + a
		p[i] = (b-a+13)%(z-a+1) + a
	}
	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
