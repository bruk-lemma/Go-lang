// package main

// import (
// 	"fmt"
// 	"sort"
// )

func main() {
	mapIntialize()

}

func mapIntialize() {

	//map declaration
	var employee = map[string]int{"Mark": 12, "sandy": 43}
	fmt.Println(employee)

	//empty map declaration
	var students = map[int]string{}
	fmt.Println(students)         //map[]
	fmt.Printf("%T \n", students) //map[int]string

	//declaration using make
	var colors = make(map[int]string)
	fmt.Printf("Colors is %T and value is %v\n", colors, colors)

	//length of a map
	employee["bruk"] = 23
	employee["lemma"] = 45
	fmt.Println("length of employee map is :", len(employee))

	//Accessing Items
	fmt.Println(employee["bruk"])
	fmt.Println(employee["sandy"])

	//Adding Items
	employee["kahsay"] = 34
	fmt.Printf("The updated map employee is : %v\n", employee)

	//update Items
	employee["kahsay"] = 67
	fmt.Printf("The updated employee map is \n : %v \n", employee)

	//delete items
	delete(employee, "kahsay")
	fmt.Printf("The updated employee map is: %v \n", employee)

	//iterator over map
	for key, value := range employee {
		fmt.Printf("The age of %v is : %v \n", key, value)
	}

	//truncate a map
	var Ballendior = map[string]int{"messi": 8, "ronaldo": 5, "modric": 1, "neymar": 0, "kaka": 1}
	fmt.Println("Ballendior map before truncate: ", Ballendior)

	//Methode - I
	for k := range Ballendior {
		delete(Ballendior, k)
	}

	fmt.Println("Ballendior after truncate: ", Ballendior)

	//Methode II
	Ballendior = make(map[string]int)

	//sort map keys
	var unsortedMap = map[string]int{"India": 20, "Canada": 70, "Germany": 15}
	keys := make([]string, 0, len(unsortedMap))
	for k := range unsortedMap {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		fmt.Println(k, unsortedMap[k])
	}

	var unsortedStudents = map[int]string{1: "messi", 2: "bosquets", 3: "inesta", 4: "xavi", 5: "suarez", 6: "neymar"}
	var stud = make([]string, 0, len(unsortedStudents))
	fmt.Printf("The unsorted map of students\n", unsortedStudents)
	for _, value := range unsortedStudents {
		stud = append(stud, value)
	}

	//sort slice of values
	sort.Strings(stud)
	fmt.Println("Sorted ------- values\n")
	for _, value := range stud {
		fmt.Println(value)
	}

	//Merge maps

	var first = map[string]int{"a": 1, "b": 2, "c": 3}
	var second = map[string]int{"a": 4, "e": 5, "c": 6, "d": 8}

	for k, v := range second {
		first[k] = v
	}

	fmt.Println("Merged----------Maps-----")
	fmt.Println(first)

}
