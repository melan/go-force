// This file was generated for SObject Event, API Version v43.0 at 2018-07-30 03:48:12.132523876 -0400 EDT m=+58.477324536

package sobjects

import (
	"fmt"
	"strings"
)

type Event struct {
	BaseSObject
	AccountId                string `force:",omitempty"`
	ActivityDate             string `force:",omitempty"`
	ActivityDateTime         string `force:",omitempty"`
	CreatedById              string `force:",omitempty"`
	CreatedDate              string `force:",omitempty"`
	Description              string `force:",omitempty"`
	DurationInMinutes        int    `force:",omitempty"`
	EndDateTime              string `force:",omitempty"`
	EventSubtype             string `force:",omitempty"`
	GroupEventType           string `force:",omitempty"`
	Id                       string `force:",omitempty"`
	IsAllDayEvent            bool   `force:",omitempty"`
	IsArchived               bool   `force:",omitempty"`
	IsChild                  bool   `force:",omitempty"`
	IsDeleted                bool   `force:",omitempty"`
	IsGroupEvent             bool   `force:",omitempty"`
	IsPrivate                bool   `force:",omitempty"`
	IsRecurrence             bool   `force:",omitempty"`
	IsReminderSet            bool   `force:",omitempty"`
	LastModifiedById         string `force:",omitempty"`
	LastModifiedDate         string `force:",omitempty"`
	Location                 string `force:",omitempty"`
	OwnerId                  string `force:",omitempty"`
	RecurrenceActivityId     string `force:",omitempty"`
	RecurrenceDayOfMonth     int    `force:",omitempty"`
	RecurrenceDayOfWeekMask  int    `force:",omitempty"`
	RecurrenceEndDateOnly    string `force:",omitempty"`
	RecurrenceInstance       string `force:",omitempty"`
	RecurrenceInterval       int    `force:",omitempty"`
	RecurrenceMonthOfYear    string `force:",omitempty"`
	RecurrenceStartDateTime  string `force:",omitempty"`
	RecurrenceTimeZoneSidKey string `force:",omitempty"`
	RecurrenceType           string `force:",omitempty"`
	ReminderDateTime         string `force:",omitempty"`
	ShowAs                   string `force:",omitempty"`
	StartDateTime            string `force:",omitempty"`
	Subject                  string `force:",omitempty"`
	SystemModstamp           string `force:",omitempty"`
	WhatId                   string `force:",omitempty"`
	WhoId                    string `force:",omitempty"`
}

func (t *Event) ApiName() string {
	return "Event"
}

func (t *Event) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("Event #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tAccountId: %v\n", t.AccountId))
	builder.WriteString(fmt.Sprintf("\tActivityDate: %v\n", t.ActivityDate))
	builder.WriteString(fmt.Sprintf("\tActivityDateTime: %v\n", t.ActivityDateTime))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tDescription: %v\n", t.Description))
	builder.WriteString(fmt.Sprintf("\tDurationInMinutes: %v\n", t.DurationInMinutes))
	builder.WriteString(fmt.Sprintf("\tEndDateTime: %v\n", t.EndDateTime))
	builder.WriteString(fmt.Sprintf("\tEventSubtype: %v\n", t.EventSubtype))
	builder.WriteString(fmt.Sprintf("\tGroupEventType: %v\n", t.GroupEventType))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tIsAllDayEvent: %v\n", t.IsAllDayEvent))
	builder.WriteString(fmt.Sprintf("\tIsArchived: %v\n", t.IsArchived))
	builder.WriteString(fmt.Sprintf("\tIsChild: %v\n", t.IsChild))
	builder.WriteString(fmt.Sprintf("\tIsDeleted: %v\n", t.IsDeleted))
	builder.WriteString(fmt.Sprintf("\tIsGroupEvent: %v\n", t.IsGroupEvent))
	builder.WriteString(fmt.Sprintf("\tIsPrivate: %v\n", t.IsPrivate))
	builder.WriteString(fmt.Sprintf("\tIsRecurrence: %v\n", t.IsRecurrence))
	builder.WriteString(fmt.Sprintf("\tIsReminderSet: %v\n", t.IsReminderSet))
	builder.WriteString(fmt.Sprintf("\tLastModifiedById: %v\n", t.LastModifiedById))
	builder.WriteString(fmt.Sprintf("\tLastModifiedDate: %v\n", t.LastModifiedDate))
	builder.WriteString(fmt.Sprintf("\tLocation: %v\n", t.Location))
	builder.WriteString(fmt.Sprintf("\tOwnerId: %v\n", t.OwnerId))
	builder.WriteString(fmt.Sprintf("\tRecurrenceActivityId: %v\n", t.RecurrenceActivityId))
	builder.WriteString(fmt.Sprintf("\tRecurrenceDayOfMonth: %v\n", t.RecurrenceDayOfMonth))
	builder.WriteString(fmt.Sprintf("\tRecurrenceDayOfWeekMask: %v\n", t.RecurrenceDayOfWeekMask))
	builder.WriteString(fmt.Sprintf("\tRecurrenceEndDateOnly: %v\n", t.RecurrenceEndDateOnly))
	builder.WriteString(fmt.Sprintf("\tRecurrenceInstance: %v\n", t.RecurrenceInstance))
	builder.WriteString(fmt.Sprintf("\tRecurrenceInterval: %v\n", t.RecurrenceInterval))
	builder.WriteString(fmt.Sprintf("\tRecurrenceMonthOfYear: %v\n", t.RecurrenceMonthOfYear))
	builder.WriteString(fmt.Sprintf("\tRecurrenceStartDateTime: %v\n", t.RecurrenceStartDateTime))
	builder.WriteString(fmt.Sprintf("\tRecurrenceTimeZoneSidKey: %v\n", t.RecurrenceTimeZoneSidKey))
	builder.WriteString(fmt.Sprintf("\tRecurrenceType: %v\n", t.RecurrenceType))
	builder.WriteString(fmt.Sprintf("\tReminderDateTime: %v\n", t.ReminderDateTime))
	builder.WriteString(fmt.Sprintf("\tShowAs: %v\n", t.ShowAs))
	builder.WriteString(fmt.Sprintf("\tStartDateTime: %v\n", t.StartDateTime))
	builder.WriteString(fmt.Sprintf("\tSubject: %v\n", t.Subject))
	builder.WriteString(fmt.Sprintf("\tSystemModstamp: %v\n", t.SystemModstamp))
	builder.WriteString(fmt.Sprintf("\tWhatId: %v\n", t.WhatId))
	builder.WriteString(fmt.Sprintf("\tWhoId: %v\n", t.WhoId))

	return builder.String()
}

type EventQueryResponse struct {
	BaseQuery
	Records []Event `json:"Records" force:"records"`
}
