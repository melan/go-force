// This file was generated for SObject ContentDistributionView, API Version v43.0 at 2018-07-30 03:48:12.24558985 -0400 EDT m=+58.590394752

package sobjects

import (
	"fmt"
	"strings"
)

type ContentDistributionView struct {
	BaseSObject
	CreatedById    string `force:",omitempty"`
	CreatedDate    string `force:",omitempty"`
	DistributionId string `force:",omitempty"`
	Id             string `force:",omitempty"`
	IsDeleted      bool   `force:",omitempty"`
	IsDownload     bool   `force:",omitempty"`
	IsInternal     bool   `force:",omitempty"`
	ParentViewId   string `force:",omitempty"`
	SystemModstamp string `force:",omitempty"`
}

func (t *ContentDistributionView) ApiName() string {
	return "ContentDistributionView"
}

func (t *ContentDistributionView) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("ContentDistributionView #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tDistributionId: %v\n", t.DistributionId))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tIsDeleted: %v\n", t.IsDeleted))
	builder.WriteString(fmt.Sprintf("\tIsDownload: %v\n", t.IsDownload))
	builder.WriteString(fmt.Sprintf("\tIsInternal: %v\n", t.IsInternal))
	builder.WriteString(fmt.Sprintf("\tParentViewId: %v\n", t.ParentViewId))
	builder.WriteString(fmt.Sprintf("\tSystemModstamp: %v\n", t.SystemModstamp))

	return builder.String()
}

type ContentDistributionViewQueryResponse struct {
	BaseQuery
	Records []ContentDistributionView `json:"Records" force:"records"`
}
