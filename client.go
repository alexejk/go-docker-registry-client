package registry

import (
	"crypto/tls"
	"net/http"
	"strings"

	"github.com/alexejk/go-docker-registry-client/transport"
	"github.com/sirupsen/logrus"
)

type RegistryClient struct {
	URL    string
	Client *http.Client
	Log    *logrus.Logger
}

func NewClient(registryUrl, username, password string) (*RegistryClient, error) {

	return newClientWithTransport(registryUrl, username, password, http.DefaultTransport)
}

func NewInsecureClient(registryUrl, username, password string) (*RegistryClient, error) {

	transport := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}

	return newClientWithTransport(registryUrl, username, password, transport)
}

func newClientWithTransport(registryUrl, username, password string, httpTransport http.RoundTripper) (*RegistryClient, error) {

	url := strings.TrimSuffix(registryUrl, "/")
	transportChain := transport.TransportChain(httpTransport, registryUrl, username, password)
	client := &RegistryClient{
		URL: url,
		Client: &http.Client{
			Transport: transportChain,
		},
		Log: logrus.New(),
	}

	return client, nil
}
