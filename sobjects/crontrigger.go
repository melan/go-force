// This file was generated for SObject CronTrigger, API Version v43.0 at 2018-07-30 03:47:23.502154494 -0400 EDT m=+9.845130349

package sobjects

import (
	"fmt"
	"strings"
)

type CronTrigger struct {
	BaseSObject
	CreatedById      string `force:",omitempty"`
	CreatedDate      string `force:",omitempty"`
	CronExpression   string `force:",omitempty"`
	CronJobDetailId  string `force:",omitempty"`
	EndTime          string `force:",omitempty"`
	Id               string `force:",omitempty"`
	LastModifiedById string `force:",omitempty"`
	NextFireTime     string `force:",omitempty"`
	OwnerId          string `force:",omitempty"`
	PreviousFireTime string `force:",omitempty"`
	StartTime        string `force:",omitempty"`
	State            string `force:",omitempty"`
	TimeZoneSidKey   string `force:",omitempty"`
	TimesTriggered   int    `force:",omitempty"`
}

func (t *CronTrigger) ApiName() string {
	return "CronTrigger"
}

func (t *CronTrigger) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("CronTrigger #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tCronExpression: %v\n", t.CronExpression))
	builder.WriteString(fmt.Sprintf("\tCronJobDetailId: %v\n", t.CronJobDetailId))
	builder.WriteString(fmt.Sprintf("\tEndTime: %v\n", t.EndTime))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tLastModifiedById: %v\n", t.LastModifiedById))
	builder.WriteString(fmt.Sprintf("\tNextFireTime: %v\n", t.NextFireTime))
	builder.WriteString(fmt.Sprintf("\tOwnerId: %v\n", t.OwnerId))
	builder.WriteString(fmt.Sprintf("\tPreviousFireTime: %v\n", t.PreviousFireTime))
	builder.WriteString(fmt.Sprintf("\tStartTime: %v\n", t.StartTime))
	builder.WriteString(fmt.Sprintf("\tState: %v\n", t.State))
	builder.WriteString(fmt.Sprintf("\tTimeZoneSidKey: %v\n", t.TimeZoneSidKey))
	builder.WriteString(fmt.Sprintf("\tTimesTriggered: %v\n", t.TimesTriggered))

	return builder.String()
}

type CronTriggerQueryResponse struct {
	BaseQuery
	Records []CronTrigger `json:"Records" force:"records"`
}
