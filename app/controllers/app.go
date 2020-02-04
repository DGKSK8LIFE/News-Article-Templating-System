package controllers

import (
	"github.com/revel/revel"
)

// Controller Instance 
type App struct {
	*revel.Controller
}

// home page receiver
func (c App) Index() revel.Result {
	return c.RenderTemplate("App/Index.html")
}

// Article Template receiver
func (c App) Template_Render() revel.Result {
	return c.RenderTemplate("App/Template_Render.html")
}
