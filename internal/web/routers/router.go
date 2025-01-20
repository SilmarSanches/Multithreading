package routers

import (
	"github.com/go-chi/chi/v5"
	"github.com/silmarsanches/multithreading/internal/web/controllers"
	"net/http"
)

func CepRouters(controllers *controllers.CepController) http.Handler {
	r := chi.NewRouter()
	r.Get("/", controllers.GetCepController)
	return r
}
