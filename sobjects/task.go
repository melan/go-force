// This file was generated for SObject Task, API Version v43.0 at 2018-07-30 03:47:52.388184052 -0400 EDT m=+38.732243825

package sobjects

import (
	"fmt"
	"strings"
)

type Task struct {
	BaseSObject
	AccountId                 string `force:",omitempty"`
	ActivityDate              string `force:",omitempty"`
	CallDisposition           string `force:",omitempty"`
	CallDurationInSeconds     int    `force:",omitempty"`
	CallObject                string `force:",omitempty"`
	CallType                  string `force:",omitempty"`
	CreatedById               string `force:",omitempty"`
	CreatedDate               string `force:",omitempty"`
	Description               string `force:",omitempty"`
	Id                        string `force:",omitempty"`
	IsArchived                bool   `force:",omitempty"`
	IsClosed                  bool   `force:",omitempty"`
	IsDeleted                 bool   `force:",omitempty"`
	IsHighPriority            bool   `force:",omitempty"`
	IsRecurrence              bool   `force:",omitempty"`
	IsReminderSet             bool   `force:",omitempty"`
	LastModifiedById          string `force:",omitempty"`
	LastModifiedDate          string `force:",omitempty"`
	OwnerId                   string `force:",omitempty"`
	Priority                  string `force:",omitempty"`
	RecurrenceActivityId      string `force:",omitempty"`
	RecurrenceDayOfMonth      int    `force:",omitempty"`
	RecurrenceDayOfWeekMask   int    `force:",omitempty"`
	RecurrenceEndDateOnly     string `force:",omitempty"`
	RecurrenceInstance        string `force:",omitempty"`
	RecurrenceInterval        int    `force:",omitempty"`
	RecurrenceMonthOfYear     string `force:",omitempty"`
	RecurrenceRegeneratedType string `force:",omitempty"`
	RecurrenceStartDateOnly   string `force:",omitempty"`
	RecurrenceTimeZoneSidKey  string `force:",omitempty"`
	RecurrenceType            string `force:",omitempty"`
	ReminderDateTime          string `force:",omitempty"`
	Status                    string `force:",omitempty"`
	Subject                   string `force:",omitempty"`
	SystemModstamp            string `force:",omitempty"`
	TaskSubtype               string `force:",omitempty"`
	WhatId                    string `force:",omitempty"`
	WhoId                     string `force:",omitempty"`
}

func (t *Task) ApiName() string {
	return "Task"
}

func (t *Task) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("Task #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tAccountId: %v\n", t.AccountId))
	builder.WriteString(fmt.Sprintf("\tActivityDate: %v\n", t.ActivityDate))
	builder.WriteString(fmt.Sprintf("\tCallDisposition: %v\n", t.CallDisposition))
	builder.WriteString(fmt.Sprintf("\tCallDurationInSeconds: %v\n", t.CallDurationInSeconds))
	builder.WriteString(fmt.Sprintf("\tCallObject: %v\n", t.CallObject))
	builder.WriteString(fmt.Sprintf("\tCallType: %v\n", t.CallType))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tDescription: %v\n", t.Description))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tIsArchived: %v\n", t.IsArchived))
	builder.WriteString(fmt.Sprintf("\tIsClosed: %v\n", t.IsClosed))
	builder.WriteString(fmt.Sprintf("\tIsDeleted: %v\n", t.IsDeleted))
	builder.WriteString(fmt.Sprintf("\tIsHighPriority: %v\n", t.IsHighPriority))
	builder.WriteString(fmt.Sprintf("\tIsRecurrence: %v\n", t.IsRecurrence))
	builder.WriteString(fmt.Sprintf("\tIsReminderSet: %v\n", t.IsReminderSet))
	builder.WriteString(fmt.Sprintf("\tLastModifiedById: %v\n", t.LastModifiedById))
	builder.WriteString(fmt.Sprintf("\tLastModifiedDate: %v\n", t.LastModifiedDate))
	builder.WriteString(fmt.Sprintf("\tOwnerId: %v\n", t.OwnerId))
	builder.WriteString(fmt.Sprintf("\tPriority: %v\n", t.Priority))
	builder.WriteString(fmt.Sprintf("\tRecurrenceActivityId: %v\n", t.RecurrenceActivityId))
	builder.WriteString(fmt.Sprintf("\tRecurrenceDayOfMonth: %v\n", t.RecurrenceDayOfMonth))
	builder.WriteString(fmt.Sprintf("\tRecurrenceDayOfWeekMask: %v\n", t.RecurrenceDayOfWeekMask))
	builder.WriteString(fmt.Sprintf("\tRecurrenceEndDateOnly: %v\n", t.RecurrenceEndDateOnly))
	builder.WriteString(fmt.Sprintf("\tRecurrenceInstance: %v\n", t.RecurrenceInstance))
	builder.WriteString(fmt.Sprintf("\tRecurrenceInterval: %v\n", t.RecurrenceInterval))
	builder.WriteString(fmt.Sprintf("\tRecurrenceMonthOfYear: %v\n", t.RecurrenceMonthOfYear))
	builder.WriteString(fmt.Sprintf("\tRecurrenceRegeneratedType: %v\n", t.RecurrenceRegeneratedType))
	builder.WriteString(fmt.Sprintf("\tRecurrenceStartDateOnly: %v\n", t.RecurrenceStartDateOnly))
	builder.WriteString(fmt.Sprintf("\tRecurrenceTimeZoneSidKey: %v\n", t.RecurrenceTimeZoneSidKey))
	builder.WriteString(fmt.Sprintf("\tRecurrenceType: %v\n", t.RecurrenceType))
	builder.WriteString(fmt.Sprintf("\tReminderDateTime: %v\n", t.ReminderDateTime))
	builder.WriteString(fmt.Sprintf("\tStatus: %v\n", t.Status))
	builder.WriteString(fmt.Sprintf("\tSubject: %v\n", t.Subject))
	builder.WriteString(fmt.Sprintf("\tSystemModstamp: %v\n", t.SystemModstamp))
	builder.WriteString(fmt.Sprintf("\tTaskSubtype: %v\n", t.TaskSubtype))
	builder.WriteString(fmt.Sprintf("\tWhatId: %v\n", t.WhatId))
	builder.WriteString(fmt.Sprintf("\tWhoId: %v\n", t.WhoId))

	return builder.String()
}

type TaskQueryResponse struct {
	BaseQuery
	Records []Task `json:"Records" force:"records"`
}
