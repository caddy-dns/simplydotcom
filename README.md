Simply.com module for Caddy
===========================

This package contains a DNS provider module for [Caddy](https://github.com/caddyserver/caddy). It can be used to manage DNS records with Simply.com.

## Caddy module name

```
dns.providers.simplydotcom
```

## Config examples

To use this module for the ACME DNS challenge, [configure the ACME issuer in your Caddy JSON](https://caddyserver.com/docs/json/apps/tls/automation/policies/issuer/acme/) like so:

```json
{
	"module": "acme",
	"challenges": {
		"dns": {
			"provider": {
				"name": "simplydotcom",
				"api_key": "YourSecretKey"
			}
		}
	}
}
```

or with the Caddyfile:

```
# globally
{
	acme_dns simplydotcom {
		api_key "YourSecretKey"
	}
}
```

```
# one site
secure.example.com {
	tls {
		dns simplydotcom {
			api_key "YourSecretKey"
		}
	}
}

# reusable import
(dnschallenge) {
	tls {
		dns simplydotcom <api_key>
	}
}

example.com {
	respond "Hello World" 
	import dnschallenge
}
```

The API key alone identifies the Simply.com account, see the [Simply.com API documentation](https://www.simply.com/en/docs/api/). The `account_name` option, and the `<account_name> <api_key>` form, are still accepted so existing configurations keep working, but the account name is no longer needed.
