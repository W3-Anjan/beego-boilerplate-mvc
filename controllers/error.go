package controllers

type ErrorController struct {
	BaseController
}

// this is the error 404
func (c *ErrorController) Error404() {
	c.Data["content"] = "page not found"
	c.TplName = "404.tpl"
}

func (c *ErrorController) Error500() {
	c.Data["content"] = "internal server error"
	c.TplName = "500.tpl"
}

func (c *ErrorController) ErrorDb() {
	c.Data["content"] = "database is now down"
	c.TplName = "dberror.tpl"
}
