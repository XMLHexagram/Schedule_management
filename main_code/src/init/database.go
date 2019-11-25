//连接数据库
package init

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"main/deal_error"
)

type affair struct {
	gorm.Model
	EventsTitle    string `json:"events_title"`
	EventsDeadline string `json:"events_deadline"`
	ExtraTips      string `json:"extra_tips"`
}

func (s *Serve) DBInit()  {
	strDb := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		s.Config.DB.User,
		s.Config.DB.Password,
		s.Config.DB.Address,
		s.Config.DB.DBName)
	db,err := gorm.Open("mysql", strDb)
	dealError.DealError(err)
	fmt.Println("success connect to database")

	db.AutoMigrate(&affair{})
	s.DB = db
	//fmt.Println(s.DB)
}
