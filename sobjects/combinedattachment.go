// This file was generated for SObject CombinedAttachment, API Version v43.0 at 2018-07-30 03:47:18.073651053 -0400 EDT m=+4.416423208

package sobjects

import (
	"fmt"
	"strings"
)

type CombinedAttachment struct {
	BaseSObject
	ContentSize            int    `force:",omitempty"`
	ContentUrl             string `force:",omitempty"`
	CreatedById            string `force:",omitempty"`
	CreatedDate            string `force:",omitempty"`
	ExternalDataSourceName string `force:",omitempty"`
	ExternalDataSourceType string `force:",omitempty"`
	FileExtension          string `force:",omitempty"`
	FileType               string `force:",omitempty"`
	Id                     string `force:",omitempty"`
	IsDeleted              bool   `force:",omitempty"`
	LastModifiedById       string `force:",omitempty"`
	LastModifiedDate       string `force:",omitempty"`
	ParentId               string `force:",omitempty"`
	RecordType             string `force:",omitempty"`
	SharingOption          string `force:",omitempty"`
	Title                  string `force:",omitempty"`
}

func (t *CombinedAttachment) ApiName() string {
	return "CombinedAttachment"
}

func (t *CombinedAttachment) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("CombinedAttachment #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tContentSize: %v\n", t.ContentSize))
	builder.WriteString(fmt.Sprintf("\tContentUrl: %v\n", t.ContentUrl))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tExternalDataSourceName: %v\n", t.ExternalDataSourceName))
	builder.WriteString(fmt.Sprintf("\tExternalDataSourceType: %v\n", t.ExternalDataSourceType))
	builder.WriteString(fmt.Sprintf("\tFileExtension: %v\n", t.FileExtension))
	builder.WriteString(fmt.Sprintf("\tFileType: %v\n", t.FileType))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tIsDeleted: %v\n", t.IsDeleted))
	builder.WriteString(fmt.Sprintf("\tLastModifiedById: %v\n", t.LastModifiedById))
	builder.WriteString(fmt.Sprintf("\tLastModifiedDate: %v\n", t.LastModifiedDate))
	builder.WriteString(fmt.Sprintf("\tParentId: %v\n", t.ParentId))
	builder.WriteString(fmt.Sprintf("\tRecordType: %v\n", t.RecordType))
	builder.WriteString(fmt.Sprintf("\tSharingOption: %v\n", t.SharingOption))
	builder.WriteString(fmt.Sprintf("\tTitle: %v\n", t.Title))

	return builder.String()
}

type CombinedAttachmentQueryResponse struct {
	BaseQuery
	Records []CombinedAttachment `json:"Records" force:"records"`
}
