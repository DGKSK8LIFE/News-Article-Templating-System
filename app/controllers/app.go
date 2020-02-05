package controllers

import (
	"fmt"
	"log"

	"github.com/News-Article-Templating-System/app"

	"time"

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
	text := c.Params.Get("text")
	query := fmt.Sprintf("INSERT INTO article (content, timestamp) VALUES (%s, %d);", text, time.UTC())
	execQuery, err := app.DB.Exec(query)
	if err != nil {
		log.Fatalf("Query failed: %s\n", err)
	}
	return c.Redirect(App.Index)
}
