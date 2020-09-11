package server

import (
	"compress/flate"
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/unrolled/secure"
)

// Helmet provides essential security extensions
func Helmet() func(http.Handler) http.Handler {
	return secure.New(secure.Options{
		// AllowedHosts:          []string{"example\\.com", ".*\\.example\\.com"},
		// AllowedHostsAreRegex:  true,
		HostsProxyHeaders: []string{"X-Forwarded-Host"},
		// SSLRedirect:       true,
		// SSLHost:               "ssl.example.com",
		SSLProxyHeaders:       map[string]string{"X-Forwarded-Proto": "https"},
		STSSeconds:            31536000,
		STSIncludeSubdomains:  true,
		STSPreload:            true,
		FrameDeny:             true,
		ContentTypeNosniff:    true,
		BrowserXssFilter:      true,
		ReferrerPolicy:        "same-origin",
		ContentSecurityPolicy: "script-src $NONCE",
	}).Handler
}

// Compression provides gzip compression
func Compression() func(http.Handler) http.Handler {
	return middleware.NewCompressor(flate.DefaultCompression).Handler
}
