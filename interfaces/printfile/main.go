package main

import (
	"fmt"
	"io"
	"os"
)

type logWriter struct {
}

// Custome writer
func (logWriter) Write(p []byte) (n int, err error) {
	fmt.Println("Hey I am here len: ", len(p))
	return len(p), nil
}
func main() {

	//  fmt.Println(os.Args)
	//	fmt.Println(os.Args[1])

	file, err := os.Open(os.Args[1])

	// lw := logWriter{}

	if err != nil {
		fmt.Println("Error:", err)
	}
	// io.Copy(lw, file)
	io.Copy(os.Stdout, file)

}
