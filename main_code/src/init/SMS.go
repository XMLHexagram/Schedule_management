//router中方法的细节
package init

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	dealError "main/deal_error"
)

type input struct {
	EventsTitle    string `json:"events_title"`
	EventsDeadline string `json:"events_deadline"`
	ExtraTips      string `json:"extra_tips"`
}

func (s *Serve) getAllAffairs(c *gin.Context) {
	data := s.DB.Where(&affair{}).Find(&affair{})
	c.JSON(makeSuccessReturn(200, data))
	return
}

func (s *Serve) addAffair(c *gin.Context) {
	temp := new(input)
	err := c.BindJSON(temp)
	dealError.DealError(err)
	if err != nil {
		c.JSON(makeErrorReturn(300, 30000, "wrong format"))
		return
	}

	tx := s.DB.Begin()
	if tx.Create(&affair{
		EventsTitle:    temp.EventsTitle,
		EventsDeadline: temp.EventsDeadline,
		ExtraTips:      temp.ExtraTips,
	}).RowsAffected != 1 {
		tx.Rollback()
		c.JSON(makeErrorReturn(400, 40000, "can't add it"))
		return
	}
	tx.Commit()
	c.JSON(makeSuccessReturn(200, ""))
	return
}

func (s *Serve) deleteAffair(c *gin.Context) {
	id,err := analysiID(c)
	if  err != nil{
		return
	}
	temp := new(affair)
	s.DB.Where(gorm.Model{ID:id}).Find(temp)
	if temp.EventsTitle == "" {
		c.JSON(makeErrorReturn(300,30000,"affair doesn't exist"))
		return
	}
	
	tx := s.DB.Begin()
	if tx.Where(gorm.Model{ID:id}).Delete(&affair{}).RowsAffected != 1 {
		tx.Rollback()
		c.JSON(makeErrorReturn(400,40000,"can't delete it"))
		return
	}
	tx.Commit()
	c.JSON(makeSuccessReturn(200,""))
	return
}

func (s *Serve) modifyAffair(c *gin.Context) {

}

func (s *Serve) findAffair(c *gin.Context) {
	id,err := analysiID(c)
	if err != nil{
		return
	}
	temp:= new(affair)
	s.DB.Where(gorm.Model{ID: id}).Find(temp)
	if temp.EventsTitle == "" {
		c.JSON(makeErrorReturn(300,30000,""))
		return
	}

	c.JSON(makeSuccessReturn(200,temp))
	return
}

func makeSuccessReturn(status int, data interface{}) (int, interface{}) {
	return status, gin.H{
		"error": 0,
		"msg":   "success",
		"data":  data,
	}
}

func makeErrorReturn(status int, error int, msg string) (int, interface{}) {
	return status, gin.H{
		"error": error,
		"msg":   msg,
	}
}

func analysiID(c *gin.Context)  (uint,error){
	var id uint
	err := c.ShouldBindUri(&id)
	if err != nil {
		c.JSON(makeErrorReturn(300, 30000, "wrong format"))
		return id,err
	} else {
		return id,err
	}
}