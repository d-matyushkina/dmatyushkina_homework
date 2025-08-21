package lesson3

import (
	"bufio"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func RemoveElement(nums []int, index int) ([]int, error) {
	if index < 0 || index >= len(nums) {
		return nil, errors.New("Индекс должен быть <= размеру слайса и >= 0")
	}
	newNums := make([]int, len(nums)-1)

	for i := 0; i < len(nums); i++ {
		if i < index {
			newNums[i] = nums[i]
		} else if i > index {
			newNums[i-1] = nums[i]
		}
	}
	return newNums, nil
}

func Task5Run(reader *bufio.Reader) {
	fmt.Print("\n- - - - - - - - \nЗадача 5\n- - - - - - - - \n")
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
	fmt.Print("Введите индекс: ")
	index, _ := reader.ReadString('\n')
	index = strings.TrimSuffix(index, "\n")
	intIndex, _ := strconv.Atoi(index)
	results, _ := RemoveElement(nums, intIndex)
	if len(results) == 0 {
		fmt.Print("nil")
	} else {
		fmt.Print(results)
	}
}
