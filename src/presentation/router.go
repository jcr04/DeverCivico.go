// presentation/router.go
package presentation

import (
	"github.com/gorilla/mux"
	"github.com/jcr04/DeverCivico.go/src/application"
)

func NewRouter(cidadaoService *application.CidadaoService) *mux.Router {
	router := mux.NewRouter()
	cidadaoHandler := &CidadaoHandler{
		service: *cidadaoService,
	}

	router.HandleFunc("/cadastrar", cidadaoHandler.CadastrarHandler).Methods("POST")
	// ... outros endpoints

	return router
}
