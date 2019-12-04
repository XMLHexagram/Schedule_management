//router中方法的细节
package init

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	dealError "main/deal_error"
	"strconv"
)

type input struct {
	Title    string `json:"title"`
	Deadline string `json:"deadline"`
	Extra    string `json:"tips"`
}

func (s *Serve) getDailyEvents(c *gin.Context) {
	data := make([]*dailyEvent,100)
	s.DB.Find(&data)
	c.JSON(makeSuccessReturn(200,data))
	return
}

func (s *Serve) getAllAffairs(c *gin.Context) {
	data := make([]*affair, 100)
	s.DB.Find(&data)
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
		Title:    temp.Title,
		Deadline: temp.Deadline,
		Extra:    temp.Extra,
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
	id, err := analysisID(c)
	if err != nil {
		return
	}
	temp := new(affair)
	s.DB.Where(&affair{Model: gorm.Model{ID: id}}).Find(temp)
	if temp.Title == "" {
		c.JSON(makeErrorReturn(300, 30000, "affair doesn't exist"))
		return
	}

	tx := s.DB.Begin()
	if tx.Where(&affair{Model: gorm.Model{ID: id}}).Delete(&affair{}).RowsAffected != 1 {
		tx.Rollback()
		c.JSON(makeErrorReturn(400, 40000, "can't delete it"))
		return
	}
	tx.Commit()
	c.JSON(makeSuccessReturn(200, ""))
	return
}

func (s *Serve) modifyAffair(c *gin.Context) {
	id, err := analysisID(c)
	if err != nil {
		return
	}

	temp := new(affair)
	s.DB.Model(&affair{}).Where(&affair{Model: gorm.Model{ID: id}}).Find(temp)
	//s.DB.Where(&affair{Model: gorm.Model{ID: id}}).Find(temp)
	if temp.Title == "" {
		c.JSON(makeErrorReturn(300, 30000, "affair doesn't exist"))
		return
	}

	err = c.BindJSON(temp)
	if err != nil {
		c.JSON(makeErrorReturn(300, 30000, "wrong format"))
		return
	}
	tx := s.DB.Begin()
	if tx.Model(&affair{}).Where(&affair{Model:gorm.Model{ID: id}}).Updates(affair{
	//if tx.Where(&affair{Model: gorm.Model{ID: id}}).Updates(affair{
		Title:    temp.Title,
		Deadline: temp.Deadline,
		Extra:    temp.Extra,
	}).RowsAffected != 1 {
		tx.Rollback()
		c.JSON(makeErrorReturn(400, 40000, "can't update it"))
		return
	}
	tx.Commit()
	c.JSON(makeSuccessReturn(200, 20000))
	return
}

func (s *Serve) findAffair(c *gin.Context) {
	id, err := analysisID(c)
	if err != nil {
		return
	}
	fmt.Println(id)
	temp := new(affair)
	//s.DB.Where(&affair{Model: gorm.Model{ID: id}}).Find(temp)
	s.DB.Where("ID = ?",id).Find(temp)
	if temp.Title == "" {
		c.JSON(makeErrorReturn(300, 30000, "affair not exist"))
		return
	}

	c.JSON(makeSuccessReturn(200, temp))
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

func analysisID(c *gin.Context) (uint, error) {
	var id int
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(makeErrorReturn(300, 30000, "wrong format"))
		return uint(id), err
	} else {
		return uint(id), err
	}
}
