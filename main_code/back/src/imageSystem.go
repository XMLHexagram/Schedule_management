package main

import (
	"github.com/gin-gonic/gin"
	"math/rand"
	"time"
)


func (s *Service) getImage(c *gin.Context) (int, interface{}) {
	var sum int
	s.DB.Table("image_urls").Count(&sum)
	//fmt.Println(sum)

	rand.Seed(time.Now().Unix())
	tempID := rand.Intn(sum) + 1

	//fmt.Println(tempID)

	tempImageURL := new(imageURL)
	s.DB.Table("image_urls").Where("id = ?",tempID).Find(tempImageURL)
	return makeSuccessReturn(200,gin.H{"URL":tempImageURL.URL,"created_at":tempImageURL.CreatedAt})
}

