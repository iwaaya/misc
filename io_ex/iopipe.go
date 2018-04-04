// io.Pipeを使用した際にデッドロックが起きたため理解のための殴り書き

package main

import (
	"bytes"
	"fmt"
	"io"
	"strings"
	"time"
)

// デッドロックのヒントがここに
// http://ascii.jp/elem/000/001/260/1260449/index-2.html

func main() {
	var writers = make([]io.Writer, 3)
	var readers = make([]io.Reader, 3)

	for i := 0; i < 3; i++ {
		readers[i], writers[i] = io.Pipe()
		defer writers[i].(*io.PipeWriter).Close()
	}

	w := io.MultiWriter(writers...)
	go io.Copy(w, strings.NewReader("test"))
	time.Sleep(5 * time.Second)
	var buf bytes.Buffer
	go io.Copy(&buf, readers[0])
	time.Sleep(5 * time.Second)
	fmt.Println(buf.String())
}
