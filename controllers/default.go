package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"proj/models"
)

type MainController struct {
	beego.Controller
}


type Respond struct {
	State bool `json:"state"`
	Msg string `json:"msg"`
	Id int64 `json:"id,omitempty"`
	Count int64 `json:"count"`
	CategoryRows []models.Category `json:"category_list,omitempty"`
	QuestionRows []models.Question `json:"question_list,omitempty"`
}

func (c *MainController) Post() {
	var res Respond
	o := orm.NewOrm()
	o.Using("default")

	op := c.GetString("op")
	switch op {
	case "add_category":
		title := c.GetString("title")
		if (title != "") {
			item := models.Category{Title:title}
			id, err := o.Insert(&item)
			if (err == nil) {
				res.State = true
				res.Id = id
			} else {
				res.State = false
				res.Msg = err.Error()
			}
		} else {
			res.State = false
			res.Msg = "invalid argument"
		}
	case "add":
		category_id, err := c.GetInt("category_id");
		if err == nil {
			var q models.Question
			q.Category_id = category_id
			q.Title = c.GetString("title")
			q.Content = c.GetString("content")
			q.Thumb = c.GetString("thumb")
			if (q.Title == "" || q.Content == "" || q.Thumb == "") {
				res.State = false
				res.Msg = "invalid argument"
			} else {
				id, err := o.Insert(&q)
				if err == nil {
					res.State = true
					res.Id = id
				} else {
					res.State = false
					res.Msg = err.Error()
				}
			}
		} else {
			res.State = false
			res.Msg = err.Error()
		}
	case "edit":

	default:
		res.State = false
		res.Msg = "invalid argument"
	}

	c.Data["json"] = &res
	c.Ctx.Output.Header("Access-Control-Allow-Origin", "*")
	c.Ctx.Output.Header("Access-Control-Allow-Methods", "POST")
	//c.Ctx.Output.Header("Content-Type", "text/html; charset=utf-8")
	c.ServeJSON()
}

func (c *MainController) Get() {
	var res Respond
	o := orm.NewOrm()
	o.Using("default")

	op := c.GetString("op")

	switch op {
	case "list":
		list_type := c.GetString("type")
		if (list_type == "category") {
			var items []models.Category
			num, err := o.Raw("SELECT * FROM category").QueryRows(&items)
			if err == nil {
				res.State = true
				res.Count = num
				res.CategoryRows = items
				res.Msg = "更多题目正在录入中，敬请期待"
			} else {
				res.State = false
				res.Msg = err.Error()
			}
		} else if (list_type == "question") {
			var items []models.Question
			category_id, err := c.GetInt("category_id")
			if (err == nil) {
				num, err := o.Raw("SELECT id,category_id,title,thumb FROM question where category_id = ?", category_id).QueryRows(&items)
				if err == nil {
					res.State = true
					res.Count = num
					res.QuestionRows = items

				} else {
					res.State = false
					res.Msg = err.Error()
				}
			} else {

			}


		}
	case "add_category":
		title := c.GetString("title")
		if (title != "") {
			item := models.Category{Title:title}
			id, err := o.Insert(&item)
			if (err == nil) {
				res.State = true
				res.Id = id
			} else {
				res.State = false
				res.Msg = err.Error()
			}
		} else {
			res.State = false
			res.Msg = "invalid argument"
		}

	default:
		res.State = false
		res.Msg = "invalid argument"
	}

	c.Data["json"] = &res
	c.ServeJSON()
}
