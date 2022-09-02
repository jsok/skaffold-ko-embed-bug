package main

import (
	_ "embed"
	"fmt"
)

//go:embed hello.txt
var b []byte

func main() {
	fmt.Printf("%s", b)
}
