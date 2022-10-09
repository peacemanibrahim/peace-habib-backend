package api

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/peace-habib-exchange/backend/application"
	"github.com/peace-habib-exchange/backend/domain"
	"github.com/peace-habib-exchange/backend/services"
)

func PeaceApi(handler *mux.Router, db domain.PeaceRepository) *mux.Router {
	service := &services.PeaceService{
		PeaceRepository: db,
	}

	peaceApplication := &application.PeaceController{
		PeaceService: service,
	}

	// You use "handler.HandleFunc" when you are Not adding middlewares. You use "handler.Handle" when you are
	// addding middlewares
	handler.HandleFunc("/signup", peaceApplication.SignUp).Methods(http.MethodPost)
	return handler
}
