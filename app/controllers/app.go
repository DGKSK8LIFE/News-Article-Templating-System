package controllers

import (
	"News-Article-Templating-System-Revel-ok/app"
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gomarkdown/markdown"
	"github.com/revel/revel"
)

// Controller struct
type App struct {
	*revel.Controller
}

// Article struct containing its database fields
type Article struct {
	Content   string `json:"Content"`
	Timestamp string `json:"Timestamp"`
	Title     string `json:"Title"`
	Id        int    `json:"Id"`
}

// Article Creator Template renderer
func (c App) ArticleTemplate() revel.Result {
	return c.RenderTemplate("App/Article.html")
}

// home page renderer
func (c App) Index() revel.Result {
	articles := []Article{}
	query := fmt.Sprintf("SELECT title, id FROM article order by id desc limit 25")
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

// queries the database with a wildcard query and then submits a struct slice -> frontend iterates over and displays results
func (c App) Search() revel.Result {
	query := c.Params.Form.Get("query")
	articles := []Article{}
	qstring := fmt.Sprintf("SELECT title, id FROM article WHERE title LIKE '%%%v%%' OR content LIKE '%%%v%%' order by id desc", query, query)
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
		c.ViewArgs["search"] = query
		return c.RenderTemplate("App/SearchResults.html")
	}

	nothingFound := fmt.Sprintf("No results found for search: %s\n", query)
	c.ViewArgs["message"] = nothingFound
	return c.RenderTemplate("App/SearchResults.html")

}

// Handles Get Request To Desired Article
func (c App) GetArticle(id int, title ...string) revel.Result {
	article := Article{Id: id}
	query := fmt.Sprintf("SELECT content, title, timestamp FROM article WHERE id=%v", id)
	err := app.DB.QueryRow(query).Scan(&article.Content, &article.Title, &article.Timestamp)
	if err != sql.ErrNoRows {
		c.ViewArgs["title"] = article.Title
		c.ViewArgs["text"] = template.HTML(markdown.ToHTML([]byte(article.Content), nil, nil))
		c.ViewArgs["timestamp"] = article.Timestamp
	} else if err == sql.ErrNoRows {
		c.Response.Status = 404
		return c.Render()
	}

	return c.RenderTemplate("App/Post.html")
}

// Marshals article data to JSON and then serves it (works as an API)
func (c App) GetArticleJSON(id int, title ...string) revel.Result {
	article := Article{Id: id}
	query := fmt.Sprintf("SELECT content, title, timestamp FROM article WHERE id=%v", id)
	err := app.DB.QueryRow(query).Scan(&article.Content, &article.Title, &article.Timestamp)
	if err != sql.ErrNoRows {
		data := make(map[string]interface{})
		data["error"] = nil
		data["article"] = article
		c.Response.ContentType = "application/json"
		return c.RenderJSON(data)
	}

	c.Response.Status = 404
	return c.Render()

}

// Scans Article Data to an Article instance
func (a *Article) ScanArticleDBData() error {
	query := fmt.Sprintf("SELECT content, title, timestamp FROM article WHERE id=%v", a.Id)
	err := app.DB.QueryRow(query).Scan(&a.Content, &a.Title, &a.Timestamp)
	return err
}

// Article Template data receiver; going to implement model interaction soon
func (c App) SubmitArticle() revel.Result {
	content := []string{c.Params.Get("text"), c.Params.Get("title")}
	query, err := app.DB.Prepare("INSERT INTO article (content, timestamp, title) VALUES (?, ?, ?)")
	if err != nil {
		log.Fatalf("query err: %s\n", err)
	}
	_, err = query.Exec(content[0], time.Now().UTC().String(), content[1])
	if err != nil {
		log.Fatalf("Query failed: %s\n", err)
	}
	return c.Redirect(App.Index)
}
