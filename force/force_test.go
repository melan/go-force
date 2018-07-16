package force

import (
	"github.com/nimajalali/go-force/sobjects"
	"testing"
)

func TestCreateWithAccessToken(t *testing.T) {
	var logger ForceApiLogger
	// Manually grab an OAuth token, so that we can pass it into CreateWithAccessToken
	oauth, err := CreateOAuth(testClientId, testClientSecret, testUserName, testPassword, testSecurityToken,
		testEnvironment, logger)

	if err != nil {
		t.Fatalf("unable to create oauth: %v", err)
	}

	forceApi, err := CreateApi(testVersion, oauth)
	if err != nil {
		t.Fatalf("unable to create force API: %v", err)
	}

	// We should be able to make a basic query now with the newly created object (i.e. the oauth details should be correctly usable).
	_, err = forceApi.DescribeSObject(&sobjects.Account{})
	if err != nil {
		t.Fatalf("Failed to retrieve description of sobject: %v", err)
	}
}
