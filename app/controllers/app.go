package controllers

import (
	"github.com/revel/revel"
	"github.com/wangboo/asset"
	"github.com/wangboo/rest/app/model"
	"gopkg.in/mgo.v2/bson"
	"io/ioutil"
	"log"
	"os"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	txt := asset.GetHTMLText("index.html")
	return c.RenderText(txt)
}

func (c App) Create() revel.Result {
	joke := &model.Joke{Content: "haha", From: "qiubai"}
	s := model.Session()
	defer s.Close()
	col := model.CollectionJoke(s)
	col.Insert(joke)
	return c.RenderText("ok")
}

func (c App) All() revel.Result {
	s := model.Session()
	defer s.Close()
	col := model.CollectionJoke(s)
	jokes := []model.Joke{}
	col.Find(bson.M{}).All(&jokes)
	return c.RenderJson(&jokes)
}

func (c App) CreateImage() revel.Result {
	s := model.Session()
	defer s.Close()
	col := model.CollectionImage(s)
	file, _ := os.Open("/Users/wangboo/Desktop/a.jpg")
	defer file.Close()
	data, _ := ioutil.ReadAll(file)
	image := &model.Image{Data: data}
	col.Insert(image)
	log.Println("id = ", image.Id)
	return c.RenderText("ok")
}

func (c App) Image() revel.Result {
	s := model.Session()
	defer s.Close()
	col := model.CollectionImage(s)
	image := &model.Image{}
	col.Find(bson.M{}).One(image)
	return ImageRenderer{Data: image.Data}
}

func (c App) Download() revel.Result {
	model.DownloadHtml()
	return c.RenderText("ok")
}

func (c App) Grape() revel.Result {
	go model.Grape(1, 10)
	return c.RenderText("ok")
}

func (c App) Joke() revel.Result {
	s := model.Session()
	defer s.Close()
	jokeCol := model.CollectionJoke(s)
	joke := &[]model.Joke{}
	jokeCol.Find(nil).All(joke)
	return c.RenderJson(joke)
}
