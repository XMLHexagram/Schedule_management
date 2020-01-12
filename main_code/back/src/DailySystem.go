package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type dailyInput struct {
	Title string `json:"title" binding:"required"`
	Extra string `json:"extra"`
}

type dailyOutput struct {
	ID        uint   `json:"id"`
	Title     string `json:"title"`
	Extra     string `json:"extra"`
	CreatedAt string `json:"created_at"`
}

func (s *Service) addDailyEvents(c *gin.Context) (int, interface{}) {
	temp := new(dailyInput)
	err := c.BindJSON(temp)
	if err != nil {
		return makeErrorReturn(400, 40000, "Wrong Format Of JSON")

	}

	tx := s.DB.Begin()
	if tx.Create(&dailyEvent{
		Title: temp.Title,
		Extra: temp.Extra,
	}).RowsAffected != 1 {
		tx.Rollback()
		return makeErrorReturn(500, 50000, "Can't Insert Into Database")
	}
	tx.Commit()
	return makeSuccessReturn(200, "")
}

func (s *Service) deleteDailyEvents(c *gin.Context) (int, interface{}) {
	id := c.Query("id")
	if id == "" {
		return makeErrorReturn(404, 40400, "Unable To Parse Parameters")

	}
	temp := new(dailyEvent)

	s.DB.Where("id = ?", id).Find(temp)
	if temp.Title == "" {
		return makeErrorReturn(404, 40410, "Not Found")

	}

	tx := s.DB.Begin()
	if tx.Where("id = ?", id).Delete(&dailyEvent{}).RowsAffected != 1 {
		tx.Rollback()
		return makeErrorReturn(500, 50000, "Can't Insert Into Database")
	}
	tx.Commit()
	return makeSuccessReturn(200, "")
}

func (s *Service) modifyDailyEvents(c *gin.Context) (int, interface{}) {
	id := c.Query("id")
	if id == "" {
		return makeErrorReturn(404, 40400, "Unable To Parse Parameters")
	}

	temp := new(dailyEvent)
	s.DB.Where("id = ?", id).Find(temp)
	//s.DB.Where(&affair{Model: gorm.Model{ID: id}}).Find(temp)
	if temp.Title == "" {
		return makeErrorReturn(404, 40410, "Not Found")
	}

	err := c.BindJSON(temp)
	//fmt.Println(temp)
	if err != nil {
		fmt.Println(err)
		return makeErrorReturn(400, 40000, "Wrong Format Of JSON") //
	}
	tx := s.DB.Begin()
	if tx.Model(&dailyEvent{}).Where("id = ?", id).Updates(&dailyEvent{
		Title: temp.Title,
		Extra: temp.Extra,
	}).RowsAffected != 1 {
		tx.Rollback()
		return makeErrorReturn(500, 50000, "Can't Insert Into Database")

	}
	tx.Commit()
	return makeSuccessReturn(200, 20000)
}
