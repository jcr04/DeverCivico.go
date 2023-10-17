// presentation/handlers.go
package presentation

import (
	"encoding/json"
	"net/http"

	"github.com/jcr04/DeverCivico.go/src/application"
	"github.com/jcr04/DeverCivico.go/src/domain"
)

type CidadaoHandler struct {
	service application.CidadaoService
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
