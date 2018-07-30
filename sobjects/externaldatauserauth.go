// This file was generated for SObject ExternalDataUserAuth, API Version v43.0 at 2018-07-30 03:47:31.283675829 -0400 EDT m=+17.626943678

package sobjects

import (
	"fmt"
	"strings"
)

type ExternalDataUserAuth struct {
	BaseSObject
	AuthProviderId       string `force:",omitempty"`
	CreatedById          string `force:",omitempty"`
	CreatedDate          string `force:",omitempty"`
	ExternalDataSourceId string `force:",omitempty"`
	Id                   string `force:",omitempty"`
	IsDeleted            bool   `force:",omitempty"`
	LastModifiedById     string `force:",omitempty"`
	LastModifiedDate     string `force:",omitempty"`
	Password             string `force:",omitempty"`
	Protocol             string `force:",omitempty"`
	SystemModstamp       string `force:",omitempty"`
	UserId               string `force:",omitempty"`
	Username             string `force:",omitempty"`
}

func (t *ExternalDataUserAuth) ApiName() string {
	return "ExternalDataUserAuth"
}

func (t *ExternalDataUserAuth) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("ExternalDataUserAuth #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tAuthProviderId: %v\n", t.AuthProviderId))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tExternalDataSourceId: %v\n", t.ExternalDataSourceId))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tIsDeleted: %v\n", t.IsDeleted))
	builder.WriteString(fmt.Sprintf("\tLastModifiedById: %v\n", t.LastModifiedById))
	builder.WriteString(fmt.Sprintf("\tLastModifiedDate: %v\n", t.LastModifiedDate))
	builder.WriteString(fmt.Sprintf("\tPassword: %v\n", t.Password))
	builder.WriteString(fmt.Sprintf("\tProtocol: %v\n", t.Protocol))
	builder.WriteString(fmt.Sprintf("\tSystemModstamp: %v\n", t.SystemModstamp))
	builder.WriteString(fmt.Sprintf("\tUserId: %v\n", t.UserId))
	builder.WriteString(fmt.Sprintf("\tUsername: %v\n", t.Username))

	return builder.String()
}

type ExternalDataUserAuthQueryResponse struct {
	BaseQuery
	Records []ExternalDataUserAuth `json:"Records" force:"records"`
}
