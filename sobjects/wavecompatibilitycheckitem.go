// This file was generated for SObject WaveCompatibilityCheckItem, API Version v43.0 at 2018-07-30 03:47:49.418832935 -0400 EDT m=+35.762781287

package sobjects

import (
	"fmt"
	"strings"
)

type WaveCompatibilityCheckItem struct {
	BaseSObject
	CreatedById              string `force:",omitempty"`
	CreatedDate              string `force:",omitempty"`
	Id                       string `force:",omitempty"`
	IsDeleted                bool   `force:",omitempty"`
	LastModifiedById         string `force:",omitempty"`
	LastModifiedDate         string `force:",omitempty"`
	Name                     string `force:",omitempty"`
	Payload                  string `force:",omitempty"`
	SystemModstamp           string `force:",omitempty"`
	TaskName                 string `force:",omitempty"`
	TaskResult               string `force:",omitempty"`
	TemplateApiName          string `force:",omitempty"`
	TemplateVersion          string `force:",omitempty"`
	WaveAutoInstallRequestId string `force:",omitempty"`
}

func (t *WaveCompatibilityCheckItem) ApiName() string {
	return "WaveCompatibilityCheckItem"
}

func (t *WaveCompatibilityCheckItem) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("WaveCompatibilityCheckItem #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tIsDeleted: %v\n", t.IsDeleted))
	builder.WriteString(fmt.Sprintf("\tLastModifiedById: %v\n", t.LastModifiedById))
	builder.WriteString(fmt.Sprintf("\tLastModifiedDate: %v\n", t.LastModifiedDate))
	builder.WriteString(fmt.Sprintf("\tName: %v\n", t.Name))
	builder.WriteString(fmt.Sprintf("\tPayload: %v\n", t.Payload))
	builder.WriteString(fmt.Sprintf("\tSystemModstamp: %v\n", t.SystemModstamp))
	builder.WriteString(fmt.Sprintf("\tTaskName: %v\n", t.TaskName))
	builder.WriteString(fmt.Sprintf("\tTaskResult: %v\n", t.TaskResult))
	builder.WriteString(fmt.Sprintf("\tTemplateApiName: %v\n", t.TemplateApiName))
	builder.WriteString(fmt.Sprintf("\tTemplateVersion: %v\n", t.TemplateVersion))
	builder.WriteString(fmt.Sprintf("\tWaveAutoInstallRequestId: %v\n", t.WaveAutoInstallRequestId))

	return builder.String()
}

type WaveCompatibilityCheckItemQueryResponse struct {
	BaseQuery
	Records []WaveCompatibilityCheckItem `json:"Records" force:"records"`
}
