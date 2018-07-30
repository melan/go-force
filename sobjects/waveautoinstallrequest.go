// This file was generated for SObject WaveAutoInstallRequest, API Version v43.0 at 2018-07-30 03:48:04.559104543 -0400 EDT m=+50.903621018

package sobjects

import (
	"fmt"
	"strings"
)

type WaveAutoInstallRequest struct {
	BaseSObject
	Configuration    string `force:",omitempty"`
	CreatedById      string `force:",omitempty"`
	CreatedDate      string `force:",omitempty"`
	FailedReason     string `force:",omitempty"`
	FolderId         string `force:",omitempty"`
	Id               string `force:",omitempty"`
	IsDeleted        bool   `force:",omitempty"`
	LastModifiedById string `force:",omitempty"`
	LastModifiedDate string `force:",omitempty"`
	Name             string `force:",omitempty"`
	RequestLog       string `force:",omitempty"`
	RequestStatus    string `force:",omitempty"`
	RequestType      string `force:",omitempty"`
	SystemModstamp   string `force:",omitempty"`
	TemplateApiName  string `force:",omitempty"`
	TemplateVersion  string `force:",omitempty"`
}

func (t *WaveAutoInstallRequest) ApiName() string {
	return "WaveAutoInstallRequest"
}

func (t *WaveAutoInstallRequest) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("WaveAutoInstallRequest #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tConfiguration: %v\n", t.Configuration))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tFailedReason: %v\n", t.FailedReason))
	builder.WriteString(fmt.Sprintf("\tFolderId: %v\n", t.FolderId))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tIsDeleted: %v\n", t.IsDeleted))
	builder.WriteString(fmt.Sprintf("\tLastModifiedById: %v\n", t.LastModifiedById))
	builder.WriteString(fmt.Sprintf("\tLastModifiedDate: %v\n", t.LastModifiedDate))
	builder.WriteString(fmt.Sprintf("\tName: %v\n", t.Name))
	builder.WriteString(fmt.Sprintf("\tRequestLog: %v\n", t.RequestLog))
	builder.WriteString(fmt.Sprintf("\tRequestStatus: %v\n", t.RequestStatus))
	builder.WriteString(fmt.Sprintf("\tRequestType: %v\n", t.RequestType))
	builder.WriteString(fmt.Sprintf("\tSystemModstamp: %v\n", t.SystemModstamp))
	builder.WriteString(fmt.Sprintf("\tTemplateApiName: %v\n", t.TemplateApiName))
	builder.WriteString(fmt.Sprintf("\tTemplateVersion: %v\n", t.TemplateVersion))

	return builder.String()
}

type WaveAutoInstallRequestQueryResponse struct {
	BaseQuery
	Records []WaveAutoInstallRequest `json:"Records" force:"records"`
}
