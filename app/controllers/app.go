package controllers

import (
	"News-Article-Templating-System-Revel-ok/app"
	"fmt"
	"log"
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

/*
// Parses SearchPatterns and Returns Database Results Accordingly
func (c App) SearchQuery() revel.Result {

}
*/

// Handles Post Request To Desired Article
func (c App) PostToArticle(id int, title string) revel.Result {
	query := fmt.Sprintf("SELECT content FROM article WHERE id='%v' AND title='%v';", id, title)
	result, err := app.DB.Query(query)
	if err != nil {
		log.Fatalf("Query failed: %s\n", err)
	}
	fieldData := struct {
		title string
		text  interface{}
	}{
		title: title,
		text:  result,
	}
	c.ViewArgs['title'] = fieldData.title
	c.ViewArgs['text'] = fieldData.text
	return c.Execute("App/Post.html")
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
