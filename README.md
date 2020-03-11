# News-Article-Templating-System 

A full-stack Go Revel webapplication that enables users to create news articles and provides a platform for sharing them (with more style :)); essentially a stripped-down, mono-functional CMS (Content Management System).

## Dependent Technologies

- Go Revel 
- Go SQL Driver
- Go MySQL Driver
- MVC Architecture
- MySQL
- Godep for dependency management
- Go HTML Templates 
- gomarkdown Engine for article formatting (finished)
- CSS Bootstrap
- JSON (integrated with revel)

## Features that are implemented and desired features

- An Article template that has a grid for images, then allocates the rest of the space for text
- User would upload images to the grid portion and is currently able to write markdown text (uses markdown engine)/HTML; for now the user can just use html to upload images, can implement the uploading of images when I get access to a CDN
- Users, after creating news pages, submit them to the server, then have their article available to access as a path with the syntax: `/article/:id/*title` or `/article/:id/`
- Users are able to search through all created articles (uses an sql `LIKE` statement for now; if I want better search results and database scanning I can implement elastisearch) (search results are comprised of article previews that hyperlink to their corresponding articles)
- Database could use binary search to reduce time spent indexing the database server (unnescessary at this scale)
- Eventually would make a template design/creation system for aesthetic diversity among articles (unsure about this for now)
- Very nice interface utilizing html5, css3, and bootstrap/js/jquery

## Completion Status

Working on it...

## Task Checklist

[Checklist](checklist.md)

## Architecture Visualization (will be updated eventually to show this app's specific architecture)

### MVC (Model View Controller)

![MVC](https://upload.wikimedia.org/wikipedia/commons/thumb/a/a0/MVC-Process.svg/500px-MVC-Process.svg.png)
