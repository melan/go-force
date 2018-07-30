// This file was generated for SObject ConnectedApplication, API Version v43.0 at 2018-07-30 03:48:09.545623955 -0400 EDT m=+55.890327543

package sobjects

import (
	"fmt"
	"strings"
)

type ConnectedApplication struct {
	BaseSObject
	CreatedById                         string `force:",omitempty"`
	CreatedDate                         string `force:",omitempty"`
	Id                                  string `force:",omitempty"`
	LastModifiedById                    string `force:",omitempty"`
	LastModifiedDate                    string `force:",omitempty"`
	MobileSessionTimeout                string `force:",omitempty"`
	MobileStartUrl                      string `force:",omitempty"`
	Name                                string `force:",omitempty"`
	OptionsAllowAdminApprovedUsersOnly  bool   `force:",omitempty"`
	OptionsFullContentPushNotifications bool   `force:",omitempty"`
	OptionsHasSessionLevelPolicy        bool   `force:",omitempty"`
	OptionsIsInternal                   bool   `force:",omitempty"`
	OptionsRefreshTokenValidityMetric   bool   `force:",omitempty"`
	PinLength                           string `force:",omitempty"`
	RefreshTokenValidityPeriod          int    `force:",omitempty"`
	StartUrl                            string `force:",omitempty"`
	SystemModstamp                      string `force:",omitempty"`
}

func (t *ConnectedApplication) ApiName() string {
	return "ConnectedApplication"
}

func (t *ConnectedApplication) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("ConnectedApplication #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tLastModifiedById: %v\n", t.LastModifiedById))
	builder.WriteString(fmt.Sprintf("\tLastModifiedDate: %v\n", t.LastModifiedDate))
	builder.WriteString(fmt.Sprintf("\tMobileSessionTimeout: %v\n", t.MobileSessionTimeout))
	builder.WriteString(fmt.Sprintf("\tMobileStartUrl: %v\n", t.MobileStartUrl))
	builder.WriteString(fmt.Sprintf("\tName: %v\n", t.Name))
	builder.WriteString(fmt.Sprintf("\tOptionsAllowAdminApprovedUsersOnly: %v\n", t.OptionsAllowAdminApprovedUsersOnly))
	builder.WriteString(fmt.Sprintf("\tOptionsFullContentPushNotifications: %v\n", t.OptionsFullContentPushNotifications))
	builder.WriteString(fmt.Sprintf("\tOptionsHasSessionLevelPolicy: %v\n", t.OptionsHasSessionLevelPolicy))
	builder.WriteString(fmt.Sprintf("\tOptionsIsInternal: %v\n", t.OptionsIsInternal))
	builder.WriteString(fmt.Sprintf("\tOptionsRefreshTokenValidityMetric: %v\n", t.OptionsRefreshTokenValidityMetric))
	builder.WriteString(fmt.Sprintf("\tPinLength: %v\n", t.PinLength))
	builder.WriteString(fmt.Sprintf("\tRefreshTokenValidityPeriod: %v\n", t.RefreshTokenValidityPeriod))
	builder.WriteString(fmt.Sprintf("\tStartUrl: %v\n", t.StartUrl))
	builder.WriteString(fmt.Sprintf("\tSystemModstamp: %v\n", t.SystemModstamp))

	return builder.String()
}

type ConnectedApplicationQueryResponse struct {
	BaseQuery
	Records []ConnectedApplication `json:"Records" force:"records"`
}
