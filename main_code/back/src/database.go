//连接数据库
package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

type User struct {
	gorm.Model
	Username string
	Password string
}

type affair struct {
	gorm.Model
	Title    string
	Deadline time.Time
	Extra    string
	Owner string
}

type dailyEvent struct {
	gorm.Model
	Title string
	Extra string
	Owner string
}

type InvitationCode struct {
	gorm.Model
	Code string
}

func (s *Service) DBInit()  {
	strDb := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		s.Config.DB.User,
		s.Config.DB.Password,
		s.Config.DB.Address,
		s.Config.DB.DBName)
	db,err := gorm.Open("mysql", strDb)
	DealError(err)
	fmt.Println("success connect to database")

	db.AutoMigrate(&affair{})
	db.AutoMigrate(&dailyEvent{})
	db.AutoMigrate(&User{})
	db.AutoMigrate(&InvitationCode{})
	s.DB = db
	//fmt.Println(s.DB)
}
