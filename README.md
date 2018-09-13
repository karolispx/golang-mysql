# GoLang and MySQL. Very simple examples for inserting, selecting and deleting from MySQL DB.

## Installing MySQL on your machine/server should be straight forward. You should be able to find tutorials on google/youtube. 

### Install MySQL Client on your machine to make your life easier, I'm using Sequel Pro on a macOS.

You will need to get the following package to make it work:
* `go get -u github.com/go-sql-driver/mysql`

## Books table:
```
CREATE TABLE IF NOT EXISTS books (
    id int(11) NOT NULL AUTO_INCREMENT,
    bookID varchar(50) DEFAULT NULL,
    bookName varchar(255) DEFAULT NULL,
    PRIMARY KEY(id)
);
```

More examples and information for go-sql-driver/mysql package can be found here: https://github.com/go-sql-driver/mysql
