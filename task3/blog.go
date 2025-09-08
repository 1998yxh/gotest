package main

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// User 用户模型
// 新增 ArticleCount 字段
type User struct {
	ID           uint   `gorm:"primaryKey"`
	Name         string `gorm:"size:64;not null"`
	Email        string `gorm:"size:128;unique;not null"`
	Posts        []Post // 一对多关系：一个用户有多篇文章
	ArticleCount int    `gorm:"default:0"` // 文章数量统计
}

// Post 文章模型
// 新增 CommentStatus 字段
type Post struct {
	ID            uint      `gorm:"primaryKey"`
	Title         string    `gorm:"size:128;not null"`
	Content       string    `gorm:"type:text"`
	UserID        uint      // 外键，关联User
	Comments      []Comment // 一对多关系：一篇文章有多个评论
	CommentStatus string    `gorm:"size:16;default:'有评论'"` // 评论状态
}

// Comment 评论模型
type Comment struct {
	ID      uint   `gorm:"primaryKey"`
	Content string `gorm:"type:text;not null"`
	PostID  uint   // 外键，关联Post
}

func main() {
	dsn := "root:1998yxh@tcp(127.0.0.1:3306)/testdb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln("数据库连接失败:", err)
	}

	//自动迁移，创建表
	err = db.AutoMigrate(&User{}, &Post{}, &Comment{})
	if err != nil {
		log.Fatalln("自动迁移失败:", err)
	}

	//插入测试数据
	user := User{Name: "张三", Email: "zhangsan@example.com"}
	db.Create(&user)

	post1 := Post{Title: "第一篇博客", Content: "这是第一篇博客内容", UserID: user.ID}
	post2 := Post{Title: "第二篇博客", Content: "这是第二篇博客内容", UserID: user.ID}
	db.Create(&post1)
	db.Create(&post2)

	comment1 := Comment{Content: "写得不错！", PostID: post1.ID}
	comment2 := Comment{Content: "很有帮助，谢谢！", PostID: post1.ID}
	comment3 := Comment{Content: "期待更多内容", PostID: post2.ID}
	db.Create(&comment1)
	db.Create(&comment2)
	db.Create(&comment3)

	log.Println("表创建成功并插入了测试数据")

	//var users []User
	//db.Preload("Posts.Comments").Find(&users, 1)
	//fmt.Println(users)
	//
	//// 查询评论数量最多的文章信息
	//var postWithMostComments Post
	//var commentCount int64
	//db.Model(&Post{}).
	//	Select("posts.*, COUNT(comments.id) as comment_count").
	//	Joins("left join comments on comments.post_id = posts.id").
	//	Group("posts.id").
	//	Order("comment_count desc").
	//	Limit(1).
	//	Scan(&postWithMostComments)
	//
	//// 查询该文章的评论数量
	//db.Model(&Comment{}).Where("post_id = ?", postWithMostComments.ID).Count(&commentCount)
	//
	//fmt.Printf("评论最多的文章: ID=%d, 标题=%s, 评论数=%d\n", postWithMostComments.ID, postWithMostComments.Title, commentCount)

}

// Post 创建钩子：创建文章时自动更新用户的文章数量
func (p *Post) AfterCreate(tx *gorm.DB) (err error) {
	err = tx.Model(&User{}).Where("id = ?", p.UserID).UpdateColumn("article_count", gorm.Expr("article_count + 1")).Error
	return
}

// Comment 删除钩子：删除评论后检查评论数并更新文章评论状态
func (c *Comment) AfterDelete(tx *gorm.DB) (err error) {
	var count int64
	err = tx.Model(&Comment{}).Where("post_id = ?", c.PostID).Count(&count).Error
	if err != nil {
		return
	}
	if count == 0 {
		err = tx.Model(&Post{}).Where("id = ?", c.PostID).Update("comment_status", "无评论").Error
	}
	return
}
