// This file was generated for SObject AssetShare, API Version v43.0 at 2018-07-30 03:47:47.233779262 -0400 EDT m=+33.577645622

package sobjects

import (
	"fmt"
	"strings"
)

type AssetShare struct {
	BaseSObject
	AssetAccessLevel string `force:",omitempty"`
	AssetId          string `force:",omitempty"`
	Id               string `force:",omitempty"`
	IsDeleted        bool   `force:",omitempty"`
	LastModifiedById string `force:",omitempty"`
	LastModifiedDate string `force:",omitempty"`
	RowCause         string `force:",omitempty"`
	UserOrGroupId    string `force:",omitempty"`
}

func (t *AssetShare) ApiName() string {
	return "AssetShare"
}

func (t *AssetShare) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("AssetShare #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tAssetAccessLevel: %v\n", t.AssetAccessLevel))
	builder.WriteString(fmt.Sprintf("\tAssetId: %v\n", t.AssetId))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tIsDeleted: %v\n", t.IsDeleted))
	builder.WriteString(fmt.Sprintf("\tLastModifiedById: %v\n", t.LastModifiedById))
	builder.WriteString(fmt.Sprintf("\tLastModifiedDate: %v\n", t.LastModifiedDate))
	builder.WriteString(fmt.Sprintf("\tRowCause: %v\n", t.RowCause))
	builder.WriteString(fmt.Sprintf("\tUserOrGroupId: %v\n", t.UserOrGroupId))

	return builder.String()
}

type AssetShareQueryResponse struct {
	BaseQuery
	Records []AssetShare `json:"Records" force:"records"`
}
