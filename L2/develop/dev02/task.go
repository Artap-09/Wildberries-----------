package main

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
)

func main() {
	var str string
	fmt.Print("Введите строку для распоковки: ")
	fmt.Scan(&str)
	result, err := Unpacking(str)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(result)
}

func Unpacking(str string) (string, error) {
	var err error
	// result - переменная для результата распоковки
	var result strings.Builder
	// prev - переменная для хранения предедущего символа
	var prev string
	var count int
	n := len(str)

	for i := 0; i < n; i++ { // Итерируемся по символам строки

		val := string(str[i]) // Текущий символ.
		switch {              // Проверяем val на спец символы ("\" или числа)
		case val == "\\":

			if i+1 >= n { // Проверяем если еще сиволы, если нет возвращаем ошибку
				err = errors.New("некорректная строка")
				return "", err
			}

			i++ // Переходим к следующему сиволу и рассматриваем его уже не как спец символ
			val = string(str[i])

			if i+1 < n && str[i+1] >= '0' && str[i+1] <= '9' { // Проверяем идет ли следом число, если да записываем в значение в prev и переходим к следующей итерации. Если нет добавляем значение в результат.
				prev = string(val)
				continue
			}

			result.Grow(1)                  // Увеличиваем cap на 1 для конкатенации строк
			result.WriteString(string(val)) // Добавляем значение к результирующей строке.

		case val >= "0" && val <= "9":
			num, err := strconv.Atoi(string(val)) // Преобразуем строку в int
			if err != nil {
				return "", err
			}

			count = count*10 + num // Счетчик сколько раз надо повторить символ из prev
			if i+1 < n && string(str[i+1]) >= "0" && string(str[i+1]) <= "9" { // Проверяем является ли следующий символ цифрой. Если да переходим к нему.
				continue
			}

			if prev == "" { // Проверяем есть ли значение для распаковки, если нет возвращаем пустую строку и ошибку.
				err = errors.New("некорректная строка")
				return "", err
			}

			result.Grow(count) // Увеличиваем cap на count для конкатенации строк.
			for count != 0 { // Добавляем символ из prev к результату count раз.
				result.WriteString(prev)
				count--
			}

			prev = "" // Показываем что нет символа для распаковки.

		default:
			if i+1 < n && str[i+1] >= '0' && str[i+1] <= '9' { // Проверяем идет ли следом число, если да записываем в значение в prev и переходим к следующей итерации. Если нет добавляем значение в результат.
				prev = string(val)
				continue
			}

			result.Grow(1)// Увеличиваем cap на 1 для конкатенации строк
			result.WriteString(string(val)) // Добавляем значение к результирующей строке.
		}
	}
	return result.String(), err // Возвращаем результирующую строку и ошибку. Если ошибки нет, err == nil.
}
