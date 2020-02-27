package controllers

import (
	"News-Article-Templating-System-Revel-ok/app"
	"database/sql"
	"fmt"
	"log"
	"net/url"
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
	articles := []Article{}
	query := fmt.Sprintf("SELECT title, id FROM article order by id desc limit 25;")
	results, err := app.DB.Query(query)

	if results == nil {
		return c.RenderTemplate("App/Index.html")
	}

	if err != nil {
		log.Fatalln(err)
	}

	for results.Next() {
		article := Article{}
		err := results.Scan(&article.Title, &article.Id)
		if err != nil {
			log.Fatalln(err)
		}
		articles = append(articles, article)
	}

	if len(articles) != 0 && results != nil {
		c.ViewArgs["articles"] = articles
		return c.RenderTemplate("App/Index.html")
	}

	return c.RenderTemplate("App/Index.html")
}

// will use this to query the database with a wildcard query and then (via frontend gohtml templates), will iterate over results
func (c App) Search() revel.Result {
	query := c.Params.Form.Get("query")
	articles := []Article{}
	qstring := fmt.Sprintf("SELECT title, id FROM article WHERE title LIKE '%%%v%%';", query)
	results, err := app.DB.Query(qstring)

	if results == nil {
		c.ViewArgs["message"] = "No matching results"
		return c.RenderTemplate("App/SearchResults.html")
	}

	if err != nil {
		log.Fatalln(err)
	}

	for results.Next() {
		article := Article{}
		err := results.Scan(&article.Title, &article.Id)
		if err != nil {
			log.Fatalln(err)
		}
		articles = append(articles, article)
	}

	if len(articles) != 0 && results != nil {
		message := fmt.Sprintf("Results for search %s:\n", query)
		c.ViewArgs["message"] = message
		c.ViewArgs["articles"] = articles
		return c.RenderTemplate("App/SearchResults.html")
	}

	c.ViewArgs["message"] = "No results found"
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
	query := fmt.Sprintf("INSERT INTO article (content, timestamp, title) VALUES ('%v', '%v', '%v');", content[0], time.Now().UTC().String(), url.QueryEscape(content[1]))
	_, err := app.DB.Exec(query)
	if err != nil {
		log.Fatalf("Query failed: %s\n", err)
	}
	return c.Redirect(App.Index)
}
