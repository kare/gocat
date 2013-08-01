package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	if 2 == len(os.Args) && os.Args[1] == "-h" {
		fmt.Printf("usage: %s [file]\n", os.Args[0])
		os.Exit(0)
	}
	if 1 == len(os.Args) {
		io.Copy(os.Stdout, os.Stdin)
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
		reader := bufio.NewReader(file)
		io.Copy(os.Stdout, reader)
	}
}
