package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	nameFile := os.Args[1]
	var column int
	var separated bool
	delim := " "
	for i := 2; i < len(os.Args); i++ {
		switch os.Args[i] {
		case "-f":
			i++
			if i >= len(os.Args) {
				log.Fatalln("укажите поле (колонку)")
			}

			var err error
			column, err = strconv.Atoi(os.Args[i])
			if err != nil {
				log.Fatalln(err)
			}
		case "-d":
			i++
			if i >= len(os.Args) {
				log.Fatalln("укажите разделитель (по умолчанию пробел)")
			}

			delim = os.Args[i]
		case "-s":
			separated = true
		default:
			log.Fatalln("не корректно указанна команда")
		}
	}

	file, err := os.Open(nameFile)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		str := strings.Split(scanner.Text(), delim)
		if column >= len(str) {
			if separated {
				continue
			}
			fmt.Println()
			continue
		} 

		fmt.Println(str[column-1])
	}
}
