package controllers

import (
	"News-Article-Templating-System-Revel-ok/app"
	"fmt"
	"log"
	"time"

	"database/sql"

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

// Handles Post Request To Desired Article
func (c App) PostToArticle(id int, title string) revel.Result {
	query := fmt.Sprintf("SELECT content FROM article WHERE id='%v' AND title='%v';", id, title)
	result, err := app.DB.QueryRow(query).Scan()
	if err != nil && err != sql.ErrNoRows {
		log.Fatalf("Query failed: %s\n", err)
	} else if err == sql.ErrNoRows {
		c.Response.Status = 404
		return c.Render()
	}
	c.ViewArgs["title"] = title
	c.ViewArgs["text"] = result
	return c.RenderTemplate("App/Post.html")
}

// Article Template renderer
func (c App) ArticleTemplate() revel.Result {
	return c.RenderTemplate("App/Article.html")
}

// Article Template data receiver; going to implement model interaction soon
func (c App) SubmitArticle() revel.Result {
	content := []string{c.Params.Get("text"), c.Params.Get("title")}
	query := fmt.Sprintf("INSERT INTO article (content, timestamp, title) VALUES ('%s', '%v', '%v');", content[0], time.Now().UTC().String(), content[1])
	_, err := app.DB.Exec(query)
	if err != nil {
		log.Fatalf("Query failed: %s\n", err)
	}
	return c.Redirect(App.Index)
}
