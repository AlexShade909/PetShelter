package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

// readMenuChoice запрашивает целое число в диапазоне [min, max] включительно.
// Используется и для "1. Забрать / 2. Добавить", и для выбора приюта/поликлиники по номеру.
func readMenuChoice(prompt string, min, max int) int {
	for {
		fmt.Print(prompt)
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Ошибка чтения ввода, попробуйте снова")
			continue
		}

		line = strings.TrimSpace(line)
		choice, err := strconv.Atoi(line)
		if err != nil {
			fmt.Println("Введите целое число")
			continue
		}

		if choice < min || choice > max {
			fmt.Printf("Число должно быть от %d до %d\n", min, max)
			continue
		}

		return choice
	}
}

// readNonEmptyString запрашивает строку и не даёт продолжить, пока она пустая.
// Используется для клички, даты поступления и т.п.
func readNonEmptyString(prompt string) string {
	for {
		fmt.Print(prompt)
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Ошибка чтения ввода, попробуйте снова")
			continue
		}

		line = strings.TrimSpace(line)
		if line == "" {
			fmt.Println("Строка не может быть пустой")
			continue
		}

		return line
	}
}

// readFloat запрашивает дробное число (например, вес собаки).
func readFloat(prompt string) float64 {
	for {
		fmt.Print(prompt)
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Ошибка чтения ввода, попробуйте снова")
			continue
		}

		line = strings.TrimSpace(line)
		line = strings.ReplaceAll(line, ",", ".") // на случай ввода "7,1" вместо "7.1"

		value, err := strconv.ParseFloat(line, 64)
		if err != nil {
			fmt.Println("Введите число, например 7.1")
			continue
		}

		return value
	}
}

// readInt запрашивает целое число без верхней/нижней границы (например, возраст).
func readInt(prompt string) int {
	for {
		fmt.Print(prompt)
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Ошибка чтения ввода, попробуйте снова")
			continue
		}

		line = strings.TrimSpace(line)
		value, err := strconv.Atoi(line)
		if err != nil {
			fmt.Println("Введите целое число")
			continue
		}

		return value
	}
}

// readYesNo запрашивает подтверждение да/нет.
// Принимает "да"/"нет", а также короткие "д"/"н" и латиницу "y"/"n" для удобства.
func readYesNo(prompt string) bool {
	for {
		fmt.Print(prompt)
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Ошибка чтения ввода, попробуйте снова")
			continue
		}

		line = strings.ToLower(strings.TrimSpace(line))

		switch line {
		case "да", "д", "y", "yes":
			return true
		case "нет", "н", "n", "no":
			return false
		default:
			fmt.Println("Введите 'да' или 'нет'")
		}
	}

}
