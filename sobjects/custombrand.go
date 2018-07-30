// This file was generated for SObject CustomBrand, API Version v43.0 at 2018-07-30 03:47:19.253513334 -0400 EDT m=+5.596329762

package sobjects

import (
	"fmt"
	"strings"
)

type CustomBrand struct {
	BaseSObject
	CreatedById      string `force:",omitempty"`
	CreatedDate      string `force:",omitempty"`
	Id               string `force:",omitempty"`
	LastModifiedById string `force:",omitempty"`
	LastModifiedDate string `force:",omitempty"`
	ParentId         string `force:",omitempty"`
}

func (t *CustomBrand) ApiName() string {
	return "CustomBrand"
}

func (t *CustomBrand) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("CustomBrand #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tLastModifiedById: %v\n", t.LastModifiedById))
	builder.WriteString(fmt.Sprintf("\tLastModifiedDate: %v\n", t.LastModifiedDate))
	builder.WriteString(fmt.Sprintf("\tParentId: %v\n", t.ParentId))

	return builder.String()
}

type CustomBrandQueryResponse struct {
	BaseQuery
	Records []CustomBrand `json:"Records" force:"records"`
}
