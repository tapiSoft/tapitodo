package main
import (
	"log"
)

const debug = true

func main() {
	db, err := OpenDatabase("db.sqlite3")
	defer db.Close()
	if err != nil {
		log.Fatal(err)
		return
	}
	db.InsertTodo("first todo")
}
