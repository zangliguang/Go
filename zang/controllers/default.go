package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	this.Data["Website"] = "www.baidu.com"
	this.Data["Email"] = "astaxie@gmail.com"
	this.Data["truecondition"] = true
	this.Data["faleconditon"] = false
	this.TplNames = "index.tpl"

	type u struct {
		Name string
		Age  int
		Sex  string
	}
	user := &u{
		Name: "zang",
		Age:  25,
		Sex:  "male",
	}
	this.Data["user"] = user
}
