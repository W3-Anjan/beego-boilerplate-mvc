package controllers

import beego "github.com/beego/beego/v2/server/web"

type NestPreparer interface {
	NestPrepare()
}
type BaseController struct {
	beego.Controller
}

// Runs after Init before request function execution
func (c *BaseController) Prepare() {
	c.Data["RequestUrl"] = c.Ctx.Input.URL()

	// It will test if the executing Controller is an implementation of NestPreparer.
	// If it is it calls the method of subclass.
	if app, ok := c.AppController.(NestPreparer); ok {
		app.NestPrepare()
	}
}

func (c *BaseController) NestPrepare() {

}

// Runs after request function execution
func (c *BaseController) Finish() {
	// Any cleanup logic common to all requests goes here. Logging or metrics, for example.
}
