package main

import (
	"database/sql"
	"fmt"

	_ "github.com/gorilla/mux"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	vt, _ := sql.Open("sqlite3", "veritabani.db")

	vt.Exec("CREATE TABLE IF NOT EXISTS tablo(Id INTEGER PRIMARY KEY,isim TEXT);")

	vt.QueryRow("SELECT * FROM tablo")
	fmt.Println(vt)

}
