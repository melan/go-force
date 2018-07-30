// This file was generated for SObject Vote, API Version v43.0 at 2018-07-30 03:48:04.252611828 -0400 EDT m=+50.597116801

package sobjects

import (
	"fmt"
	"strings"
)

type Vote struct {
	BaseSObject
	CreatedById      string `force:",omitempty"`
	CreatedDate      string `force:",omitempty"`
	Id               string `force:",omitempty"`
	IsDeleted        bool   `force:",omitempty"`
	LastModifiedById string `force:",omitempty"`
	LastModifiedDate string `force:",omitempty"`
	ParentId         string `force:",omitempty"`
	SystemModstamp   string `force:",omitempty"`
	Type             string `force:",omitempty"`
}

func (t *Vote) ApiName() string {
	return "Vote"
}

func (t *Vote) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("Vote #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tIsDeleted: %v\n", t.IsDeleted))
	builder.WriteString(fmt.Sprintf("\tLastModifiedById: %v\n", t.LastModifiedById))
	builder.WriteString(fmt.Sprintf("\tLastModifiedDate: %v\n", t.LastModifiedDate))
	builder.WriteString(fmt.Sprintf("\tParentId: %v\n", t.ParentId))
	builder.WriteString(fmt.Sprintf("\tSystemModstamp: %v\n", t.SystemModstamp))
	builder.WriteString(fmt.Sprintf("\tType: %v\n", t.Type))

	return builder.String()
}

type VoteQueryResponse struct {
	BaseQuery
	Records []Vote `json:"Records" force:"records"`
}
