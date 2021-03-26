package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Geth onward")
	now := time.Now()
	year := now.Year()
	fmt.Println(year)
}
