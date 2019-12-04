//连接数据库
package init

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"main/deal_error"
	"time"
)

type affair struct {
	gorm.Model
	Title    string
	Deadline time.Time
	Extra    string
}

type dailyEvent struct {
	Title string
	Extra string
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
