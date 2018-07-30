// This file was generated for SObject DatacloudPurchaseUsage, API Version v43.0 at 2018-07-30 03:47:47.41234017 -0400 EDT m=+33.756213230

package sobjects

import (
	"fmt"
	"strings"
)

type DatacloudPurchaseUsage struct {
	BaseSObject
	CreatedById         string  `force:",omitempty"`
	CreatedDate         string  `force:",omitempty"`
	DatacloudEntityType string  `force:",omitempty"`
	Description         string  `force:",omitempty"`
	Id                  string  `force:",omitempty"`
	IsDeleted           bool    `force:",omitempty"`
	LastModifiedById    string  `force:",omitempty"`
	LastModifiedDate    string  `force:",omitempty"`
	Name                string  `force:",omitempty"`
	PurchaseType        string  `force:",omitempty"`
	SystemModstamp      string  `force:",omitempty"`
	Usage               float64 `force:",omitempty"`
	UserId              string  `force:",omitempty"`
	UserType            string  `force:",omitempty"`
}

func (t *DatacloudPurchaseUsage) ApiName() string {
	return "DatacloudPurchaseUsage"
}

func (t *DatacloudPurchaseUsage) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("DatacloudPurchaseUsage #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tDatacloudEntityType: %v\n", t.DatacloudEntityType))
	builder.WriteString(fmt.Sprintf("\tDescription: %v\n", t.Description))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tIsDeleted: %v\n", t.IsDeleted))
	builder.WriteString(fmt.Sprintf("\tLastModifiedById: %v\n", t.LastModifiedById))
	builder.WriteString(fmt.Sprintf("\tLastModifiedDate: %v\n", t.LastModifiedDate))
	builder.WriteString(fmt.Sprintf("\tName: %v\n", t.Name))
	builder.WriteString(fmt.Sprintf("\tPurchaseType: %v\n", t.PurchaseType))
	builder.WriteString(fmt.Sprintf("\tSystemModstamp: %v\n", t.SystemModstamp))
	builder.WriteString(fmt.Sprintf("\tUsage: %v\n", t.Usage))
	builder.WriteString(fmt.Sprintf("\tUserId: %v\n", t.UserId))
	builder.WriteString(fmt.Sprintf("\tUserType: %v\n", t.UserType))

	return builder.String()
}

type DatacloudPurchaseUsageQueryResponse struct {
	BaseQuery
	Records []DatacloudPurchaseUsage `json:"Records" force:"records"`
}
