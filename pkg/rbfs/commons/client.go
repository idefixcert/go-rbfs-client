package commons

import (
	"net/http"
	"net/url"
	"sync"

	"gitlab.rtbrick.net/martin/bricklet/pkg/rbfs/state"
)

var clients = make(map[string]*state.APIClient)
var x = &sync.Mutex{}

// GetAPIClient creates a new API client for the given endpoint.
func GetAPIClient(client *http.Client, endpoint *url.URL) *state.APIClient {
	// TODO Cache API client, because it is supposed to be only a few per pod?
	host := endpoint.Host

	apiClient := clients[host]

	if apiClient == nil {
		x.Lock()
		defer x.Unlock()
		apiClient = clients[host]
		if apiClient == nil {
			// If client is still nil, i.e. was not created in the meanwhile by another request,
			// then create a new client.
			config := state.NewConfiguration()
			config.BasePath = endpoint.String()
			config.Host = endpoint.Host
			config.HTTPClient = client
			config.UserAgent = "Diagnostic Actor"
			apiClient = state.NewAPIClient(config)
			clients[host] = apiClient
		}
	}

	return apiClient
}
