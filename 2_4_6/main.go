package main

import (
	"fmt"
)

func main() {
	var k, n int
	fmt.Scan(&n, &k)
	fmt.Print(k % n)
}
