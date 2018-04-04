package main

import (
	"bytes"
	"fmt"
	"io"
	"strings"
)

func main() {
	var buffer1 bytes.Buffer
	w1 := io.MultiWriter(&buffer1)
	r1 := strings.NewReader("test")
	io.Copy(w1, r1)
	fmt.Println(buffer1.String())

	buffer2 := new(bytes.Buffer)
	w2 := io.MultiWriter(buffer2)
	r2 := strings.NewReader("test")
	io.Copy(w2, r2)
	fmt.Println(buffer2.String())

	/*
		var buffers []*bytes.Buffer
		for i := 0; i < 3; i++ {
			b := new(bytes.Buffer)
			buffers = append(buffers, b)
		}
		w3 := io.MultiWriter(buffers0...)
		r3 := strings.NewReader("test")
		io.Copy(w3, r3)
		fmt.Println(buffers[0].String())

		以下のエラ-でコンパイルできない
				# command-line-arguments
			./main.go:28:22: cannot use buffers (type []*bytes.Buffer) as type []io.Writer in argument to io.MultiWriter
	*/

	var buffers []*bytes.Buffer
	for i := 0; i < 3; i++ {
		b := new(bytes.Buffer)
		buffers = append(buffers, b)
	}
	w4 := io.MultiWriter(buffers[0], buffers[1])
	r4 := strings.NewReader("test")
	io.Copy(w4, r4)
	fmt.Println(buffers[0].String())
}
