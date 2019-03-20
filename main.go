package main

import (
	"github.com/astaxie/beego"
	"beego_test/controllers"
	"beego_test/util"
	"beego_test/menu"
)

func main() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/v1/shorten", &controllers.ShortController{})
	beego.Router("/v1/expand", &controllers.ExpandController{})
	beego.Router("/v1/log", &controllers.LogController{})
	go beego.Run()
	util.CommandOutput()
	menu.MainMenu()
}


