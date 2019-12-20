package main

import (
	"github.com/gin-gonic/gin"
	"strconv"
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

func (s *Service) addDailyEvents(c *gin.Context) {
	temp := new(dailyInput)
	err := c.BindJSON(temp)
	if err != nil {
		c.JSON(makeErrorReturn(400, 40000, "Bad Request"))
		return
	}

	tx := s.DB.Begin()
	if tx.Create(dailyEvent{
		Title: "",
		Extra: "",
	}).RowsAffected != 1 {
		tx.Rollback()
		c.JSON(makeErrorReturn(500,50000,"Internal Server Error"))
		return
	}
	tx.Commit()
	c.JSON(makeSuccessReturn(200,""))
	return
}

func (s *Service) deleteDailyEvents(c *gin.Context) {
	tempID := c.Query("id")
	if tempID == "" {
		c.JSON(makeErrorReturn(400,40000,"Bad Request"))
		return
	}
	id,err := strconv.ParseUint(tempID,10,32)
	if err != nil {
		c.JSON(makeErrorReturn(400,40000,"Bad Request"))
		return
	}

	
}

func (s *Service) modifyDailyEvents(c *gin.Context) {
}
