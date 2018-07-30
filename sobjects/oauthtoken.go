// This file was generated for SObject OauthToken, API Version v43.0 at 2018-07-30 03:47:58.10228324 -0400 EDT m=+44.446557429

package sobjects

import (
	"fmt"
	"strings"
)

type OauthToken struct {
	BaseSObject
	AccessToken   string `force:",omitempty"`
	AppMenuItemId string `force:",omitempty"`
	AppName       string `force:",omitempty"`
	CreatedDate   string `force:",omitempty"`
	DeleteToken   string `force:",omitempty"`
	Id            string `force:",omitempty"`
	LastUsedDate  string `force:",omitempty"`
	RequestToken  string `force:",omitempty"`
	UseCount      int    `force:",omitempty"`
	UserId        string `force:",omitempty"`
}

func (t *OauthToken) ApiName() string {
	return "OauthToken"
}

func (t *OauthToken) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("OauthToken #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tAccessToken: %v\n", t.AccessToken))
	builder.WriteString(fmt.Sprintf("\tAppMenuItemId: %v\n", t.AppMenuItemId))
	builder.WriteString(fmt.Sprintf("\tAppName: %v\n", t.AppName))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tDeleteToken: %v\n", t.DeleteToken))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tLastUsedDate: %v\n", t.LastUsedDate))
	builder.WriteString(fmt.Sprintf("\tRequestToken: %v\n", t.RequestToken))
	builder.WriteString(fmt.Sprintf("\tUseCount: %v\n", t.UseCount))
	builder.WriteString(fmt.Sprintf("\tUserId: %v\n", t.UserId))

	return builder.String()
}

type OauthTokenQueryResponse struct {
	BaseQuery
	Records []OauthToken `json:"Records" force:"records"`
}
