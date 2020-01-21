package main

import "github.com/gin-gonic/gin"


func (s *Service) getImage(c *gin.Context) (int, interface{}) {
	tempImageURL := new(imageURL)
	s.DB.Take(tempImageURL)
	return makeSuccessReturn(200,gin.H{"URL":tempImageURL.URL,"created_at":tempImageURL.CreatedAt})
}

