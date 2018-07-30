// This file was generated for SObject OpportunityHistory, API Version v43.0 at 2018-07-30 03:47:27.29579206 -0400 EDT m=+13.638910268

package sobjects

import (
	"fmt"
	"strings"
)

type OpportunityHistory struct {
	BaseSObject
	Amount           string `force:",omitempty"`
	CloseDate        string `force:",omitempty"`
	CreatedById      string `force:",omitempty"`
	CreatedDate      string `force:",omitempty"`
	ExpectedRevenue  string `force:",omitempty"`
	ForecastCategory string `force:",omitempty"`
	Id               string `force:",omitempty"`
	IsDeleted        bool   `force:",omitempty"`
	OpportunityId    string `force:",omitempty"`
	Probability      string `force:",omitempty"`
	StageName        string `force:",omitempty"`
	SystemModstamp   string `force:",omitempty"`
}

func (t *OpportunityHistory) ApiName() string {
	return "OpportunityHistory"
}

func (t *OpportunityHistory) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("OpportunityHistory #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tAmount: %v\n", t.Amount))
	builder.WriteString(fmt.Sprintf("\tCloseDate: %v\n", t.CloseDate))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tExpectedRevenue: %v\n", t.ExpectedRevenue))
	builder.WriteString(fmt.Sprintf("\tForecastCategory: %v\n", t.ForecastCategory))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tIsDeleted: %v\n", t.IsDeleted))
	builder.WriteString(fmt.Sprintf("\tOpportunityId: %v\n", t.OpportunityId))
	builder.WriteString(fmt.Sprintf("\tProbability: %v\n", t.Probability))
	builder.WriteString(fmt.Sprintf("\tStageName: %v\n", t.StageName))
	builder.WriteString(fmt.Sprintf("\tSystemModstamp: %v\n", t.SystemModstamp))

	return builder.String()
}

type OpportunityHistoryQueryResponse struct {
	BaseQuery
	Records []OpportunityHistory `json:"Records" force:"records"`
}
