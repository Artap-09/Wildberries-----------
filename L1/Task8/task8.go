package main

import (
	"fmt"
)

func main() {
	var i64, i, result int64
	var oneOrZero string

	fmt.Print("Введите число: ")
	fmt.Scanln(&i64)

	module := false
	if i64 < 0 {
		i64 = -i64
		module = true
	}

	for {
		fmt.Printf("Введите порядковый номер бита, числа %d, который нужно поменять: ", i64)
		fmt.Scanln(&i)
		if i > 0 {
			break
		}
		fmt.Println("Порядковый номер должен быть больше 0")
	}

Choose: //Это label
	fmt.Printf("Введите на что поменять %d бит на 1 или на 0: ", i)
	fmt.Scanln(&oneOrZero)

	switch oneOrZero {
	case "1":
		result = i64 | (1 << (i - 1))
		if len(fmt.Sprintf("%b", i64)) < len(fmt.Sprintf("%b", result)) {
			var zero string
			for n := 0; n < len(fmt.Sprintf("%b", result))-len(fmt.Sprintf("%b", i64)); n++ {
				zero += "0"
			}
			if module {
				fmt.Printf("Было:  -%s%b (%d)\nСтало: %b (%d)\n", zero, i64, -i64, -result, -result)
			} else {
				fmt.Printf("Было:  %s%b (%d)\nСтало: %b (%d)\n", zero, i64, i64, result, result)
			}
		} else {
			if module {
				result = -result
				i64 = -i64
			}
			fmt.Printf("Было:  %b (%d)\nСтало: %b (%d)\n", i64, i64, result, result)
		}
	case "0":
		result = i64 &^ (1 << (i - 1))
		if len(fmt.Sprintf("%b", i64)) > len(fmt.Sprintf("%b", result)) {
			if module {
				fmt.Printf("Было:  %b (%d)\nСтало: -0%b (%d)\n", -i64, -i64, result, -result)
			} else {
				fmt.Printf("Было:  %b (%d)\nСтало: 0%b (%d)\n", i64, i64, result, result)
			}
		} else {
			if module {
				result = -result
				i64 = -i64
			}
			fmt.Printf("Было:  %b (%d)\nСтало: %b (%d)\n", i64, i64, result, result)
		}
	default:
		fmt.Println("Может быть только 1 или 0")
		goto Choose // Переходим к label Choose (Так не принято делать лучше оформлять через for. Пример выше.)
	}
}
