# News-Article-Templating-System 

A full stack Go Revel webapplication that enables users to create short and simple news articles and provides a means of searching through every single submitted article; essentially a stripped-down, mono-functional CMS (Content Management System).

## Dependent Technologies

- Go Revel 
- Go SQL Driver
- Go MySQL Driver
- MVC Architecture
- Elastisearch (go-elastisearch) (potentially; for now can just send an SQL LIKE query) 
- MySQL
- Godep for dependency management
- Redis (will only be used if live template saving is going to be implemented)
- Go HTML Templates 
- ReactJS (eventually)
- Markdown Engine for article formatting (eventually)

## Desired Features

- Client-Side rendering (potentially)
- An Article template that has a grid for images, then allocates the rest of the space for text
- User would upload images to the grid portion and would be able to write markdown text (uses markdown engine) 
- Users, after creative news pages, would submit them to the server, then they'd have their article available as a path with the syntax: `:id/:title`
- Users would be able to search through all created articles (use an sql Like statement for now, then if I want better search results and database scanning: elastisearch) (shows users article preview which hyperlinks them to the article as a search result)
- Database could use binary search to reduce bandwidth usage/time spent indexing the database server (unnescessary at this scale)
- Eventually would make a template design/creation system for aesthetic diversity among articles (still unsure about this)
- Would eventually implement article crawling for search engine optimization (potentially)
- Eventually implement React for more than just client side rendering (ie: sexy frontend)

## Completion Status

Working on it...

## Task Checklist

[Checklist](checklist.md)

## Architecture Visualization (will be updated eventually to show this app's specific architecture)

### MVC (Model View Controller)

![MVC](https://upload.wikimedia.org/wikipedia/commons/thumb/a/a0/MVC-Process.svg/500px-MVC-Process.svg.png)

### MySQL (SQL server IE: data store that serves as model)

![MYSQL LOGO](https://upload.wikimedia.org/wikipedia/en/thumb/6/62/MySQL.svg/1200px-MySQL.svg.png)
