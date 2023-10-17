// presentation/router.go
package presentation

import (
	"github.com/gorilla/mux"
	"github.com/jcr04/DeverCivico.go/src/application"
)

func NewRouter(cidadaoService *application.CidadaoService,
	problemaService *application.ProblemaService,
	discussaoService *application.DiscussaoService,
	informacoesService *application.InformacoesService) *mux.Router {

	router := mux.NewRouter()
	cidadaoHandler := &CidadaoHandler{service: *cidadaoService}
	problemaHandler := &ProblemaHandler{service: *problemaService}
	discussaoHandler := &DiscussaoHandler{service: *discussaoService}
	informacoesHandler := &InformacoesHandler{service: *informacoesService}

	router.HandleFunc("/cadastrar", cidadaoHandler.CadastrarHandler).Methods("POST")
	router.HandleFunc("/login", cidadaoHandler.LoginHandler).Methods("POST")
	router.HandleFunc("/reportar-problema", problemaHandler.ReportarProblemaHandler).Methods("POST")
	router.HandleFunc("/discussoes", discussaoHandler.DiscussoesHandler).Methods("GET")
	router.HandleFunc("/criar-discussao", discussaoHandler.CriarDiscussaoHandler).Methods("POST")
	router.HandleFunc("/informacoes-governamentais", informacoesHandler.InformacoesGovernamentaisHandler).Methods("GET")

	return router
}
