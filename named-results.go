package main

import "fmt"

//         inputs    named return values
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return   // returns the named return values
}

func main() {
	fmt.Println(split(17))
}
