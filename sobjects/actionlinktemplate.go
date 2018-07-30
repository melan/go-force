// This file was generated for SObject ActionLinkTemplate, API Version v43.0 at 2018-07-30 03:47:58.55727601 -0400 EDT m=+44.901567272

package sobjects

import (
	"fmt"
	"strings"
)

type ActionLinkTemplate struct {
	BaseSObject
	ActionLinkGroupTemplateId string `force:",omitempty"`
	ActionUrl                 string `force:",omitempty"`
	CreatedById               string `force:",omitempty"`
	CreatedDate               string `force:",omitempty"`
	Headers                   string `force:",omitempty"`
	Id                        string `force:",omitempty"`
	IsConfirmationRequired    bool   `force:",omitempty"`
	IsDeleted                 bool   `force:",omitempty"`
	IsGroupDefault            bool   `force:",omitempty"`
	Label                     string `force:",omitempty"`
	LabelKey                  string `force:",omitempty"`
	LastModifiedById          string `force:",omitempty"`
	LastModifiedDate          string `force:",omitempty"`
	LinkType                  string `force:",omitempty"`
	Method                    string `force:",omitempty"`
	Position                  int    `force:",omitempty"`
	RequestBody               string `force:",omitempty"`
	SystemModstamp            string `force:",omitempty"`
	UserAlias                 string `force:",omitempty"`
	UserVisibility            string `force:",omitempty"`
}

func (t *ActionLinkTemplate) ApiName() string {
	return "ActionLinkTemplate"
}

func (t *ActionLinkTemplate) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("ActionLinkTemplate #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tActionLinkGroupTemplateId: %v\n", t.ActionLinkGroupTemplateId))
	builder.WriteString(fmt.Sprintf("\tActionUrl: %v\n", t.ActionUrl))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tHeaders: %v\n", t.Headers))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tIsConfirmationRequired: %v\n", t.IsConfirmationRequired))
	builder.WriteString(fmt.Sprintf("\tIsDeleted: %v\n", t.IsDeleted))
	builder.WriteString(fmt.Sprintf("\tIsGroupDefault: %v\n", t.IsGroupDefault))
	builder.WriteString(fmt.Sprintf("\tLabel: %v\n", t.Label))
	builder.WriteString(fmt.Sprintf("\tLabelKey: %v\n", t.LabelKey))
	builder.WriteString(fmt.Sprintf("\tLastModifiedById: %v\n", t.LastModifiedById))
	builder.WriteString(fmt.Sprintf("\tLastModifiedDate: %v\n", t.LastModifiedDate))
	builder.WriteString(fmt.Sprintf("\tLinkType: %v\n", t.LinkType))
	builder.WriteString(fmt.Sprintf("\tMethod: %v\n", t.Method))
	builder.WriteString(fmt.Sprintf("\tPosition: %v\n", t.Position))
	builder.WriteString(fmt.Sprintf("\tRequestBody: %v\n", t.RequestBody))
	builder.WriteString(fmt.Sprintf("\tSystemModstamp: %v\n", t.SystemModstamp))
	builder.WriteString(fmt.Sprintf("\tUserAlias: %v\n", t.UserAlias))
	builder.WriteString(fmt.Sprintf("\tUserVisibility: %v\n", t.UserVisibility))

	return builder.String()
}

type ActionLinkTemplateQueryResponse struct {
	BaseQuery
	Records []ActionLinkTemplate `json:"Records" force:"records"`
}
