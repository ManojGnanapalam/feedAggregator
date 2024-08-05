package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/ManojGnanapalam/feedAggregator/internal/database"
	"github.com/ManojGnanapalam/feedAggregator/respond"
	"github.com/go-chi/chi"
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

func (apiCfg *LocalApiconfig) HandlerGetFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {

	feed_follows, err := apiCfg.DB.GetFeedFollow(r.Context(), user.ID)
	if err != nil {
		respond.ResponseWithError(w, 400, fmt.Sprintf("could not get feed follow:%s", err))
		return
	}

	respond.RespondJSON(w, 201, feed_follows)
}

func (apiCfg *LocalApiconfig) HandlerDeleteFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	feedFollowIDStr := chi.URLParam(r, "feedFolloID")

	feedFollowID, err := uuid.Parse(feedFollowIDStr)

	if err != nil {
		respond.ResponseWithError(w, 400, fmt.Sprintf("could't find the feed follow id %v", err))
	}

	err = apiCfg.DB.DeleteFeedFollow(r.Context(), database.DeleteFeedFollowParams{ID: feedFollowID, UserID: user.ID})
	if err != nil {
		respond.ResponseWithError(w, 400, fmt.Sprintf("couldn't  delete feed follow%v", err))
	}
	respond.RespondJSON(w, 200, struct{}{})
}
