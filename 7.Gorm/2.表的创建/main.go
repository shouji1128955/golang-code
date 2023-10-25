package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

type BaseModel struct {
	ID         int        `gorm:"primaryKey"`
	CreateTime *time.Time `gorm:"autoCreateTime"`
	UpdateTime *time.Time `gorm:"autoCreateTime"`
	Name       string     `gorm:"type:varchar(32);unique;not null"`
}

type Teacher struct {
	BaseModel //继承来自BaseModel中的字段
	Tno       int
	Pwd       string `gorm:"type:varchar(100);not null"`
	Tel       string `gorm:"type:char(11);"`
	Birth     *time.Time
	Remark    string `gorm:"type:varchar(255);"`
}

type Class struct {
	BaseModel
	Num       int
	TeacherID int
	Teacher   Teacher //默认情况下,Teacher+ID,然后关联Teacher表的ID字段

}

type Course struct {
	BaseModel
	Credit int
	Period int

	// 多对一
	TeacherID int
	Teacher   Teacher
}

type Student struct {
	BaseModel
	Sno    int
	Pwd    string `gorm:"type:varchar(100);not null"`
	Tel    string `gorm:"type:char(11);"`
	Gender byte   `gorm:"default:1"`
	Birth  *time.Time
	Remark string `gorm:"type:varchar(255);"`

	// 多对一
	ClassID int
	Class   Class `gorm:"foreignKey:ClassID"`
	// 多对多
	Courses []Course `gorm:"many2many:student2course;constraint:OnDelete:CASCADE;"`
}

// constraint:OnDelete:CASCADE 添加此配置,就是连带外键的数据就自动删除

func main() {
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	dsn := "root:Zlq1802909990.@tcp(www.zlqit.com:6666)/grom?charset=utf8mb4&parseTime=True&loc=Local"

	//创建日志对象
	// 创建日志对象
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			//SlowThreshold: time.Second,   // 慢 SQL 阈值
			LogLevel: logger.Info, // Log level
		},
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger, // 日志配置
	})
	if err != nil {
		panic("数据库连接失败!")
	}
	fmt.Println(db, err)

	//迁移表
	db.AutoMigrate(&Teacher{})
	db.AutoMigrate(&Class{})
	db.AutoMigrate(&Course{})
	db.AutoMigrate(&Student{})
}
