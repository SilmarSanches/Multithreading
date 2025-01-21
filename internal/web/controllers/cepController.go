package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/silmarsanches/multithreading/internal/useCase"
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

	fmt.Println("Cep: ", cep.Cep)
	fmt.Println("Street: ", cep.Street)
	fmt.Println("Neighborhood: ", cep.Neighborhood)
	fmt.Println("City: ", cep.City)
	fmt.Println("State: ", cep.State)
	fmt.Println("Source: ", cep.Source)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(cep)
	if err != nil {
		return
	}
}
