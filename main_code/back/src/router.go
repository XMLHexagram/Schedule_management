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
		all.GET("/allAffairs", s.getAllAffairs)
		all.GET("/dailyEvents", s.getDailyEvents)
	}


	operaDaily := r.Group("/operaDaily")
	{
		operaDaily.POST("/add",s.addDailyEvents)
		operaDaily.DELETE("",s.deleteDailyEvents)
		operaDaily.PUT("",s.modifyDailyEvents)
		//operaDaily.GET()
	}

	opera := r.Group("/opera")
	{
		opera.POST("/add", s.addAffair)      //增
		opera.DELETE("", s.deleteAffair) //删
		opera.PUT("", s.modifyAffair)    //改
		//opera.GET("/find", s.findAffair) //查
	}

	s.Router = r
	err := s.Router.Run(s.Config.Web.Port)
	DealError(err)
}
