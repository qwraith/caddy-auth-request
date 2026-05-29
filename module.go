package caddyauthrequest

import (
	"github.com/caddyserver/caddy/v2"
)

func init() {
	caddy.RegisterModule(AuthRequest{})
}

func (AuthRequest) CaddyModule() caddy.ModuleInfo {
	return caddy.ModuleInfo{
		ID:  "http.handlers.auth_request",
		New: func() caddy.Module { return new(AuthRequest) },
	}
}
