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
	var before, after int
	var ignore, fix, lineNum bool
	count := -1

	if os.Args[1] == "-h" || os.Args[1] == "--help" {
		fmt.Println("")
		return
	}

	if strings.HasPrefix(os.Args[len(os.Args)-1], "-") {
		log.Fatal("Укажите файл")
	}

	file, err := os.Open(os.Args[len(os.Args)-1])
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for i := 1; i < len(os.Args)-2; i++ {
		switch os.Args[i] {
		case "-A":
			i++
			if strings.HasPrefix(os.Args[i], "-") {
				log.Fatal("Укажите количество строк")
			}

			after, err = strconv.Atoi(os.Args[i])
			if err != nil {
				log.Fatal(err)
			}
		case "-B":
			i++
			if strings.HasPrefix(os.Args[i], "-") {
				log.Fatal("Укажите количество строк")
			}

			before, err = strconv.Atoi(os.Args[i])
			if err != nil {
				log.Fatal(err)
			}
		case "-C":
			i++
			if strings.HasPrefix(os.Args[i], "-") {
				log.Fatal("Укажите количество строк")
			}

			str := strings.Split(os.Args[i], "+")

			after, err = strconv.Atoi(str[0])
			if err != nil {
				log.Fatal(err)
			}

			before, err = strconv.Atoi(str[1])
			if err != nil {
				log.Fatal(err)
			}
		case "-c":
			i++
			if strings.HasPrefix(os.Args[i], "-") {
				log.Fatal("Укажите количество строк")
			}

			count, err = strconv.Atoi(os.Args[i])
			if err != nil {
				log.Fatal(err)
			}
		case "-i":
			ignore=true
		case "-F":
			fix=true
		case "-n":
			lineNum=true
		default:
			log.Fatalln("не допустимые флаги. Воспользуйтесь -h или --help")
		}
	}

	queueBefore:= NewQueue(before+1)
	beforeS,afterS := Grep(queueBefore,os.Args[len(os.Args)-2],before, after, count, ignore, fix, lineNum, *scanner)
	for i, val := range beforeS {
		for val.Empty() {
			str,_ := val.Dequeue()
			fmt.Println(str)
		}

		for afterS[i].Empty() {
			str,_ := afterS[i].Dequeue()
			fmt.Println(str)
		}
	}
}

func Grep(queueBefore *Queue, pattern string, before, after, count int, ignore, fix, lineNum bool, scanner bufio.Scanner) ([]*Queue, []*Queue) {
	var ig bool

	num := 1

	switch count {
	case 0:
		return nil, nil
	default:
		var resultAfter, resultBefore []*Queue
		queueAfter := NewQueue(after)
		for scanner.Scan() {
			text := scanner.Text()
			if lineNum {
				text = "[" + string(rune(num)) + "]" + text
			}
			queueBefore.Enqueue(text)
			num++

			if fix {
				if ignore {
					if strings.Compare(strings.ToLower(queueBefore.LastEl.Value), strings.ToLower(pattern)) == 0 {
						ig = true
					} else {
						ig = false
					}
				} else {
					if strings.Compare(queueBefore.LastEl.Value, pattern) == 0 {
						ig = true
					} else {
						ig = false
					}
				}
			} else {
				if ignore {
					ig = strings.Contains(strings.ToLower(queueBefore.LastEl.Value), strings.ToLower(pattern))
				} else {
					ig = strings.Contains(queueBefore.LastEl.Value, pattern)
				}
			}

			if ig {
				count--
				resultBefore = append(resultBefore, queueBefore)
				q := *queueBefore
				q.Dequeue()
				rB, rA := Grep(&q, pattern, before, after, count, ignore, fix, lineNum, scanner)
				resultBefore = append(resultBefore, rB...)
				for i := 0; i < after; i++ {
					if !scanner.Scan() {
						break
					}
					text = scanner.Text()
					if lineNum {
						text = "[" + string(rune(num)) + "]" + text
					}
					num++
					queueAfter.Enqueue(text)
				}
				resultAfter = append(resultAfter, queueAfter)
				resultAfter = append(resultAfter, rA...)
				break
			}
		}
		return resultBefore, resultAfter
	}
}
