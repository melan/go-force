// This file was generated for SObject AssetHistory, API Version v43.0 at 2018-07-30 03:48:07.136075919 -0400 EDT m=+53.480689092

package sobjects

import (
	"fmt"
	"strings"
)

type AssetHistory struct {
	BaseSObject
	AssetId     string `force:",omitempty"`
	CreatedById string `force:",omitempty"`
	CreatedDate string `force:",omitempty"`
	Field       string `force:",omitempty"`
	Id          string `force:",omitempty"`
	IsDeleted   bool   `force:",omitempty"`
	NewValue    string `force:",omitempty"`
	OldValue    string `force:",omitempty"`
}

func (t *AssetHistory) ApiName() string {
	return "AssetHistory"
}

func (t *AssetHistory) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("AssetHistory #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tAssetId: %v\n", t.AssetId))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tField: %v\n", t.Field))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tIsDeleted: %v\n", t.IsDeleted))
	builder.WriteString(fmt.Sprintf("\tNewValue: %v\n", t.NewValue))
	builder.WriteString(fmt.Sprintf("\tOldValue: %v\n", t.OldValue))

	return builder.String()
}

type AssetHistoryQueryResponse struct {
	BaseQuery
	Records []AssetHistory `json:"Records" force:"records"`
}
