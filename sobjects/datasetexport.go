// This file was generated for SObject DatasetExport, API Version v43.0 at 2018-07-30 03:47:34.877823567 -0400 EDT m=+21.221226283

package sobjects

import (
	"fmt"
	"strings"
)

type DatasetExport struct {
	BaseSObject
	CompressedMetadataLength int    `force:",omitempty"`
	CreatedById              string `force:",omitempty"`
	CreatedDate              string `force:",omitempty"`
	Id                       string `force:",omitempty"`
	IsDeleted                bool   `force:",omitempty"`
	LastModifiedById         string `force:",omitempty"`
	LastModifiedDate         string `force:",omitempty"`
	Metadata                 string `force:",omitempty"`
	MetadataLength           int    `force:",omitempty"`
	Owner                    string `force:",omitempty"`
	PublisherInfo            string `force:",omitempty"`
	PublisherType            string `force:",omitempty"`
	Status                   string `force:",omitempty"`
	SystemModstamp           string `force:",omitempty"`
}

func (t *DatasetExport) ApiName() string {
	return "DatasetExport"
}

func (t *DatasetExport) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("DatasetExport #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tCompressedMetadataLength: %v\n", t.CompressedMetadataLength))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tIsDeleted: %v\n", t.IsDeleted))
	builder.WriteString(fmt.Sprintf("\tLastModifiedById: %v\n", t.LastModifiedById))
	builder.WriteString(fmt.Sprintf("\tLastModifiedDate: %v\n", t.LastModifiedDate))
	builder.WriteString(fmt.Sprintf("\tMetadata: %v\n", t.Metadata))
	builder.WriteString(fmt.Sprintf("\tMetadataLength: %v\n", t.MetadataLength))
	builder.WriteString(fmt.Sprintf("\tOwner: %v\n", t.Owner))
	builder.WriteString(fmt.Sprintf("\tPublisherInfo: %v\n", t.PublisherInfo))
	builder.WriteString(fmt.Sprintf("\tPublisherType: %v\n", t.PublisherType))
	builder.WriteString(fmt.Sprintf("\tStatus: %v\n", t.Status))
	builder.WriteString(fmt.Sprintf("\tSystemModstamp: %v\n", t.SystemModstamp))

	return builder.String()
}

type DatasetExportQueryResponse struct {
	BaseQuery
	Records []DatasetExport `json:"Records" force:"records"`
}
