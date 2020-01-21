package main

import (
	"github.com/gin-gonic/gin"
)

type InputUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Code     string `json:"code"`
}

func (s *Service) register(c *gin.Context) (int, interface{}) {
	tempUser := new(InputUser)
	err := c.BindJSON(tempUser)
	if err != nil || tempUser.Username == "" || tempUser.Password == "" {
		return makeErrorReturn(400, 40000, "Wrong Format of JSON")
	}

	tempInvitationCode := new(InvitationCode)
	s.DB.Table("invitation_codes").Where("code = ?", tempUser.Code).Find(tempInvitationCode)
	if tempInvitationCode.ID <= 0 {
		return makeErrorReturn(404,40420,"Invitation Code Wrong")
	}

	dbUser := new(User)
	s.DB.Table("users").Where("username = ?", tempUser.Username).Find(dbUser)
	if dbUser.ID > 0 {
		return makeErrorReturn(400, 40030, "Duplicate username")
	}

	tx := s.DB.Begin()
	if tx.Create(&User{
		Username: tempUser.Username,
		Password: tempUser.Password,
	}).RowsAffected != 1 {
		tx.Rollback()
		return makeErrorReturn(500, 50000, "Can't Insert Into Database")
	}
	tx.Commit()

	token, err := GenerateToken(tempUser.Username)
	if err != nil {
		//fmt.Println(err)
		return makeErrorReturn(500, 50010, "Can't Generate Token")
	}
	return makeSuccessReturn(200, gin.H{"token":token})
}

func (s *Service) login(c *gin.Context) (int, interface{}) {
	tempUser := new(InputUser)
	err := c.BindJSON(tempUser)
	if err != nil || tempUser.Username == "" || tempUser.Password == "" {
		return makeErrorReturn(400, 40000, "Wrong Format of JSON")
	}

	dbUser := new(User)
	s.DB.Table("users").Where("username = ? AND password = ?", tempUser.Username, tempUser.Password).Find(dbUser)
	if dbUser.ID <= 0 {
		return makeErrorReturn(404, 40410, "Username or Password Wrong")
	}

	token, err := GenerateToken(tempUser.Username)
	if err != nil {
		return makeErrorReturn(500, 50010, "Can't Generate Token")
	}

	return makeSuccessReturn(200, gin.H{"token":token})
}
