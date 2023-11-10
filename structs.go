// package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"reflect"
// 	//"golang.org/x/tools/go/analysis/passes/assign"
// )

func main() {
	// structFundamntals()
	// creatingStrucLiterals()
	//structIntializeWithPointer()
	//nestedStructs()
	// tag_definition()
	// e := Employee{
	// 	FirstName: "Mark",
	// 	LastName:  "Jones",
	// 	Email:     "mark@gmail.com",
	// 	Age:       25,
	// 	MonthlySalary: []Salary{
	// 		Salary{
	// 			Basic: 15000.00,
	// 			HRA:   5000.00,
	// 			TA:    2000.00,
	// 		},
	// 		Salary{
	// 			Basic: 16000.00,
	// 			HRA:   5000.00,
	// 			TA:    2100.00,
	// 		},
	// 		Salary{
	// 			Basic: 17000.00,
	// 			HRA:   5000.00,
	// 			TA:    2200.00,
	// 		},
	// 	},
	// }
	//fmt.Println(e.EmpInfo())
	// type Employee struct {
	// 	Name string
	// 	Age int
	//  }
	//  func (obj *Employee) Info() {
	// 	if obj.Name == ""{
	// 		obj.Name = "john doe"
	// 	}
	// 	if obj.Age == 0 {
	// 		obj.Age = 25
	// 	}
	//  }
	//assignDefaultValuesToStructFields()
	//comparingStructs()
	copyStructTypeUsingValueAndPointer()
}

func structFundamntals() {
	type rectangle struct {
		length  float64
		breadth float64
		color   string

		geometry struct {
			area      float64
			perimeter float64
		}
	}

	var rect rectangle //dot notation
	rect.length = 10
	rect.breadth = 23
	rect.color = "Green"

	rect.geometry.area = rect.length * rect.breadth
	rect.geometry.perimeter = 2 * (rect.breadth + rect.length)

	fmt.Println(rect)
	fmt.Println("Area:\t", rect.geometry.area, "meter square")
	fmt.Println("Permeter:\t", rect.geometry.perimeter, "meters")
}

func creatingStrucLiterals() {

	type rectangle struct {
		length  int
		breadth int
		color   string
	}

	var rect1 = rectangle{10, 23, "green"}
	fmt.Println(rect1)

	var rect2 = rectangle{length: 34, color: "blue"} //breadth vlaue skipped
	fmt.Println(rect2)

	var rect3 = rectangle{12, 23, "Gteen"}
	fmt.Println(rect3)

	//creating structs using new key word
	rectInstance := new(rectangle) //is a pointer to an instace of rectangle
	rectInstance.color = "black"
	rectInstance.breadth = 4
	rectInstance.length = 5
	fmt.Println("new Rectangle : ", rectInstance)

	var rectInstance2 = new(rectangle) //is an instance of rectangle
	rectInstance2.breadth = 34
	rectInstance2.length = 45
	rectInstance2.color = "purple"

	fmt.Println("Rectangle instace secod:", rectInstance2)
}

func structIntializeWithPointer() {
	type rectangle struct {
		length  int
		breadth int
		color   string
	}
	var rect1 = &rectangle{12, 34, "tree"} //cant skip any value
	fmt.Println(rect1)

	var rect2 = &rectangle{}
	rect2.length = 89
	//rect2.breadth = 56 //breadth skipped
	rect2.color = "Industry black"
	fmt.Println(rect2)

	var rect3 = &rectangle{}
	//(*rect3).breadth = 10 //length skipped
	(*rect3).length = 78
	(*rect3).color = "yellow"
	fmt.Println(rect3)

}

func nestedStructs() {

	type Salary struct {
		Basic, HRA, TA float64
	}

	type Employee struct {
		FirstName, LastName, Email string
		Age                        int
		MonthlySalary              []Salary
	}
	e := Employee{
		FirstName: "Mark",
		LastName:  "Jones",
		Email:     "mark@gmail.com",
		Age:       25,
		MonthlySalary: []Salary{
			Salary{
				Basic: 15000.00,
				HRA:   5000.00,
				TA:    2000.00,
			},
			Salary{
				Basic: 16000.00,
				HRA:   5000.00,
				TA:    2100.00,
			},
			Salary{
				Basic: 17000.00,
				HRA:   5000.00,
				TA:    2200.00,
			},
		},
	}
	fmt.Println(e.FirstName, e.LastName)
	fmt.Println(e.Age)
	fmt.Println(e.Email)
	fmt.Println(e.MonthlySalary[0])
	fmt.Println(e.MonthlySalary[1])
	fmt.Println(e.MonthlySalary[2])

}

func tag_definition() {

	type Employee struct {
		FirstName string `json:"firstname"`
		LastName  string `json:"lastname"`
		City      string `json:"city"`
	}
	json_string := `
    {
        "firstname": "Rocky",
        "lastname": "Sting",
        "city": "London"
    }`

	emp1 := new(Employee)
	json.Unmarshal([]byte(json_string), emp1)
	fmt.Println(emp1)

	emp2 := new(Employee)
	emp2.FirstName = "Ramesh"
	emp2.LastName = "Soni"
	emp2.City = "Mumbai"
	jsonStr, _ := json.Marshal(emp2)
	fmt.Printf("%s\n", jsonStr)
}

// func addMethodToStruct() {
// 	type Salary struct {
// 		Basic, HRA, TA float64
// 	}

// 	type Employee struct {
// 		FirstName, LastName, Email string
// 		Age                        int32
// 		MonthlySalary              []Salary
// 	}
// 	func (e Employee) EmpInfo() string {
// 		fmt.Println(e.FirstName,e.LastName)
// 		fmt.Println(e.Age)
// 		fmt.Println(e.Email)
// 		for _, info := range e.MonthlySalary{
// 			fmt.Println("================")
// 			fmt.Println(info.Basic)
// 			fmt.Println(info.HRA)
// 			fmt.Println(info.TA)
// 		}
// 		return "----------------------"

// 	}
// }

func assignDefaultValuesToStructFields() {
	/*
	 Method of assigning a custom default value can be achieve by using
	 constructor function. Instead of creating a struct directly,
	 the Info function can be used to create an Employee struct with a custom default
	 value for the Name and Age field.*/

	// emp1 := Employee{Name: "Mr.fred"}
	// emp1.Info()
	// fmt.Println(emp1)

	// emp2 := Employee{Age: 25}
	// emp2.Info
	// fmt.Println(emp2)

}

func comparingStructs() {
	type rectangel struct {
		length  float64
		breadth float64
		color   string
	}
	var rect1 = rectangel{10, 20, "Green"}
	rect2 := rectangel{length: 34, breadth: 55, color: "black"}

	if rect1 == rect2 {
		fmt.Println("True")
	} else {
		fmt.Printf("False react1: %v, rect2: %v\n", reflect.ValueOf(rect1).Kind(), reflect.ValueOf(rect2).Kind())
	}
	rect3 := new(rectangel)
	var rect4 = &rectangel{}

	if rect3 == rect4 {
		fmt.Println("True")
	} else {
		fmt.Printf("False rect3: %v, rect4: %V\n", reflect.ValueOf(rect3).Kind(), reflect.ValueOf(rect4).Kind())
	}

}

func copyStructTypeUsingValueAndPointer() {

	type rectangle struct {
		lenght float64
		width  float64
		color  string
	}

	r1 := rectangle{10, 20, "Green"}
	fmt.Println(r1)

	r2 := r1
	r2.color = "Pink"
	fmt.Println(r2)

	r3 := &r1
	r3.color = "Red" //uses reference of r1 so
	//when this line executes it
	//changes the color of r1 to red because r3 is a pointer to r1
	fmt.Println(r3)
	fmt.Println(r1)
}
