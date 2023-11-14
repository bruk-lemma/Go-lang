// package main

// import (
// 	"fmt"
// )

func main() {
	typeThatSatisfiesMultipleInterfaces()
	InterfaceWithSimilarMethods()
}

//	func createInterface() {
//		type Employee interface {
//			PrintName(name string) string
//			PrintAddress(id int)
//			PrintSalary(basic int, tax int) float64
//		}
//	}
type Pentagon int

func (p Pentagon) Perimeter() {
	fmt.Println("Perimeter of Pentagon", 5*p)
}

func (p Pentagon) NumberOfSide() {
	fmt.Println("Pentagon has 5 sides")
}

func typeThatSatisfiesMultipleInterfaces() {
	type Polygons interface {
		Perimeter()
	}

	type Object interface {
		NumberOfSide()
	}

	var p Polygons = Pentagon(50)
	p.Perimeter()

	var o Pentagon = Pentagon(50)
	o.NumberOfSide()

	var obj Object = Pentagon(50)
	obj.NumberOfSide()

	// Type assertion to access Pentagon-specific methods
	var pent Pentagon = obj.(Pentagon)
	pent.Perimeter()
}

type Car string

func (c Car) Structure() []string {
	var parts = []string{"ECU", "Engine", "Air Filters", "Wipers", "Gas Task"}
	return parts
}

func (c Car) Speed() string {
	return "200 Km/Hrs"
}

type Man string

func (m Man) Structure() []string {
	var parts = []string{"Brain", "Heart", "Nose", "Eyelashes", "Stomach"}
	return parts
}

func (m Man) Performance() string {
	return "8 Hrs/Day"
}
func InterfaceWithSimilarMethods() {

	type Vehicle interface {
		Structure() []string // Common Method
		Speed() string
	}

	type Human interface {
		Structure() []string // Common Method
		Performance() string
	}
	var bmw Vehicle
	bmw = Car("World Top Brand")

	var labour Human
	labour = Man("Software Developer")

	for i, j := range bmw.Structure() {
		fmt.Printf("%-15s <=====> %15s\n", j, labour.Structure()[i])
	}
}
