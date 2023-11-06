//package main

import (
	"fmt"
	"reflect"
)

func main() {
	var intArray [5]int
	var strArray [5]string
	intArray[0] = 5
	strArray[3] = "bruk"

	// fmt.Println(reflect.ValueOf(intArray).Kind())
	// fmt.Println(reflect.ValueOf(strArray).Kind())
	// fmt.Println(strArray[3])
	// fmt.Println(intArray[0])
	// intializeElements()
	// initializewithellipse()
	// intializeSpecifcElements()
	// loopThroughIndexedArray()
	copyAnArray()
	strArraycheck := [5]string{"India", "canada", "japan", "italy"}
	//var testString = "bruk"
	fmt.Println(checkElementExists(strArraycheck, "canada"))
	fmt.Println(checkElementExists(strArraycheck, "kenya"))
	filterArray()

}

func intializeElements() {
	x := [5]int{12, 45, 67, 89, 23}   //intialize with values
	var y [5]int = [5]int{43, 67, 78} //partial asssignment
	fmt.Println(x)
	fmt.Println(y)
}

func initializewithellipse() {
	var newArray = [...]int{1, 2, 3, 4, 5, 6}
	fmt.Println(len(newArray))
	fmt.Println(newArray)
}

func intializeSpecifcElements() {
	var intArray = [4]int{1: 2, 4, 3: 5}
	fmt.Println(intArray)
}
func loopThroughIndexedArray() {
	var intArray = [5]int{10, 20, 30, 40, 50}
	fmt.Println("\n -----------Example 1---------\n")
	for i := 0; i < len(intArray); i++ {
		fmt.Println(intArray[i])
	}

	fmt.Println("\n--------Example 2--------------\n")
	for index, value := range intArray {
		fmt.Println(index, "=>", value)
	}

	fmt.Println("\n-----------Example 3 --------------\n")
	for _, value := range intArray {
		fmt.Println(value)
	}

	j := 0
	fmt.Println("\n-----------Example 4 --------------\n")
	for range intArray {
		fmt.Println(intArray[j])
		j++
	}

}

func copyAnArray() {
	var strArray1 = [3]string{"bruk", "lemma", "kahsay"}
	var strArray2 = strArray1  //data passed by value
	var strArray3 = &strArray1 //data passesd by referece

	fmt.Printf("strArray 2: %v\n", strArray2)
	fmt.Printf("strArray 1: %v\n", strArray1)

	strArray1[0] = "leo Messi"

	fmt.Printf("strArray1: %v \n", strArray1)
	fmt.Printf("strArray2: %v \n", strArray2)
	fmt.Printf("strArray3: %v \n", strArray3)
}

func checkElementExists(arrayType interface{}, item interface{}) bool {
	arr := reflect.ValueOf(arrayType)
	fmt.Println("arra is of type ,%v", arr.Kind())
	if arr.Kind() != reflect.Array {
		panic("invalid data-type")
	}
	for i := 0; i < arr.Len(); i++ {
		if arr.Index(i).Interface() == item {
			return true
		}
	}
	return false

}
func filterArray() {
	var countries = [...]string{"India", "Canada", "Japan", "Germany", "Italy"}
	fmt.Printf("Countries: %v \n", countries)

	fmt.Printf(":2 %v \n", countries[:2])
	fmt.Printf("1:3 %v \n", countries[1:3])
	fmt.Printf("2:5 %v \n", countries[2:5])
	fmt.Printf("0:3 %v \n", countries[0:3])
	fmt.Printf("Last element: %v \n", countries[len(countries)-1])
	fmt.Printf("All elements: %v \n", countries[:])
	fmt.Println(countries[0:])
	fmt.Println(countries[0:len(countries)])
}
