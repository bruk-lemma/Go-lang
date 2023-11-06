//package main

// import (
// 	"fmt"
// 	"reflect"
// )

// func main() {
// 	//fmt.Println("Hello world")
// 	// varibles()
// 	// outPut()
// 	// arrays()
// 	// fmt.Println(linearSearch([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 79))
// 	// fmt.Println(binarySearch(98, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}))
// 	// forloop()
// 	//infiniteLoop()
// 	//fmt.Println(functionWithReturnType(2, 4))
// 	// fmt.Println(functionWithnoSpecificReturn(2, 4))
// 	// fmt.Println(functWithMultipleReturns(2, 4))
// 	// var age = 20
// 	// var text = "bruk"
// 	// fmt.Println("Before:", text, age)
// 	// passingAdderssValue(&age, &text)
// 	// fmt.Println("After:", text, age)
// 	// fmt.Println(area(9, 9))
// 	// parent()
// 	// anotherParent()
// 	// partial := higher_order_func(2)
// 	// fmt.Println("higerorder", partial(3))
// 	variadicFunction("red", "yellow", "green", "orange")
// 	vardiacfuncWithMultipleArguments("red", "yellow", "green", "orange")
// 	vardiacfuncWithMultipleArguments("red", "yellow", "green", "orange", "blue", "black")
// 	vardiac_with_different_typeof_inputs(1, "red", true, 10.5, []string{"foo", "bar", "baz"},
// 		map[string]int{"apple": 23, "tomato": 13})
// 	defer second()
// 	first()

// }

func varibles() {

	var age int = 34
	var name string = "bruk Lemma"
	var x = 4
	fmt.Printf("The age is %v and the name is %v finally x is %v\n", age, name, x)

	//variable declaration with out intial value
	var unknown string
	var status bool
	var number int
	/*when a varible is declared but not u=intialized the default value
	will be the default value of its type example if it is bool it would be false
	if string "" and if int 0*/

	fmt.Printf("the value of the unknow varible is %v\n", unknown)
	fmt.Println(unknown)
	fmt.Println(status)
	println(number)

	/* the operator = can be used inside or out side a fucntion
	but := can only be used inside a function and a value must be provided
	*/
	//and also when using := we nedd not to sepcift type typ is infered by go
	newValue := "new string"
	fmt.Println(newValue)

	/*
		declaring multiple variables
	*/
	var w, y, z int = 1, 2, 3
	fmt.Printf("the value of the varibles is w: %v  y: %v and z: %v \n", w, y, z)

	/*
		groupes assignment for better readablity
	*/
	var (
		firstName string = "bruk"
		lastName  string = "Lemma"
	)

	fmt.Printf("The value of FirstName : %v and lastName: %v\n", firstName, lastName)
	/*
		This is how to declare constants
		constant can not be changed and we can declare multiple
		constants
	*/
	const kelvinTemp float32 = 273.89
	const (
		celiciosTemp  float32 = 34.891
		feranhiteTemp float32 = 29.45
	)
	fmt.Printf("The temprature out side is %v  kelvin\n", kelvinTemp)
	fmt.Printf("The temprature out side is %v  celicios\n", celiciosTemp)
	fmt.Printf("The temprature out side is %v  feranhite\n", feranhiteTemp)

	/*
		we can have unsigned integers
	*/
	var myage uint32 = 67
	fmt.Printf("My Age is %v\n", myage)

}

func outPut() {
	var age = 778

	/*
		Print prints argumet
	*/
	fmt.Print(age, "\n")
	fmt.Println(age)
	fmt.Printf("the age is %v and its a type of %T\n", age, age)
	fmt.Printf(" age is %v and syntax is %#v\n", age, age)

}

func arrays() {
	var names = [7]string{"name1", "name2", "name3"}
	var ages = []int32{12, 34, 56, 78}
	grades := []int32{3, 7, 8}

	fmt.Printf("The values of the array are %v \n", names)
	fmt.Printf("The ages are %v \n", ages)
	fmt.Printf("The value of grades is %v \n", grades)
	//change the value of an array member
	names[0] = "Bruk lemma"
	/*
		this is how to access the values of an array
	*/
	fmt.Printf("The first value of names is %v\n", names[0])

	/*
		this is how to assieg a certain value to a certain index of an
		array
	*/
	var animals = [6]string{1: "Lion", 3: "tiger", 4: "elephant"}
	fmt.Printf("tha values of animals array is %v\n", animals)
	var primes = [4]int32{3, 7, 3: 11}
	fmt.Printf("Primes: %v\n", primes)

}

func linearSearch(dataList []int, key int) bool {
	for _, item := range dataList {
		if item == key {
			return true
		}

	}
	return false

}

func binarySearch(needle int, haystack []int) bool {
	low := 0
	high := len(haystack) - 1
	for low <= high {
		median := (low + high) / 2
		if haystack[median] < needle {
			low = median + 1
		} else {
			high = median - 1
		}
	}
	if low == len(haystack) || haystack[low] != needle {
		return false
	}
	return true

}

func forloop() {
	var cities = [3]string{"Addis", "Dire", "Hawassa"}
	for index, element := range cities {
		fmt.Printf("the cityIndex  is %v and city : %v\n", index, element)
	}

	/*
		loop over a string
	*/

	name := "bruk"
	for index, element := range name {
		fmt.Printf("The char in the string is %v  and value is %c\n", index, element)
	}
}

func infiniteLoop() {
	var i int32 = 0
	for {
		fmt.Println("hello")
		if i == 10 {
			break
		}
		i++
	}
}

func functionWithReturnType(x int32, y int32) int32 {
	var sum int32 = x + y
	return sum
}

func functionWithnoSpecificReturn(x int32, y int32) (area int32) {
	//here we havenot declared area but we have declared
	//it in the return type so we can use it
	area = (x * y)
	return
}

func functWithMultipleReturns(x int32, y int32) (lateralarea int32, area int32) {
	lateralarea = 2 * (x + y)
	area = x * y
	return
}

func passingAdderssValue(a *int, y *string) {
	*a = *a + 34
	*y = *y + "lemma"
	return
}

var (
	area = func(l int, b int) int {
		return l * b
	}
)

func parent() {
	var age int32 = 34
	func() {
		var myage int32
		myage = age + 67
		fmt.Println("My age is ", myage)
	}()
}

func anotherParent() {
	for i := 10.0; i < 12; i += 10.0 {
		rad := func() float64 {
			return i * 39.370
		}()
		fmt.Printf("%.2f meter = %.2f Inch\n", i, rad)
	}
}

func parameter_to_Higher_order_func(x int32, y int32) int32 {
	return x + y
}

func higher_order_func(x int32) func(int32) int32 {
	return func(y int32) int32 {
		return parameter_to_Higher_order_func(x, y)
	}

}

// func customeDefinedFucntion(){
// 	type First func(int) int
// 	type Second func(int) First

// func squareSum(x int) Second{
// 	return func(y int) first{
// 		return func(z int) int{
// 			return x*x + y*y + z*z
// 			}
// 		}
// 	}
// fmt.Println(squareSum(2)(3)(4))
// }

func variadicFunction(s ...string) {
	fmt.Println(s[0])
	fmt.Println(s[3])

}

func vardiacfuncWithMultipleArguments(s ...string) {
	fmt.Printf("The inputs are: %v\n", s)
}

func vardiac_with_different_typeof_inputs(i ...interface{}) {
	for _, v := range i {
		fmt.Println(v, "--", reflect.ValueOf(v).Kind())
	}

}

func first() {
	fmt.Println("First")
}
func second() {
	fmt.Println("Second")
}
