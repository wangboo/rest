package mjobs

import (
	"github.com/wangboo/rest/app/model"
)

type QiubaiGraper struct {
}

func (q *QiubaiGraper) Run() {
	model.Grape(1, 10)
}
