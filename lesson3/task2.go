package lesson3

import (
	"bufio"
	"fmt"
	"sort"
	"strings"
)

func AreAnagrams(s1, s2 string) bool {
	s1New := deleteTabs(strings.ToLower(s1))
	s2New := deleteTabs(strings.ToLower(s2))
	if len(s1New) != len(s2New) {
		return false
	}
	sort.SliceStable(s1New, func(i, j int) bool {
		return s1New[i] < s1New[j]
	})
	sort.SliceStable(s2New, func(i, j int) bool {
		return s2New[i] < s2New[j]
	})
	for i := 0; i < len(s1New); i++ {
		if s1New[i] != s2New[i] {
			return false
		}
	}
	return true
}

func deleteTabs(withTabs string) []rune {
	woTabs := make([]rune, 0, len(withTabs))
	for _, r := range withTabs {
		switch r {
		case ' ', '\n', '\t', '\f', '\v', '\r':
			continue
		default:
			woTabs = append(woTabs, r)
		}
	}
	return woTabs
}

func Task2Run(reader *bufio.Reader) {
	fmt.Print("\n- - - - - - - - \nЗадача 2\n- - - - - - - - \n")
	fmt.Print("Введите первое слово/фразу: ")
	firstPhrase, _ := reader.ReadString('\n')
	fmt.Print("Введите второе слово/фразу: ")
	secondPhrase, _ := reader.ReadString('\n')
	fmt.Print(AreAnagrams(firstPhrase, secondPhrase))
}
