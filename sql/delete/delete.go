package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:development@/learning_go")

	if err != nil {
		panic(err)
	}

	defer db.Close()

	stmt, _ := db.Prepare("DELETE FROM users WHERE id = ?")

	stmt.Exec(3)
}
