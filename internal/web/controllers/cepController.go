package controllers

import (
	"encoding/json"
	"github.com/silmarsanches/multithreading/internal/useCase"
	"net/http"
)

type CepController struct {
	useCase.UseCaseGetCep
}

func NewCepController(useCase *useCase.UseCaseGetCep) *CepController {
	return &CepController{
		UseCaseGetCep: *useCase,
	}
}

func (c *CepController) GetCepController(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	cepUrl := r.URL.Query().Get("cep")
	if cepUrl == "" {
		http.Error(w, "O CEP é obrigatório", http.StatusBadRequest)
		return
	}
	cep, err := c.GetCep(ctx, cepUrl)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(cep)
	if err != nil {
		return
	}
}
