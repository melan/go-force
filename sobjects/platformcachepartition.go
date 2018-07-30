// This file was generated for SObject PlatformCachePartition, API Version v43.0 at 2018-07-30 03:48:06.082240498 -0400 EDT m=+52.426814127

package sobjects

import (
	"fmt"
	"strings"
)

type PlatformCachePartition struct {
	BaseSObject
	CreatedById        string `force:",omitempty"`
	CreatedDate        string `force:",omitempty"`
	Description        string `force:",omitempty"`
	DeveloperName      string `force:",omitempty"`
	Id                 string `force:",omitempty"`
	IsDefaultPartition bool   `force:",omitempty"`
	IsDeleted          bool   `force:",omitempty"`
	Language           string `force:",omitempty"`
	LastModifiedById   string `force:",omitempty"`
	LastModifiedDate   string `force:",omitempty"`
	MasterLabel        string `force:",omitempty"`
	NamespacePrefix    string `force:",omitempty"`
	SystemModstamp     string `force:",omitempty"`
}

func (t *PlatformCachePartition) ApiName() string {
	return "PlatformCachePartition"
}

func (t *PlatformCachePartition) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("PlatformCachePartition #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tDescription: %v\n", t.Description))
	builder.WriteString(fmt.Sprintf("\tDeveloperName: %v\n", t.DeveloperName))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tIsDefaultPartition: %v\n", t.IsDefaultPartition))
	builder.WriteString(fmt.Sprintf("\tIsDeleted: %v\n", t.IsDeleted))
	builder.WriteString(fmt.Sprintf("\tLanguage: %v\n", t.Language))
	builder.WriteString(fmt.Sprintf("\tLastModifiedById: %v\n", t.LastModifiedById))
	builder.WriteString(fmt.Sprintf("\tLastModifiedDate: %v\n", t.LastModifiedDate))
	builder.WriteString(fmt.Sprintf("\tMasterLabel: %v\n", t.MasterLabel))
	builder.WriteString(fmt.Sprintf("\tNamespacePrefix: %v\n", t.NamespacePrefix))
	builder.WriteString(fmt.Sprintf("\tSystemModstamp: %v\n", t.SystemModstamp))

	return builder.String()
}

type PlatformCachePartitionQueryResponse struct {
	BaseQuery
	Records []PlatformCachePartition `json:"Records" force:"records"`
}
