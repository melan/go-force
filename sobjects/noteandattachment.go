// This file was generated for SObject NoteAndAttachment, API Version v43.0 at 2018-07-30 03:47:49.659649499 -0400 EDT m=+36.003606887

package sobjects

import (
	"fmt"
	"strings"
)

type NoteAndAttachment struct {
	BaseSObject
	CreatedById      string `force:",omitempty"`
	CreatedDate      string `force:",omitempty"`
	Id               string `force:",omitempty"`
	IsDeleted        bool   `force:",omitempty"`
	IsNote           bool   `force:",omitempty"`
	IsPrivate        bool   `force:",omitempty"`
	LastModifiedById string `force:",omitempty"`
	LastModifiedDate string `force:",omitempty"`
	OwnerId          string `force:",omitempty"`
	ParentId         string `force:",omitempty"`
	SystemModstamp   string `force:",omitempty"`
	Title            string `force:",omitempty"`
}

func (t *NoteAndAttachment) ApiName() string {
	return "NoteAndAttachment"
}

func (t *NoteAndAttachment) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("NoteAndAttachment #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tIsDeleted: %v\n", t.IsDeleted))
	builder.WriteString(fmt.Sprintf("\tIsNote: %v\n", t.IsNote))
	builder.WriteString(fmt.Sprintf("\tIsPrivate: %v\n", t.IsPrivate))
	builder.WriteString(fmt.Sprintf("\tLastModifiedById: %v\n", t.LastModifiedById))
	builder.WriteString(fmt.Sprintf("\tLastModifiedDate: %v\n", t.LastModifiedDate))
	builder.WriteString(fmt.Sprintf("\tOwnerId: %v\n", t.OwnerId))
	builder.WriteString(fmt.Sprintf("\tParentId: %v\n", t.ParentId))
	builder.WriteString(fmt.Sprintf("\tSystemModstamp: %v\n", t.SystemModstamp))
	builder.WriteString(fmt.Sprintf("\tTitle: %v\n", t.Title))

	return builder.String()
}

type NoteAndAttachmentQueryResponse struct {
	BaseQuery
	Records []NoteAndAttachment `json:"Records" force:"records"`
}
