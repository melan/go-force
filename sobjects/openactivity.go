// This file was generated for SObject OpenActivity, API Version v43.0 at 2018-07-30 03:47:41.026956861 -0400 EDT m=+27.370590317

package sobjects

import (
	"fmt"
	"strings"
)

type OpenActivity struct {
	BaseSObject
	AccountId              string `force:",omitempty"`
	ActivityDate           string `force:",omitempty"`
	ActivitySubtype        string `force:",omitempty"`
	ActivityType           string `force:",omitempty"`
	AlternateDetailId      string `force:",omitempty"`
	CallDisposition        string `force:",omitempty"`
	CallDurationInSeconds  int    `force:",omitempty"`
	CallObject             string `force:",omitempty"`
	CallType               string `force:",omitempty"`
	CreatedById            string `force:",omitempty"`
	CreatedDate            string `force:",omitempty"`
	Description            string `force:",omitempty"`
	DurationInMinutes      int    `force:",omitempty"`
	EndDateTime            string `force:",omitempty"`
	Id                     string `force:",omitempty"`
	IsAllDayEvent          bool   `force:",omitempty"`
	IsClosed               bool   `force:",omitempty"`
	IsDeleted              bool   `force:",omitempty"`
	IsHighPriority         bool   `force:",omitempty"`
	IsReminderSet          bool   `force:",omitempty"`
	IsTask                 bool   `force:",omitempty"`
	IsVisibleInSelfService bool   `force:",omitempty"`
	LastModifiedById       string `force:",omitempty"`
	LastModifiedDate       string `force:",omitempty"`
	Location               string `force:",omitempty"`
	OwnerId                string `force:",omitempty"`
	Priority               string `force:",omitempty"`
	ReminderDateTime       string `force:",omitempty"`
	StartDateTime          string `force:",omitempty"`
	Status                 string `force:",omitempty"`
	Subject                string `force:",omitempty"`
	SystemModstamp         string `force:",omitempty"`
	WhatId                 string `force:",omitempty"`
	WhoId                  string `force:",omitempty"`
}

func (t *OpenActivity) ApiName() string {
	return "OpenActivity"
}

func (t *OpenActivity) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("OpenActivity #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tAccountId: %v\n", t.AccountId))
	builder.WriteString(fmt.Sprintf("\tActivityDate: %v\n", t.ActivityDate))
	builder.WriteString(fmt.Sprintf("\tActivitySubtype: %v\n", t.ActivitySubtype))
	builder.WriteString(fmt.Sprintf("\tActivityType: %v\n", t.ActivityType))
	builder.WriteString(fmt.Sprintf("\tAlternateDetailId: %v\n", t.AlternateDetailId))
	builder.WriteString(fmt.Sprintf("\tCallDisposition: %v\n", t.CallDisposition))
	builder.WriteString(fmt.Sprintf("\tCallDurationInSeconds: %v\n", t.CallDurationInSeconds))
	builder.WriteString(fmt.Sprintf("\tCallObject: %v\n", t.CallObject))
	builder.WriteString(fmt.Sprintf("\tCallType: %v\n", t.CallType))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tDescription: %v\n", t.Description))
	builder.WriteString(fmt.Sprintf("\tDurationInMinutes: %v\n", t.DurationInMinutes))
	builder.WriteString(fmt.Sprintf("\tEndDateTime: %v\n", t.EndDateTime))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tIsAllDayEvent: %v\n", t.IsAllDayEvent))
	builder.WriteString(fmt.Sprintf("\tIsClosed: %v\n", t.IsClosed))
	builder.WriteString(fmt.Sprintf("\tIsDeleted: %v\n", t.IsDeleted))
	builder.WriteString(fmt.Sprintf("\tIsHighPriority: %v\n", t.IsHighPriority))
	builder.WriteString(fmt.Sprintf("\tIsReminderSet: %v\n", t.IsReminderSet))
	builder.WriteString(fmt.Sprintf("\tIsTask: %v\n", t.IsTask))
	builder.WriteString(fmt.Sprintf("\tIsVisibleInSelfService: %v\n", t.IsVisibleInSelfService))
	builder.WriteString(fmt.Sprintf("\tLastModifiedById: %v\n", t.LastModifiedById))
	builder.WriteString(fmt.Sprintf("\tLastModifiedDate: %v\n", t.LastModifiedDate))
	builder.WriteString(fmt.Sprintf("\tLocation: %v\n", t.Location))
	builder.WriteString(fmt.Sprintf("\tOwnerId: %v\n", t.OwnerId))
	builder.WriteString(fmt.Sprintf("\tPriority: %v\n", t.Priority))
	builder.WriteString(fmt.Sprintf("\tReminderDateTime: %v\n", t.ReminderDateTime))
	builder.WriteString(fmt.Sprintf("\tStartDateTime: %v\n", t.StartDateTime))
	builder.WriteString(fmt.Sprintf("\tStatus: %v\n", t.Status))
	builder.WriteString(fmt.Sprintf("\tSubject: %v\n", t.Subject))
	builder.WriteString(fmt.Sprintf("\tSystemModstamp: %v\n", t.SystemModstamp))
	builder.WriteString(fmt.Sprintf("\tWhatId: %v\n", t.WhatId))
	builder.WriteString(fmt.Sprintf("\tWhoId: %v\n", t.WhoId))

	return builder.String()
}

type OpenActivityQueryResponse struct {
	BaseQuery
	Records []OpenActivity `json:"Records" force:"records"`
}
