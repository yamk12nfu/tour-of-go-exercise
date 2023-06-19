package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (r13 *rot13Reader) Read(b []byte) (int, error) {
	n, err := r13.r.Read(b)
	if err != nil {
		return 0, err
	}

	for i, v := range b {
		switch {
			case v >= 'A' && v <= 'M', v >= 'a' && v <= 'm':
				b[i] += 13
			case v >= 'N' && v <= 'Z', v >= 'n' && v <= 'z':
				b[i] -= 13
			}
		}

		return n, nil
	}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
