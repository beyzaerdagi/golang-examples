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

	for i := 0; i < len(text); i++ {
		buf = append(buf, text[i])

		if len(text[i:]) >= linkSize && text[i:i+linkSize] == link {

		}
	}
	fmt.Println(string(buf))

}
