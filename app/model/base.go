package model

import (
	// "gopkg.in/mgo.v2/bson"
	"labix.org/v2/mgo"
	"log"
	"time"
)

type BaseModel struct {
	CreatedAt time.Time
	UpdateAt  time.Time
}

var (
	session *mgo.Session
)

func init() {
	sess, err := mgo.Dial("localhost/qiushibaike")
	if err != nil {
		log.Printf("mongo dial err %s\n", err.Error())
		return
	}
	sess.SetMode(mgo.Monotonic, true)
	session = sess
}

func Session() *mgo.Session {
	return session.Copy()
}

func DB(s *mgo.Session) *mgo.Database {
	return s.DB("qiushibaike")
}

func (b *BaseModel) Init() {
	b.CreatedAt = time.Now()
	b.UpdateAt = time.Now()
}
