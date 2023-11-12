package forwarder

import (
	"net/http/cgi"
	"net/http/httputil"
	"net/url"
)

func StartForwarderFor(targetURL *url.URL) error {

	proxy := httputil.NewSingleHostReverseProxy(targetURL)

	return cgi.Serve(proxy)
}
