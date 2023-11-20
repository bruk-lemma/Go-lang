//package main

// import "fmt"

func main() {
	fmt.Println(quickSort([]int{-317, -381, -14, -215, -590, -243, -412, 380, -312, 925, 158, -46, 177, 22, -482, 273, 217, 514, -392, 424}))
}

func quickSort(array []int) []int {
	if len(array) < 2 {
		return array
	}
	left, right := 0, len(array)-1

	pivot := array[len(array)-1]

	for i, _ := range array {

		if array[i] < pivot {
			array[left], array[i] = array[i], array[left]
			left++
		}
	}
	array[left], array[right] = array[right], array[left]
	quickSort(array[:left])
	quickSort(array[left+1:])
	return array
}
