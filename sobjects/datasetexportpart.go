// This file was generated for SObject DatasetExportPart, API Version v43.0 at 2018-07-30 03:47:37.142947133 -0400 EDT m=+23.486434845

package sobjects

import (
	"fmt"
	"strings"
)

type DatasetExportPart struct {
	BaseSObject
	CompressedDataFileLength int    `force:",omitempty"`
	CreatedById              string `force:",omitempty"`
	CreatedDate              string `force:",omitempty"`
	DataFile                 string `force:",omitempty"`
	DataFileLength           int    `force:",omitempty"`
	DatasetExportId          string `force:",omitempty"`
	Id                       string `force:",omitempty"`
	IsDeleted                bool   `force:",omitempty"`
	LastModifiedById         string `force:",omitempty"`
	LastModifiedDate         string `force:",omitempty"`
	Owner                    string `force:",omitempty"`
	PartNumber               int    `force:",omitempty"`
	SystemModstamp           string `force:",omitempty"`
}

func (t *DatasetExportPart) ApiName() string {
	return "DatasetExportPart"
}

func (t *DatasetExportPart) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("DatasetExportPart #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tCompressedDataFileLength: %v\n", t.CompressedDataFileLength))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tDataFile: %v\n", t.DataFile))
	builder.WriteString(fmt.Sprintf("\tDataFileLength: %v\n", t.DataFileLength))
	builder.WriteString(fmt.Sprintf("\tDatasetExportId: %v\n", t.DatasetExportId))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tIsDeleted: %v\n", t.IsDeleted))
	builder.WriteString(fmt.Sprintf("\tLastModifiedById: %v\n", t.LastModifiedById))
	builder.WriteString(fmt.Sprintf("\tLastModifiedDate: %v\n", t.LastModifiedDate))
	builder.WriteString(fmt.Sprintf("\tOwner: %v\n", t.Owner))
	builder.WriteString(fmt.Sprintf("\tPartNumber: %v\n", t.PartNumber))
	builder.WriteString(fmt.Sprintf("\tSystemModstamp: %v\n", t.SystemModstamp))

	return builder.String()
}

type DatasetExportPartQueryResponse struct {
	BaseQuery
	Records []DatasetExportPart `json:"Records" force:"records"`
}
