package main

import (
	"fmt"
	"time"
)

func main() {
	date := time.Now()

	fmt.Printf("Today is : %s", date.Local())
}
