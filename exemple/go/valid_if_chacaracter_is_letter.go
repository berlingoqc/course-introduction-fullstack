package main

import (
	"fmt"
	"os"
	"strings"
)

// La valeur ascii pour la lettre a en miniscule
const A_ASCII_INDEX = 97

// La valeur ascii pour la lettre z en minuscule
const Z_ASCII_INDEX = 122

func main() {

	// Crée une variable pour stocker ce que l'user entre
	chars_readed := ""

	for {
		// scan une chaine de charactere
		fmt.Scan(&chars_readed)

		// transform en lower case
		chars_readed = strings.ToLower(chars_readed)

		// iter sur chaqu'un des character de ma string
		// on peut utiliser la fonction len() sur les strings comme sur les array
		for i := 0; i < len(chars_readed); i++ {
			// une string peux etre acceder a un index comme avec un array
			char_read := chars_readed[i]
			char_int := int(char_read)
			// Condition si la valeur ascii n'est pas une des lettres de l'alphabet
			// si chart_int est plus petit que A_ASCII_INDEX ou si char_int est plus grand que Z_ASCII_INDEX
			if char_int < A_ASCII_INDEX || char_int > Z_ASCII_INDEX {
				fmt.Println("not a character you stink")
				// os.Exit(0) permet d'arreter le programme
				os.Exit(1)
			}

			fmt.Printf("%c %d \n", char_read, char_int)
		}
	}

}
