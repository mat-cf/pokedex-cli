package pokeapi

import (
	"net/http"
	"time"

	"github.com/mat-cf/pokedexcli/internal"
)

type Client struct {
	httpClient http.Client
	cache *internal.Cache
}

func NewClient(timeout, cacheInterval time.Duration) Client {
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
		cache: internal.NewCache(cacheInterval),
	}
}