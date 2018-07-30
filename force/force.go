// A Go package that provides bindings to the force.com REST API
//
// See http://www.salesforce.com/us/developer/docs/api_rest/
package force

import (
	"net/http"
)

func CreateApi(version string, oauth *OAuth) (*Api, error) {
	client := &http.Client{}

	forceApi := &Api{
		traceable:              &traceable{},
		apiResources:           make(map[string]string),
		apiSObjects:            make(map[string]*SObjectMetaData),
		apiSObjectDescriptions: make(map[string]*SObjectDescription),
		apiVersion:             version,
		oauth:                  oauth,
		client:                 client,
	}

	// Init Api Resources
	err := forceApi.getApiResources()
	if err != nil {
		return nil, err
	}
	err = forceApi.getApiSObjects()
	if err != nil {
		return nil, err
	}

	return forceApi, nil
}

type ForceApiLogger interface {
	Printf(format string, v ...interface{})
}

// TraceOn turns on logging for this Api. After this is called, all
// requests, responses, and raw response bodies will be sent to the logger.
// If prefix is a non-empty string, it will be written to the front of all
// logged strings, which can aid in filtering log lines.
//
// Use TraceOn if you want to spy on the Api requests and responses.
//
// Note that the base log.Logger type satisfies ForceApiLogger, but adapters
// can easily be written for other logging packages (e.g., the
// golang-sanctioned glog framework).
func (forceApi *Api) TraceOn(prefix string, logger ForceApiLogger) {
	forceApi.traceable.TraceOn(prefix, logger)
	forceApi.oauth.TraceOn(prefix, logger)
}

// TraceOff turns off tracing. It is idempotent.
func (forceApi *Api) TraceOff() {
	forceApi.traceable.TraceOff()
	forceApi.oauth.TraceOff()
}
