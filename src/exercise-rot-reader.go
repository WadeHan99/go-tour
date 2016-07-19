package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type rot13Reader1 struct {
	r io.Reader
}

func (rot rot13Reader1) Read(p []byte) (n int, err error) {
	n, err = rot.r.Read(p)

	// for each letter read from io.Reader
	for i := 0; i < len(p); i++ {
		// if the letter's index is between A -N, add 13 to its index
		if (p[i] >= 'A' && p[i] < 'N') || (p[i] > 'a' && p[i] < 'n') {
			p[i] += 13
		} else if (p[i] > 'M' && p[i] <= 'Z') || (p[i] > 'm' && p[i] <= 'z') {
			p[i] -= 13
		}
	}
	return
}

type rot13Reader2 struct {
	r io.Reader
}

func (rot rot13Reader2) Read(p []byte) (n int, err error) {
	n, err = rot.r.Read(p)

	leng := len(p)
	for i := 0; i < leng; i++ {
		switch {
		case p[i] >= 'a' && p[i] < 'n':
			fallthrough
		case p[i] >= 'A' && p[i] < 'N':
			p[i] = p[i] + 13
		case p[i] >= 'n' && p[i] <= 'z':
			fallthrough
		case p[i] >= 'N' && p[i] <= 'Z':
			p[i] = p[i] - 13
		}
	}

	return
}

func main() {
	fmt.Println("Lbh penpxrq gur pbqr!")
	s1 := strings.NewReader("Lbh penpxrq gur pbqr!")
	r1 := rot13Reader1{s1}
	io.Copy(os.Stdout, &r1)
	fmt.Println()

	s2 := strings.NewReader("Lbh penpxrq gur pbqr!")
	r2 := rot13Reader2{s2}
	io.Copy(os.Stdout, &r2)
	fmt.Println()
}
