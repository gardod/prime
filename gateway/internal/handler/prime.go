package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"tratnik.net/gateway/internal/model"
	"tratnik.net/gateway/internal/service"
	"tratnik.net/gateway/pkg/http/response"
)

type Prime struct {
	primeService service.IPrime
}

func NewPrime(router *mux.Router, primeService service.IPrime) *Prime {
	h := &Prime{
		primeService: primeService,
	}
	h.registerRoutes(router)
	return h
}

func (h *Prime) registerRoutes(r *mux.Router) {
	r.Path("/").Methods("POST").HandlerFunc(h.validate)
}

func (h *Prime) validate(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	prime := model.Prime{}

	err := json.NewDecoder(r.Body).Decode(&prime)
	if err != nil {
		err = fmt.Errorf("malformed body: %w", err)
		response.JSON(w, err, http.StatusBadRequest)
		return
	}

	prime.IsPrime, err = h.primeService.Validate(ctx, prime.N)
	if err != nil {
		response.JSON(w, nil, http.StatusInternalServerError)
		return
	}

	response.JSON(w, prime, http.StatusOK)
}
