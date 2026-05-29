package caddyauthrequest

import (
	"net/http"

	"github.com/caddyserver/caddy/v2/modules/caddyhttp"
)

type AuthRequest struct {
	URI string `json:"uri,omitempty"`
}

func (a AuthRequest) ServeHTTP(w http.ResponseWriter, r *http.Request, next caddyhttp.Handler) error {
	req, err := http.NewRequest("GET", a.URI, nil)
	if err != nil {
		return err
	}

	req.Header = r.Header.Clone()

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		return next.ServeHTTP(w, r)
	}

	w.WriteHeader(resp.StatusCode)
	return nil
}
