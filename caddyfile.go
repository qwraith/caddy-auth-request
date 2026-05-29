package caddyauthrequest

import (
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

func (AuthRequest) CaddyfileModule() caddyfile.ModuleInfo {
	return caddyfile.ModuleInfo{
		ID:  "http.handlers.auth_request",
		New: func() caddyfile.Unmarshaler { return new(AuthRequest) },
	}
}
