package main

import (
	"fmt"
)

func main() {
	const kopeckInRouble = 100
	var roubles, kopeck, amount, price int
	fmt.Scan(&roubles, &kopeck, &amount)
	kopeck = roubles*kopeckInRouble + kopeck
	price = kopeck * amount
	fmt.Print(price/kopeckInRouble, price%kopeckInRouble)
}
