package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Print(n%10 + n/100 + (n/10)%10)
}
