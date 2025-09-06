package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

type Employee struct {
	Person
	EmployeeID int
}

func (e Employee) PrintInfo() {
	fmt.Printf("姓名: %s, 年龄: %d, 员工ID: %d\n", e.Name, e.Age, e.EmployeeID)
}

//func main() {
//	emp := Employee{
//		Person:     Person{Name: "张三", Age: 30},
//		EmployeeID: 1001,
//	}
//	emp.PrintInfo()
//}
