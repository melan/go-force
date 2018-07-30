// This file was generated for SObject PlatformAction, API Version v43.0 at 2018-07-30 03:47:59.444331809 -0400 EDT m=+45.788656357

package sobjects

import (
	"fmt"
	"strings"
)

type PlatformAction struct {
	BaseSObject
	ActionListContext   string `force:",omitempty"`
	ActionTarget        string `force:",omitempty"`
	ActionTargetType    string `force:",omitempty"`
	ApiName             string `force:",omitempty"`
	Category            string `force:",omitempty"`
	ConfirmationMessage string `force:",omitempty"`
	DeviceFormat        string `force:",omitempty"`
	ExternalId          string `force:",omitempty"`
	GroupId             string `force:",omitempty"`
	IconContentType     string `force:",omitempty"`
	IconHeight          int    `force:",omitempty"`
	IconUrl             string `force:",omitempty"`
	IconWidth           int    `force:",omitempty"`
	Id                  string `force:",omitempty"`
	InvocationStatus    string `force:",omitempty"`
	InvokedByUserId     string `force:",omitempty"`
	IsGroupDefault      bool   `force:",omitempty"`
	IsMassAction        bool   `force:",omitempty"`
	Label               string `force:",omitempty"`
	LastModifiedDate    string `force:",omitempty"`
	PrimaryColor        string `force:",omitempty"`
	RelatedListRecordId string `force:",omitempty"`
	RelatedSourceEntity string `force:",omitempty"`
	Section             string `force:",omitempty"`
	SourceEntity        string `force:",omitempty"`
	Subtype             string `force:",omitempty"`
	TargetObject        string `force:",omitempty"`
	TargetUrl           string `force:",omitempty"`
	Type                string `force:",omitempty"`
}

func (t *PlatformAction) ApiName() string {
	return "PlatformAction"
}

func (t *PlatformAction) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("PlatformAction #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tActionListContext: %v\n", t.ActionListContext))
	builder.WriteString(fmt.Sprintf("\tActionTarget: %v\n", t.ActionTarget))
	builder.WriteString(fmt.Sprintf("\tActionTargetType: %v\n", t.ActionTargetType))
	builder.WriteString(fmt.Sprintf("\tApiName: %v\n", t.ApiName))
	builder.WriteString(fmt.Sprintf("\tCategory: %v\n", t.Category))
	builder.WriteString(fmt.Sprintf("\tConfirmationMessage: %v\n", t.ConfirmationMessage))
	builder.WriteString(fmt.Sprintf("\tDeviceFormat: %v\n", t.DeviceFormat))
	builder.WriteString(fmt.Sprintf("\tExternalId: %v\n", t.ExternalId))
	builder.WriteString(fmt.Sprintf("\tGroupId: %v\n", t.GroupId))
	builder.WriteString(fmt.Sprintf("\tIconContentType: %v\n", t.IconContentType))
	builder.WriteString(fmt.Sprintf("\tIconHeight: %v\n", t.IconHeight))
	builder.WriteString(fmt.Sprintf("\tIconUrl: %v\n", t.IconUrl))
	builder.WriteString(fmt.Sprintf("\tIconWidth: %v\n", t.IconWidth))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tInvocationStatus: %v\n", t.InvocationStatus))
	builder.WriteString(fmt.Sprintf("\tInvokedByUserId: %v\n", t.InvokedByUserId))
	builder.WriteString(fmt.Sprintf("\tIsGroupDefault: %v\n", t.IsGroupDefault))
	builder.WriteString(fmt.Sprintf("\tIsMassAction: %v\n", t.IsMassAction))
	builder.WriteString(fmt.Sprintf("\tLabel: %v\n", t.Label))
	builder.WriteString(fmt.Sprintf("\tLastModifiedDate: %v\n", t.LastModifiedDate))
	builder.WriteString(fmt.Sprintf("\tPrimaryColor: %v\n", t.PrimaryColor))
	builder.WriteString(fmt.Sprintf("\tRelatedListRecordId: %v\n", t.RelatedListRecordId))
	builder.WriteString(fmt.Sprintf("\tRelatedSourceEntity: %v\n", t.RelatedSourceEntity))
	builder.WriteString(fmt.Sprintf("\tSection: %v\n", t.Section))
	builder.WriteString(fmt.Sprintf("\tSourceEntity: %v\n", t.SourceEntity))
	builder.WriteString(fmt.Sprintf("\tSubtype: %v\n", t.Subtype))
	builder.WriteString(fmt.Sprintf("\tTargetObject: %v\n", t.TargetObject))
	builder.WriteString(fmt.Sprintf("\tTargetUrl: %v\n", t.TargetUrl))
	builder.WriteString(fmt.Sprintf("\tType: %v\n", t.Type))

	return builder.String()
}

type PlatformActionQueryResponse struct {
	BaseQuery
	Records []PlatformAction `json:"Records" force:"records"`
}
