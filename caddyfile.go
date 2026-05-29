package caddyauthrequest

import (
	"github.com/caddyserver/caddy/v2"
	"github.com/caddyserver/caddy/v2/caddyconfig/caddyfile"
)

func (a *AuthRequest) UnmarshalCaddyfile(d *caddyfile.Dispenser) error {
	for d.Next() {
		if !d.Args(&a.URI) {
			return d.ArgErr()
		}
	}
	return nil
}

func (AuthRequest) CaddyModule() caddy.ModuleInfo {
    return caddy.ModuleInfo{
        ID:  "http.handlers.auth_request",
        New: func() caddy.Module { return new(AuthRequest) },
    }
}


