package router

import (
    "fmt"

    "github.com/go-chi/chi/v5"
    httpSwagger "github.com/swaggo/http-swagger/v2"

    "github.com/Mlstermass/task2/api/controller"
    // _ "github.com/Mlstermass/task2/swagger"
    "github.com/Mlstermass/task2/pkg/env"
)

func New(ctl controller.LogService, conf env.Config) *chi.Mux {
    r := chi.NewRouter()

    r.Get("/swagger/*", httpSwagger.Handler(
        httpSwagger.URL(fmt.Sprintf("%s/swagger/doc.json", conf.AppHost))))

    r.Route("/", func(r chi.Router) {

    })

    return r
}