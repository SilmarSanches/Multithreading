package entities

type CommonCepDto struct {
    Cep          string `json:"cep"`
    Street       string `json:"street"`
    Neighborhood string `json:"neighborhood"`
    City         string `json:"city"`
    State        string `json:"state"`
    Source       string `json:"source"`
}