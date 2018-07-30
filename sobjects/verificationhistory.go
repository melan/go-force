// This file was generated for SObject VerificationHistory, API Version v43.0 at 2018-07-30 03:47:20.597971974 -0400 EDT m=+6.940838852

package sobjects

import (
	"fmt"
	"strings"
)

type VerificationHistory struct {
	BaseSObject
	Activity           string `force:",omitempty"`
	CreatedById        string `force:",omitempty"`
	CreatedDate        string `force:",omitempty"`
	EventGroup         int    `force:",omitempty"`
	Id                 string `force:",omitempty"`
	IsDeleted          bool   `force:",omitempty"`
	LastModifiedById   string `force:",omitempty"`
	LastModifiedDate   string `force:",omitempty"`
	LoginGeoId         string `force:",omitempty"`
	LoginHistoryId     string `force:",omitempty"`
	Policy             string `force:",omitempty"`
	Remarks            string `force:",omitempty"`
	ResourceId         string `force:",omitempty"`
	SourceIp           string `force:",omitempty"`
	Status             string `force:",omitempty"`
	SystemModstamp     string `force:",omitempty"`
	UserId             string `force:",omitempty"`
	VerificationMethod string `force:",omitempty"`
	VerificationTime   string `force:",omitempty"`
}

func (t *VerificationHistory) ApiName() string {
	return "VerificationHistory"
}

func (t *VerificationHistory) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("VerificationHistory #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tActivity: %v\n", t.Activity))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tEventGroup: %v\n", t.EventGroup))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tIsDeleted: %v\n", t.IsDeleted))
	builder.WriteString(fmt.Sprintf("\tLastModifiedById: %v\n", t.LastModifiedById))
	builder.WriteString(fmt.Sprintf("\tLastModifiedDate: %v\n", t.LastModifiedDate))
	builder.WriteString(fmt.Sprintf("\tLoginGeoId: %v\n", t.LoginGeoId))
	builder.WriteString(fmt.Sprintf("\tLoginHistoryId: %v\n", t.LoginHistoryId))
	builder.WriteString(fmt.Sprintf("\tPolicy: %v\n", t.Policy))
	builder.WriteString(fmt.Sprintf("\tRemarks: %v\n", t.Remarks))
	builder.WriteString(fmt.Sprintf("\tResourceId: %v\n", t.ResourceId))
	builder.WriteString(fmt.Sprintf("\tSourceIp: %v\n", t.SourceIp))
	builder.WriteString(fmt.Sprintf("\tStatus: %v\n", t.Status))
	builder.WriteString(fmt.Sprintf("\tSystemModstamp: %v\n", t.SystemModstamp))
	builder.WriteString(fmt.Sprintf("\tUserId: %v\n", t.UserId))
	builder.WriteString(fmt.Sprintf("\tVerificationMethod: %v\n", t.VerificationMethod))
	builder.WriteString(fmt.Sprintf("\tVerificationTime: %v\n", t.VerificationTime))

	return builder.String()
}

type VerificationHistoryQueryResponse struct {
	BaseQuery
	Records []VerificationHistory `json:"Records" force:"records"`
}
