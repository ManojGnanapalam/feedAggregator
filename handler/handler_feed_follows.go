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

func (apiCfg *LocalApiconfig) HandlerCreateFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		FeedID uuid.UUID `json:"feed_id"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respond.ResponseWithError(w, 400, fmt.Sprintf("Error parsing JSON:%s", err))
		return
	}
	feed_follows, err := apiCfg.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdateAt:  time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    params.FeedID,
	})
	if err != nil {
		respond.ResponseWithError(w, 400, fmt.Sprintf("could not create feed follow:%s", err))
		return
	}

	respond.RespondJSON(w, 201, feed_follows)
}
