package routers

import (
	"beetpl/controllers"
	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	beego.Router("/index", &controllers.MainController{})
}
