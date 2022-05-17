package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/yqlbu/go-rest-gin/entity"
	"github.com/yqlbu/go-rest-gin/service"
)

type VideoController interface {
	FindAll() []entity.Video
	Save(ctx *gin.Context) entity.Video
}

type controller struct {
	service service.VideoService
}

// constructor function
func New(service service.VideoService) VideoController {
	return &controller{
		service: service,
	}
}

func (c *controller) FindAll() []entity.Video {
	return c.service.FindAll()

}

func (c *controller) Save(ctx *gin.Context) entity.Video {
	var video entity.Video
	// unmarshal the video
	ctx.BindJSON(&video)
	c.service.Save(video)
	return video
}
