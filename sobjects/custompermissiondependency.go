// This file was generated for SObject CustomPermissionDependency, API Version v43.0 at 2018-07-30 03:47:24.059760542 -0400 EDT m=+10.402757321

package sobjects

import (
	"fmt"
	"strings"
)

type CustomPermissionDependency struct {
	BaseSObject
	CreatedById                string `force:",omitempty"`
	CreatedDate                string `force:",omitempty"`
	CustomPermissionId         string `force:",omitempty"`
	Id                         string `force:",omitempty"`
	IsDeleted                  bool   `force:",omitempty"`
	LastModifiedById           string `force:",omitempty"`
	LastModifiedDate           string `force:",omitempty"`
	RequiredCustomPermissionId string `force:",omitempty"`
	SystemModstamp             string `force:",omitempty"`
}

func (t *CustomPermissionDependency) ApiName() string {
	return "CustomPermissionDependency"
}

func (t *CustomPermissionDependency) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("CustomPermissionDependency #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tCustomPermissionId: %v\n", t.CustomPermissionId))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tIsDeleted: %v\n", t.IsDeleted))
	builder.WriteString(fmt.Sprintf("\tLastModifiedById: %v\n", t.LastModifiedById))
	builder.WriteString(fmt.Sprintf("\tLastModifiedDate: %v\n", t.LastModifiedDate))
	builder.WriteString(fmt.Sprintf("\tRequiredCustomPermissionId: %v\n", t.RequiredCustomPermissionId))
	builder.WriteString(fmt.Sprintf("\tSystemModstamp: %v\n", t.SystemModstamp))

	return builder.String()
}

type CustomPermissionDependencyQueryResponse struct {
	BaseQuery
	Records []CustomPermissionDependency `json:"Records" force:"records"`
}
