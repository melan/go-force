// This file was generated for SObject CategoryData, API Version v43.0 at 2018-07-30 03:48:04.136338384 -0400 EDT m=+50.480838995

package sobjects

import (
	"fmt"
	"strings"
)

type CategoryData struct {
	BaseSObject
	CategoryNodeId   string `force:",omitempty"`
	CreatedById      string `force:",omitempty"`
	CreatedDate      string `force:",omitempty"`
	Id               string `force:",omitempty"`
	IsDeleted        bool   `force:",omitempty"`
	LastModifiedById string `force:",omitempty"`
	LastModifiedDate string `force:",omitempty"`
	RelatedSobjectId string `force:",omitempty"`
	SystemModstamp   string `force:",omitempty"`
}

func (t *CategoryData) ApiName() string {
	return "CategoryData"
}

func (t *CategoryData) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("CategoryData #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tCategoryNodeId: %v\n", t.CategoryNodeId))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tIsDeleted: %v\n", t.IsDeleted))
	builder.WriteString(fmt.Sprintf("\tLastModifiedById: %v\n", t.LastModifiedById))
	builder.WriteString(fmt.Sprintf("\tLastModifiedDate: %v\n", t.LastModifiedDate))
	builder.WriteString(fmt.Sprintf("\tRelatedSobjectId: %v\n", t.RelatedSobjectId))
	builder.WriteString(fmt.Sprintf("\tSystemModstamp: %v\n", t.SystemModstamp))

	return builder.String()
}

type CategoryDataQueryResponse struct {
	BaseQuery
	Records []CategoryData `json:"Records" force:"records"`
}
