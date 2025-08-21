package main

import (
	"bufio"
	"dmatyushkina_homework/lesson3"
	"os"
)

// TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>
var reader *bufio.Reader

func main() {

	reader := bufio.NewReader(os.Stdin)
	// Task 1
	lesson3.Task1Run(reader)

	// Task 2
	lesson3.Task2Run(reader)

	// Task 3
	lesson3.Task3Run(reader)

	// Task 4
	lesson3.Task4Run(reader)

	// Task 5
	lesson3.Task5Run(reader)

	// Task 6
	lesson3.Task6Run(reader)

	// Task 7
	lesson3.Task7Run(reader)
}
