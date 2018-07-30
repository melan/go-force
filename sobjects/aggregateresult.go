// This file was generated for SObject AggregateResult, API Version v43.0 at 2018-07-30 03:47:30.51652835 -0400 EDT m=+16.859767412

package sobjects

import (
	"fmt"
	"strings"
)

type AggregateResult struct {
	BaseSObject
	Id string `force:",omitempty"`
}

func (t *AggregateResult) ApiName() string {
	return "AggregateResult"
}

func (t *AggregateResult) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("AggregateResult #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))

	return builder.String()
}

type AggregateResultQueryResponse struct {
	BaseQuery
	Records []AggregateResult `json:"Records" force:"records"`
}
