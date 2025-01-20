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

type ExternalServiceBrasilApiInterface interface {
	GetCep(ctx context.Context, cep string) (map[string]interface{}, error)
}

type HttpExternalServiceBrasilApi struct {
	HttpClient *http.Client
	appConfig  config.Config
}

func NewHttpExternalServiceBrasilApi(appConfig *config.Config) *HttpExternalServiceBrasilApi {
	return &HttpExternalServiceBrasilApi{
		HttpClient: &http.Client{},
		appConfig:  *appConfig,
	}
}

func (e *HttpExternalServiceBrasilApi) GetCep(ctx context.Context, cep string) (map[string]interface{}, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	url := e.appConfig.URLBrasilAPI + "/api/cep/v1/" + cep

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	res, err := e.HttpClient.Do(req)
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			return nil, fmt.Errorf("timeout de 3s excedido ao consultar o serviço brasilapi: %v", err)
		}
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Printf("erro ao fechar o corpo da resposta brasilapi: %v", err)
		}
	}(res.Body)

	if res.StatusCode != http.StatusOK {
		return nil, errors.New("erro ao consultar o serviço brasilapi: " + res.Status)
	}

	var data map[string]interface{}
	err = json.NewDecoder(res.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
