// This file was generated for SObject Holiday, API Version v43.0 at 2018-07-30 03:47:29.872015491 -0400 EDT m=+16.215230369

package sobjects

import (
	"fmt"
	"strings"
)

type Holiday struct {
	BaseSObject
	ActivityDate            string `force:",omitempty"`
	CreatedById             string `force:",omitempty"`
	CreatedDate             string `force:",omitempty"`
	Description             string `force:",omitempty"`
	EndTimeInMinutes        int    `force:",omitempty"`
	Id                      string `force:",omitempty"`
	IsAllDay                bool   `force:",omitempty"`
	IsRecurrence            bool   `force:",omitempty"`
	LastModifiedById        string `force:",omitempty"`
	LastModifiedDate        string `force:",omitempty"`
	Name                    string `force:",omitempty"`
	RecurrenceDayOfMonth    int    `force:",omitempty"`
	RecurrenceDayOfWeekMask int    `force:",omitempty"`
	RecurrenceEndDateOnly   string `force:",omitempty"`
	RecurrenceInstance      string `force:",omitempty"`
	RecurrenceInterval      int    `force:",omitempty"`
	RecurrenceMonthOfYear   string `force:",omitempty"`
	RecurrenceStartDate     string `force:",omitempty"`
	RecurrenceType          string `force:",omitempty"`
	StartTimeInMinutes      int    `force:",omitempty"`
	SystemModstamp          string `force:",omitempty"`
}

func (t *Holiday) ApiName() string {
	return "Holiday"
}

func (t *Holiday) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("Holiday #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tActivityDate: %v\n", t.ActivityDate))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tDescription: %v\n", t.Description))
	builder.WriteString(fmt.Sprintf("\tEndTimeInMinutes: %v\n", t.EndTimeInMinutes))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tIsAllDay: %v\n", t.IsAllDay))
	builder.WriteString(fmt.Sprintf("\tIsRecurrence: %v\n", t.IsRecurrence))
	builder.WriteString(fmt.Sprintf("\tLastModifiedById: %v\n", t.LastModifiedById))
	builder.WriteString(fmt.Sprintf("\tLastModifiedDate: %v\n", t.LastModifiedDate))
	builder.WriteString(fmt.Sprintf("\tName: %v\n", t.Name))
	builder.WriteString(fmt.Sprintf("\tRecurrenceDayOfMonth: %v\n", t.RecurrenceDayOfMonth))
	builder.WriteString(fmt.Sprintf("\tRecurrenceDayOfWeekMask: %v\n", t.RecurrenceDayOfWeekMask))
	builder.WriteString(fmt.Sprintf("\tRecurrenceEndDateOnly: %v\n", t.RecurrenceEndDateOnly))
	builder.WriteString(fmt.Sprintf("\tRecurrenceInstance: %v\n", t.RecurrenceInstance))
	builder.WriteString(fmt.Sprintf("\tRecurrenceInterval: %v\n", t.RecurrenceInterval))
	builder.WriteString(fmt.Sprintf("\tRecurrenceMonthOfYear: %v\n", t.RecurrenceMonthOfYear))
	builder.WriteString(fmt.Sprintf("\tRecurrenceStartDate: %v\n", t.RecurrenceStartDate))
	builder.WriteString(fmt.Sprintf("\tRecurrenceType: %v\n", t.RecurrenceType))
	builder.WriteString(fmt.Sprintf("\tStartTimeInMinutes: %v\n", t.StartTimeInMinutes))
	builder.WriteString(fmt.Sprintf("\tSystemModstamp: %v\n", t.SystemModstamp))

	return builder.String()
}

type HolidayQueryResponse struct {
	BaseQuery
	Records []Holiday `json:"Records" force:"records"`
}
