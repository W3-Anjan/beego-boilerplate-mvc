package controllers

import "time"

type MainController struct {
	BaseController
}

// MainController will execute and checks whether there is a Prepare method.
// If not it keeps searching baseController. If there is it will execute the logic this.AppController in baseRouter
// is the currently executing Controller MainController. Next, it will execute MainController.NestPrepare method.
func (c *MainController) NestPrepare() {

	c.Data["IsAdmin"] = true
	/*if app, ok := c.AppController.(ModelPreparer); ok{
		app.ModelPrepare()
		return
	}*/

}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.Data["t"] = time.Now()
	c.Data["m"] = map[string]interface{}{
		"a": 1,
		"1": map[string]float64{
			"c": 4,
		},
	}
	c.TplName = "index.tpl"
}
