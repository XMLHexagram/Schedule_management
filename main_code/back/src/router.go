//初始化并运行服务器，将细节抽离到SMS.go中
package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type toReturn func(*gin.Context)

func (s *Service) RouterInit() {
	r := gin.Default()
	r.Use(cors.Default())

	all := r.Group("/all")
	{
		all.GET("/affairs", requestEntry(s.getAllAffairs))
		all.GET("/dailyAffairs", requestEntry(s.getDailyEvents))
	}

	operaDaily := r.Group("/operaDaily")
	{
		operaDaily.POST("/add", requestEntry(s.addDailyEvents))
		operaDaily.DELETE("", requestEntry(s.deleteDailyEvents))
		operaDaily.PUT("", requestEntry(s.modifyDailyEvents))
		//operaDaily.GET()
	}

	opera := r.Group("/opera")
	{
		opera.POST("/add", requestEntry(s.addAffair))  //增
		opera.DELETE("", requestEntry(s.deleteAffair)) //删
		opera.PUT("", requestEntry(s.modifyAffair))    //改
		//opera.GET("/find", s.findAffair) //查
	}

	s.Router = r
	err := s.Router.Run(s.Config.Web.Port)
	DealError(err)
}

func requestEntry(f func(c *gin.Context) (int, interface{})) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(f(c))
	}
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
