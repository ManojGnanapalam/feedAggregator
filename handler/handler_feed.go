package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/ManojGnanapalam/feedAggregator/internal/database"
	"github.com/ManojGnanapalam/feedAggregator/respond"
	"github.com/google/uuid"
)

func (apiCfg *LocalApiconfig) HandlerCreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respond.ResponseWithError(w, 400, fmt.Sprintf("Error parsing JSON:%s", err))
		return
	}
	feed, err := apiCfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdateAt:  time.Now().UTC(),
		Name:      params.Name,
		Url:       params.URL,
		UserID:    user.ID,
	})
	if err != nil {
		respond.ResponseWithError(w, 400, fmt.Sprintf("could not create user:%s", err))
		return
	}

	respond.RespondJSON(w, 201, feed)
}

func (apiCfg *LocalApiconfig) HandlerGetFeed(w http.ResponseWriter, r *http.Request, feed database.Feed) {
	respond.RespondJSON(w, 200, feed)
}
