// This file was generated for SObject ExternalDataSource, API Version v43.0 at 2018-07-30 03:47:45.141370149 -0400 EDT m=+31.485157994

package sobjects

import (
	"fmt"
	"strings"
)

type ExternalDataSource struct {
	BaseSObject
	AuthProviderId      string `force:",omitempty"`
	CreatedById         string `force:",omitempty"`
	CreatedDate         string `force:",omitempty"`
	CustomConfiguration string `force:",omitempty"`
	DeveloperName       string `force:",omitempty"`
	Endpoint            string `force:",omitempty"`
	Id                  string `force:",omitempty"`
	IsDeleted           bool   `force:",omitempty"`
	IsWritable          bool   `force:",omitempty"`
	Language            string `force:",omitempty"`
	LargeIconId         string `force:",omitempty"`
	LastModifiedById    string `force:",omitempty"`
	LastModifiedDate    string `force:",omitempty"`
	MasterLabel         string `force:",omitempty"`
	NamespacePrefix     string `force:",omitempty"`
	PrincipalType       string `force:",omitempty"`
	Protocol            string `force:",omitempty"`
	Repository          string `force:",omitempty"`
	SmallIconId         string `force:",omitempty"`
	SystemModstamp      string `force:",omitempty"`
	Type                string `force:",omitempty"`
}

func (t *ExternalDataSource) ApiName() string {
	return "ExternalDataSource"
}

func (t *ExternalDataSource) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("ExternalDataSource #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tAuthProviderId: %v\n", t.AuthProviderId))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tCustomConfiguration: %v\n", t.CustomConfiguration))
	builder.WriteString(fmt.Sprintf("\tDeveloperName: %v\n", t.DeveloperName))
	builder.WriteString(fmt.Sprintf("\tEndpoint: %v\n", t.Endpoint))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tIsDeleted: %v\n", t.IsDeleted))
	builder.WriteString(fmt.Sprintf("\tIsWritable: %v\n", t.IsWritable))
	builder.WriteString(fmt.Sprintf("\tLanguage: %v\n", t.Language))
	builder.WriteString(fmt.Sprintf("\tLargeIconId: %v\n", t.LargeIconId))
	builder.WriteString(fmt.Sprintf("\tLastModifiedById: %v\n", t.LastModifiedById))
	builder.WriteString(fmt.Sprintf("\tLastModifiedDate: %v\n", t.LastModifiedDate))
	builder.WriteString(fmt.Sprintf("\tMasterLabel: %v\n", t.MasterLabel))
	builder.WriteString(fmt.Sprintf("\tNamespacePrefix: %v\n", t.NamespacePrefix))
	builder.WriteString(fmt.Sprintf("\tPrincipalType: %v\n", t.PrincipalType))
	builder.WriteString(fmt.Sprintf("\tProtocol: %v\n", t.Protocol))
	builder.WriteString(fmt.Sprintf("\tRepository: %v\n", t.Repository))
	builder.WriteString(fmt.Sprintf("\tSmallIconId: %v\n", t.SmallIconId))
	builder.WriteString(fmt.Sprintf("\tSystemModstamp: %v\n", t.SystemModstamp))
	builder.WriteString(fmt.Sprintf("\tType: %v\n", t.Type))

	return builder.String()
}

type ExternalDataSourceQueryResponse struct {
	BaseQuery
	Records []ExternalDataSource `json:"Records" force:"records"`
}
