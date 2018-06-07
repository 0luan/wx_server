package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type Records struct {
	Id int64 `orm:"pk;auto" json:"id"`
	Content string	`json:"content,omitempty"`
	Images string `json:"images,omitempty"`
	Time time.Time `json:"time"`
}

func init() {
	orm.RegisterModel(new(Records))
}