package main

import (
    "io"
    "os"
)

func main() {
    if 1 == len(os.Args) {
        io.Copy(os.Stdout, os.Stdin)
    }
}
