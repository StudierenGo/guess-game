package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func main() {
	const attempts = 10
	success := false

	target := rand.Intn(100) + 1
	fmt.Println("Я загадал число от 1 до 100. Попробуй угадать его?")
	// fmt.Println("Введите ваше предположение: ")
	fmt.Println(target) // Для тестирования, убрать в реальной игре

	for i := 1; i < attempts; i += 1 {
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("Введите ваше предположение: ")
		input, err := reader.ReadString('\n')

		if err != nil {
			fmt.Println("Ошибка при чтении ввода:", err)
			return
		}

		input = strings.TrimSpace(input)
		guess, err := strconv.Atoi(input)

		if err != nil {
			fmt.Println("Пожалуйста, введите корректное число.")
			return
		}

		if guess < target {
			fmt.Println("Ваше число меньше загаданного мною числа!")
		} else if guess > target {
			fmt.Println("Ваше число больше загаданного мною числа!")
		} else {
			fmt.Printf("Поздравляю! Вы угадали число %d с %d попытки!\n", target, i)
			success = true
			break
		}
	}

	if !success {
		fmt.Printf("К сожалению, вы исчерпали все попытки. Загаданное число было %d.\n", target)
	}
}
