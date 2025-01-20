package services

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/silmarsanches/multithreading/config"
	"io"
	"net/http"
	"time"
)

type ExternalServiceViaCepInterface interface {
	GetCep(ctx context.Context, cep string) (map[string]interface{}, error)
}

type HttpExternalServiceViaCep struct {
	HttpClient *http.Client
	appConfig  config.Config
}

func NewHttpExternalServiceViaCep(appConfig *config.Config) *HttpExternalServiceViaCep {
	return &HttpExternalServiceViaCep{
		HttpClient: &http.Client{},
		appConfig:  *appConfig,
	}
}

func (e *HttpExternalServiceViaCep) GetCep(ctx context.Context, cep string) (map[string]interface{}, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	url := e.appConfig.URLViaCep + "/" + cep + "/json"

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	res, err := e.HttpClient.Do(req)
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			return nil, fmt.Errorf("Timeout de 3s excedido ao consultar o serviço ViaCep: %v", err)
		}
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Printf("Erro ao fechar o corpo da resposta ViaCep: %v", err)
		}
	}(res.Body)

	if res.StatusCode != http.StatusOK {
		return nil, errors.New("erro ao consultar o serviço ViaCep: " + res.Status)
	}

	var data map[string]interface{}
	err = json.NewDecoder(res.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
