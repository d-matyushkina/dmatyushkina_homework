package lesson3

import (
	"bufio"
	"fmt"
	"strings"
)

func IsPalindrome(s string) bool {
	newStr := deleteTabs(strings.ToLower(s))

	i, j := 0, len(newStr)-1
	for i < j {
		if newStr[i] != newStr[j] {
			return false
		}
		i++
		j--
	}
	return true

}

func Task6Run(reader *bufio.Reader) {
	fmt.Print("\n- - - - - - - - \nЗадача 6\n- - - - - - - - \n")
	fmt.Print("Введите фразу: ")
	phrase, _ := reader.ReadString('\n')
	fmt.Print(IsPalindrome(strings.TrimRight(phrase, "\r\n")))
}
