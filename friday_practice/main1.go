package main

import "fmt"

/*
	type geometry interface {
		area() float64
		perim() float64
	}

	type rect struct {
		width, height float64
	}

	type circle struct {
		radius float64
	}

	func (r rect) area() float64 {
		return r.width * r.height
	}

	func (r rect) perim() float64 {
		return 2*r.width + 2*r.height
	}

	func (c circle) area() float64 {
		return math.Pi * c.radius * c.radius
	}

	func (c circle) perim() float64 {
		return 2 * math.Pi * c.radius
	}

	func measure(g geometry) {
		fmt.Println(g)
		fmt.Println(g.area())
		fmt.Println(g.perim())
	}

	func detectCircle(g geometry) {
		if c, ok := g.(circle); ok {
			fmt.Println("circle with radius", c.radius)
		}
	}

	func main() {
		r := rect{width: 3, height: 4}
		c := circle{radius: 5}

		measure(r)
		measure(c)

		detectCircle(r)
		detectCircle(c)
	}
*/
type Address struct {
	City    string
	Country string
}
type employee struct {
	Name        string
	Designation string
	Salary      int
	Street      string
	Address
}

func main() {
	//var variable int
	//variable = 5
	//	var variable1 string
	//	variable1 = "muhammad sohail"
	//	fmt.Printf("my name is %v & my phone number starts from the %v", variable1, variable)
	employeedata := employee{
		Name:        "Muhammad Sohail",
		Designation: "Software Engineer",
		Salary:      100000,
		Street:      "Civil lines, Street 1",
		Address: Address{
			City:    "Lahore",
			Country: "Pakistan",
		},
	}
	fmt.Println("Our Employee Details:\n", employeedata.Name, "\n", employeedata.Designation, "\n", employeedata.Salary, "\n", employeedata.Street, "\n", employeedata.Address.City, "\n", employeedata.Address.Country)

}
