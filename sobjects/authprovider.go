// This file was generated for SObject AuthProvider, API Version v43.0 at 2018-07-30 03:47:40.903710767 -0400 EDT m=+27.247339598

package sobjects

import (
	"fmt"
	"strings"
)

type AuthProvider struct {
	BaseSObject
	AuthorizeUrl                         string `force:",omitempty"`
	ConsumerKey                          string `force:",omitempty"`
	ConsumerSecret                       string `force:",omitempty"`
	CreatedDate                          string `force:",omitempty"`
	CustomMetadataTypeRecord             string `force:",omitempty"`
	DefaultScopes                        string `force:",omitempty"`
	DeveloperName                        string `force:",omitempty"`
	ErrorUrl                             string `force:",omitempty"`
	ExecutionUserId                      string `force:",omitempty"`
	FriendlyName                         string `force:",omitempty"`
	IconUrl                              string `force:",omitempty"`
	Id                                   string `force:",omitempty"`
	IdTokenIssuer                        string `force:",omitempty"`
	LinkKickoffUrl                       string `force:",omitempty"`
	LogoutUrl                            string `force:",omitempty"`
	OauthKickoffUrl                      string `force:",omitempty"`
	OptionsIncludeOrgIdInId              bool   `force:",omitempty"`
	OptionsSendAccessTokenInHeader       bool   `force:",omitempty"`
	OptionsSendClientCredentialsInHeader bool   `force:",omitempty"`
	PluginId                             string `force:",omitempty"`
	ProviderType                         string `force:",omitempty"`
	RegistrationHandlerId                string `force:",omitempty"`
	SsoKickoffUrl                        string `force:",omitempty"`
	TokenUrl                             string `force:",omitempty"`
	UserInfoUrl                          string `force:",omitempty"`
}

func (t *AuthProvider) ApiName() string {
	return "AuthProvider"
}

func (t *AuthProvider) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("AuthProvider #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tAuthorizeUrl: %v\n", t.AuthorizeUrl))
	builder.WriteString(fmt.Sprintf("\tConsumerKey: %v\n", t.ConsumerKey))
	builder.WriteString(fmt.Sprintf("\tConsumerSecret: %v\n", t.ConsumerSecret))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tCustomMetadataTypeRecord: %v\n", t.CustomMetadataTypeRecord))
	builder.WriteString(fmt.Sprintf("\tDefaultScopes: %v\n", t.DefaultScopes))
	builder.WriteString(fmt.Sprintf("\tDeveloperName: %v\n", t.DeveloperName))
	builder.WriteString(fmt.Sprintf("\tErrorUrl: %v\n", t.ErrorUrl))
	builder.WriteString(fmt.Sprintf("\tExecutionUserId: %v\n", t.ExecutionUserId))
	builder.WriteString(fmt.Sprintf("\tFriendlyName: %v\n", t.FriendlyName))
	builder.WriteString(fmt.Sprintf("\tIconUrl: %v\n", t.IconUrl))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tIdTokenIssuer: %v\n", t.IdTokenIssuer))
	builder.WriteString(fmt.Sprintf("\tLinkKickoffUrl: %v\n", t.LinkKickoffUrl))
	builder.WriteString(fmt.Sprintf("\tLogoutUrl: %v\n", t.LogoutUrl))
	builder.WriteString(fmt.Sprintf("\tOauthKickoffUrl: %v\n", t.OauthKickoffUrl))
	builder.WriteString(fmt.Sprintf("\tOptionsIncludeOrgIdInId: %v\n", t.OptionsIncludeOrgIdInId))
	builder.WriteString(fmt.Sprintf("\tOptionsSendAccessTokenInHeader: %v\n", t.OptionsSendAccessTokenInHeader))
	builder.WriteString(fmt.Sprintf("\tOptionsSendClientCredentialsInHeader: %v\n", t.OptionsSendClientCredentialsInHeader))
	builder.WriteString(fmt.Sprintf("\tPluginId: %v\n", t.PluginId))
	builder.WriteString(fmt.Sprintf("\tProviderType: %v\n", t.ProviderType))
	builder.WriteString(fmt.Sprintf("\tRegistrationHandlerId: %v\n", t.RegistrationHandlerId))
	builder.WriteString(fmt.Sprintf("\tSsoKickoffUrl: %v\n", t.SsoKickoffUrl))
	builder.WriteString(fmt.Sprintf("\tTokenUrl: %v\n", t.TokenUrl))
	builder.WriteString(fmt.Sprintf("\tUserInfoUrl: %v\n", t.UserInfoUrl))

	return builder.String()
}

type AuthProviderQueryResponse struct {
	BaseQuery
	Records []AuthProvider `json:"Records" force:"records"`
}
