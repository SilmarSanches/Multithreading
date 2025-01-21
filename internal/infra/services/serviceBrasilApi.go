package services

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/silmarsanches/multithreading/config"
	"github.com/silmarsanches/multithreading/internal/entities"
)

type ExternalServiceBrasilApiInterface interface {
	GetCep(ctx context.Context, cep string) (entities.BrasilApiDto, error)
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

func (e *HttpExternalServiceBrasilApi) GetCep(ctx context.Context, cep string) (entities.BrasilApiDto, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	url := e.appConfig.URLBrasilAPI + "/api/cep/v1/" + cep

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return entities.BrasilApiDto{}, err
	}

	res, err := e.HttpClient.Do(req)
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			return entities.BrasilApiDto{}, fmt.Errorf("timeout de 3s excedido ao consultar o serviço brasilapi: %v", err)
		}
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Printf("erro ao fechar o corpo da resposta brasilapi: %v", err)
		}
	}(res.Body)

	if res.StatusCode != http.StatusOK {
		return entities.BrasilApiDto{}, errors.New("erro ao consultar o serviço brasilapi: " + res.Status)
	}

	var data entities.BrasilApiDto
	err = json.NewDecoder(res.Body).Decode(&data)
	if err != nil {
		return entities.BrasilApiDto{}, err
	}

	return data, nil
}
