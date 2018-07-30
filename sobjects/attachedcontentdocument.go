// This file was generated for SObject AttachedContentDocument, API Version v43.0 at 2018-07-30 03:47:17.075439499 -0400 EDT m=+3.418174197

package sobjects

import (
	"fmt"
	"strings"
)

type AttachedContentDocument struct {
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
	LinkedEntityId         string `force:",omitempty"`
	SharingOption          string `force:",omitempty"`
	Title                  string `force:",omitempty"`
}

func (t *AttachedContentDocument) ApiName() string {
	return "AttachedContentDocument"
}

func (t *AttachedContentDocument) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("AttachedContentDocument #%s - %s\n", t.Id, t.Name))
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
	builder.WriteString(fmt.Sprintf("\tLinkedEntityId: %v\n", t.LinkedEntityId))
	builder.WriteString(fmt.Sprintf("\tSharingOption: %v\n", t.SharingOption))
	builder.WriteString(fmt.Sprintf("\tTitle: %v\n", t.Title))

	return builder.String()
}

type AttachedContentDocumentQueryResponse struct {
	BaseQuery
	Records []AttachedContentDocument `json:"Records" force:"records"`
}
