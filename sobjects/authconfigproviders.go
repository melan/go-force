// This file was generated for SObject AuthConfigProviders, API Version v43.0 at 2018-07-30 03:47:43.499331006 -0400 EDT m=+29.843057235

package sobjects

import (
	"fmt"
	"strings"
)

type AuthConfigProviders struct {
	BaseSObject
	AuthConfigId     string `force:",omitempty"`
	AuthProviderId   string `force:",omitempty"`
	CreatedById      string `force:",omitempty"`
	CreatedDate      string `force:",omitempty"`
	Id               string `force:",omitempty"`
	IsDeleted        bool   `force:",omitempty"`
	LastModifiedById string `force:",omitempty"`
	LastModifiedDate string `force:",omitempty"`
	SystemModstamp   string `force:",omitempty"`
}

func (t *AuthConfigProviders) ApiName() string {
	return "AuthConfigProviders"
}

func (t *AuthConfigProviders) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("AuthConfigProviders #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tAuthConfigId: %v\n", t.AuthConfigId))
	builder.WriteString(fmt.Sprintf("\tAuthProviderId: %v\n", t.AuthProviderId))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tIsDeleted: %v\n", t.IsDeleted))
	builder.WriteString(fmt.Sprintf("\tLastModifiedById: %v\n", t.LastModifiedById))
	builder.WriteString(fmt.Sprintf("\tLastModifiedDate: %v\n", t.LastModifiedDate))
	builder.WriteString(fmt.Sprintf("\tSystemModstamp: %v\n", t.SystemModstamp))

	return builder.String()
}

type AuthConfigProvidersQueryResponse struct {
	BaseQuery
	Records []AuthConfigProviders `json:"Records" force:"records"`
}
