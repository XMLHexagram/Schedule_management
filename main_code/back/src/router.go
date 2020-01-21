//初始化并运行服务器，将细节抽离到SMS.go中
package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

//type toReturn func(*gin.Context)

func (s *Service) RouterInit() {
	r := gin.Default()
	r.Use(CorsMiddleware())

	auth := r.Group("/auth")
	{
		auth.POST("/login",requestEntryDefault(s.login))
		auth.POST("/register",requestEntryDefault(s.register))
	}

	all := r.Group("/all")
	all.Use(JWT())
	{
		all.GET("/affairs", requestEntryWithStatus(s.getAllAffairs))
		all.GET("/dailyAffairs", requestEntryWithStatus(s.getDailyEvents))
	}

	operaDaily := r.Group("/operaDaily")
	operaDaily.Use(JWT())
	{
		operaDaily.POST("/add", requestEntryWithStatus(s.addDailyEvents))
		operaDaily.DELETE("", requestEntryWithStatus(s.deleteDailyEvents))
		operaDaily.PUT("", requestEntryWithStatus(s.modifyDailyEvents))
		//operaDaily.GET()
	}

	opera := r.Group("/opera")
	opera.Use(JWT())
	{
		opera.POST("/add", requestEntryWithStatus(s.addAffair))  //增
		opera.DELETE("", requestEntryWithStatus(s.deleteAffair)) //删
		opera.PUT("", requestEntryWithStatus(s.modifyAffair))    //改
		//opera.GET("/find", s.findAffair) //查
	}

	image := r.Group("/backgroundImage")
	//image.Use(JWT())
	{
		image.GET("",requestEntryDefault(s.getImage))
	}

	s.Router = r
	err := s.Router.Run(s.Config.Web.Port)
	DealError(err)
}

func requestEntryWithStatus(f func(c *gin.Context,owner string) (int, interface{})) gin.HandlerFunc {
	return func(c *gin.Context) {
		tempOwner,exist:=c.Get("owner")
		strOwner := fmt.Sprintf("%v",tempOwner)
		if !exist {
			c.JSON(makeErrorReturn(500,50020,"Middleware Wrong"))
		}

		c.JSON(f(c,strOwner))
	}
}

func requestEntryDefault(f func(c *gin.Context) (int, interface{})) gin.HandlerFunc {
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
