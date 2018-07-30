// This file was generated for SObject ChatterExtensionConfig, API Version v43.0 at 2018-07-30 03:47:47.114866354 -0400 EDT m=+33.458728252

package sobjects

import (
	"fmt"
	"strings"
)

type ChatterExtensionConfig struct {
	BaseSObject
	CanCreate          bool   `force:",omitempty"`
	CanRead            bool   `force:",omitempty"`
	ChatterExtensionId string `force:",omitempty"`
	CreatedById        string `force:",omitempty"`
	CreatedDate        string `force:",omitempty"`
	Id                 string `force:",omitempty"`
	IsDeleted          bool   `force:",omitempty"`
	LastModifiedById   string `force:",omitempty"`
	LastModifiedDate   string `force:",omitempty"`
	Position           int    `force:",omitempty"`
	SystemModstamp     string `force:",omitempty"`
}

func (t *ChatterExtensionConfig) ApiName() string {
	return "ChatterExtensionConfig"
}

func (t *ChatterExtensionConfig) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("ChatterExtensionConfig #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tCanCreate: %v\n", t.CanCreate))
	builder.WriteString(fmt.Sprintf("\tCanRead: %v\n", t.CanRead))
	builder.WriteString(fmt.Sprintf("\tChatterExtensionId: %v\n", t.ChatterExtensionId))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tIsDeleted: %v\n", t.IsDeleted))
	builder.WriteString(fmt.Sprintf("\tLastModifiedById: %v\n", t.LastModifiedById))
	builder.WriteString(fmt.Sprintf("\tLastModifiedDate: %v\n", t.LastModifiedDate))
	builder.WriteString(fmt.Sprintf("\tPosition: %v\n", t.Position))
	builder.WriteString(fmt.Sprintf("\tSystemModstamp: %v\n", t.SystemModstamp))

	return builder.String()
}

type ChatterExtensionConfigQueryResponse struct {
	BaseQuery
	Records []ChatterExtensionConfig `json:"Records" force:"records"`
}
