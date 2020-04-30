# News-Article-Templating-System 

A full-stack Go Revel webapplication that enables users to create news articles and provides a platform for sharing them (with more style :)); essentially a stripped-down, mono-functional CMS (Content Management System).

## Installation

1. `cd $GOPATH/go/src`
2. git pull the project
3. Install dependencies (go modules are already handled with dep); just need to install the revel cli
4. Run the `setup.sql` script on an sql database with the name `articles`
5. Put a valid database user in `DB_user.yaml` file
6. Have fun and run with `revel run`
7. *Note that this is the development code; it's not ready for production out of the box* (some stuff may be tweaked in configs)

## Dependencies

- Go Revel 
- Go SQL Driver
- Go MySQL Driver
- MVC Architecture
- MySQL
- Godep for dependency management
- Go HTML Templates 
- gomarkdown
- CSS Bootstrap
- JSON (integrated with revel)
- Moment.JS (imported on client-side with cloudflare min)
- YAML

## Features that are implemented/desired features

- An Article template that has a grid for images, then allocates the rest of the space for text
- User would upload images to the grid portion and is currently able to write markdown text (uses markdown engine)/HTML; for now the user can just use html to upload images, can implement the uploading of images when I get access to a CDN
- Users, after creating news pages, submit them to the server, then have their article available to access as a path with the syntax: `/article/:id/*title` or `/article/:id/`
- Users are able to search through all created articles (uses an sql `LIKE` statement for now; if I want better search results and database scanning I can implement elastisearch) (search results are comprised of article previews that hyperlink to their corresponding articles)
- Database could use binary search to reduce time spent indexing the database server (unnescessary at this scale)
- Eventually would make a template design/creation system for aesthetic diversity among articles (unsure about this for now)
- Very nice interface utilizing html5, css3, and bootstrap
- REST API that serves json and is accessable with the routes: `/api/article/:id/` or `/api/article/:id/*title`
- Articles, when loaded, show their timestamps in the users' local time (client-side UTC conversion with Moment.JS)

## Completion Status

Working on it... (essential development version done)

## Task Checklist

[Checklist](checklist.md)

## Architecture Visualization (will be updated eventually to show system's specific applications of architecture)

### MVC (Model View Controller)

![MVC](https://upload.wikimedia.org/wikipedia/commons/thumb/a/a0/MVC-Process.svg/500px-MVC-Process.svg.png)
