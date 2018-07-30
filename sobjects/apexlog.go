// This file was generated for SObject ApexLog, API Version v43.0 at 2018-07-30 03:47:53.326123547 -0400 EDT m=+39.670218516

package sobjects

import (
	"fmt"
	"strings"
)

type ApexLog struct {
	BaseSObject
	Application          string `force:",omitempty"`
	DurationMilliseconds int    `force:",omitempty"`
	Id                   string `force:",omitempty"`
	LastModifiedDate     string `force:",omitempty"`
	Location             string `force:",omitempty"`
	LogLength            int    `force:",omitempty"`
	LogUserId            string `force:",omitempty"`
	Operation            string `force:",omitempty"`
	Request              string `force:",omitempty"`
	StartTime            string `force:",omitempty"`
	Status               string `force:",omitempty"`
	SystemModstamp       string `force:",omitempty"`
}

func (t *ApexLog) ApiName() string {
	return "ApexLog"
}

func (t *ApexLog) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("ApexLog #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tApplication: %v\n", t.Application))
	builder.WriteString(fmt.Sprintf("\tDurationMilliseconds: %v\n", t.DurationMilliseconds))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tLastModifiedDate: %v\n", t.LastModifiedDate))
	builder.WriteString(fmt.Sprintf("\tLocation: %v\n", t.Location))
	builder.WriteString(fmt.Sprintf("\tLogLength: %v\n", t.LogLength))
	builder.WriteString(fmt.Sprintf("\tLogUserId: %v\n", t.LogUserId))
	builder.WriteString(fmt.Sprintf("\tOperation: %v\n", t.Operation))
	builder.WriteString(fmt.Sprintf("\tRequest: %v\n", t.Request))
	builder.WriteString(fmt.Sprintf("\tStartTime: %v\n", t.StartTime))
	builder.WriteString(fmt.Sprintf("\tStatus: %v\n", t.Status))
	builder.WriteString(fmt.Sprintf("\tSystemModstamp: %v\n", t.SystemModstamp))

	return builder.String()
}

type ApexLogQueryResponse struct {
	BaseQuery
	Records []ApexLog `json:"Records" force:"records"`
}
