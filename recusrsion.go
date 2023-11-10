package main

import "fmt"

func main() {
	fmt.Println(recurssion(500000))

}

func recurssion(number uint) uint {
	if number == 1 {
		return 1
	}
	return (number + recurssion(number-1))
}
