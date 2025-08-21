package lesson3

import (
	"bufio"
	"fmt"
	"strings"
)

func WordCount(s string) map[string]int {
	counts := make(map[string]int)
	words := strings.Fields(s)
	for _, word := range words {
		counts[word]++
	}
	return counts
}

func Task1Run(reader *bufio.Reader) {
	fmt.Print("\n- - - - - - - - \nЗадача 1\n- - - - - - - - \n")
	fmt.Print("Введите фразу: ")
	phrase, _ := reader.ReadString('\n')
	fmt.Print(WordCount(phrase))
}
