package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var book string
	scanner := bufio.NewScanner(os.Stdin)
	_ = scanner.Scan()
	book = scanner.Text()
	fmt.Println(book, "- лучшая книга!")
}
