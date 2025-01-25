package main

import "fmt"

func main() {
	var num int
	fmt.Scan(&num)
	t1 := num * num * num
	t2 := t1 * t1
	fmt.Print(t2)
}
