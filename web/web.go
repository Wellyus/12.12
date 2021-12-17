package web

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type Post struct {
	Id      int
	Content string
	Author  string
}

var Db *sql.DB //create a handle for connection of database

func init() {
	var err error
	//connect with database
	Db, err = sql.Open("postgres", "user=wellyus dbname=wellyus password=ln2.7182818285 sslmode=disable")
	if err != nil {
		panic(err)
	}
}

// create a post accroding to a Post structure
func (post *Post) Create() (err error) {
	//sql statement model
	statement := "insert into posts (content, author) values ($1, $2) returning id"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()
	// query a row
	err = stmt.QueryRow(post.Content, post.Author).Scan(&post.Id)
	if err != nil {
		return
	}
	return
}

//get a post in database according to a post id (int)
func GetPost(id int) (post Post, err error) {
	post = Post{}
	err = Db.QueryRow("select id, content, author from posts where id = $1", id).Scan(&post.Id, &post.Content, &post.Author)
	return
}

func (post *Post) Update() (err error) {
	_, err = Db.Exec("Update posts set content = $2, author = $3  where id = $1", post.Id, post.Content, post.Author)
	return
}

func (post *Post) Delete() (err error) {
	_, err = Db.Exec("delete from posts where id = $1", post.Id)
	return
}

func Posts(limit int) (posts []Post, err error) {
	// return some row up to limit
	rows, err := Db.Query("select id, content, author from posts limit $1", limit)
	if err != nil {
		return
	}
	//iterate rows by row
	for rows.Next() {
		post := Post{}
		//write content of Post from database into post Structure
		err = rows.Scan(&post.Id, &post.Content, &post.Author)
		if err != nil {
			return
		}
		posts = append(posts, post)
	}
	rows.Close()
	return
}
