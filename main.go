package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/agustin-sarasua/rs-model"

	"github.com/agustin-sarasua/rs-user-api/app"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	db, err := gorm.Open("mysql", app.ConnectionString)
	app.Db = db
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()
	db.LogMode(true)

	// Migrate the schema
	var user m.User
	db.DropTableIfExists(&user)
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&user)

	router := mux.NewRouter()
	router.HandleFunc("/user", app.CreateUserEndpoint).Methods("POST")
	router.HandleFunc("/user/{id:.+}", app.GetUserEndpoint).Methods("GET")
	router.HandleFunc("/user/{id:.+}", app.UpdateUserEndpoint).Methods("PUT")

	fmt.Println("Hello User API")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func use(h http.HandlerFunc, middleware ...func(http.HandlerFunc) http.HandlerFunc) http.HandlerFunc {
	for _, m := range middleware {
		h = m(h)
	}
	return h
}
