// This file was generated for SObject OwnedContentDocument, API Version v43.0 at 2018-07-30 03:48:08.826337128 -0400 EDT m=+55.171013726

package sobjects

import (
	"fmt"
	"strings"
)

type OwnedContentDocument struct {
	BaseSObject
	ContentDocumentId      string `force:",omitempty"`
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
	OwnerId                string `force:",omitempty"`
	Title                  string `force:",omitempty"`
}

func (t *OwnedContentDocument) ApiName() string {
	return "OwnedContentDocument"
}

func (t *OwnedContentDocument) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("OwnedContentDocument #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tContentDocumentId: %v\n", t.ContentDocumentId))
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
	builder.WriteString(fmt.Sprintf("\tOwnerId: %v\n", t.OwnerId))
	builder.WriteString(fmt.Sprintf("\tTitle: %v\n", t.Title))

	return builder.String()
}

type OwnedContentDocumentQueryResponse struct {
	BaseQuery
	Records []OwnedContentDocument `json:"Records" force:"records"`
}
