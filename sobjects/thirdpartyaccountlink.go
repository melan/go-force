// This file was generated for SObject ThirdPartyAccountLink, API Version v43.0 at 2018-07-30 03:47:35.591711306 -0400 EDT m=+21.935140809

package sobjects

import (
	"fmt"
	"strings"
)

type ThirdPartyAccountLink struct {
	BaseSObject
	Handle                   string `force:",omitempty"`
	Id                       string `force:",omitempty"`
	IsNotSsoUsable           bool   `force:",omitempty"`
	Provider                 string `force:",omitempty"`
	RemoteIdentifier         string `force:",omitempty"`
	SsoProviderId            string `force:",omitempty"`
	SsoProviderName          string `force:",omitempty"`
	ThirdPartyAccountLinkKey string `force:",omitempty"`
	UserId                   string `force:",omitempty"`
}

func (t *ThirdPartyAccountLink) ApiName() string {
	return "ThirdPartyAccountLink"
}

func (t *ThirdPartyAccountLink) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("ThirdPartyAccountLink #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tHandle: %v\n", t.Handle))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tIsNotSsoUsable: %v\n", t.IsNotSsoUsable))
	builder.WriteString(fmt.Sprintf("\tProvider: %v\n", t.Provider))
	builder.WriteString(fmt.Sprintf("\tRemoteIdentifier: %v\n", t.RemoteIdentifier))
	builder.WriteString(fmt.Sprintf("\tSsoProviderId: %v\n", t.SsoProviderId))
	builder.WriteString(fmt.Sprintf("\tSsoProviderName: %v\n", t.SsoProviderName))
	builder.WriteString(fmt.Sprintf("\tThirdPartyAccountLinkKey: %v\n", t.ThirdPartyAccountLinkKey))
	builder.WriteString(fmt.Sprintf("\tUserId: %v\n", t.UserId))

	return builder.String()
}

type ThirdPartyAccountLinkQueryResponse struct {
	BaseQuery
	Records []ThirdPartyAccountLink `json:"Records" force:"records"`
}
