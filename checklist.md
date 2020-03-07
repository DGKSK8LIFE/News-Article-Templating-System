# Checklist

## Completed tasks 

1. Template creator (with basic interface; to be improved) correctly writes information to the database.
2. Prototype user interface works.
3. Templates, once created, have established routes with this form: /id/title.
4. `setup.sql` script for users who want to clone and run the working app with its database
5. Removed bugs allowing the serving of non-existent routes (ie: requested article doesn't exist -> now gives 404 instead of serving it)
6. Improved documentation and organization
7. Get stylesheet to work in Post.html
8. Working Search System
9. Working Article Previews that hyperlink to their designated articles
10. Working Article Previews as Search Results
11. Articles with dangerous characters are escaped 
12. Articles with dangerous characters are url encoded when hyperlinking 
13. Encoded URLs route properly
14. Removed need to manually escape html characters (made method change so that html is escaped automagically and is stored unescaped in database)
15. Implemented route redundancy of article title
16. Implemented markdown text rendering
17. Implemented html/css rendering with security and xss mitigation (js is escaped)
18. Implemented user friendly UI with bootstrap/min-css (will be improved)

## TODO

 - Fix all bugs/Satisfy all feature requests that linger in uncompleted tasks sections (IE: Github issues)
 - Attain eventual goals (contained in README).
 - Implement image serving and uploading ability in templates (with use of CDN). (am complacent with just html image loading for now)
 - Add proper formatting to files for readability


