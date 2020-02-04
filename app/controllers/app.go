package controllers

import (
	// "database/sql"
	"github.com/revel/revel"
	_ "github.com/go-sql-driver/mysql"
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
func (c App) TemplateRender() revel.Result {
	return c.RenderTemplate("App/Template_Render.html")
}

func (c App) SubmitArticle()  {

}
