package handler

import (
	"net/http"

	"github.com/ManojGnanapalam/feedAggregator/respond"
)

func HandlerReadiness(w http.ResponseWriter, r *http.Request) {
	respond.RespondJSON(w, 200, struct{}{})
}

func HandlerErr(w http.ResponseWriter, r *http.Request) {
	respond.ResponseWithError(w, 400, "something wrong")
}
