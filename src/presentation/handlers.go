// presentation/handlers.go
package presentation

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

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

	err = h.service.CadastrarCidadao(&cidadao)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	response := map[string]interface{}{
		"message": "Cidadão cadastrado com sucesso",
		"cidadao": cidadao,
	}
	json.NewEncoder(w).Encode(response)
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

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(cidadao) // Serializa o objeto cidadão em JSON e escreve na resposta
}

func (h *ProblemaHandler) ReportarProblemaHandler(w http.ResponseWriter, r *http.Request) {
	var problema domain.ProblemaReportado
	err := decodeJSON(r, &problema)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if problema.DataHora == "" {
		problema.DataHora = time.Now().Format(time.RFC3339) // ou outro formato que você deseja usar
	}

	log.Printf("Reportando problema: %+v\n", problema)

	err = h.service.ReportarProblema(&problema)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json") // Defina o tipo de conteúdo da resposta para JSON
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(problema) // Serializa o objeto problema em JSON e escreve na resposta
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
	json.NewEncoder(w).Encode(discussoes)
}

func (h *DiscussaoHandler) CriarDiscussaoHandler(w http.ResponseWriter, r *http.Request) {
	var discussao domain.Discussao
	err := decodeJSON(r, &discussao)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.service.CriarDiscussao(&discussao)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(discussao)
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
	json.NewEncoder(w).Encode(info)
}
