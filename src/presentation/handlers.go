// presentation/handlers.go
package presentation

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/jcr04/DeverCivico.go/src/application"
	"github.com/jcr04/DeverCivico.go/src/domain"
)

type CidadaoHandler struct {
	service application.CidadaoService
}

type ProblemaHandler struct {
	service application.ProblemaService
}

type DiscussaoHandler struct {
	service application.DiscussaoService
}

type InformacoesHandler struct {
	service application.InformacoesService
}

func (h *CidadaoHandler) CadastrarHandler(w http.ResponseWriter, r *http.Request) {
	var cidadao domain.Cidadao
	err := json.NewDecoder(r.Body).Decode(&cidadao)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.service.CadastrarCidadao(cidadao)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *CidadaoHandler) LoginHandler(w http.ResponseWriter, r *http.Request) {
	var credentials domain.Credentials
	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	cidadao, err := h.service.Login(credentials)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	// Exemplo: configurando um cookie com o ID do cidadão
	http.SetCookie(w, &http.Cookie{
		Name:  "session",
		Value: strconv.Itoa(cidadao.ID), // Converta o ID do cidadão para uma string
	})

	w.WriteHeader(http.StatusOK)
}

func (h *ProblemaHandler) ReportarProblemaHandler(w http.ResponseWriter, r *http.Request) {
	var problema domain.ProblemaReportado
	err := decodeJSON(r, &problema)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.service.ReportarProblema(problema)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *DiscussaoHandler) DiscussoesHandler(w http.ResponseWriter, r *http.Request) {
	discussoes, err := h.service.ListarDiscussoes()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	responseData, err := json.Marshal(discussoes)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseData)
}

func (h *DiscussaoHandler) CriarDiscussaoHandler(w http.ResponseWriter, r *http.Request) {
	var discussao domain.Discussao
	err := decodeJSON(r, &discussao)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.service.CriarDiscussao(discussao)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *InformacoesHandler) InformacoesGovernamentaisHandler(w http.ResponseWriter, r *http.Request) {
	info, err := h.service.ObterInformacoes()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	responseData, err := json.Marshal(info)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseData)
}
