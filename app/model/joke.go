package model

import (
	"github.com/revel/revel"
	"gopkg.in/mgo.v2/bson"
	"labix.org/v2/mgo"
)

// 笑话
type Joke struct {
	BaseModel
	Id       bson.ObjectId   `bson:"_id"`
	Jid      int             // jokeId
	Content  string          // 内容
	From     string          // 来源
	ImageIds []bson.ObjectId // 图片
}

// 图片
type Image struct {
	Id   bson.ObjectId `bson:"_id"`
	Data []byte
}

// 回复
type Reply struct {
	BaseModel
	Id      bson.ObjectId `bson:"_id" 		json:"-"`
	JokeId  bson.ObjectId `bson:"jokeId" 	json:"-"`
	Content string        `json:"msg"`
}

// 返回给前台对象
type JokeVO struct {
	Id      bson.ObjectId `json:"id"`
	Jid     int           `json:"jid"`
	Content string        `json:"msg"`
	ImageId string        `json:"imgId"`
	Replies *[]Reply      `json:"replies"`
}

const (
	DB_NAME   = "qiushibaike"
	COL_JOKE  = "joke"
	COL_IMAGE = "image"
	COL_REPLY = "reply"
)

func CollectionJoke(s *mgo.Session) *mgo.Collection {
	return s.DB(DB_NAME).C(COL_JOKE)
}

func CollectionImage(s *mgo.Session) *mgo.Collection {
	return s.DB(DB_NAME).C(COL_IMAGE)
}

func Collection(s *mgo.Session, name string) *mgo.Collection {
	return s.DB(DB_NAME).C(name)
}

func FindJokeByPage(page int) []JokeVO {
	s := Session()
	defer s.Close()
	jokeCol := CollectionJoke(s)
	imageCol := CollectionImage(s)
	replyCol := Collection(s, COL_REPLY)
	jokes := []Joke{}
	jokeCol.Find(nil).Sort("-_id").Limit(10).Skip(page * 10).All(&jokes)
	vos := []JokeVO{}
	for i, joke := range jokes {
		vo := JokeVO{Id: joke.Id, Jid: joke.Jid, Content: joke.Content}
		// 查询回复
		replies := []Reply{}
		replyCol.Find(bson.M{"jokeId": joke.Id}).All(&replies)
		vo.Replies = &replies
		vos = append(vos, vo)
		if len(joke.ImageIds) > 0 {
			image := &Image{}
			imageCol.Find(bson.M{"_id": joke.ImageIds[0]}).One(image)
			vos[i].ImageId = image.Id.Hex()
		}
	}
	return vos
}

func GetImageDataByImageId(id string) []byte {
	s := Session()
	defer s.Close()
	col := CollectionImage(s)
	image := Image{}
	err := col.FindId(bson.ObjectIdHex(id)).One(&image)
	if err != nil {
		revel.INFO.Println("ERROR ", err.Error())
	}
	return image.Data
}

// 创建一个回复
func CreateReply(id, msg string) {
	s := Session()
	defer s.Close()
	jokeCol := CollectionJoke(s)
	objectId := bson.ObjectIdHex(id)
	count, err := jokeCol.FindId(objectId).Count()
	if err != nil || count < 0 {
		return
	}
	replyCol := Collection(s, COL_REPLY)
	reply := &Reply{Id: bson.NewObjectId(), Content: msg, JokeId: objectId}
	reply.Init()
	replyCol.Insert(reply)
}
