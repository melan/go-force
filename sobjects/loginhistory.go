// This file was generated for SObject LoginHistory, API Version v43.0 at 2018-07-30 03:47:21.907311351 -0400 EDT m=+8.250227361

package sobjects

import (
	"fmt"
	"strings"
)

type LoginHistory struct {
	BaseSObject
	ApiType                 string `force:",omitempty"`
	ApiVersion              string `force:",omitempty"`
	Application             string `force:",omitempty"`
	AuthenticationServiceId string `force:",omitempty"`
	Browser                 string `force:",omitempty"`
	CipherSuite             string `force:",omitempty"`
	ClientVersion           string `force:",omitempty"`
	CountryIso              string `force:",omitempty"`
	Id                      string `force:",omitempty"`
	LoginGeoId              string `force:",omitempty"`
	LoginTime               string `force:",omitempty"`
	LoginType               string `force:",omitempty"`
	LoginUrl                string `force:",omitempty"`
	Platform                string `force:",omitempty"`
	SourceIp                string `force:",omitempty"`
	Status                  string `force:",omitempty"`
	TlsProtocol             string `force:",omitempty"`
	UserId                  string `force:",omitempty"`
}

func (t *LoginHistory) ApiName() string {
	return "LoginHistory"
}

func (t *LoginHistory) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("LoginHistory #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tApiType: %v\n", t.ApiType))
	builder.WriteString(fmt.Sprintf("\tApiVersion: %v\n", t.ApiVersion))
	builder.WriteString(fmt.Sprintf("\tApplication: %v\n", t.Application))
	builder.WriteString(fmt.Sprintf("\tAuthenticationServiceId: %v\n", t.AuthenticationServiceId))
	builder.WriteString(fmt.Sprintf("\tBrowser: %v\n", t.Browser))
	builder.WriteString(fmt.Sprintf("\tCipherSuite: %v\n", t.CipherSuite))
	builder.WriteString(fmt.Sprintf("\tClientVersion: %v\n", t.ClientVersion))
	builder.WriteString(fmt.Sprintf("\tCountryIso: %v\n", t.CountryIso))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tLoginGeoId: %v\n", t.LoginGeoId))
	builder.WriteString(fmt.Sprintf("\tLoginTime: %v\n", t.LoginTime))
	builder.WriteString(fmt.Sprintf("\tLoginType: %v\n", t.LoginType))
	builder.WriteString(fmt.Sprintf("\tLoginUrl: %v\n", t.LoginUrl))
	builder.WriteString(fmt.Sprintf("\tPlatform: %v\n", t.Platform))
	builder.WriteString(fmt.Sprintf("\tSourceIp: %v\n", t.SourceIp))
	builder.WriteString(fmt.Sprintf("\tStatus: %v\n", t.Status))
	builder.WriteString(fmt.Sprintf("\tTlsProtocol: %v\n", t.TlsProtocol))
	builder.WriteString(fmt.Sprintf("\tUserId: %v\n", t.UserId))

	return builder.String()
}

type LoginHistoryQueryResponse struct {
	BaseQuery
	Records []LoginHistory `json:"Records" force:"records"`
}
