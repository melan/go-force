// This file was generated for SObject DatasetExportEvent, API Version v43.0 at 2018-07-30 03:47:34.312612657 -0400 EDT m=+20.655994163

package sobjects

import (
	"fmt"
	"strings"
)

type DatasetExportEvent struct {
	BaseSObject
	CreatedById        string `force:",omitempty"`
	CreatedDate        string `force:",omitempty"`
	DataflowInstanceId string `force:",omitempty"`
	DatasetExportId    string `force:",omitempty"`
	Message            string `force:",omitempty"`
	Owner              string `force:",omitempty"`
	PublisherInfo      string `force:",omitempty"`
	PublisherType      string `force:",omitempty"`
	ReplayId           string `force:",omitempty"`
	Status             string `force:",omitempty"`
}

func (t *DatasetExportEvent) ApiName() string {
	return "DatasetExportEvent"
}

func (t *DatasetExportEvent) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("DatasetExportEvent #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tDataflowInstanceId: %v\n", t.DataflowInstanceId))
	builder.WriteString(fmt.Sprintf("\tDatasetExportId: %v\n", t.DatasetExportId))
	builder.WriteString(fmt.Sprintf("\tMessage: %v\n", t.Message))
	builder.WriteString(fmt.Sprintf("\tOwner: %v\n", t.Owner))
	builder.WriteString(fmt.Sprintf("\tPublisherInfo: %v\n", t.PublisherInfo))
	builder.WriteString(fmt.Sprintf("\tPublisherType: %v\n", t.PublisherType))
	builder.WriteString(fmt.Sprintf("\tReplayId: %v\n", t.ReplayId))
	builder.WriteString(fmt.Sprintf("\tStatus: %v\n", t.Status))

	return builder.String()
}

type DatasetExportEventQueryResponse struct {
	BaseQuery
	Records []DatasetExportEvent `json:"Records" force:"records"`
}
