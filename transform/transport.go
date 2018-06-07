package transform

import (
	"net/http"
)

type Transport = func(rt http.RoundTripper) http.RoundTripper

func ComposeTransports(transports ...Transport) Transport {
	return func(rt http.RoundTripper) http.RoundTripper {
		for i := range transports {
			transport := transports[i]
			if transport == nil {
				continue
			}
			rt = transport(rt)
		}
		return rt
	}
}
