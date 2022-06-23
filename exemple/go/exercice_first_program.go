// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func main() {

	// Ma reponse
	reponse := "William Quintall"
	// Je transforme ma string en array de rune (des characteres (une lettre) compatible unicode, le chinoix marche la dedans)
	// mais ca affectera pas notre code
	chars := []rune(reponse)

	// crée une variable pour faire la sum
	var sum int
	// On iter sur toute les characters et on les imprimes
	for i := 0; i < len(chars); i++ {
		// Imprime la valeur ascii de la lettre , tu peux regarder avec le tableau https://duckduckgo.com/?q=tableau+ascii&t=newext&atb=v314-1&iax=images&ia=images&iai=https%3A%2F%2Fi.pinimg.com%2Foriginals%2F75%2F28%2Fb1%2F7528b199208cc9078adfa6830be7f072.jpg
		fmt.Println(chars[i])
		// Utilise un cast pour convertir la charactère en integer int(chars[i]) et on l'ajouter a notre sum
		sum = sum + int(chars[i])
	}

	fmt.Printf("Sum du mot %d \n", sum)

	// On peut convertir notre sum en un index dans notre list
	list := []int{1, 2, 3, 43, 50, 6, 23, 35, 25, 23, 53, 51, 145, 61}

	index := sum % len(list)

	fmt.Printf("Index %d : %d \n", index, list[index])
}
