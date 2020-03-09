# todo-list-6

## Setting up MySQL server

If you haven't already, download MySQL Community Server at https://dev.mysql.com/downloads/mysql/. 
Then run the following commands to set up a Todo database.
```sql
create database Todo_db;

use Todo_db;

create table Todo (
ID int NOT NULL AUTO_INCREMENT,
Title varchar(100),
Complete bool,
PRIMARY KEY (ID)
);

insert into Todo (Title, Complete)
values ('create app', TRUE);
```

In db/mySQL/mySql.go, update the ```ConnectMySQL()``` method to use your MySQL user and password.


## GoLang Development server
### Dependencies
* MySQL
  * github.com/go-sql-driver/mysql
* CORS
  * github.com/rs/cors
  
### Run Go Server
Run ```go run main.go``` to start the app. It will open up a connection on ```http://localhost:8080/``` to communicate with the Angular project.

## Angular Development server
Navigate to the ```/ui``` directory and run ```ng serve``` for a dev server. Navigate to ```http://localhost:4200/```.
