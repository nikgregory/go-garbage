package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	s, sep := "",""
	for i, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
        fmt.Println("Index: " + strconv.Itoa(i) + " Args: " + arg )
	}
	fmt.Println(s)
}
