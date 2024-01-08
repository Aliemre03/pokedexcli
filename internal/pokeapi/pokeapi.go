package pokeapi

import (
	"github.com/Aliemre03/pokedexcli/pokecache"
	"net/http"
	"time"
)

const baseURL = "https://pokeapi.co/api/v2/"

type Client struct {
	cache      *pokecache.Cache
	httpClient http.Client
}

func NewClient() *Client {
	return &Client{
		cache: pokecache.NewCache(),
		httpClient: http.Client{
			Timeout: time.Minute,
		},
	}
}
