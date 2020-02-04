package controllers

import (
	// "database/sql"
	_ "github.com/go-sql-driver/mysql"
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
func (c App) ArticleTemplate() revel.Result {
	return c.RenderTemplate("App/Article.html")
}

func (c App) SubmitArticle() revel.Result {
	text := c.Params.Get("text")
}
