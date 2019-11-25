//初始化并运行服务器，将细节抽离到SMS.go中
package init

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"main/deal_error"
)

type toReturn func(*gin.Context)

func (s *Serve) RouterInit() {
	r := gin.Default()
	r.Use(cors.Default())

	r.GET("/allAffairs", s.getAllAffairs)
	temp := r.Group("/opera")
	{
		temp.POST("/add",s.addAffair) //增
		temp.DELETE("/:id",s.deleteAffair) //删
		temp.PUT("/:id",s.modifyAffair) //改
		temp.GET("/find/:id",s.findAffair) //查
	}

	s.Router = r
	err := s.Router.Run(s.Config.Web.Port)
	dealError.DealError(err)
}