package controllers

import (
	"github.com/astaxie/beego"
)

type LogController struct {
	beego.Controller
}

func (this *LogController) Get() {
	var result ShortResult
	name := this.Input().Get("name")
	result.UrlShort = name
	if urlcache.IsExist(name) {
		result.UrlLong = urlcache.Get(name).(string)
	} else {
		result.UrlLong = ""
	}
	this.Data["json"] = result
	this.ServeJSON()
}