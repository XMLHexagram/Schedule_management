//router中方法的细节
package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"strconv"
	"time"
)

type input struct {
	Title    string `json:"title" binding:"required"`
	Deadline time.Time `json:"deadline"`
	Extra    string `json:"extra"`
}

type output struct {
	ID uint `json:"id"`
	Title    string `json:"title"`
	Deadline string `json:"deadline"`
	Extra    string `json:"extra"`
	CreatedAt string `json:"created_at"`
}


func (s *Service) addAffair(c *gin.Context) {
	temp := new(input)
	err := c.BindJSON(temp)
	//DealError(err)
	if err != nil {
		c.JSON(makeErrorReturn(400, 40000, "Bad Request"))
		return
	}

	tx := s.DB.Begin()
	if tx.Create(&affair{
		Title:    temp.Title,
		Deadline: temp.Deadline,
		Extra:    temp.Extra,
	}).RowsAffected != 1 {
		tx.Rollback()
		c.JSON(makeErrorReturn(500, 50000, "can't add it"))
		return
	}
	tx.Commit()
	c.JSON(makeSuccessReturn(200, ""))
	return
}

func (s *Service) deleteAffair(c *gin.Context) {
	tempID := c.Query("id")
	if tempID=="" {
		c.JSON(makeErrorReturn(400,40000,"Bad Request"))
		return
	}
	id,err := strconv.ParseUint(tempID,10,32)
	if err != nil {
		c.JSON(makeErrorReturn(400,40000,"Bad Request"))
		return
	}

	temp := new(affair)
	s.DB.Where(&affair{Model: gorm.Model{ID: id}}).Find(temp)
	if temp.Title == "" {
		c.JSON(makeErrorReturn(400, 40010, "affair doesn't exist"))
		return
	}

	tx := s.DB.Begin()
	if tx.Where(&affair{Model: gorm.Model{ID: id}}).Delete(&affair{}).RowsAffected != 1 {
		tx.Rollback()
		c.JSON(makeErrorReturn(500, 50000, "can't delete it"))
		return
	}
	tx.Commit()
	c.JSON(makeSuccessReturn(200, ""))
	return
}

func (s *Service) modifyAffair(c *gin.Context) {
	tempID := c.Query("id")
	if tempID=="" {
		c.JSON(makeErrorReturn(400,40000,"Bad Request"))
		return
	}
	id,err := strconv.ParseUint(tempID,10,32)
	if err != nil {
		c.JSON(makeErrorReturn(400,40000,"Bad Request"))
		return
	}

	temp := new(affair)
	s.DB.Model(&affair{}).Where(&affair{Model: gorm.Model{ID: id}}).Find(temp)
	//s.DB.Where(&affair{Model: gorm.Model{ID: id}}).Find(temp)
	if temp.Title == "" {
		c.JSON(makeErrorReturn(400, 40000, "affair doesn't exist"))
		return
	}

	err = c.BindJSON(temp)
	if err != nil {
		fmt.Println(err)
		c.JSON(makeErrorReturn(400, 40000, "wrong format"))
		return
	}
	tx := s.DB.Begin()
	if tx.Model(&affair{}).Where(&affair{Model: gorm.Model{ID: id}}).Updates(affair{
	//if tx.Where(&affair{Model: gorm.Model{ID: id}}).Updates(affair{
		Title:    temp.Title,
		Deadline: temp.Deadline,
		Extra:    temp.Extra,
	}).RowsAffected != 1 {
		tx.Rollback()
		c.JSON(makeErrorReturn(500, 50000, "can't update it"))
		return
	}
	tx.Commit()
	c.JSON(makeSuccessReturn(200, 20000))
	return
}

func (s *Service) findAffair(c *gin.Context) {
	tempID := c.Query("id")
	if tempID=="" {
		c.JSON(makeErrorReturn(400,40000,"Bad Request"))
		return
	}
	id,err := strconv.ParseUint(tempID,10,32)
	if err != nil {
		c.JSON(makeErrorReturn(400,40000,"Bad Request"))
		return
	}

	temp := new(affair)
	//s.DB.Where(&affair{Model: gorm.Model{ID: id}}).Find(temp)
	s.DB.Where("ID = ?",id).Find(temp)
	if temp.Title == "" {
		c.JSON(makeErrorReturn(400, 40010, "affair not exist"))
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

