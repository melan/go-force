package force

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

const (
	grantType    = "password"
	loginUri     = "https://login.salesforce.com/services/oauth2/token"
	testLoginUri = "https://test.salesforce.com/services/oauth2/token"

	invalidSessionErrorCode = "INVALID_SESSION_ID"
)

type OAuth struct {
	*traceable
	AccessToken string `json:"access_token"`
	InstanceUrl string `json:"instance_url"`
	Id          string `json:"id"`
	IssuedAt    string `json:"issued_at"`
	Signature   string `json:"signature"`

	client        *http.Client
	clientId      string
	clientSecret  string
	refreshToken  string
	userName      string
	password      string
	securityToken string
	environment   string
}

func CreateOAuth(clientId, clientSecret, userName, password, securityToken, environment string) (*OAuth, error) {
	oauth := &OAuth{
		traceable:     &traceable{},
		clientId:      clientId,
		clientSecret:  clientSecret,
		userName:      userName,
		password:      password,
		securityToken: securityToken,
		environment:   environment,
		client:        &http.Client{},
	}

	// Init oauth
	err := oauth.Authenticate()
	if err != nil {
		return nil, err
	}

	return oauth, nil
}

func CreateOAuthWithAccessToken(clientId, accessToken, instanceUrl string) (*OAuth, error) {
	oauth := &OAuth{
		traceable:   &traceable{},
		clientId:    clientId,
		AccessToken: accessToken,
		InstanceUrl: instanceUrl,
		client:      &http.Client{},
	}

	// We need to check for oath correctness here, since we are not generating the token ourselves.
	if err := oauth.Validate(); err != nil {
		return nil, err
	}

	return oauth, nil
}

func CreateOAuthWithRefreshToken(clientId, accessToken, instanceUrl string) (*OAuth, error) {
	oauth := &OAuth{
		traceable:   &traceable{},
		clientId:    clientId,
		AccessToken: accessToken,
		InstanceUrl: instanceUrl,
		client:      &http.Client{},
	}

	// obtain access token
	if err := oauth.RefreshToken(); err != nil {
		return nil, err
	}

	// We need to check for oath correctness here, since we are not generating the token ourselves.
	if err := oauth.Validate(); err != nil {
		return nil, err
	}

	return oauth, nil
}

func (oauth *OAuth) Validate() error {
	if oauth == nil || len(oauth.InstanceUrl) == 0 || len(oauth.AccessToken) == 0 {
		return fmt.Errorf("invalid Force OAuth Object: %#v", oauth)
	}

	return nil
}

func (oauth *OAuth) Expired(apiErrors ApiErrors) bool {
	for _, err := range apiErrors {
		if err.ErrorCode == invalidSessionErrorCode {
			return true
		}
	}

	return false
}

func (oauth *OAuth) Authenticate() error {
	payload := url.Values{
		"grant_type":    {grantType},
		"client_id":     {oauth.clientId},
		"client_secret": {oauth.clientSecret},
		"username":      {oauth.userName},
		"password":      {fmt.Sprintf("%v%v", oauth.password, oauth.securityToken)},
	}

	// Build Uri
	uri := loginUri
	if oauth.environment == "sandbox" {
		uri = testLoginUri
	}

	// Build Body
	body := strings.NewReader(payload.Encode())

	// Build Request
	req, err := http.NewRequest("POST", uri, body)
	if err != nil {
		return fmt.Errorf("error creating authentication request: %v", err)
	}

	// Add Headers
	req.Header.Set("User-Agent", userAgent)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Accept", responseType)

	resp, err := oauth.client.Do(req)
	if err != nil {
		return fmt.Errorf("error sending authentication request: %v", err)
	}
	defer resp.Body.Close()

	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("error reading authentication response bytes: %v", err)
	}

	// Attempt to parse response as a force.com api error
	apiError := &ApiError{}
	if err := json.Unmarshal(respBytes, apiError); err == nil {
		// Check if api error is valid
		if apiError.Validate() {
			return apiError
		}
	}

	if err := json.Unmarshal(respBytes, oauth); err != nil {
		return fmt.Errorf("unable to unmarshal authentication response: %v", err)
	}

	return nil
}

func (oauth *OAuth) RefreshToken() error {
	res := &RefreshTokenResponse{}
	payload := map[string]string{
		"grant_type":    "refresh_token",
		"refresh_token": oauth.refreshToken,
		"client_id":     oauth.clientId,
		"client_secret": oauth.clientSecret,
	}

	err := oauth.request("POST", "/services/oauth2/token", nil, payload, res)
	if err != nil {
		return err
	}

	oauth.AccessToken = res.AccessToken
	return nil
}

func (oauth *OAuth) request(method, path string, params url.Values, payload, out interface{}) error {
	if err := oauth.Validate(); err != nil {
		return fmt.Errorf("error creating %v request: %v", method, err)
	}

	// Build Uri
	var uri bytes.Buffer
	uri.WriteString(oauth.InstanceUrl)
	uri.WriteString(path)
	if params != nil && len(params) != 0 {
		uri.WriteString("?")
		uri.WriteString(params.Encode())
	}

	// Build body
	var body io.Reader
	if payload != nil {

		jsonBytes, err := json.Marshal(payload)
		if err != nil {
			return fmt.Errorf("error marshaling encoded payload: %v", err)
		}

		body = bytes.NewReader(jsonBytes)
	}

	// Build Request
	req, err := http.NewRequest(method, uri.String(), body)
	if err != nil {
		return fmt.Errorf("error creating %v request: %v", method, err)
	}

	// Add Headers
	req.Header.Set("User-Agent", userAgent)
	req.Header.Set("Content-Type", contentType)
	req.Header.Set("Accept", responseType)
	req.Header.Set("Authorization", fmt.Sprintf("%v %v", "Bearer", oauth.AccessToken))

	// Send
	oauth.traceRequest(req)
	resp, err := oauth.client.Do(req)
	if err != nil {
		return fmt.Errorf("error sending %v request: %v", method, err)
	}
	defer resp.Body.Close()
	oauth.traceResponse(resp)

	// Sometimes the force API returns no body, we should catch this early
	if resp.StatusCode == http.StatusNoContent {
		return nil
	}

	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("error reading response bytes: %v", err)
	}
	oauth.traceResponseBody(respBytes)

	// Attempt to parse response into out
	var objectUnmarshalErr error
	if out != nil {
		objectUnmarshalErr = json.Unmarshal(respBytes, out)
		if objectUnmarshalErr == nil {
			return nil
		}
	}

	// Attempt to parse response as a force.com api error before returning object unmarshal err
	apiErrors := ApiErrors{}
	if marshalErr := json.Unmarshal(respBytes, &apiErrors); marshalErr == nil {
		if apiErrors.Validate() {
			// Check if error is oauth token expired
			if oauth.Expired(apiErrors) {
				// Reauthenticate then attempt query again
				oauthErr := oauth.Authenticate()
				if oauthErr != nil {
					return oauthErr
				}

				return oauth.request(method, path, params, payload, out)
			}

			return apiErrors
		}
	}

	if objectUnmarshalErr != nil {
		// Not a force.com api error. Just an unmarshalling error.
		return fmt.Errorf("unable to unmarshal response to object: %v", objectUnmarshalErr)
	}

	// Sometimes no response is expected. For example delete and update. We still have to make sure an error wasn't returned.
	return nil
}

func (oauth *OAuth) traceRequest(req *http.Request) {
	oauth.traceable.trace("Request:", req, "%v")
}

func (oauth *OAuth) traceResponse(resp *http.Response) {
	oauth.traceable.trace("Response:", resp, "%v")
}

func (oauth *OAuth) traceResponseBody(body []byte) {
	oauth.traceable.trace("Response Body:", string(body), "%s")
}
