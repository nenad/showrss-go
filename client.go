package showrss

import (
    "net/http"
)

const ApiBaseUrl = "https://showrss.info"

type Client struct {
    internal *http.Client
}

func NewClient(internal *http.Client) *Client {
    return &Client{ internal: internal}
}
