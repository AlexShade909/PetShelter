package cli

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func ReadMenuChoice(prompt string, min, max int) int {
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

func ReadNonEmptyString(prompt string) string {
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

func ReadFloat(prompt string) float64 {
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

func ReadInt(prompt string) int {
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

func ReadYesNo(prompt string) bool {
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
