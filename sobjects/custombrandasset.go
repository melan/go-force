// This file was generated for SObject CustomBrandAsset, API Version v43.0 at 2018-07-30 03:47:59.182511733 -0400 EDT m=+45.526826456

package sobjects

import (
	"fmt"
	"strings"
)

type CustomBrandAsset struct {
	BaseSObject
	AssetCategory    string `force:",omitempty"`
	AssetSourceId    string `force:",omitempty"`
	CreatedById      string `force:",omitempty"`
	CreatedDate      string `force:",omitempty"`
	CustomBrandId    string `force:",omitempty"`
	Id               string `force:",omitempty"`
	LastModifiedById string `force:",omitempty"`
	LastModifiedDate string `force:",omitempty"`
	TextAsset        string `force:",omitempty"`
}

func (t *CustomBrandAsset) ApiName() string {
	return "CustomBrandAsset"
}

func (t *CustomBrandAsset) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("CustomBrandAsset #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tAssetCategory: %v\n", t.AssetCategory))
	builder.WriteString(fmt.Sprintf("\tAssetSourceId: %v\n", t.AssetSourceId))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tCustomBrandId: %v\n", t.CustomBrandId))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tLastModifiedById: %v\n", t.LastModifiedById))
	builder.WriteString(fmt.Sprintf("\tLastModifiedDate: %v\n", t.LastModifiedDate))
	builder.WriteString(fmt.Sprintf("\tTextAsset: %v\n", t.TextAsset))

	return builder.String()
}

type CustomBrandAssetQueryResponse struct {
	BaseQuery
	Records []CustomBrandAsset `json:"Records" force:"records"`
}
