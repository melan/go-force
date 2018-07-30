// This file was generated for SObject IdpEventLog, API Version v43.0 at 2018-07-30 03:47:53.200361776 -0400 EDT m=+39.544452025

package sobjects

import (
	"fmt"
	"strings"
)

type IdpEventLog struct {
	BaseSObject
	AppId               string `force:",omitempty"`
	AuthSessionId       string `force:",omitempty"`
	ErrorCode           string `force:",omitempty"`
	Id                  string `force:",omitempty"`
	IdentityUsed        string `force:",omitempty"`
	InitiatedBy         string `force:",omitempty"`
	OptionsHasLogoutUrl bool   `force:",omitempty"`
	SamlEntityUrl       string `force:",omitempty"`
	SsoType             string `force:",omitempty"`
	Timestamp           string `force:",omitempty"`
	UserId              string `force:",omitempty"`
}

func (t *IdpEventLog) ApiName() string {
	return "IdpEventLog"
}

func (t *IdpEventLog) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("IdpEventLog #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tAppId: %v\n", t.AppId))
	builder.WriteString(fmt.Sprintf("\tAuthSessionId: %v\n", t.AuthSessionId))
	builder.WriteString(fmt.Sprintf("\tErrorCode: %v\n", t.ErrorCode))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tIdentityUsed: %v\n", t.IdentityUsed))
	builder.WriteString(fmt.Sprintf("\tInitiatedBy: %v\n", t.InitiatedBy))
	builder.WriteString(fmt.Sprintf("\tOptionsHasLogoutUrl: %v\n", t.OptionsHasLogoutUrl))
	builder.WriteString(fmt.Sprintf("\tSamlEntityUrl: %v\n", t.SamlEntityUrl))
	builder.WriteString(fmt.Sprintf("\tSsoType: %v\n", t.SsoType))
	builder.WriteString(fmt.Sprintf("\tTimestamp: %v\n", t.Timestamp))
	builder.WriteString(fmt.Sprintf("\tUserId: %v\n", t.UserId))

	return builder.String()
}

type IdpEventLogQueryResponse struct {
	BaseQuery
	Records []IdpEventLog `json:"Records" force:"records"`
}
