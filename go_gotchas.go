// package main

// import (
// 	"fmt"
// 	"unicode/utf8"
// )

func main() {

	//The Go type system will not allow any mathematical operation between integer and float variables.
	var x, y = 13.0, 3.5
	fmt.Println(x / y) //.\error-1.go:7:16: invalid operation: x / y (mismatched types int and float64)

	var z = 12 / 7.7 /*
		In the above program, the right side of = are constants, not variables. Hence compiler will convert 13 into a float type to execute the program.
		This program will print the output */
	fmt.Println(z) //1

	///////////////////////////////////////////////////////////////////////////////////////////////////////////////////

	//Assignment to entry in nil map
	/*
	   Map types are reference types,
	   like pointers or slices,
	   and so the value of rect is nil;
	   it doesn't point to an initialized map. A nil map behaves like an empty map when reading,
	   but attempts to write to a nil map will cause a runtime panic; don't do that.*/

	// fails
	//	var rect1 map[string]int
	//	rect1["height"] = 10
	//	fmt.Println(rect1["height"]) //panic: assignment to entry in nil map

	//works
	var rect map[string]int
	fmt.Println(rect["height"])
	fmt.Println(len(rect))

	idx, key := rect["height"]
	fmt.Println(idx)
	fmt.Println(key)

	//in go raw strings can be enclosed in backticks like `somestring\n\t` and the slash
	//will not have a meaning. This is useful when you want to print some string as it is

	//but if u want to print string with / slash meaning s included use "java\tpaython\n"
	//and the  \t and \n wil beinterepreted as tab and newline respectively

	s := "Go\tJava\nPython"
	fmt.Println(s)

	new := `Go\tJava\nPython`
	fmt.Println(new)

	//The operands to numeric operations must have the
	//same type unless the operation involves shifts or untyped constants.

	var timeout = 3
	fmt.Println(timeout)
	//fmt.Println(timeout * time.Millisecond)

	///string lengthbytes and runes

	data := "We♥Go"
	fmt.Println(len(data)) //the emojo takes three bytes
	fmt.Println(utf8.RuneCountInString(data))
	//7

	//intializing a variable with a nil
	//var data2 string = nil //cannot use nil as type string in assignment

	var data3 *string = nil
	if data3 == nil {
		fmt.Println(data3)
	}

	//floating point multiplication
	var m = 1.39
	fmt.Println(m * m)

	const n = 1.39
	fmt.Println(n * n)

	const r = 1.39
	fmt.Println(r * r)

	//string type conversion

	k := 105
	s = string(k)
	fmt.Println(s)

}
