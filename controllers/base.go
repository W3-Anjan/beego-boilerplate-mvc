package controllers

import beego "github.com/beego/beego/v2/server/web"

type BaseController struct {
	beego.Controller
}

// Runs after Init before request function execution
func (c *BaseController) Prepare() {
	c.Data["RequestUrl"] = c.Ctx.Input.URL()
}

// Runs after request function execution
func (c *BaseController) Finish() {
	// Any cleanup logic common to all requests goes here. Logging or metrics, for example.
}
