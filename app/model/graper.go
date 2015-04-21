package model

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"gopkg.in/mgo.v2/bson"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func DownloadHtml() {
	resp, _ := http.Get("http://www.qiushibaike.com/")
	data, _ := ioutil.ReadAll(resp.Body)
	file, _ := os.Open("/Users/wangboo/Desktop/a.html")
	defer file.Close()
	file.Write(data)
	log.Println("data len = ", len(data))
}

func GetUrl(page int) string {
	return fmt.Sprintf("http://www.qiushibaike.com/8hr/page/%d?s=4765285", page)
}

func Grape(page int, max int) {
	if page > max {
		return
	}
	url := GetUrl(page)
	log.Println("grape ", url)
	doc, err := goquery.NewDocument(url)
	if err != nil {
		log.Println("GetUrl failed : ", url)
		Grape(page+1, max)
		return
	}
	s := Session()
	defer s.Close()
	jokeCol := CollectionJoke(s)
	imageCol := CollectionImage(s)
	doc.Find(".article").Each(func(i int, s *goquery.Selection) {
		all, _ := s.Html()
		idStr := regexp.MustCompile(`qiushi_counts_(\d+)`).FindStringSubmatch(all)[1]
		jid, _ := strconv.Atoi(idStr)
		// 查找id是否存在
		count, _ := jokeCol.Find(bson.M{"jid": jid}).Count()
		if count > 0 {
			log.Printf("id %d exist!\n", count)
			return
		}
		txt := s.Find(".content").Text()
		txt = strings.TrimSpace(txt)
		joke := &Joke{Id: bson.NewObjectId(), Content: txt, From: "qiubai", Jid: jid}
		imgDiv := s.Find("div .thumb img").First()
		if imgDiv != nil {
			src, has := imgDiv.Attr("src")
			if has {
				id := bson.NewObjectId()
				image := &Image{Id: id, Data: GetImageData(src)}
				imageCol.Insert(image)
				joke.ImageIds = append(joke.ImageIds, id)
				log.Println("grape image")
			}
		}
		log.Println("save ", joke.Content)
		err := jokeCol.Insert(joke)
		if err != nil {
			log.Fatalf("save joke error : %s\n", err.Error())
		}
	})
	time.Sleep(10 * time.Second)
	Grape(page+1, max)
}

func GetImageData(url string) []byte {
	resp, _ := http.Get(url)
	data, _ := ioutil.ReadAll(resp.Body)
	return data
}
