package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	sortChar = iota
	sortNum
	sortMonth
)

func main() {

	var (
		idx    int
		sf     = sortChar
		resort = false
		unique = false
		err    error
	)

	if os.Args[1] == "-h" || os.Args[1] == "--help" {
		fmt.Printf("-k — указание колонки для сортировки (слова в строке могут выступать в качестве колонок, по умолчанию разделитель — пробел)\n-n — сортировать по числовому значению\n-r — сортировать в обратном порядке\n-u — не выводить повторяющиеся строки\n-M — сортировать по названию месяца\n")
		return
	}

	for i := 2; i < len(os.Args); i++ {
		switch os.Args[i] {
		case "-k":
			if idx != 0 {
				log.Fatalf("%d - повторяющейся ключ.\n", i+1)
			}

			i++

			idx, err = strconv.Atoi(os.Args[i])
			if err != nil {
				log.Fatalln(err)
			}
			idx--
		case "-n":
			if sf != sortChar {
				log.Fatalf("%d - повторяющейся ключ.\n", i+1)
			}

			sf = sortNum
		case "-r":
			if resort {
				log.Fatalf("%d - повторяющейся ключ.\n", i+1)
			}

			resort = true
		case "-u":
			if unique {
				log.Fatalf("%d - повторяющейся ключ.\n", i+1)
			}

			unique = true
		case "-M":
			if sf != sortChar {
				log.Fatalf("%d - повторяющейся ключ.\n", i+1)
			}

			sf = sortMonth
		default:
			log.Fatalf("%s - неизвестный ключ. Можете узнать доступные ключи по ключу -h или --help\n", os.Args[i])
		}
	}

	err=Sort(os.Args[1], idx, sf, resort, unique)
	if err != nil {
		log.Fatalln(err)
	}
}

func Sort(fileName string, idx, sort int, resort, unique bool) error {
	var sf = sortFunc{idx, SortChar}
	tree := NewTree()

	if sort == sortNum {
		sf = sortFunc{idx, SortInt}
	} else if sort == sortMonth {
		sf = sortFunc{idx, SortMonth}
	}

	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	newFile, err := os.Create("Sort" + fileName)
	if err != nil {
		return err
	}
	defer newFile.Close()

	readWriter := bufio.NewReadWriter(bufio.NewReader(file), bufio.NewWriter(newFile))
	header, err := readWriter.ReadString('\n')
	if err != nil {
		return err
	}

	_, err = readWriter.WriteString(header)
	if err != nil {
		return err
	}

	for {
		str, err1 := readWriter.ReadString('\n')
		if err1 != nil && err1 != io.EOF {
			return err1
		}

		strSlice := strings.Split(str, " ")

		if str != "" {
			err = tree.Insert(strSlice, sf, unique)
			if err != nil {
				return err
			}
		}
		
		if err1 == io.EOF {
			break
		}

	}

	for _, val := range tree.Println(resort) {
		_, err = readWriter.WriteString(val)
		if err != nil {
			return err
		}

		err = readWriter.Flush()
		if err != nil {
			return err
		}
	}

	return nil
}
