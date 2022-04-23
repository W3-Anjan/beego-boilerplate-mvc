package controllers

import "time"

type MainController struct {
	BaseController
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
