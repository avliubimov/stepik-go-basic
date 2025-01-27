package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var divider, string1, string2, string3 string
	scanner := bufio.NewScanner(os.Stdin)
	_ = scanner.Scan()
	divider = scanner.Text()
	_ = scanner.Scan()
	string1 = scanner.Text()
	_ = scanner.Scan()
	string2 = scanner.Text()
	_ = scanner.Scan()
	string3 = scanner.Text()
	fmt.Print(string1 + divider + string2 + divider + string3)
}
