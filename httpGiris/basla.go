package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	var satir string

	file, err := os.OpenFile("dosya.txt", os.O_RDONLY, 0755)
	if err != nil {
		panic(err)
	}

	// Blok sonunda dosyayı kapat
	defer file.Close()

	// Satır satır oku
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		satir = scanner.Text()
		fmt.Println(satir)

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	db, _ := sql.Open("sqlite3", "veritabani.db")
	ekle, _ := db.Prepare("INSERT INTO tablo(isim) values(?)")
	veri, _ := ekle.Exec(satir)
	id, _ := veri.LastInsertId()
	fmt.Println("Kisi id : ", id)
	defer db.Close()

}
