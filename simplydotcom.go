package simplydotcom

import (
	"github.com/caddyserver/caddy/v2"
	"github.com/caddyserver/caddy/v2/caddyconfig/caddyfile"
	libdnssimplydotcom "github.com/libdns/simplydotcom"
)

// Provider lets Caddy read and manipulate DNS records hosted by this DNS provider.
type Provider struct{ *libdnssimplydotcom.Provider }

func init() {
	caddy.RegisterModule(Provider{})
}

// CaddyModule returns the Caddy module information.
func (Provider) CaddyModule() caddy.ModuleInfo {
	return caddy.ModuleInfo{
		ID:  "dns.providers.simplydotcom",
		New: func() caddy.Module { return &Provider{new(libdnssimplydotcom.Provider)} },
	}
}

// TODO: This is just an example. Useful to allow env variable placeholders; update accordingly.
// Provision sets up the module. Implements caddy.Provisioner.
func (p *Provider) Provision(ctx caddy.Context) error {
	p.Provider.APIKey = caddy.NewReplacer().ReplaceAll(p.Provider.APIKey, "")
	p.Provider.AccountName = caddy.NewReplacer().ReplaceAll(p.Provider.AccountName, "")
	p.Provider.BaseURL = caddy.NewReplacer().ReplaceAll(p.Provider.BaseURL, "")
	p.Provider.MaxRetries = caddy.NewReplacer().ReplaceAll(p.Provider.MaxRetries, "")

	return nil
}

// UnmarshalCaddyfile sets up the DNS provider from Caddyfile tokens. Syntax:
//
//	simplydotcom [<api_key>] {
//	    api_key <api_key>
//	}
//
// The older form with an account name before the API key is still accepted,
// but the account name is no longer needed.
func (p *Provider) UnmarshalCaddyfile(d *caddyfile.Dispenser) error {
	for d.Next() {
		switch args := d.RemainingArgs(); len(args) {
		case 0:
		case 1:
			p.Provider.APIKey = args[0]
		case 2:
			p.Provider.AccountName = args[0]
			p.Provider.APIKey = args[1]
		default:
			return d.ArgErr()
		}
		for nesting := d.Nesting(); d.NextBlock(nesting); {
			switch d.Val() {
			case "api_key":
				if p.Provider.APIKey != "" {
					return d.Err("API key already set")
				}
				if d.NextArg() {
					p.Provider.APIKey = d.Val()
				}
				if d.NextArg() {
					return d.ArgErr()
				}

			case "account_name":
				if p.Provider.AccountName != "" {
					return d.Err("Account name already set")
				}
				if d.NextArg() {
					p.Provider.AccountName = d.Val()
				}
				if d.NextArg() {
					return d.ArgErr()
				}

			case "base_url":
				if p.Provider.BaseURL != "" {
					return d.Err("Base URL already set")
				}
				if d.NextArg() {
					p.Provider.BaseURL = d.Val()
				}
				if d.NextArg() {
					return d.ArgErr()
				}

			case "max_retries":
				if p.Provider.MaxRetries != "" {
					return d.Err("Max retries already set")
				}
				if d.NextArg() {
					p.Provider.MaxRetries = d.Val()
				}
				if d.NextArg() {
					return d.ArgErr()
				}

			default:
				return d.Errf("unrecognized subdirective '%s'", d.Val())
			}
		}
	}
	if p.Provider.APIKey == "" {
		return d.Err("missing API key")
	}
	return nil
}

// Interface guards
var (
	_ caddyfile.Unmarshaler = (*Provider)(nil)
	_ caddy.Provisioner     = (*Provider)(nil)
)
