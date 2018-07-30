// This file was generated for SObject SecurityCustomBaseline, API Version v43.0 at 2018-07-30 03:47:48.576754648 -0400 EDT m=+34.920671402

package sobjects

import (
	"fmt"
	"strings"
)

type SecurityCustomBaseline struct {
	BaseSObject
	Baseline         string `force:",omitempty"`
	CreatedById      string `force:",omitempty"`
	CreatedDate      string `force:",omitempty"`
	DeveloperName    string `force:",omitempty"`
	Id               string `force:",omitempty"`
	IsDefault        bool   `force:",omitempty"`
	IsDeleted        bool   `force:",omitempty"`
	Language         string `force:",omitempty"`
	LastModifiedById string `force:",omitempty"`
	LastModifiedDate string `force:",omitempty"`
	MasterLabel      string `force:",omitempty"`
	NamespacePrefix  string `force:",omitempty"`
	SystemModstamp   string `force:",omitempty"`
}

func (t *SecurityCustomBaseline) ApiName() string {
	return "SecurityCustomBaseline"
}

func (t *SecurityCustomBaseline) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("SecurityCustomBaseline #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tBaseline: %v\n", t.Baseline))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tDeveloperName: %v\n", t.DeveloperName))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tIsDefault: %v\n", t.IsDefault))
	builder.WriteString(fmt.Sprintf("\tIsDeleted: %v\n", t.IsDeleted))
	builder.WriteString(fmt.Sprintf("\tLanguage: %v\n", t.Language))
	builder.WriteString(fmt.Sprintf("\tLastModifiedById: %v\n", t.LastModifiedById))
	builder.WriteString(fmt.Sprintf("\tLastModifiedDate: %v\n", t.LastModifiedDate))
	builder.WriteString(fmt.Sprintf("\tMasterLabel: %v\n", t.MasterLabel))
	builder.WriteString(fmt.Sprintf("\tNamespacePrefix: %v\n", t.NamespacePrefix))
	builder.WriteString(fmt.Sprintf("\tSystemModstamp: %v\n", t.SystemModstamp))

	return builder.String()
}

type SecurityCustomBaselineQueryResponse struct {
	BaseQuery
	Records []SecurityCustomBaseline `json:"Records" force:"records"`
}
