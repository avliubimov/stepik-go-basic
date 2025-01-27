package main

import "fmt"

func main() {
	var phonePrice, casePrice, chargerPrice, earphonesPrice int
	fmt.Scan(&phonePrice, &casePrice, &chargerPrice, &earphonesPrice)
	fmt.Print((phonePrice + casePrice + chargerPrice + earphonesPrice) * 3)
}
