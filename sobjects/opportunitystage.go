// This file was generated for SObject OpportunityStage, API Version v43.0 at 2018-07-30 03:47:45.513248948 -0400 EDT m=+31.857050747

package sobjects

import (
	"fmt"
	"strings"
)

type OpportunityStage struct {
	BaseSObject
	ApiName              string `force:",omitempty"`
	CreatedById          string `force:",omitempty"`
	CreatedDate          string `force:",omitempty"`
	DefaultProbability   string `force:",omitempty"`
	Description          string `force:",omitempty"`
	ForecastCategory     string `force:",omitempty"`
	ForecastCategoryName string `force:",omitempty"`
	Id                   string `force:",omitempty"`
	IsActive             bool   `force:",omitempty"`
	IsClosed             bool   `force:",omitempty"`
	IsWon                bool   `force:",omitempty"`
	LastModifiedById     string `force:",omitempty"`
	LastModifiedDate     string `force:",omitempty"`
	MasterLabel          string `force:",omitempty"`
	SortOrder            int    `force:",omitempty"`
	SystemModstamp       string `force:",omitempty"`
}

func (t *OpportunityStage) ApiName() string {
	return "OpportunityStage"
}

func (t *OpportunityStage) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("OpportunityStage #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tApiName: %v\n", t.ApiName))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tDefaultProbability: %v\n", t.DefaultProbability))
	builder.WriteString(fmt.Sprintf("\tDescription: %v\n", t.Description))
	builder.WriteString(fmt.Sprintf("\tForecastCategory: %v\n", t.ForecastCategory))
	builder.WriteString(fmt.Sprintf("\tForecastCategoryName: %v\n", t.ForecastCategoryName))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tIsActive: %v\n", t.IsActive))
	builder.WriteString(fmt.Sprintf("\tIsClosed: %v\n", t.IsClosed))
	builder.WriteString(fmt.Sprintf("\tIsWon: %v\n", t.IsWon))
	builder.WriteString(fmt.Sprintf("\tLastModifiedById: %v\n", t.LastModifiedById))
	builder.WriteString(fmt.Sprintf("\tLastModifiedDate: %v\n", t.LastModifiedDate))
	builder.WriteString(fmt.Sprintf("\tMasterLabel: %v\n", t.MasterLabel))
	builder.WriteString(fmt.Sprintf("\tSortOrder: %v\n", t.SortOrder))
	builder.WriteString(fmt.Sprintf("\tSystemModstamp: %v\n", t.SystemModstamp))

	return builder.String()
}

type OpportunityStageQueryResponse struct {
	BaseQuery
	Records []OpportunityStage `json:"Records" force:"records"`
}
