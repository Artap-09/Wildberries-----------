package main

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
	"testing"
)

const (
	fileName    = "text.txt"
	fileOutName = "Sorttext.txt"
	col1        = 0
	col2        = 1
	col6        = 5
)

var col int

func TestSort(t *testing.T) {
	col = col2
	err := Sort(fileName, col, sortChar, false, false)
	if err != nil {
		t.Error(err)
	}

	file, err := os.Open(fileOutName)
	if err != nil {
		t.Error(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	_, err = reader.ReadString('\n')
	if err != nil {
		t.Error(err)
	}

	str, err1 := reader.ReadString('\n')
	if err != nil && err1 != io.EOF {
		t.Error(err)
	}

	strSlice := strings.Split(str, " ")
	var prev string

	if str != "" {
		prev = strSlice[col]
		if err != nil {
			t.Error(err)
		}
	}

	for {
		str, err := reader.ReadString('\n')
		if err != nil && err != io.EOF {
			t.Error(err)
		}

		strSlice := strings.Split(str, " ")

		if str != "" {

			if strings.Compare(prev, strSlice[col]) > 0 {
				t.Errorf("нарушен порядок на строке %s", str)
			}
		}

		if err == io.EOF {
			break
		}

	}
}

func TestSortNum(t *testing.T) {
	col = col1
	err := Sort(fileName, col, sortNum, false, false)
	if err != nil {
		t.Error(err)
	}

	file, err := os.Open(fileOutName)
	if err != nil {
		t.Error(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	_, err = reader.ReadString('\n')
	if err != nil {
		t.Error(err)
	}

	str, err1 := reader.ReadString('\n')
	if err != nil && err1 != io.EOF {
		t.Error(err)
	}

	strSlice := strings.Split(str, " ")
	var prev, now int

	if str != "" {
		prev, err = strconv.Atoi(strSlice[col])
		if err != nil {
			t.Error(err)
		}
	}

	for {
		str, err := reader.ReadString('\n')
		if err != nil && err != io.EOF {
			t.Error(err)
		}

		strSlice := strings.Split(str, " ")

		if str != "" {
			now, err = strconv.Atoi(strSlice[0])
			if err != nil {
				t.Error(err)
			}

			if prev > now {
				t.Errorf("нарушен порядок на строке %s", str)
			}
		}

		if err == io.EOF {
			break
		}

	}
}

func TestSortResortNum(t *testing.T) {
	col = col1
	err := Sort(fileName, col, sortNum, true, false)
	if err != nil {
		t.Error(err)
	}

	file, err := os.Open(fileOutName)
	if err != nil {
		t.Error(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	_, err = reader.ReadString('\n')
	if err != nil {
		t.Error(err)
	}

	str, err1 := reader.ReadString('\n')
	if err != nil && err1 != io.EOF {
		t.Error(err)
	}

	strSlice := strings.Split(str, " ")
	var prev, now int

	if str != "" {
		prev, err = strconv.Atoi(strSlice[col])
		if err != nil {
			t.Error(err)
		}
	}

	for {
		str, err := reader.ReadString('\n')
		if err != nil && err != io.EOF {
			t.Error(err)
		}

		strSlice := strings.Split(str, " ")

		if str != "" {
			now, err = strconv.Atoi(strSlice[0])
			if err != nil {
				t.Error(err)
			}

			if prev < now {
				t.Errorf("нарушен порядок на строке %s", str)
			}
		}

		if err == io.EOF {
			break
		}

	}
}

func TestSortUnication(t *testing.T) {
	col = col2
	err := Sort(fileName, col, sortChar, false, true)
	if err != nil {
		t.Error(err)
	}

	file, err := os.Open(fileOutName)
	if err != nil {
		t.Error(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	_, err = reader.ReadString('\n')
	if err != nil {
		t.Error(err)
	}

	str, err1 := reader.ReadString('\n')
	if err != nil && err1 != io.EOF {
		t.Error(err)
	}
	var prev string

	if str != "" {
		prev = str
		if err != nil {
			t.Error(err)
		}
	}

	for {
		str, err := reader.ReadString('\n')
		if err != nil && err != io.EOF {
			t.Error(err)
		}

		if str != "" {
			if strings.Compare(prev, str) == 0 {
				t.Errorf("нарушен порядок на строке %s", str)
			}
		}

		if err == io.EOF {
			break
		}

	}
}

func TestSortMonth(t *testing.T) {
	col = col6
	err := Sort(fileName, col, sortMonth, false, false)
	if err != nil {
		t.Error(err)
	}

	file, err := os.Open(fileOutName)
	if err != nil {
		t.Error(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	_, err = reader.ReadString('\n')
	if err != nil {
		t.Error(err)
	}

	str, err1 := reader.ReadString('\n')
	if err != nil && err1 != io.EOF {
		t.Error(err)
	}

	strSlice := strings.Split(str, " ")
	var prev, now int

	if str != "" {
		prev = ParseRuMonth(strSlice[col])
		if prev == 0 {
			t.Error("не месяц", strSlice[col])
		}
	}

	for {
		str, err := reader.ReadString('\n')
		if err != nil && err != io.EOF {
			t.Error(err)
		}

		strSlice := strings.Split(str, " ")

		if str != "" {

			now = ParseRuMonth(strSlice[col])
			if now == 0 {
				t.Error("не месяц", strSlice[col])
			}

			if prev > now {
				t.Errorf("нарушен порядок на строке %s", str)
			}
		}

		if err == io.EOF {
			break
		}

	}
}
