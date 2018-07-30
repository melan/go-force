// This file was generated for SObject EventLogFile, API Version v43.0 at 2018-07-30 03:47:24.579447617 -0400 EDT m=+10.922463897

package sobjects

import (
	"fmt"
	"strings"
)

type EventLogFile struct {
	BaseSObject
	ApiVersion         float64 `force:",omitempty"`
	CreatedById        string  `force:",omitempty"`
	CreatedDate        string  `force:",omitempty"`
	EventType          string  `force:",omitempty"`
	Id                 string  `force:",omitempty"`
	Interval           string  `force:",omitempty"`
	IsDeleted          bool    `force:",omitempty"`
	LastModifiedById   string  `force:",omitempty"`
	LastModifiedDate   string  `force:",omitempty"`
	LogDate            string  `force:",omitempty"`
	LogFile            string  `force:",omitempty"`
	LogFileContentType string  `force:",omitempty"`
	LogFileFieldNames  string  `force:",omitempty"`
	LogFileFieldTypes  string  `force:",omitempty"`
	LogFileLength      float64 `force:",omitempty"`
	Sequence           int     `force:",omitempty"`
	SystemModstamp     string  `force:",omitempty"`
}

func (t *EventLogFile) ApiName() string {
	return "EventLogFile"
}

func (t *EventLogFile) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("EventLogFile #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tApiVersion: %v\n", t.ApiVersion))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tEventType: %v\n", t.EventType))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tInterval: %v\n", t.Interval))
	builder.WriteString(fmt.Sprintf("\tIsDeleted: %v\n", t.IsDeleted))
	builder.WriteString(fmt.Sprintf("\tLastModifiedById: %v\n", t.LastModifiedById))
	builder.WriteString(fmt.Sprintf("\tLastModifiedDate: %v\n", t.LastModifiedDate))
	builder.WriteString(fmt.Sprintf("\tLogDate: %v\n", t.LogDate))
	builder.WriteString(fmt.Sprintf("\tLogFile: %v\n", t.LogFile))
	builder.WriteString(fmt.Sprintf("\tLogFileContentType: %v\n", t.LogFileContentType))
	builder.WriteString(fmt.Sprintf("\tLogFileFieldNames: %v\n", t.LogFileFieldNames))
	builder.WriteString(fmt.Sprintf("\tLogFileFieldTypes: %v\n", t.LogFileFieldTypes))
	builder.WriteString(fmt.Sprintf("\tLogFileLength: %v\n", t.LogFileLength))
	builder.WriteString(fmt.Sprintf("\tSequence: %v\n", t.Sequence))
	builder.WriteString(fmt.Sprintf("\tSystemModstamp: %v\n", t.SystemModstamp))

	return builder.String()
}

type EventLogFileQueryResponse struct {
	BaseQuery
	Records []EventLogFile `json:"Records" force:"records"`
}
