package main

import (
	"fmt"
	"os"
	"strings"
)

const A_ASCII_INDEX = 97
const Z_ASCII_INDEX = 122

func main() {

	chars_readed := ""

	for {
		fmt.Scan(&chars_readed)

		chars_readed = strings.ToLower(chars_readed)

		for i := 0; i < len(chars_readed); i++ {
			char_read := chars_readed[i]
			char_int := int(char_read)
			if char_int < A_ASCII_INDEX || char_int > Z_ASCII_INDEX {
				fmt.Println("not a character you stink")
				os.Exit(1)
			}

			fmt.Printf("%c %d \n", char_read, char_int)
		}
	}

}
