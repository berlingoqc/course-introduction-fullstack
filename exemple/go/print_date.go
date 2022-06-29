package main

import (
	"fmt"
	"time"
)

func main() {
	date := time.Now()

	fmt.Printf("Today is : %s \n", date.Local())

	year, month, day := date.Date()

	fmt.Printf("%d %d %d \n", year, month, day)

}
