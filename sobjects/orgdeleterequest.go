// This file was generated for SObject OrgDeleteRequest, API Version v43.0 at 2018-07-30 03:47:50.716509943 -0400 EDT m=+37.060506989

package sobjects

import (
	"fmt"
	"strings"
)

type OrgDeleteRequest struct {
	BaseSObject
	CreatedById      string `force:",omitempty"`
	CreatedDate      string `force:",omitempty"`
	Id               string `force:",omitempty"`
	IsDeleted        bool   `force:",omitempty"`
	LastModifiedById string `force:",omitempty"`
	LastModifiedDate string `force:",omitempty"`
	Name             string `force:",omitempty"`
	OwnerId          string `force:",omitempty"`
	RequestType      string `force:",omitempty"`
	SystemModstamp   string `force:",omitempty"`
}

func (t *OrgDeleteRequest) ApiName() string {
	return "OrgDeleteRequest"
}

func (t *OrgDeleteRequest) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("OrgDeleteRequest #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tIsDeleted: %v\n", t.IsDeleted))
	builder.WriteString(fmt.Sprintf("\tLastModifiedById: %v\n", t.LastModifiedById))
	builder.WriteString(fmt.Sprintf("\tLastModifiedDate: %v\n", t.LastModifiedDate))
	builder.WriteString(fmt.Sprintf("\tName: %v\n", t.Name))
	builder.WriteString(fmt.Sprintf("\tOwnerId: %v\n", t.OwnerId))
	builder.WriteString(fmt.Sprintf("\tRequestType: %v\n", t.RequestType))
	builder.WriteString(fmt.Sprintf("\tSystemModstamp: %v\n", t.SystemModstamp))

	return builder.String()
}

type OrgDeleteRequestQueryResponse struct {
	BaseQuery
	Records []OrgDeleteRequest `json:"Records" force:"records"`
}
