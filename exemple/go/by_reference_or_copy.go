package main

import "fmt"

type Person struct {
	age int
}

// Three different way to modify the inner value of a sturcture with a function

// this func is attach the a structe that you initialize
func (p *Person) modify_age(age int) {
	fmt.Printf("Addresse of variable persone with function attach to the structure is %p \n", p)
	p.age = age
}

// this function can be call with a pointer to a struct Person to modify the age
// in the address pointed by the pointer
func modify_age_by_reference(person *Person, age int) {
	fmt.Printf("Addresse of variable persone with function by reference is %p \n", person)
	person.age = age
}

// this function can by call with a struct Person by value instead of reference (copy)
// the function will return back the modify value that you can reassign the your structure
func modify_age_by_copy(person Person, age int) Person {
	fmt.Printf("Addresse of variable persone with function by copy is %p \n", &person)
	person.age = age
	return person
}

func main() {
	person := Person{age: 30}

	fmt.Printf("Addresse of variable persone is %p \n", &person)

	fmt.Println(person)

	person.modify_age(25)

	fmt.Println(person)

	modify_age_by_reference(&person, 40)

	fmt.Println(person)

	person = modify_age_by_copy(person, 55)

	fmt.Printf("Addresse of variable persone is still the same but the value have been copy inside from the other memory space from the other scope %p \n", &person)

	fmt.Println(person)
}
