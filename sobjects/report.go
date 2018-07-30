// This file was generated for SObject Report, API Version v43.0 at 2018-07-30 03:47:49.297499486 -0400 EDT m=+35.641443285

package sobjects

import (
	"fmt"
	"strings"
)

type Report struct {
	BaseSObject
	CreatedById        string `force:",omitempty"`
	CreatedDate        string `force:",omitempty"`
	Description        string `force:",omitempty"`
	DeveloperName      string `force:",omitempty"`
	FolderName         string `force:",omitempty"`
	Format             string `force:",omitempty"`
	Id                 string `force:",omitempty"`
	IsDeleted          bool   `force:",omitempty"`
	LastModifiedById   string `force:",omitempty"`
	LastModifiedDate   string `force:",omitempty"`
	LastReferencedDate string `force:",omitempty"`
	LastRunDate        string `force:",omitempty"`
	LastViewedDate     string `force:",omitempty"`
	Name               string `force:",omitempty"`
	NamespacePrefix    string `force:",omitempty"`
	OwnerId            string `force:",omitempty"`
	SystemModstamp     string `force:",omitempty"`
}

func (t *Report) ApiName() string {
	return "Report"
}

func (t *Report) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("Report #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tDescription: %v\n", t.Description))
	builder.WriteString(fmt.Sprintf("\tDeveloperName: %v\n", t.DeveloperName))
	builder.WriteString(fmt.Sprintf("\tFolderName: %v\n", t.FolderName))
	builder.WriteString(fmt.Sprintf("\tFormat: %v\n", t.Format))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tIsDeleted: %v\n", t.IsDeleted))
	builder.WriteString(fmt.Sprintf("\tLastModifiedById: %v\n", t.LastModifiedById))
	builder.WriteString(fmt.Sprintf("\tLastModifiedDate: %v\n", t.LastModifiedDate))
	builder.WriteString(fmt.Sprintf("\tLastReferencedDate: %v\n", t.LastReferencedDate))
	builder.WriteString(fmt.Sprintf("\tLastRunDate: %v\n", t.LastRunDate))
	builder.WriteString(fmt.Sprintf("\tLastViewedDate: %v\n", t.LastViewedDate))
	builder.WriteString(fmt.Sprintf("\tName: %v\n", t.Name))
	builder.WriteString(fmt.Sprintf("\tNamespacePrefix: %v\n", t.NamespacePrefix))
	builder.WriteString(fmt.Sprintf("\tOwnerId: %v\n", t.OwnerId))
	builder.WriteString(fmt.Sprintf("\tSystemModstamp: %v\n", t.SystemModstamp))

	return builder.String()
}

type ReportQueryResponse struct {
	BaseQuery
	Records []Report `json:"Records" force:"records"`
}
