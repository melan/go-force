// This file was generated for SObject DatacloudOwnedEntity, API Version v43.0 at 2018-07-30 03:47:26.982223499 -0400 EDT m=+13.325329941

package sobjects

import (
	"fmt"
	"strings"
)

type DatacloudOwnedEntity struct {
	BaseSObject
	CreatedById         string `force:",omitempty"`
	CreatedDate         string `force:",omitempty"`
	DataDotComKey       string `force:",omitempty"`
	DatacloudEntityType string `force:",omitempty"`
	Id                  string `force:",omitempty"`
	IsDeleted           bool   `force:",omitempty"`
	LastModifiedById    string `force:",omitempty"`
	LastModifiedDate    string `force:",omitempty"`
	Name                string `force:",omitempty"`
	PurchaseType        string `force:",omitempty"`
	PurchaseUsageId     string `force:",omitempty"`
	SystemModstamp      string `force:",omitempty"`
	UserId              string `force:",omitempty"`
}

func (t *DatacloudOwnedEntity) ApiName() string {
	return "DatacloudOwnedEntity"
}

func (t *DatacloudOwnedEntity) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("DatacloudOwnedEntity #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tDataDotComKey: %v\n", t.DataDotComKey))
	builder.WriteString(fmt.Sprintf("\tDatacloudEntityType: %v\n", t.DatacloudEntityType))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tIsDeleted: %v\n", t.IsDeleted))
	builder.WriteString(fmt.Sprintf("\tLastModifiedById: %v\n", t.LastModifiedById))
	builder.WriteString(fmt.Sprintf("\tLastModifiedDate: %v\n", t.LastModifiedDate))
	builder.WriteString(fmt.Sprintf("\tName: %v\n", t.Name))
	builder.WriteString(fmt.Sprintf("\tPurchaseType: %v\n", t.PurchaseType))
	builder.WriteString(fmt.Sprintf("\tPurchaseUsageId: %v\n", t.PurchaseUsageId))
	builder.WriteString(fmt.Sprintf("\tSystemModstamp: %v\n", t.SystemModstamp))
	builder.WriteString(fmt.Sprintf("\tUserId: %v\n", t.UserId))

	return builder.String()
}

type DatacloudOwnedEntityQueryResponse struct {
	BaseQuery
	Records []DatacloudOwnedEntity `json:"Records" force:"records"`
}
