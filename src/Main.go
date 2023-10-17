package main

import (
	"log"
	"net/http"

	"github.com/jcr04/DeverCivico.go/src/application"
	"github.com/jcr04/DeverCivico.go/src/infrastructure"
	"github.com/jcr04/DeverCivico.go/src/presentation"
)

func main() {
	db, err := infrastructure.NewDB()
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	cidadaoRepo := infrastructure.NewCidadaoRepository(db)
	cidadaoService := application.NewCidadaoService(*cidadaoRepo)

	problemaRepo := infrastructure.NewProblemaRepository(db)
	problemaService := application.NewProblemaService(*problemaRepo)

	discussaoRepo := infrastructure.NewDiscussaoRepository(db)
	discussaoService := application.NewDiscussaoService(*discussaoRepo)

	informacoesRepo := infrastructure.NewInformacoesRepository(db)
	informacoesService := application.NewInformacoesService(*informacoesRepo)

	router := presentation.NewRouter(cidadaoService, problemaService, discussaoService, informacoesService)

	http.ListenAndServe(":8080", router)
}
