package transport

import "net/http"

func TransportChain(transport http.RoundTripper, url, username, password string) http.RoundTripper {

	basicAuthTransport := &BasicAuthTransport{
		Transport: transport,
		URL:       url,
		Username:  username,
		Password:  password,
	}

	errorTransport := &ErrorTransport{
		Transport: basicAuthTransport,
	}

	return errorTransport
}
