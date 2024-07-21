package mytype

import (
	"net/http"

	"github.com/ManojGnanapalam/feedAggregator/internal/database"
)

type ApiConfig struct {
	DB *database.Queries
}

type AuthedHandler func(http.ResponseWriter, *http.Request, database.User)
