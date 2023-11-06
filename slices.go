package main

import (
	"fmt"
	"reflect"
)

func main() {
	//slice()
	//declareSlice()
	//initializeSliceWithSliceLiterals()
	//declareSliceWithNewKEyword()
	//addToSlice()
	//operationSonSlice()
	//copySlice()
	//loopThroughSlice()
	var intSlice = []int{1, 2, 3, 4, 5, 6}
	fmt.Println(checkIfItemExistsInSlice(intSlice, 3))

}

func slice() {
	var intSlice []int
	var strSlice []string
	fmt.Println(reflect.ValueOf(intSlice).Kind())
	fmt.Println(reflect.ValueOf(strSlice).Kind())
}

func declareSlice() {
	var intSlice = make([]int, 10)        //when length and capacity is the same
	var strSlice = make([]string, 12, 20) //when length and capacity are different
	fmt.Printf("intslice \tLen: %v \tCap: %v\n", len(intSlice), cap(intSlice))
	fmt.Println(reflect.ValueOf(intSlice).Kind())

	fmt.Printf("strslice \tLen: %v \tCap: %v\n", len(strSlice), cap(strSlice))
	fmt.Println(reflect.ValueOf(strSlice).Kind())
}

func initializeSliceWithSliceLiterals() {
	var intSlice = []int{12, 34, 56}
	var strSlice = []string{"usa", "canada", "mexico"}
	fmt.Printf("intSlice \tLen: %v \tCap: %v\n", len(intSlice), cap(intSlice))
	fmt.Printf("strSlice \tLen: %v \tCap: %v\n", len(strSlice), cap(strSlice))
}

func declareSliceWithNewKEyword() {
	var intSlice = new([50]int)[0:10] //define the size to be filled with 0
	fmt.Println(reflect.ValueOf(intSlice).Kind())
	fmt.Printf("intSlice \tLen: %v \tCap: %v\n", len(intSlice), cap(intSlice))
	fmt.Println(intSlice)

}

func addToSlice() {
	var intSLice = make([]int, 4, 5)
	intSLice[0] = 10
	intSLice[1] = 20
	fmt.Println("SLice intSlice: ", intSLice)
	fmt.Printf("Length is %d capcity is %d\n", len(intSLice), cap(intSLice))

	intSLice = append(intSLice, 23, 45, 67, 89, 90)
	fmt.Println("Slice A after appending data: ", intSLice)
	fmt.Printf("Length is %d capacity is %d \n", len(intSLice), cap(intSLice))
}

func operationSonSlice() {
	var strSlice = []string{"India", "Canada", "Japan", "Germany", "Italy"}
	fmt.Println(strSlice)
	//change a specific value
	strSlice[2] = "uk"
	fmt.Println(strSlice)

	//remove an item from slice
	strSlice = removeIndex(strSlice, 3)
	fmt.Println(strSlice)

}
func removeIndex(s []string, index int) []string {
	return append(s[:index], s[index+1:]...)
}

func copySlice() {
	var intSlice = []int{1, 2, 3, 4, 5}
	fmt.Printf("[Slice: intSlice] length is %d capacity is %d \n", len(intSlice), cap(intSlice))
	var intSliceBigger = make([]int, 5, 12) //createa bigger slice
	copy(intSliceBigger, intSlice)
	fmt.Printf("[Slice : intSliceBigger length is %d capacity is %d\n]", len(intSliceBigger), cap(intSliceBigger)) //copy function
	fmt.Println("SLice intSliceBIgger after copying:", intSliceBigger)
	intSliceBigger[3] = 90
	intSliceBigger[4] = 89
	fmt.Println("Slice intSliceBigger after adding elements:", intSliceBigger)
}

func loopThroughSlice() {
	var intSlice = []int{1, 2, 3, 4, 5, 6}
	fmt.Println("\n-------------Example 1 --------------\n")
	fmt.Printf("The type  of the content is %v \n", reflect.ValueOf(intSlice).Kind())
	for index, value := range intSlice {
		fmt.Printf("The index is : %v and the value is : %v\n", index, value)
	}
}

func checkIfItemExistsInSlice(slice interface{}, item interface{}) bool {
	s := reflect.ValueOf(slice)
	if s.Kind() != reflect.Slice {
		panic("Invalide data-type it is slice")

	}
	for i := 0; i < s.Len(); i++ {
		if s.Index(i).Interface() == item {
			return true
		}
	}
	return false
}
