// This file was generated for SObject AuthSession, API Version v43.0 at 2018-07-30 03:47:19.012514934 -0400 EDT m=+5.355322319

package sobjects

import (
	"fmt"
	"strings"
)

type AuthSession struct {
	BaseSObject
	CreatedDate          string `force:",omitempty"`
	Id                   string `force:",omitempty"`
	IsCurrent            bool   `force:",omitempty"`
	LastModifiedDate     string `force:",omitempty"`
	LoginGeoId           string `force:",omitempty"`
	LoginHistoryId       string `force:",omitempty"`
	LoginType            string `force:",omitempty"`
	LogoutUrl            string `force:",omitempty"`
	NumSecondsValid      int    `force:",omitempty"`
	ParentId             string `force:",omitempty"`
	SessionSecurityLevel string `force:",omitempty"`
	SessionType          string `force:",omitempty"`
	SourceIp             string `force:",omitempty"`
	UserType             string `force:",omitempty"`
	UsersId              string `force:",omitempty"`
}

func (t *AuthSession) ApiName() string {
	return "AuthSession"
}

func (t *AuthSession) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("AuthSession #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tIsCurrent: %v\n", t.IsCurrent))
	builder.WriteString(fmt.Sprintf("\tLastModifiedDate: %v\n", t.LastModifiedDate))
	builder.WriteString(fmt.Sprintf("\tLoginGeoId: %v\n", t.LoginGeoId))
	builder.WriteString(fmt.Sprintf("\tLoginHistoryId: %v\n", t.LoginHistoryId))
	builder.WriteString(fmt.Sprintf("\tLoginType: %v\n", t.LoginType))
	builder.WriteString(fmt.Sprintf("\tLogoutUrl: %v\n", t.LogoutUrl))
	builder.WriteString(fmt.Sprintf("\tNumSecondsValid: %v\n", t.NumSecondsValid))
	builder.WriteString(fmt.Sprintf("\tParentId: %v\n", t.ParentId))
	builder.WriteString(fmt.Sprintf("\tSessionSecurityLevel: %v\n", t.SessionSecurityLevel))
	builder.WriteString(fmt.Sprintf("\tSessionType: %v\n", t.SessionType))
	builder.WriteString(fmt.Sprintf("\tSourceIp: %v\n", t.SourceIp))
	builder.WriteString(fmt.Sprintf("\tUserType: %v\n", t.UserType))
	builder.WriteString(fmt.Sprintf("\tUsersId: %v\n", t.UsersId))

	return builder.String()
}

type AuthSessionQueryResponse struct {
	BaseQuery
	Records []AuthSession `json:"Records" force:"records"`
}
