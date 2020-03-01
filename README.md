![revel gopher image (logo)](https://revel.github.io/img/RevelWhiteLines.png)
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
- User would upload images to the grid portion and would be able to write markdown text which could be processed on the server side and sent through the view once submitted by the user, sent through the controller, and written to the model
- Users, after creative news pages, would submit them to the server, then they'd have their article available as a path with the syntax: `:id/:title` 
- Users would be able to search through all created articles (Can use an sql Like statement for now, then if I want better search results and database scanning, elastisearch)
- Database could use binary search to reduce bandwidth usage/time spent indexing the database server 
- Eventually would make a template design/creation system for aesthetic diversity among articles 
- Would eventually implement article crawling for search engine optimization (potentially)/implement slightly heavier frontend for template creation; css grid/flexbox can be changed while article being created?
- Potentially, after the stable build is written and works flawlessly, integrate incremental saving and updating of elements of templating files which would be facilitated by a more dynamic design (implemented with gRPC and other tools); metadata about images that were uploaded and associated with it and their location would be written to the database as the author is creating the article and text would be saved and updated incrementally so they don’t lose any work; have local storage sync with server/db (buffered)
- To serve an article upon user request, the application will query its article JSON RESTful API. (eventually)

## Completion Status

Working on it...

## Task Checklist

[Checklist](checklist.md)

## Architecture Visualization

### MVC (Model View Controller)

![MVC](https://upload.wikimedia.org/wikipedia/commons/thumb/a/a0/MVC-Process.svg/500px-MVC-Process.svg.png)

### MySQL (SQL server IE: data store that serves as model)

![MYSQL LOGO](https://upload.wikimedia.org/wikipedia/en/thumb/6/62/MySQL.svg/1200px-MySQL.svg.png)
