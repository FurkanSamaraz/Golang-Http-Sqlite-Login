package main

import (
	"bufio"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"
)

type Player struct {
	Name    string
	Surname string
}

func sayhelloName(w http.ResponseWriter, r *http.Request) {

	t, _ := template.ParseFiles("index.html")
	t.Execute(w, nil)
}

var Yes2 string
var No2 string

func login(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("login.gtpl")
	t.Execute(w, nil)

	// logic part of log in

	Yes2 = r.FormValue("password")
	No2 = r.FormValue("username")
	fmt.Printf("Adı : %s  \nsoyadı : %s \n", No2, Yes2)

	r.ParseForm()
	file, err := os.Create("dosya.txt")
	if err != nil {
		log.Fatal(err)
	}

	// Fonksiyon sonunda dosyayı kapat
	defer file.Close()

	len, err := file.WriteString(No2 + " " + Yes2 + "\n")

	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Yazılan byte boyutu: " + strconv.Itoa(len))
	}

}

func main() {
	http.HandleFunc("/", sayhelloName) // setting router rule
	http.HandleFunc("/login", login)

	err := http.ListenAndServe(":9090", nil) // setting listening port

	if err != nil {
		log.Fatal("ListenAndServe: ", err)

	}
	file, err := os.OpenFile("dosya.txt", os.O_RDONLY, 0755)
	if err != nil {
		panic(err)
	}

	// Blok sonunda dosyayı kapat
	defer file.Close()

	// Satır satır oku
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		satir := scanner.Text()
		go fmt.Println(satir)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
