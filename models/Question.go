package models

import (
	"github.com/astaxie/beego/orm"
)

type Category struct {
	Category_id int `orm:"pk;auto", json:"category_id"`
	Title string	`json:"title"`
}

type Question struct {
	Id int `orm:"pk;auto", json:"id"`
	Category_id int `json:"category_id"`
	Title string	`json:"title"`
	Content string	`json:"content,omitempty"`
	Thumb string `json:"thumb,omitempty"`
}

func init() {
	orm.RegisterModel(new(Category))
	orm.RegisterModel(new(Question))
}