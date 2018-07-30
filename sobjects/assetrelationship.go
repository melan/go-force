// This file was generated for SObject AssetRelationship, API Version v43.0 at 2018-07-30 03:47:34.465042009 -0400 EDT m=+20.808429235

package sobjects

import (
	"fmt"
	"strings"
)

type AssetRelationship struct {
	BaseSObject
	AssetId                 string `force:",omitempty"`
	AssetRelationshipNumber string `force:",omitempty"`
	CreatedById             string `force:",omitempty"`
	CreatedDate             string `force:",omitempty"`
	FromDate                string `force:",omitempty"`
	Id                      string `force:",omitempty"`
	IsDeleted               bool   `force:",omitempty"`
	LastModifiedById        string `force:",omitempty"`
	LastModifiedDate        string `force:",omitempty"`
	LastReferencedDate      string `force:",omitempty"`
	LastViewedDate          string `force:",omitempty"`
	RelatedAssetId          string `force:",omitempty"`
	RelationshipType        string `force:",omitempty"`
	SystemModstamp          string `force:",omitempty"`
	ToDate                  string `force:",omitempty"`
}

func (t *AssetRelationship) ApiName() string {
	return "AssetRelationship"
}

func (t *AssetRelationship) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("AssetRelationship #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tAssetId: %v\n", t.AssetId))
	builder.WriteString(fmt.Sprintf("\tAssetRelationshipNumber: %v\n", t.AssetRelationshipNumber))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tFromDate: %v\n", t.FromDate))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tIsDeleted: %v\n", t.IsDeleted))
	builder.WriteString(fmt.Sprintf("\tLastModifiedById: %v\n", t.LastModifiedById))
	builder.WriteString(fmt.Sprintf("\tLastModifiedDate: %v\n", t.LastModifiedDate))
	builder.WriteString(fmt.Sprintf("\tLastReferencedDate: %v\n", t.LastReferencedDate))
	builder.WriteString(fmt.Sprintf("\tLastViewedDate: %v\n", t.LastViewedDate))
	builder.WriteString(fmt.Sprintf("\tRelatedAssetId: %v\n", t.RelatedAssetId))
	builder.WriteString(fmt.Sprintf("\tRelationshipType: %v\n", t.RelationshipType))
	builder.WriteString(fmt.Sprintf("\tSystemModstamp: %v\n", t.SystemModstamp))
	builder.WriteString(fmt.Sprintf("\tToDate: %v\n", t.ToDate))

	return builder.String()
}

type AssetRelationshipQueryResponse struct {
	BaseQuery
	Records []AssetRelationship `json:"Records" force:"records"`
}
