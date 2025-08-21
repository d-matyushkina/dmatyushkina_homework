package lesson3

import (
	"bufio"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func MakeChessboard(width int, height int) ([][]rune, error) {
	if width < 1 || height < 1 {
		return nil, errors.New("Ширина и высота доски должны быть больше 0")
	}
	result := make([][]rune, height)
	for i := range result {
		result[i] = make([]rune, width)
		for j := range result[i] {
			if (i+j)%2 == 0 {
				result[i][j] = ' '
			} else {
				result[i][j] = '#'
			}
		}
	}
	return result, nil
}

func Task7Run(reader *bufio.Reader) {
	fmt.Print("\n- - - - - - - - \nЗадача 7\n- - - - - - - - \n")
	fmt.Print("Введите ширину доски: ")
	width, _ := reader.ReadString('\n')
	width = strings.TrimSuffix(width, "\n")
	intWidth, _ := strconv.Atoi(width)
	fmt.Print("Введите высоту доски: ")
	height, _ := reader.ReadString('\n')
	height = strings.TrimSuffix(height, "\n")
	intHeight, _ := strconv.Atoi(height)
	if intHeight < 1 || intWidth < 1 {
		_, err := MakeChessboard(intWidth, intHeight)
		fmt.Println(err)
	} else {
		chess, _ := MakeChessboard(intWidth, intHeight)
		for _, r := range chess {
			fmt.Println(string(r))
		}
	}
}
