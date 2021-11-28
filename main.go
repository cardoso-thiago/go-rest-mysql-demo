package main

import (
	"log"
	"net/http"
	"os"

	"github.com/cardoso-thiago/go-rest-mysql-demo/app"
	"github.com/cardoso-thiago/go-rest-mysql-demo/app/db"
)

func main() {
	app := app.New()
	app.DB = &db.DB{}
	err := app.DB.Open()
	checkErr(err)

	defer app.DB.Close()

	http.HandleFunc("/", app.Router.ServeHTTP)

	log.Println("Aplicação iniciada na porta 9000!")

	err = http.ListenAndServe(":9000", nil)
	checkErr(err)
}

func checkErr(e error) {
	if e != nil {
		log.Println(e)
		os.Exit(1)
	}
}
