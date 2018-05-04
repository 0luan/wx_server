package main

import (
	_ "proj/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	_ "proj/models"
)

func init() {
	user := beego.AppConfig.String("dbuser")
	pswd := beego.AppConfig.String("dbpswd")
	db_name := beego.AppConfig.String("dbname")	
	db_connect_str := user + ":" + pswd + "@tcp(39.108.150.51:3306)/" + db_name + "?charset=utf8"
	orm.RegisterDataBase("default", "mysql", db_connect_str, 10)
}

func main() {
	o := orm.NewOrm()
	o.Using("default")

	beego.Run()
}

