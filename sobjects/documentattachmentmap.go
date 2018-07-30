// This file was generated for SObject DocumentAttachmentMap, API Version v43.0 at 2018-07-30 03:47:46.868005293 -0400 EDT m=+33.211857928

package sobjects

import (
	"fmt"
	"strings"
)

type DocumentAttachmentMap struct {
	BaseSObject
	CreatedById      string `force:",omitempty"`
	CreatedDate      string `force:",omitempty"`
	DocumentId       string `force:",omitempty"`
	DocumentSequence int    `force:",omitempty"`
	Id               string `force:",omitempty"`
	ParentId         string `force:",omitempty"`
}

func (t *DocumentAttachmentMap) ApiName() string {
	return "DocumentAttachmentMap"
}

func (t *DocumentAttachmentMap) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("DocumentAttachmentMap #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tDocumentId: %v\n", t.DocumentId))
	builder.WriteString(fmt.Sprintf("\tDocumentSequence: %v\n", t.DocumentSequence))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tParentId: %v\n", t.ParentId))

	return builder.String()
}

type DocumentAttachmentMapQueryResponse struct {
	BaseQuery
	Records []DocumentAttachmentMap `json:"Records" force:"records"`
}
