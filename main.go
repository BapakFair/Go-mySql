package main

import (
	"fmt"
	"go_sql/config"
	"log"
	"net/http"

	. "go_sql/functions"

	"github.com/julienschmidt/httprouter"
)

const PORT string = ":8088"

func main() {

	db, err := config.MySql()
	if err != nil {
		log.Fatal(err)
	}
	errDB := db.Ping()
	if errDB != nil {
		panic(errDB.Error())
	}
	fmt.Println("gasss...")

	router := httprouter.New()
	router.GET("/getAllData", GetMovie)
	router.POST("/postData", PostMovie)
	router.PUT("/:id/putData", PutMovie)
	router.DELETE("/:id/delData", DelMovie)

	fmt.Println("server running ", PORT)
	log.Fatal(http.ListenAndServe(PORT, router))
}
