// This file was generated for SObject PlatformCachePartitionType, API Version v43.0 at 2018-07-30 03:47:25.203153075 -0400 EDT m=+11.546192758

package sobjects

import (
	"fmt"
	"strings"
)

type PlatformCachePartitionType struct {
	BaseSObject
	AllocatedCapacity          int    `force:",omitempty"`
	AllocatedPurchasedCapacity int    `force:",omitempty"`
	AllocatedTrialCapacity     int    `force:",omitempty"`
	CacheType                  string `force:",omitempty"`
	CreatedById                string `force:",omitempty"`
	CreatedDate                string `force:",omitempty"`
	Id                         string `force:",omitempty"`
	IsDeleted                  bool   `force:",omitempty"`
	LastModifiedById           string `force:",omitempty"`
	LastModifiedDate           string `force:",omitempty"`
	PlatformCachePartitionId   string `force:",omitempty"`
	SystemModstamp             string `force:",omitempty"`
}

func (t *PlatformCachePartitionType) ApiName() string {
	return "PlatformCachePartitionType"
}

func (t *PlatformCachePartitionType) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("PlatformCachePartitionType #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tAllocatedCapacity: %v\n", t.AllocatedCapacity))
	builder.WriteString(fmt.Sprintf("\tAllocatedPurchasedCapacity: %v\n", t.AllocatedPurchasedCapacity))
	builder.WriteString(fmt.Sprintf("\tAllocatedTrialCapacity: %v\n", t.AllocatedTrialCapacity))
	builder.WriteString(fmt.Sprintf("\tCacheType: %v\n", t.CacheType))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tIsDeleted: %v\n", t.IsDeleted))
	builder.WriteString(fmt.Sprintf("\tLastModifiedById: %v\n", t.LastModifiedById))
	builder.WriteString(fmt.Sprintf("\tLastModifiedDate: %v\n", t.LastModifiedDate))
	builder.WriteString(fmt.Sprintf("\tPlatformCachePartitionId: %v\n", t.PlatformCachePartitionId))
	builder.WriteString(fmt.Sprintf("\tSystemModstamp: %v\n", t.SystemModstamp))

	return builder.String()
}

type PlatformCachePartitionTypeQueryResponse struct {
	BaseQuery
	Records []PlatformCachePartitionType `json:"Records" force:"records"`
}
