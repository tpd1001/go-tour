package main

import "fmt"

// At that point you can drop the semicolons: C's while is spelled for in Go.

func main() {
	sum := 1
	for sum < 1000 {
		sum += sum
		//fmt.Println(sum)
	}
	// if note is that the default record separator is a space (as usual)
	fmt.Println("total:",sum)
}
