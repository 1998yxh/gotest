package main

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// 1. 定义 Book 结构体
type Book struct {
	ID     int     `db:"id"`
	Title  string  `db:"title"`
	Author string  `db:"author"`
	Price  float64 `db:"price"`
}

type Employee struct {
	ID         int
	Name       string
	Department string
	Salary     float64
}

func main() {
	// 2. 连接数据库
	dsn := "root:1998yxh@tcp(127.0.0.1:3306)/testdb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := sqlx.Open("mysql", dsn)
	if err != nil {
		log.Fatalln("数据库连接失败:", err)
	}
	defer db.Close()

	//插入数据
	_, err = db.Exec("INSERT INTO books (title, author, price) VALUES (?, ?, ?)", "Go语言编程", "张三", 59.9)
	if err != nil {
		log.Fatalln("插入数据失败:", err)
	}

	// 3. 查询价格大于50元的书籍
	var books []Book
	err = db.Select(&books, "SELECT id, title, author, price FROM books WHERE price > ?", 50)
	if err != nil {
		log.Fatalln("查询失败:", err)
	}

	// 4. 打印结果
	for _, book := range books {
		fmt.Printf("ID:%d, 标题:%s, 作者:%s, 价格:%.2f\n", book.ID, book.Title, book.Author, book.Price)
	}

	//查询部门为 "技术部" 的员工信息
	var employees []Employee
	err = db.Select(&employees, "SELECT id, name,department,salary FROM employees WHERE name = ?", "技术部")
	if err != nil {
		log.Fatalln("查询失败:", err)
	}

	for _, emp := range employees {
		fmt.Printf("ID:%d, 姓名:%s, 部门:%s, 工资:%.2f\n", emp.ID, emp.Name, emp.Department, emp.Salary)
	}

	//查询工资最高的员工
	var topEmployee Employee
	err = db.Get(&topEmployee, "SELECT id, name, department, salary FROM employees ORDER BY salary DESC LIMIT 1")
	if err != nil {
		log.Fatalln("查询失败:", err)
	}

	fmt.Printf("工资最高员工: ID:%d, 姓名:%s, 部门:%s, 工资:%.2f\n", topEmployee.ID, topEmployee.Name, topEmployee.Department, topEmployee.Salary)

}
