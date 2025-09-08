package main

import (
	"errors"

	"gorm.io/driver/mysql"
	"gorm.io/gorm/clause"
)
import "gorm.io/gorm"

type Student struct {
	ID    int
	Name  string
	Age   int
	Grade string
}

type Account struct {
	ID      int
	Balance float64
}

type Transaction struct {
	ID            int
	FromAccountID int
	ToAccountID   int
	Amount        float64
}

func main() {
	db, err := gorm.Open(mysql.Open("root:1998yxh@tcp(127.0.0.1:3306)/testdb?charset=utf8mb4&parseTime=True&loc=Local"))
	if err != nil {
		panic(err)
	}
	//db.AutoMigrate(&Student{})

	//student := Student{ID: 1, Name: "张三", Age: 20, Grade: "三年级"}
	//db.Create(student)
	//student2 := Student{ID: 2, Name: "李四", Age: 17, Grade: "一年级"}
	//db.Create(student2)

	//var students Student
	////db.Debug().Where("age > ?", 18).Find(&students)
	////fmt.Println(students)

	//db.Model(&Student{}).Debug().Where("name = ?", "张三").Update("Grade", "四年级")
	//db.Debug().Where("age < 18").Delete(&Student{})

	//db.AutoMigrate(&Account{})
	//db.AutoMigrate(&Transaction{})
	//// 插入账户数据
	//db.Create(&Account{ID: 1, Balance: 1000})
	//db.Create(&Account{ID: 2, Balance: 500})
	//
	//// 插入转账记录
	//db.Create(&Transaction{
	//	FromAccountID: 1,
	//	ToAccountID:   2,
	//	Amount:        100,
	//})

	db.Transaction(func(tx *gorm.DB) error {
		var fromID, toID int = 1, 2
		var amount float64 = 100
		var from, to Account

		// 查询转出账户并加锁
		if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).First(&from, fromID).Error; err != nil {
			return err
		}
		// 检查余额
		if from.Balance < amount {
			return errors.New("余额不足")
		}
		// 查询转入账户并加锁
		if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).First(&to, toID).Error; err != nil {
			return err
		}
		// 扣除转出账户余额
		if err := tx.Model(&Account{}).Where("id = ?", fromID).Update("balance", from.Balance-amount).Error; err != nil {
			return err
		}
		// 增加转入账户余额
		if err := tx.Model(&Account{}).Where("id = ?", toID).Update("balance", to.Balance+amount).Error; err != nil {
			return err
		}
		// 记录转账
		tr := Transaction{
			FromAccountID: fromID,
			ToAccountID:   toID,
			Amount:        amount,
		}
		if err := tx.Create(&tr).Error; err != nil {
			return err
		}
		return nil
	})
}
