package controllers

import (
	"github.com/revel/revel"
	// "log"
)

type ImageRenderer struct {
	ContentType string
	Data        []byte
}

func (r ImageRenderer) Apply(req *revel.Request, resp *revel.Response) {
	if len(r.ContentType) == 0 {
		resp.ContentType = "image/jpg"
	} else {
		resp.ContentType = r.ContentType
	}
	resp.Out.Write(r.Data)
}
