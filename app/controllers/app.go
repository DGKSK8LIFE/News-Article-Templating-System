package controllers

import (
	"News-Article-Templating-System-Revel-ok/app"
	"database/sql"
	"fmt"
	"log"
	"time"

	"html"

	_ "github.com/go-sql-driver/mysql"
	"github.com/revel/revel"
)

// Controller struct
type App struct {
	*revel.Controller
}

// RESTful API Article struct
type Article struct {
	Content   string
	Timestamp string
	Title     string
	Id        int
}

// home page renderer
func (c App) Index() revel.Result {
	return c.RenderTemplate("App/Index.html")
}

// will use this to query the database with a wildcard query and then (via frontend gohtml templates), will iterate over results
func (c App) Search(query string) revel.Result {
	articles := Article{}
	wildcardQuery := fmt.Sprintf("SELECT title, content FROM article WHERE title LIKE %'%v'%", query)
	err := app.DB.QueryRow(wildcardQuery).Scan(&articles.Title, &articles.Content)
	if err == sql.ErrNoRows {
		c.ViewArgs["message"] = "No matching results"
		return c.RenderTemplate("App/SearchResults.html")
	}
	message := fmt.Sprintf("Results for search: %s\n", query)
	c.ViewArgs["message"] = message
	c.ViewArgs["articles"] = articles
	return c.RenderTemplate("App/SearchResults.html")

}

// Handles Post Request To Desired Article
func (c App) GetArticle(id int, title string) revel.Result {
	article := Article{}
	query := fmt.Sprintf("SELECT content FROM article WHERE id='%v' AND title='%v';", id, title)
	err := app.DB.QueryRow(query).Scan(&article.Content)
	if err != sql.ErrNoRows {
		fmt.Println("database nil err?")
	} else if err == sql.ErrNoRows {
		c.Response.Status = 404
		return c.Render()
	}
	c.ViewArgs["title"] = title
	c.ViewArgs["text"] = article.Content
	return c.RenderTemplate("App/Post.html")
}

// Article Template renderer
func (c App) ArticleTemplate() revel.Result {
	return c.RenderTemplate("App/Article.html")
}

// Article Template data receiver; going to implement model interaction soon
func (c App) SubmitArticle() revel.Result {
	content := []string{html.EscapeString(c.Params.Get("text")), html.EscapeString(c.Params.Get("title"))}
	query := fmt.Sprintf("INSERT INTO article (content, timestamp, title) VALUES ('%v', '%v', '%v');", content[0], time.Now().UTC().String(), content[1])
	_, err := app.DB.Exec(query)
	if err != nil {
		log.Fatalf("Query failed: %s\n", err)
	}
	return c.Redirect(App.Index)
}
