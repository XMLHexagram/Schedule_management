//serve对象内部结构的定义，初始化方法和运行
package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type Service struct {
	Config Conf
	DB     *gorm.DB
	Router *gin.Engine
}

func (s *Service) Init() {
	s.ConfigInit()
	s.DBInit()
	s.RouterInit()
}
