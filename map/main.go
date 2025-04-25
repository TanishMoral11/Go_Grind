package main

import "fmt"

func main(){
	// name <-> grade

	studentGrades := make(map[string]int)

	studentGrades["Alice"] = 90
	studentGrades["Bob"] = 85
	studentGrades["Charlie"] = 92
	studentGrades["David"] = 88
	studentGrades["Eve"] = 95

	person := map[string]int{
		"John": 80,
		"Jane": 75,
		"Tom": 90,
		"Jerry": 85,
		"Anna": 95,
		
	}

	for index, value := range person {
		fmt.Println("-------------------Key is %s and value is %d", index, value)
	}
}