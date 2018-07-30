// This file was generated for SObject CronJobDetail, API Version v43.0 at 2018-07-30 03:47:38.844891075 -0400 EDT m=+25.188442651

package sobjects

import (
	"fmt"
	"strings"
)

type CronJobDetail struct {
	BaseSObject
	Id      string `force:",omitempty"`
	JobType string `force:",omitempty"`
	Name    string `force:",omitempty"`
}

func (t *CronJobDetail) ApiName() string {
	return "CronJobDetail"
}

func (t *CronJobDetail) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("CronJobDetail #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tJobType: %v\n", t.JobType))
	builder.WriteString(fmt.Sprintf("\tName: %v\n", t.Name))

	return builder.String()
}

type CronJobDetailQueryResponse struct {
	BaseQuery
	Records []CronJobDetail `json:"Records" force:"records"`
}
