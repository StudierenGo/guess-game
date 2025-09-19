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

	for i := 1; i <= attempts; i += 1 {
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
			fmt.Printf("Ваше число меньше загаданного мною числа! У Вас осталось %d попыток.\n", attempts-i)
		} else if guess > target {
			fmt.Printf("Ваше число больше загаданного мною числа! У Вас осталось %d попыток.\n", attempts-i)
		} else {
			fmt.Printf("Поздравляю! Вы угадали число %d с %d попытки!\n", target, i)
			success = true
			break
		}
	}

	if !success {
		fmt.Printf("К сожалению, Вы исчерпали все попытки. Загаданное мною число было - %d.\n", target)
	}
}
