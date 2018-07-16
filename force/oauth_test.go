package force

import (
	"testing"
)

const (
	testVersion       = "v36.0"
	testClientId      = "3MVG9A2kN3Bn17hs8MIaQx1voVGy662rXlC37svtmLmt6wO_iik8Hnk3DlcYjKRvzVNGWLFlGRH1ryHwS217h"
	testClientSecret  = "4165772184959202901"
	testUserName      = "go-force@jalali.net"
	testPassword      = "golangrocks3"
	testSecurityToken = "kAlicVmti9nWRKRiWG3Zvqtte"
	testEnvironment   = "production"
)

func TestOauth(t *testing.T) {
	var logger ForceApiLogger
	oauth, err := CreateOAuth(testClientId, testClientSecret, testUserName, testPassword, testSecurityToken,
		testEnvironment, logger)

	if err != nil {
		t.Fatalf("Unable to create Api for test: %v", err)
	}

	// Verify oauth object is valid
	if err := oauth.Validate(); err != nil {
		t.Fatalf("OAuth object is invlaid: %#v", err)
	}

	// We shouldn't hit any errors creating a new force instance and manually passing in these oauth details now.
	newForceApi, err := CreateOAuthWithAccessToken(testClientId, oauth.AccessToken, oauth.InstanceUrl, logger)
	if err != nil {
		t.Fatalf("Unable to create new force api instance using pre-defined oauth details: %#v", err)
	}

	// Verify oauth object is valid
	if err := newForceApi.Validate(); err != nil {
		t.Fatalf("OAuth object is invlaid: %#v", err)
	}
}
