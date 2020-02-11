# News-Article-Templating-System

A full stack webapplication that enables users to make short and simple news articles and provides a means of searching through every single submitted article; essentially a stripped-down, mono-functional CMS (Content Management System) with anonymous capabilities.

## Dependent Technologies

- Go Revel
- MVC Architecture
- Elastisearch (go-elastisearch)
- MySQL
- Godep
- Redis (maybe)
- gRPC (potentially)
- Go HTML Templates 
- ReactJS (eventually)
- Markdown Engine (Eventually)

## Desired Features

- Client-Side rendering (eventually)
- An Article template that has a grid for images, then allocates the rest of the space for text
- User would upload images to the grid portion and would be able to write markdown text which could be processed on the server side and sent through the view once submitted by the user, sent through the controller, and written to the model
- Users, after creative news pages, would submit them to the server, then they'd have their article available as a global resource within the webapp
- Users would be able to search through all created articles (would have to implement search engine stuff like regex)
- Database could use binary search to reduce bandwidth usage/time spent indexing the database server
- Eventually would make a template design/creation system for aesthetic diversity among articles
- Would eventually implement article crawling for search engine optimization
- Eventually, after the stable build is written and works flawlessly, integrate incremental saving and updating of elements of templating files which would be facilitated by a more dynamic design (implemented with gRPC and other tools); metadata about images that were uploaded and associated with it and their location would be written to the database as the author is creating the article and text would be saved and updated incrementally so they don’t lose any work; have local storage sync with server/db (buffered)

## Task Checklist

[Checklist](checklist.md)
## Architecture Visualization

### MVC (Model View Controller)

![MVC](https://upload.wikimedia.org/wikipedia/commons/thumb/a/a0/MVC-Process.svg/500px-MVC-Process.svg.png)


