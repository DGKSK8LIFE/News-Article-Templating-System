package controllers

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/revel/revel"
)

// Controller Instance
type App struct {
	*revel.Controller
}

// home page renderer
func (c App) Index() revel.Result {
	return c.RenderTemplate("App/Index.html")
}

// Article Template renderer
func (c App) ArticleTemplate() revel.Result {
	return c.RenderTemplate("App/Article.html")
}

// Article Template data receiver; going to implement model interaction soon
func (c App) SubmitArticle() revel.Result {
	db, err := sql.Open("mysql", "")
	if err != nil {
		log.Fatalf("database error: %s\n", err)
	}
	defer db.Close()
	text := c.Params.Get("text")
	return c.Redirect()
}
