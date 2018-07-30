// This file was generated for SObject TenantUsageEntitlement, API Version v43.0 at 2018-07-30 03:48:10.153878778 -0400 EDT m=+56.498605191

package sobjects

import (
	"fmt"
	"strings"
)

type TenantUsageEntitlement struct {
	BaseSObject
	AmountUsed           float64 `force:",omitempty"`
	CreatedById          string  `force:",omitempty"`
	CreatedDate          string  `force:",omitempty"`
	CurrentAmountAllowed float64 `force:",omitempty"`
	EndDate              string  `force:",omitempty"`
	Frequency            string  `force:",omitempty"`
	HasRollover          bool    `force:",omitempty"`
	Id                   string  `force:",omitempty"`
	IsDeleted            bool    `force:",omitempty"`
	IsPersistentResource bool    `force:",omitempty"`
	LastModifiedById     string  `force:",omitempty"`
	LastModifiedDate     string  `force:",omitempty"`
	MasterLabel          string  `force:",omitempty"`
	OverageGrace         string  `force:",omitempty"`
	ResourceGroupKey     string  `force:",omitempty"`
	Setting              string  `force:",omitempty"`
	StartDate            string  `force:",omitempty"`
	SystemModstamp       string  `force:",omitempty"`
	UsageDate            string  `force:",omitempty"`
}

func (t *TenantUsageEntitlement) ApiName() string {
	return "TenantUsageEntitlement"
}

func (t *TenantUsageEntitlement) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("TenantUsageEntitlement #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tAmountUsed: %v\n", t.AmountUsed))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tCurrentAmountAllowed: %v\n", t.CurrentAmountAllowed))
	builder.WriteString(fmt.Sprintf("\tEndDate: %v\n", t.EndDate))
	builder.WriteString(fmt.Sprintf("\tFrequency: %v\n", t.Frequency))
	builder.WriteString(fmt.Sprintf("\tHasRollover: %v\n", t.HasRollover))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tIsDeleted: %v\n", t.IsDeleted))
	builder.WriteString(fmt.Sprintf("\tIsPersistentResource: %v\n", t.IsPersistentResource))
	builder.WriteString(fmt.Sprintf("\tLastModifiedById: %v\n", t.LastModifiedById))
	builder.WriteString(fmt.Sprintf("\tLastModifiedDate: %v\n", t.LastModifiedDate))
	builder.WriteString(fmt.Sprintf("\tMasterLabel: %v\n", t.MasterLabel))
	builder.WriteString(fmt.Sprintf("\tOverageGrace: %v\n", t.OverageGrace))
	builder.WriteString(fmt.Sprintf("\tResourceGroupKey: %v\n", t.ResourceGroupKey))
	builder.WriteString(fmt.Sprintf("\tSetting: %v\n", t.Setting))
	builder.WriteString(fmt.Sprintf("\tStartDate: %v\n", t.StartDate))
	builder.WriteString(fmt.Sprintf("\tSystemModstamp: %v\n", t.SystemModstamp))
	builder.WriteString(fmt.Sprintf("\tUsageDate: %v\n", t.UsageDate))

	return builder.String()
}

type TenantUsageEntitlementQueryResponse struct {
	BaseQuery
	Records []TenantUsageEntitlement `json:"Records" force:"records"`
}
