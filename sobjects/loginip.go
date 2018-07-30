// This file was generated for SObject LoginIp, API Version v43.0 at 2018-07-30 03:48:03.148258401 -0400 EDT m=+49.492721935

package sobjects

import (
	"fmt"
	"strings"
)

type LoginIp struct {
	BaseSObject
	ChallengeMethod   string `force:",omitempty"`
	ChallengeSentDate string `force:",omitempty"`
	CreatedDate       string `force:",omitempty"`
	Id                string `force:",omitempty"`
	IsAuthenticated   bool   `force:",omitempty"`
	SourceIp          string `force:",omitempty"`
	UsersId           string `force:",omitempty"`
}

func (t *LoginIp) ApiName() string {
	return "LoginIp"
}

func (t *LoginIp) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("LoginIp #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tChallengeMethod: %v\n", t.ChallengeMethod))
	builder.WriteString(fmt.Sprintf("\tChallengeSentDate: %v\n", t.ChallengeSentDate))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tIsAuthenticated: %v\n", t.IsAuthenticated))
	builder.WriteString(fmt.Sprintf("\tSourceIp: %v\n", t.SourceIp))
	builder.WriteString(fmt.Sprintf("\tUsersId: %v\n", t.UsersId))

	return builder.String()
}

type LoginIpQueryResponse struct {
	BaseQuery
	Records []LoginIp `json:"Records" force:"records"`
}
