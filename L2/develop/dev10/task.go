package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	host := os.Args[len(os.Args)-2]
	port := os.Args[len(os.Args)-1]
	var timeout time.Duration
	var err error

	if len(os.Args) == 4 {
		if strings.HasPrefix(os.Args[1], "--timeout=") {
			timeout, err = time.ParseDuration(strings.TrimPrefix(os.Args[1], "--timeout="))
			if err != nil {
				log.Fatal(err)
			}
		}
	}

	client := &http.Client{Timeout: timeout}
	_, err = client.Head(host + ":" + port)
	if err != nil {
		log.Println(err)
	}

	scanner := bufio.NewScanner(os.Stdin)
	std := make([]string, 0, 0)
	for scanner.Scan() {
		text := scanner.Text()
		if text == "push" {
			body := make([]byte, 0, 0)
			for _, v := range std {
				body = append(body, []byte(v)...)
			}
			req, err := http.NewRequest("POST", host+":"+port, bytes.NewBuffer(body))
			if err != nil {
				log.Fatalln(err)
			}
			resp, err := client.Do(req)
			if err != nil {
				log.Fatalln(err)
			}

			fmt.Println(resp)

			std = make([]string, 0, 0)
			continue
		}
		std = append(std, text)
	}
}
