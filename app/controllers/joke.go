package controllers

import (
	"github.com/revel/revel"
	"github.com/wangboo/rest/app/model"
)

type Joke struct {
	*revel.Controller
}

func (c *Joke) Index(page int) revel.Result {
	jokes := model.FindJokeByPage(page)
	return c.RenderJson(jokes)
}

func (c *Joke) Image() revel.Result {
	revel.INFO.Println("id : ", c.Params.Get("id"))
	data := model.GetImageDataByImageId(c.Params.Get("id"))
	revel.INFO.Println("data len : ", len(data))
	return ImageRenderer{Data: data}
}

// 回复
func (c *Joke) Reply(id, msg string) revel.Result {
	model.CreateReply(id, msg)
	return c.RenderText("ok")
}
