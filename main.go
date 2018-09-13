package main

import (
	"database/sql"
	"fmt"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

const (
	DB_USER     = "dbusername"
	DB_PASSWORD = "dbpassword"
	DB_NAME     = "database"
)

func main() {
	// DB connection
	db, err := sql.Open("mysql", DB_USER+":"+DB_PASSWORD+"@/"+DB_NAME)

	checkErr(err)

	defer db.Close()

	printMessage("Delete books from DB...")
	deleteBooks(db)

	printMessage("Insert books into DB...")
	insertBooks(db)

	printMessage("Getting books...")
	getBooks(db)
}

// Insert books
func insertBooks(db *sql.DB) {
	// Insert several books
	stmtIns, err := db.Prepare("INSERT INTO books (bookID, bookName) VALUES(?, ?)")
	checkErr(err)

	for i := 1; i < 11; i++ {
		// Convert i to string - for demo purposes only
		var thisBookID_string = strconv.Itoa(i)

		var bookID = thisBookID_string

		var bookName = thisBookID_string + "_name"

		fmt.Println("Inserting new book with ID: " + bookID + " and name: " + bookName)

		_, err = stmtIns.Exec(bookID, bookName)
		checkErr(err)
	}
}

// Get books
func getBooks(db *sql.DB) {
	// Get all books from books table that don't have bookID = "1"
	rows, err := db.Query("SELECT * FROM books where bookID <> 1")

	checkErr(err)

	// Foreach book
	for rows.Next() {
		var id int
		var bookID string
		var bookName string

		err = rows.Scan(&id, &bookID, &bookName)

		checkErr(err)

		fmt.Println("Book ID: " + bookID + ", book name: " + bookName)
	}
}

// Delete books
func deleteBooks(db *sql.DB) {
	printMessage("Deleting all books...")

	_, err := db.Exec("DELETE FROM books")
	checkErr(err)

	printMessage("All books have been deleted successfully!")
}

// Function for handling messages
func printMessage(message string) {
	fmt.Println("")
	fmt.Println(message)
	fmt.Println("")
}

// Function for handling errors
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
