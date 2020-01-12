package main

import "github.com/gin-gonic/gin"

func (s *Service) getDailyEvents(c *gin.Context) (int,interface{}){
	data := make([]*dailyEvent,0)
	s.DB.Find(&data)
	return makeSuccessReturn(200,data)
}

func (s *Service) getAllAffairs(c *gin.Context) (int,interface{}){
	data := make([]*affair,0,100)
	s.DB.Find(&data)
	out := make([]*output,0,100)
	for _,v := range data {
		out = append(out, &output{
			ID:       v.ID,
			Title:    v.Title,
			Deadline: v.Deadline.Format("2006-01-02 15:04:05"),
			Extra:    v.Extra,
			CreatedAt:v.CreatedAt.Format("2006-01-02 15:04:05"),
		})
	}
	//fmt.Println(out)
	return makeSuccessReturn(200, out)
}
