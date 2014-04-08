package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func cat(dst io.Writer, src io.Reader) (written int64, err error) {
	return io.Copy(dst, src)
}

func main() {
	if 2 == len(os.Args) && os.Args[1] == "-h" {
		fmt.Fprintf(os.Stderr, "usage: %s [file]\n", os.Args[0])
		os.Exit(0)
	}
	if 1 == len(os.Args) {
		if _, err := cat(os.Stdout, os.Stdin); err != nil {
			panic(err)
		}
	}
	if 2 == len(os.Args) {
		var (
			file *os.File
			err  error
		)
		if file, err = os.Open(os.Args[1]); err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
		defer file.Close()
		if _, err := cat(os.Stdout, bufio.NewReader(file)); err != nil {
			panic(err)
		}
	}
}
