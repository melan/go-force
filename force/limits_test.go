package force

import (
	"testing"
)

func TestLimits(t *testing.T) {
	// Manually grab an OAuth token, so that we can pass it into CreateWithAccessToken
	oauth, err := CreateOAuth(testClientId, testClientSecret, testUserName, testPassword, testSecurityToken,
		testEnvironment)

	if err != nil {
		t.Fatalf("unable to create oauth: %v", err)
	}

	forceApi, err := CreateApi(testVersion, oauth)
	if err != nil {
		t.Fatalf("unable to create force API: %v", err)
	}

	limits, err := forceApi.GetLimits()
	if err != nil {
		// Developer Accounts, which the testbed uses, do not have access to the limits API. So this will always fail.
		// t.Fatalf("Failed to get Limits: %v", err)
		t.Logf("Failed to get Limits, this is expected due to the developer account: %v", err)
	}

	t.Log(limits)
}
