// This file was generated for SObject Pricebook2, API Version v43.0 at 2018-07-30 03:47:37.388242171 -0400 EDT m=+23.731739087

package sobjects

import (
	"fmt"
	"strings"
)

type Pricebook2 struct {
	BaseSObject
	CreatedById        string `force:",omitempty"`
	CreatedDate        string `force:",omitempty"`
	Description        string `force:",omitempty"`
	Id                 string `force:",omitempty"`
	IsActive           bool   `force:",omitempty"`
	IsArchived         bool   `force:",omitempty"`
	IsDeleted          bool   `force:",omitempty"`
	IsStandard         bool   `force:",omitempty"`
	LastModifiedById   string `force:",omitempty"`
	LastModifiedDate   string `force:",omitempty"`
	LastReferencedDate string `force:",omitempty"`
	LastViewedDate     string `force:",omitempty"`
	Name               string `force:",omitempty"`
	SystemModstamp     string `force:",omitempty"`
}

func (t *Pricebook2) ApiName() string {
	return "Pricebook2"
}

func (t *Pricebook2) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("Pricebook2 #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tDescription: %v\n", t.Description))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tIsActive: %v\n", t.IsActive))
	builder.WriteString(fmt.Sprintf("\tIsArchived: %v\n", t.IsArchived))
	builder.WriteString(fmt.Sprintf("\tIsDeleted: %v\n", t.IsDeleted))
	builder.WriteString(fmt.Sprintf("\tIsStandard: %v\n", t.IsStandard))
	builder.WriteString(fmt.Sprintf("\tLastModifiedById: %v\n", t.LastModifiedById))
	builder.WriteString(fmt.Sprintf("\tLastModifiedDate: %v\n", t.LastModifiedDate))
	builder.WriteString(fmt.Sprintf("\tLastReferencedDate: %v\n", t.LastReferencedDate))
	builder.WriteString(fmt.Sprintf("\tLastViewedDate: %v\n", t.LastViewedDate))
	builder.WriteString(fmt.Sprintf("\tName: %v\n", t.Name))
	builder.WriteString(fmt.Sprintf("\tSystemModstamp: %v\n", t.SystemModstamp))

	return builder.String()
}

type Pricebook2QueryResponse struct {
	BaseQuery
	Records []Pricebook2 `json:"Records" force:"records"`
}
