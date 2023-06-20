package links

import (
	"log"

	database "github.com/jordaneaster/graphql-golang/internal/pkg/db/mysql"
	"github.com/jordaneaster/graphql-golang/internal/users"
)

// #1
type Link struct {
	ID      string
	Title   string
	Address string
	User    *users.User
}

// #2
func (link Link) Save() int64 {
	//#3
	stmt, err := database.Db.Prepare("INSERT INTO Links(Title,Address) VALUES(?,?)")
	if err != nil {
		log.Fatal(err)
	}
	//#4
	res, err := stmt.Exec(link.Title, link.Address)
	if err != nil {
		log.Fatal(err)
	}
	//#5
	id, err := res.LastInsertId()
	if err != nil {
		log.Fatal("Error:", err.Error())
	}
	log.Print("Row inserted!")
	return id
}

// In users.go we just defined a struct that represent users we get from database, But let me explain links.go part by part:

// 1: definition of struct that represent a link.
// 2: function that insert a Link object into database and returns it’s ID.
// 3: our sql query to insert link into Links table. you see we used prepare here before db.Exec, the prepared statements helps you with security and also performance improvement in some cases. you can read more about it here.
// 4: execution of our sql statement.
// 5: retrieving Id of inserted Link.
