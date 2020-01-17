//router中方法的细节
package main

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

type input struct {
	Title    string    `json:"title" binding:"required"`
	Deadline time.Time `json:"deadline"`
	Extra    string    `json:"extra"`
}

type output struct {
	ID        uint   `json:"id"`
	Title     string `json:"title"`
	Deadline  string `json:"deadline"`
	Extra     string `json:"extra"`
	CreatedAt string `json:"created_at"`
}

func (s *Service) addAffair(c *gin.Context, owner string) (int, interface{}) {
	temp := new(input)
	err := c.BindJSON(temp)
	//DealError(err)
	if err != nil {
		return makeErrorReturn(400, 40000, "Wrong Format Of JSON") //

	}

	tx := s.DB.Begin()
	if tx.Create(&affair{
		Title:    temp.Title,
		Deadline: temp.Deadline,
		Extra:    temp.Extra,
		Owner:    owner,
	}).RowsAffected != 1 {
		tx.Rollback()
		return makeErrorReturn(500, 50000, "Can't Insert Into Database")
	}
	tx.Commit()
	return makeSuccessReturn(200, "")

}

func (s *Service) deleteAffair(c *gin.Context,owner string) (int, interface{}) {
	tempID := c.Query("id")
	if tempID == "" {
		return makeErrorReturn(404, 40400, "Unable To Parse Parameters")
	}
	id, err := strconv.ParseUint(tempID, 10, 32)
	if err != nil {
		return makeErrorReturn(404, 40400, "Unable To Parse Parameters")
	}

	temp := new(affair)
	s.DB.Where("id = ? AND owner = ?", id,owner).Find(temp)
	if temp.Title == "" {
		return makeErrorReturn(404, 40410, "Not Found")
	}

	tx := s.DB.Begin()
	if tx.Where("id = ? AND owner = ?", id,owner).Delete(&affair{}).RowsAffected != 1 {
		tx.Rollback()
		return makeErrorReturn(500, 50000, "Can't Insert Into Database")
	}
	tx.Commit()
	return makeSuccessReturn(200, "")
}

func (s *Service) modifyAffair(c *gin.Context,owner string) (int, interface{}) {
	tempID := c.Query("id")
	if tempID == "" {
		return makeErrorReturn(404, 40400, "Unable To Parse Parameters")
	}
	id, err := strconv.ParseUint(tempID, 10, 32)
	if err != nil {
		return makeErrorReturn(404, 40400, "Unable To Parse Parameters")
	}

	temp := new(affair)
	s.DB.Where("id = ? AND owner = ?", id,owner).Find(temp)
	//s.DB.Where(&affair{Model: gorm.Model{ID: id}}).Find(temp)
	if temp.Title == "" {
		return makeErrorReturn(404, 40410, "Not Found")
	}

	err = c.BindJSON(temp)
	if err != nil {
		//fmt.Println(err)
		return makeErrorReturn(400, 40000, "Wrong Format Of JSON") //
	}
	tx := s.DB.Begin()
	if tx.Model(&affair{}).Where("id = ? AND owner = ?", id,owner).Updates(&affair{
		//if tx.Where(&affair{Model: gorm.Model{ID: id}}).Updates(affair{
		Title:    temp.Title,
		Deadline: temp.Deadline,
		Extra:    temp.Extra,
	}).RowsAffected != 1 {
		tx.Rollback()
		return makeErrorReturn(500, 50000, "Can't Insert Into Database")
	}
	tx.Commit()
	return makeSuccessReturn(200, "")
}

//func (s *Service) findAffair(c *gin.Context) (int, interface{}) {
//	tempID := c.Query("id")
//	if tempID == "" {
//		return makeErrorReturn(404, 40400, "Unable To Parse Parameters")
//	}
//	id, err := strconv.ParseUint(tempID, 10, 32)
//	if err != nil {
//		return makeErrorReturn(404, 40400, "Unable To Parse Parameters")
//	}
//
//	temp := new(affair)
//	s.DB.Where(&affair{Model: gorm.Model{ID: id}}).Find(temp)
	//s.DB.Where("ID = ?", id).Find(temp)
	//if temp.Title == "" {
	//	return makeErrorReturn(404, 40410, "Not Found")
	//}
	//
	//return makeSuccessReturn(200, temp)
//}
