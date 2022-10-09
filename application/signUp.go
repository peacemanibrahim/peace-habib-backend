package application

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/peace-habib-exchange/backend/domain"
)

type PeaceController struct {
	PeaceService domain.PeaceService
}

func (a PeaceController) SignUp(w http.ResponseWriter, r *http.Request) {
	var user domain.User
	json.NewDecoder(r.Body).Decode(&user)
	validate := validator.New()
	err := validate.Struct(&user)
	if err != nil {
		// Return a bad request and a helpful error message if you wished, you could concat the validation error into
		// this message to help point your consumer in the right direction
		http.Error(w, "failed to validate struct", 400)
		return
	}
	ctx := context.Background()
	result, _ := a.PeaceService.CreateUser(ctx, user)
	json.NewEncoder(w).Encode(result)
}
