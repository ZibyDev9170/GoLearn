package main

import (
	"fmt"
	"strings"
)

func tree(height int) {
	for i := range height {
		fmt.Print(strings.Repeat(" ", height-i-1))
		fmt.Print(strings.Repeat("*", i*2+1))
		fmt.Println()
	}
}

func main() {
	tree(4)
	tree(7)
	tree(2)
}
