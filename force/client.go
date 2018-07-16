package force

import (
	"net/url"
)

const (
	version      = "1.0.0"
	userAgent    = "go-force/" + version
	contentType  = "application/json"
	responseType = "application/json"
)

// Get issues a GET to the specified path with the given params and put the
// umarshalled (json) result in the third parameter
func (forceApi *Api) Get(path string, params url.Values, out interface{}) error {
	return forceApi.oauth.request("GET", path, params, nil, out)
}

// Post issues a POST to the specified path with the given params and payload
// and put the unmarshalled (json) result in the third parameter
func (forceApi *Api) Post(path string, params url.Values, payload, out interface{}) error {
	return forceApi.oauth.request("POST", path, params, payload, out)
}

// Put issues a PUT to the specified path with the given params and payload
// and put the unmarshalled (json) result in the third parameter
func (forceApi *Api) Put(path string, params url.Values, payload, out interface{}) error {
	return forceApi.oauth.request("PUT", path, params, payload, out)
}

// Patch issues a PATCH to the specified path with the given params and payload
// and put the unmarshalled (json) result in the third parameter
func (forceApi *Api) Patch(path string, params url.Values, payload, out interface{}) error {
	return forceApi.oauth.request("PATCH", path, params, payload, out)
}

// Delete issues a DELETE to the specified path with the given payload
func (forceApi *Api) Delete(path string, params url.Values) error {
	return forceApi.oauth.request("DELETE", path, params, nil, nil)
}
