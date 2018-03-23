package main

import "fmt"

// The init and post statement are optional.

func main() {
	sum := 1
	for ; sum < 1000; {
		sum += sum
	}
	fmt.Println(sum)
}
