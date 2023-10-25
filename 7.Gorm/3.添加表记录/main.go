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
}

type Teacher struct {
	BaseModel        //继承来自BaseModel中的字段
	Name      string `gorm:"type:varchar(32);unique;not null"`
	Tno       int
	Pwd       string `gorm:"type:varchar(100);not null"`
	Tel       string `gorm:"type:char(11);"`
	Birth     *time.Time
	Remark    string `gorm:"type:varchar(255);"`
}

type Class struct {
	BaseModel
	Name      string `gorm:"type:varchar(32);unique;not null"`
	Num       int
	TeacherID int
	Teacher   Teacher //默认情况下,Teacher+ID,然后关联Teacher表的ID字段

}

type Course struct {
	BaseModel
	Name   string `gorm:"type:varchar(32);unique;not null"`
	Credit int
	Period int

	// 多对一
	TeacherID int
	Teacher   Teacher
}

type Student struct {
	BaseModel
	Name   string `gorm:"type:varchar(32);unique;not null"`
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

// 添加表记录
func addRecord(db *gorm.DB) {
	//结构体对象和表记录

	//添加老师对象
	//t1 := Teacher{Name: "wanggang", Tno: 1002, Pwd: "123.com"}
	/*	db.Create(&t1)
		t2 := Teacher{Name: "zhangyuan", Tno: 1003, Pwd: "123.com"}
		db.Create(&t2)
		t3 := Teacher{Name: "huxiaoguang", Tno: 1003, Pwd: "123.com"}
		db.Create(&t3)

		db.Debug().Create(&t2) //指定单条sql执行插入debug信息
		fmt.Println(t2)*/

	//批量创建班级
	/*	c1 := Class{Name: "软件1班", Num: 78, TeacherID: 3}
		c2 := Class{Name: "软件2班", Num: 79, TeacherID: 4}
		c3 := Class{Name: "软件3班", Num: 75, TeacherID: 3}
		c4 := Class{Name: "软件4班", Num: 73, TeacherID: 5}
		classes := []Class{c1, c2, c3, c4}
		db.Create(&classes)
		for i, class := range classes {
			fmt.Println(":::::", i, class)
		}*/

	/*	//批量创建课程
		course01 := Course{Name: "计算机网络技术", Credit: 3, Period: 16, TeacherID: 3}
		course02 := Course{Name: "数据结构", Credit: 2, Period: 24, TeacherID: 3}
		course03 := Course{Name: "数据库", Credit: 5, Period: 18, TeacherID: 4}
		course04 := Course{Name: "模拟电路", Credit: 4, Period: 22, TeacherID: 3}
		course05 := Course{Name: "数字电路", Credit: 2, Period: 19, TeacherID: 5}
		courses := []Course{course01, course02, course03, course04, course05}
		db.Create(&courses)
		for i, course := range courses {
			fmt.Println("::::", i, course)
		}*/

	//多对多添加记录
	//添加学生1
	/*	s1 := Student{Name: "stu001", Sno: 2005, Pwd: "abc.com", ClassID: 15}
		s2 := Student{Name: "stu002", Sno: 2002, Pwd: "abc.com", ClassID: 14}
		s3 := Student{Name: "stu003", Sno: 2003, Pwd: "abc.com", ClassID: 14}
		s4 := Student{Name: "stu004", Sno: 2004, Pwd: "abc.com", ClassID: 15}

		ss := []Student{s1, s2, s3, s4}
		db.Create(&ss)
		fmt.Println("s1::::", s1)
		fmt.Println("s1::::", s1.Class)
		fmt.Println("s1::::", s1.Courses)*/

	//方式1 创建对象即绑定多对多的关系
	/*	var courses []Course
		db.Where("name in ?", []string{"数据结构", "数据库"}).Find(&courses)
		fmt.Println("#######################")
		fmt.Println("courses:", courses)

		s1 := Student{Name: "王武", Sno: 2002, Pwd: "123", ClassID: 13, Courses: courses}
		db.Create(&s1)*/

	//方式2 创建对象即绑定多对多的关系
	/*	var courses []Course
		fmt.Println("#######################")
		db.Where("name in ?", []string{"计算机网络技术", "模拟电路"}).Find(&courses)
		fmt.Println("courses:", courses)
		s1 := Student{Name: "赵六", Sno: 2003, Pwd: "123", ClassID: 15}
		db.Create(&s1)
		db.Model(&s1).Association("Courses").Append(courses)*/

	//方式3.查询对象和既绑定多对多的关系
	/*	var courses []Course
		fmt.Println("#######################")
		db.Where("name in ?", []string{"数据结构", "模拟电路"}).Find(&courses)
		fmt.Println("courses", courses)

		var student = Student{}
		db.Where("name = ?", "李四").First(&student)
		db.Model(&student).Association("Courses").Append(courses)*/

	//解除绑定
	//1.查询课程
	var courses []Course
	fmt.Println("#######################")
	db.Where("name in ?", []string{"计算机网络技术", "模拟电路"}).Find(&courses)
	fmt.Println("courses:", courses)
	//2.查询学生
	var student = Student{}
	db.Where("name = ?", "赵六").First(&student)
	fmt.Println(student)

	//3.解除多对多的绑定关系
	//db.Model(&student).Association("Courses").Delete(courses)
	//db.Model(&student).Association("Courses").Clear()
	//4.查询赵六选修了哪些课程
	var course []Course
	db.Model(&student).Association("Courses").Find(&course)
	fmt.Println("course::::", course)
}

func selectRecord(db *gorm.DB) {
	/*	//(1) 查询全部记录
		var teachers []Teacher
		db.Find(&teachers)
		fmt.Println("###################")
		fmt.Println("teachers:", teachers)
		for i, v := range teachers {
			fmt.Println(i, v.Name, v.Tno)
		}*/

	//(2) 查询单条记录
	/*	course := Course{}
		fmt.Println("###################")
		db.Take(&course)  //limit 1
			db.First(&course) // order by xx limit 1
			db.Last(&course)  // order by xx desc limit 1

		db.Where("credit < ?", 3).Order("credit").Take(&course)
		fmt.Println("course", course)*/

	//(3) where 语句查询
	/*		var courses []Course
		fmt.Println("################")
		//db.Where("credit > ?", 3).Find(&courses)  //SELECT * FROM `courses` WHERE credit > 2 ORDER BY `courses`.`id` LIMIT 1
		//fmt.Println("courses:", courses)
		db.Where("credit between ? and ?", 1, 3).Find(&courses) //SELECT * FROM `courses` WHERE credit between 1 and 3

		fmt.Println("courses:", courses)

		var total int64
		db.Model(&Course{}).Where("credit between ? and ?", 1, 2).Count(&total)
		fmt.Println("total:", total) //SELECT count(*) FROM `courses` WHERE credit between 1 and 2



	var courses []Course
	//db.Where("credit >? and period<?", 1, 20).Find(&courses) // SELECT * FROM `courses` WHERE credit >1 and period<20
	//fmt.Println("courses:", courses)

	//结构体where语句
	fmt.Println("####################")
	//db.Where(Course{Credit: 3, Period: 0}).Find(&courses) //SELECT * FROM `courses` WHERE `courses`.`credit` = 3 AND `courses`.`period` = 16

	//fmt.Println("courses:::", courses)
	//map的where语句
	db.Where(map[string]interface{}{"Credit": 3, "Period": 16}).Find(&courses) // SELECT * FROM `courses` WHERE `Credit` = 3 AND `Period` = 16
	fmt.Println("courses:", courses)*/

	//(4)其他查询语句
	/*	var courses []Course
		fmt.Println("####################")
		db.Select("name,credit").Where("id >?  and id <?", 1, 5).Find(&courses) //SELECT name,credit FROM `courses` WHERE id >1  and id <5
		log.Println(courses)

		var course Course //只保存一个结构图体对象的内容
		fmt.Println("####################")
		db.Select("name,credit").Where("id  = ?", 5).Take(&course) //SELECT name,credit FROM `courses` WHERE id  = 5 LIMIT 1

		log.Println(course)
	*/
	/*	var courses []Course
		db.Order("create_time desc").Limit(2).Offset(2).Find(&courses)
		log.Println(courses)*/

	// 查询每一个teacher_id 的课程个数
	type Ret struct {
		TeacherID int
		Acount    int
	}
	var ret []Ret
	fmt.Println("----------------------")
	//db.Model(&Class{}).Select("teacher_id,Count(*) as acount").Group("teacher_id").Having("acount >?", 1).Scan(&ret) //需要巩固学习
	db.Model(&Class{}).Select("teacher_id,Count(*) as acount").Group("teacher_id").Having("acount >?", 1).Find(&ret) //需要巩固学习
	fmt.Println("ret::::", ret)

}

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

	addRecord(db)
	selectRecord(db)
}
