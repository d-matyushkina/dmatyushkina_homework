package lesson3

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func RemoveDuplicates(nums []int) []int {
	counts := make(map[int]int, len(nums))
	woDuplicates := make([]int, 0, len(nums))
	for _, num := range nums {
		_, exists := counts[num]
		if exists {
			continue
		}
		counts[num] = 1
		woDuplicates = append(woDuplicates, num)
	}
	return woDuplicates
}

func Task4Run(reader *bufio.Reader) {
	fmt.Print("\n- - - - - - - - \nЗадача 4\n- - - - - - - - \n")
	fmt.Print("Введите цифры через запятую: ")
	input, _ := reader.ReadString('\n')
	trimmed := strings.TrimSpace(input)
	splittedInput := strings.Split(trimmed, ",")
	nums := make([]int, 0, len(splittedInput))
	for _, strNum := range splittedInput {
		strNum = strings.TrimSpace(strNum)
		num, _ := strconv.Atoi(strNum)
		nums = append(nums, num)
	}
	fmt.Print(RemoveDuplicates(nums))
}
