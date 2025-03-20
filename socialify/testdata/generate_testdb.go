package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// Delete existing database if it exists
	os.Remove("./socialify_test.db")

	// Create a new database
	db, err := sql.Open("sqlite3", "./socialify_test.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create tables
	createTables(db)

	// Insert sample data
	insertSampleData(db)

	fmt.Println("Test database created successfully.")
}

func createTables(db *sql.DB) {
	// Create users table
	usersTable := `CREATE TABLE users (
		id TEXT PRIMARY KEY,
		name TEXT NOT NULL
	);`

	// Create posts table
	postsTable := `CREATE TABLE posts (
		id INTEGER PRIMARY KEY,
		userid TEXT NOT NULL,
		content TEXT NOT NULL,
		FOREIGN KEY (userid) REFERENCES users (id)
	);`

	// Create comments table
	commentsTable := `CREATE TABLE comments (
		id INTEGER PRIMARY KEY,
		postid INTEGER NOT NULL,
		content TEXT NOT NULL,
		FOREIGN KEY (postid) REFERENCES posts (id)
	);`

	// Execute SQL statements
	_, err := db.Exec(usersTable)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(postsTable)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(commentsTable)
	if err != nil {
		log.Fatal(err)
	}
}

func insertSampleData(db *sql.DB) {
	// Sample users
	users := []struct {
		id   string
		name string
	}{
		{"1", "John Doe"},
		{"2", "Jane Doe"},
		{"3", "Alice Smith"},
		{"4", "Bob Johnson"},
		{"5", "Charlie Brown"},
		{"6", "Diana White"},
		{"7", "Edward Davis"},
		{"8", "Fiona Miller"},
		{"9", "George Wilson"},
		{"10", "Helen Moore"},
	}

	// Insert users
	userStmt, err := db.Prepare("INSERT INTO users(id, name) VALUES(?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer userStmt.Close()

	for _, user := range users {
		_, err = userStmt.Exec(user.id, user.name)
		if err != nil {
			log.Fatal(err)
		}
	}

	// Sample posts
	posts := []struct {
		id      int
		userid  string
		content string
	}{
		{246, "1", "Post about ant"},
		{161, "1", "Post about elephant"},
		{150, "1", "Post about ocean"},
		{370, "1", "Post about monkey"},
		{344, "1", "Post about ocean"},
		{952, "1", "Post about zebra"},
		{647, "1", "Post about igloo"},
		{421, "1", "Post about house"},
		{890, "1", "Post about bat"},
		{461, "1", "Post about umbrella"},
		{247, "2", "Post about flowers"},
		{162, "2", "Post about gardens"},
		{151, "2", "Post about rivers"},
		{371, "2", "Post about mountains"},
		{345, "3", "Post about hiking"},
		{953, "3", "Post about camping"},
		{648, "4", "Post about cooking"},
		{422, "4", "Post about baking"},
		{891, "5", "Post about music"},
		{462, "5", "Post about art"},
	}

	// Insert posts
	postStmt, err := db.Prepare("INSERT INTO posts(id, userid, content) VALUES(?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer postStmt.Close()

	for _, post := range posts {
		_, err = postStmt.Exec(post.id, post.userid, post.content)
		if err != nil {
			log.Fatal(err)
		}
	}

	// Sample comments
	comments := []struct {
		id      int
		postid  int
		content string
	}{
		{3893, 150, "Old comment"},
		{4791, 150, "Boring comment"},
		{4792, 150, "Interesting comment"},
		{3894, 161, "Nice post"},
		{4793, 161, "Great observation"},
		{3895, 246, "I agree"},
		{3896, 370, "Funny post"},
		{4794, 370, "LOL"},
		{4795, 370, "ROFL"},
	}

	// Insert comments
	commentStmt, err := db.Prepare("INSERT INTO comments(id, postid, content) VALUES(?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer commentStmt.Close()

	for _, comment := range comments {
		_, err = commentStmt.Exec(comment.id, comment.postid, comment.content)
		if err != nil {
			log.Fatal(err)
		}
	}
}
