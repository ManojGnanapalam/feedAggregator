package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/ManojGnanapalam/feedAggregator/internal/database"
	"github.com/ManojGnanapalam/feedAggregator/respond"
	"github.com/ManojGnanapalam/feedAggregator/shared"
	"github.com/google/uuid"
)

type LocalApiconfig struct {
	shared.ApiConfig
}

func (apiCfg *LocalApiconfig) HandlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respond.ResponseWithError(w, 400, fmt.Sprintf("Error parsing JSON:%s", err))
		return
	}
	user, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdateAt:  time.Now().UTC(),
		Name:      params.Name,
	})
	if err != nil {
		respond.ResponseWithError(w, 400, fmt.Sprintf("could not create user:%s", err))
		return
	}

	respond.RespondJSON(w, 200, user)
}
