package main

import "fmt"

func main() {
	const minutesInHour = 60
	var minutes int
	fmt.Scan(&minutes)
	fmt.Println(minutes, "мин - это", minutes/minutesInHour, "час", minutes%minutesInHour, "минут.")
}
