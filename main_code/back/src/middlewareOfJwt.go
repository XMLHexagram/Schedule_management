package main

import (
	"github.com/gin-gonic/gin"
	"time"
)

type Header struct {
	Token string `header:"token"`
}

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		tempHeader := new(Header)
		err := c.BindHeader(tempHeader)
		if err != nil || tempHeader.Token == "" {
			c.JSON(makeErrorReturn(400, 40010, "Wrong Format Of Header"))
			c.Abort()
			return
		}
		claims, err := ParseToken(tempHeader.Token)
		if err != nil {
			c.JSON(makeErrorReturn(400,40020,"Wrong Format of Token"))
			c.Abort()
			return
		}

		if time.Now().Unix()>claims.ExpiresAt {
			c.JSON(makeErrorReturn(302,30200,"Token Expired"))
			c.Abort()
			return
		}

		c.Next()
	}
}
