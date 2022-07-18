package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	link     = "http://"
	linkSize = len(link)
)

func main() {

	fmt.Println("INPUT:")
	fmt.Print("Here: ")
	scanner := bufio.NewScanner(os.Stdin)
	text := ""
	if scanner.Scan() {
		text = scanner.Text()
	}
	buf := make([]byte, 0, len(text))

	in := true
	for i := 0; i < len(text); i++ {

		if len(text[i:]) >= linkSize && text[i:i+linkSize] == link {
			buf = append(buf, link...)
			i += linkSize
			in = true
		}

		tmp := text[i]
		if tmp == ' ' || tmp == '\n' || tmp == '\t' {
			in = false
		}

		if in {
			buf = append(buf, '*')
		} else {
			buf = append(buf, tmp)
		}
	}
	fmt.Println("OUTPUT:")
	fmt.Println("Here:", string(buf))
}
