package main

import (
	"log"
	"net/http"
	"snippetsharing/handlers"
)

func main() {
	http.HandleFunc("/", handlers.IndexHandler)
	http.HandleFunc("/share", handlers.ShareHandler)

	server := http.Server{
		Addr:    ":8080",
		Handler: nil,
	}
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

// import (
// 	"fmt"
// 	"snippetsharing/databases"
// 	"snippetsharing/models"
// )

// func main() {
// 	c := models.Content{
// 		Hash:       "kakiaf",
// 		Sourcecode: "print('hello')",
// 		Language:   "python",
// 	}
// 	cDb := databases.Content{}

// 	id, _ := cDb.Create(&c)
// 	fmt.Println(id)
// }
