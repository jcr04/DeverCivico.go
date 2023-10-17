// DeverCivico.go
package main

import (
	"net/http"

	"DeverCivico.go/infrastructure"
	"DeverCivico.go/presentation"
)

func main() {
	db, err := infrastructure.NewDB()
	if err != nil {
		//... tratamento de erro
	}
	router := presentation.NewRouter()
	http.ListenAndServe(":8080", router)
}
