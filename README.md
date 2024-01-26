# Test-CF

this Repo contain backend api CRUD in hexagonal architecture with golang 
 - use gin create RestAPI 
 - use gorm to connect sqlite

## How  to test this Repo

Setup Enviorment

 - install go on your machine
 - install sqlite
 - I use postman to curl api
 - in the db file test-cf.db already have some data but can used the script to add more additional data with your database tool

````
INSERT INTO tbl_customers (name, age) VALUES('Alan Krub', 30),('Ari Ka',50);
````
 - I prepared curl postman you can import in your postman on file 'Test-CF.postman_collection.json'
 - run the project and you can test now !

Thank you 