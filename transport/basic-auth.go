package transport

import (
	"net/http"
	"strings"
)

type BasicAuthTransport struct {
	Transport http.RoundTripper
	URL       string
	Username  string
	Password  string
}

func (t *BasicAuthTransport) RoundTrip(req *http.Request) (*http.Response, error) {

	if strings.HasPrefix(req.URL.String(), t.URL) {
		if t.Username != "" || t.Password != "" {
			req.SetBasicAuth(t.Username, t.Password)
		}
	}
	resp, err := t.Transport.RoundTrip(req)
	return resp, err
}
