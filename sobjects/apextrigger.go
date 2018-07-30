// This file was generated for SObject ApexTrigger, API Version v43.0 at 2018-07-30 03:48:05.114402876 -0400 EDT m=+51.458940188

package sobjects

import (
	"fmt"
	"strings"
)

type ApexTrigger struct {
	BaseSObject
	ApiVersion            float64 `force:",omitempty"`
	Body                  string  `force:",omitempty"`
	BodyCrc               float64 `force:",omitempty"`
	CreatedById           string  `force:",omitempty"`
	CreatedDate           string  `force:",omitempty"`
	Id                    string  `force:",omitempty"`
	IsValid               bool    `force:",omitempty"`
	LastModifiedById      string  `force:",omitempty"`
	LastModifiedDate      string  `force:",omitempty"`
	LengthWithoutComments int     `force:",omitempty"`
	Name                  string  `force:",omitempty"`
	NamespacePrefix       string  `force:",omitempty"`
	Status                string  `force:",omitempty"`
	SystemModstamp        string  `force:",omitempty"`
	TableEnumOrId         string  `force:",omitempty"`
	UsageAfterDelete      bool    `force:",omitempty"`
	UsageAfterInsert      bool    `force:",omitempty"`
	UsageAfterUndelete    bool    `force:",omitempty"`
	UsageAfterUpdate      bool    `force:",omitempty"`
	UsageBeforeDelete     bool    `force:",omitempty"`
	UsageBeforeInsert     bool    `force:",omitempty"`
	UsageBeforeUpdate     bool    `force:",omitempty"`
	UsageIsBulk           bool    `force:",omitempty"`
}

func (t *ApexTrigger) ApiName() string {
	return "ApexTrigger"
}

func (t *ApexTrigger) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("ApexTrigger #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tApiVersion: %v\n", t.ApiVersion))
	builder.WriteString(fmt.Sprintf("\tBody: %v\n", t.Body))
	builder.WriteString(fmt.Sprintf("\tBodyCrc: %v\n", t.BodyCrc))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tIsValid: %v\n", t.IsValid))
	builder.WriteString(fmt.Sprintf("\tLastModifiedById: %v\n", t.LastModifiedById))
	builder.WriteString(fmt.Sprintf("\tLastModifiedDate: %v\n", t.LastModifiedDate))
	builder.WriteString(fmt.Sprintf("\tLengthWithoutComments: %v\n", t.LengthWithoutComments))
	builder.WriteString(fmt.Sprintf("\tName: %v\n", t.Name))
	builder.WriteString(fmt.Sprintf("\tNamespacePrefix: %v\n", t.NamespacePrefix))
	builder.WriteString(fmt.Sprintf("\tStatus: %v\n", t.Status))
	builder.WriteString(fmt.Sprintf("\tSystemModstamp: %v\n", t.SystemModstamp))
	builder.WriteString(fmt.Sprintf("\tTableEnumOrId: %v\n", t.TableEnumOrId))
	builder.WriteString(fmt.Sprintf("\tUsageAfterDelete: %v\n", t.UsageAfterDelete))
	builder.WriteString(fmt.Sprintf("\tUsageAfterInsert: %v\n", t.UsageAfterInsert))
	builder.WriteString(fmt.Sprintf("\tUsageAfterUndelete: %v\n", t.UsageAfterUndelete))
	builder.WriteString(fmt.Sprintf("\tUsageAfterUpdate: %v\n", t.UsageAfterUpdate))
	builder.WriteString(fmt.Sprintf("\tUsageBeforeDelete: %v\n", t.UsageBeforeDelete))
	builder.WriteString(fmt.Sprintf("\tUsageBeforeInsert: %v\n", t.UsageBeforeInsert))
	builder.WriteString(fmt.Sprintf("\tUsageBeforeUpdate: %v\n", t.UsageBeforeUpdate))
	builder.WriteString(fmt.Sprintf("\tUsageIsBulk: %v\n", t.UsageIsBulk))

	return builder.String()
}

type ApexTriggerQueryResponse struct {
	BaseQuery
	Records []ApexTrigger `json:"Records" force:"records"`
}
