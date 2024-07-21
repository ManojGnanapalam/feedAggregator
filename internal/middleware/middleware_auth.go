package middleware

import (
	"fmt"
	"net/http"

	"github.com/ManojGnanapalam/feedAggregator/internal/auth"
	"github.com/ManojGnanapalam/feedAggregator/internal/mytype"
	"github.com/ManojGnanapalam/feedAggregator/respond"
)

type LocalApiconfig struct {
	mytype.ApiConfig
}

func MiddlewareAuth(apiCfg *LocalApiconfig, handler mytype.AuthedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetAPIKey(r.Header)
		if err != nil {
			respond.ResponseWithError(w, 403, fmt.Sprintf("Auth error:%v", err))
			return
		}

		user, err := apiCfg.DB.GetUserByAPIKey(r.Context(), apiKey)
		if err != nil {
			respond.ResponseWithError(w, 404, fmt.Sprintf("Couldn't get user:%v", err))
			return
		}
		handler(w, r, user)
	}
}
