package main

import (
	"fmt"
)

func main() {
	typeThatSatisfiesMultipleInterfaces()
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
