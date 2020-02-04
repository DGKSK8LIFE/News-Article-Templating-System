package controllers

import (
	"github.com/revel/revel"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	return c.RenderTemplate("App/Index.html")
}

func (c App) Template_Render() revel.Result {
	return c.RenderTemplate("App/Template_Render.html")
}
