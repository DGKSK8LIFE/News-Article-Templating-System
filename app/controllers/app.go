package controllers

import (
	"News-Article-Templating-System-Revel-ok/app"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/revel/revel"
)

// Controller struct
type App struct {
	*revel.Controller
}

// RESTful API POST struct (will be redundant until I start work on the API)
type Post struct {
	content   string
	timestamp string
	title     string
	id        int
}

// home page renderer
func (c App) Index() revel.Result {
	return c.RenderTemplate("App/Index.html")
}

// will use this to query the database with a wildcard query and then (via frontend gohtml templates), will iterate over results
func (c App) ServeResults(search string) revel.Result {
	return c.Render()
}

// Handles Post Request To Desired Article
func (c App) GetArticle(id int, title string) revel.Result {
	if exists := postExists(id); true {
		query := fmt.Sprintf("SELECT content FROM article WHERE id='%v' AND title='%v';", id, title)
		result, err := app.DB.Query(query)
		if err != nil {
			log.Fatalf("Query error: %s\n", err)
		}
		c.ViewArgs["title"] = title
		c.ViewArgs["text"] = result
		return c.RenderTemplate("App/Post.html")
	} else if exists == false {
		c.Response.Status = 404
		return c.Render()
	}
	return c.Redirect("App/Article.html")

}

// Article Template renderer
func (c App) ArticleTemplate() revel.Result {
	return c.RenderTemplate("App/Article.html")
}

// Article Template data receiver; going to implement model interaction soon
func (c App) SubmitArticle() revel.Result {
	content := []string{c.Params.Get("text"), c.Params.Get("title")}
	query := fmt.Sprintf("INSERT INTO article (content, timestamp, title) VALUES ('%v', '%v', '%v');", content[0], time.Now().UTC().String(), content[1])
	_, err := app.DB.Exec(query)
	if err != nil {
		log.Fatalf("Query failed: %s\n", err)
	}
	return c.Redirect(App.Index)
}

// checks whether a post exists; broken atm
func postExists(id int) bool {
	var exists bool
	query := fmt.Sprintf("SELECT * FROM article WHERE id='%v';", id)
	if err := app.DB.QueryRow(query).Scan(&id); err != nil && err != sql.ErrNoRows {
		log.Fatalf("database error: %s\n", err)
	} else if err == sql.ErrNoRows {
		exists = false
	} else {
		exists = true
	}
	return exists
}
