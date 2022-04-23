package main

import (
	"beetpl/common"
	"beetpl/controllers"
	_ "beetpl/routers"
	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	// Custom template functions will be here
	beego.AddFuncMap("LoadTimes", common.LoadTimes)
	// Use beego.ErrorController to register the error controller before beego.Run
	beego.ErrorController(&controllers.ErrorController{})
	beego.Run()
}
