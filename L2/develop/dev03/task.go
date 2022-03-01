package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	Sort("text.txt", "", 1)
}

func Sort(fileName, sort string, idx int) {
	var sf = sortFunc{idx, SortChar}
	tree := NewTree()

	if sort == "int" {
		sf = sortFunc{idx, SortInt}
	}

	file, err := os.Open(fileName)
	if err != nil {
		log.Println(err)
	}
	defer file.Close()

	newFile, err := os.Create("Sort" + fileName)
	if err != nil {
		log.Println(err)
	}
	defer newFile.Close()

	readWriter := bufio.NewReadWriter(bufio.NewReader(file), bufio.NewWriter(newFile))
	header, err := readWriter.ReadString('\n')
	if err != nil {
		log.Fatalln(err)
	}

	_, err = readWriter.WriteString(header)
	if err != nil {
		log.Println(err)
	}

	for {
		str, err1 := readWriter.ReadString('\n')
		if err1 != nil && err1 != io.EOF {
			log.Fatalln(err)
		}

		strSlice := strings.Split(str, " ")

		if strSlice[0] != "" {
			tree.Insert(strSlice, sf)
		}

		if err1 == io.EOF {
			break
		}

	}
	/*
		_, err=file.Seek(0,io.SeekStart)
		if err != nil {
			log.Println(err)
		}

		err=file.Truncate(8)
		if err != nil {
			log.Println(err)
		}
	*/

	for _, val := range tree.Println() {
		_, err = readWriter.WriteString(val)
		if err != nil {
			log.Println(err)
		}

		err = readWriter.Flush()
		if err != nil {
			log.Println(err)
		}
	}
}
