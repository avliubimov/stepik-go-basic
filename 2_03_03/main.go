/*
Напишите программу, которая считывает три строки по очереди, а затем выводит их в той же последовательности, каждую на отдельной строчке.
Формат входных данных
На вход программе подаются три строки, каждая на отдельной строке.

Формат выходных данных
Программа должна вывести введенные строки в той же последовательности, каждую на отдельной строке.

Sample Input 1:

I was
born
this way
Sample Output 1:

I was
born
this way
Sample Input 2:

I love
Go
so much
Sample Output 2:

I love
Go
so much
*/

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var string1, string2, string3 string
	scanner := bufio.NewScanner(os.Stdin)
	_ = scanner.Scan()
	string1 = scanner.Text()
	_ = scanner.Scan()
	string2 = scanner.Text()
	_ = scanner.Scan()
	string3 = scanner.Text()
	fmt.Println(string1)
	fmt.Println(string2)
	fmt.Println(string3)

}
