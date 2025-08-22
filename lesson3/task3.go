package lesson3

import (
	"bufio"
	"fmt"
	"strings"
)

func FirstUnique(s string) rune {
	counts := make(map[rune]int, len(s))
	runes := []rune(s)

	for _, r := range runes {
		counts[r]++
	}

	for _, r := range runes {
		if counts[r] == 1 {
			return r
		}
	}
	return 0
}

func Task3Run(reader *bufio.Reader) {
	fmt.Print("\n- - - - - - - - \nЗадача 3\n- - - - - - - - \n")
	fmt.Print("Введите фразу: ")
	symbols, _ := reader.ReadString('\n')
	result := FirstUnique(strings.TrimRight(symbols, "\r\n"))
	if result == 0 {
		fmt.Println("0")
	} else {
		fmt.Printf("%c", result)
	}
}
