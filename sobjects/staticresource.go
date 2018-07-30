// This file was generated for SObject StaticResource, API Version v43.0 at 2018-07-30 03:47:56.120425896 -0400 EDT m=+42.464625718

package sobjects

import (
	"fmt"
	"strings"
)

type StaticResource struct {
	BaseSObject
	Body             string `force:",omitempty"`
	BodyLength       int    `force:",omitempty"`
	CacheControl     string `force:",omitempty"`
	ContentType      string `force:",omitempty"`
	CreatedById      string `force:",omitempty"`
	CreatedDate      string `force:",omitempty"`
	Description      string `force:",omitempty"`
	Id               string `force:",omitempty"`
	LastModifiedById string `force:",omitempty"`
	LastModifiedDate string `force:",omitempty"`
	Name             string `force:",omitempty"`
	NamespacePrefix  string `force:",omitempty"`
	SystemModstamp   string `force:",omitempty"`
}

func (t *StaticResource) ApiName() string {
	return "StaticResource"
}

func (t *StaticResource) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("StaticResource #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tBody: %v\n", t.Body))
	builder.WriteString(fmt.Sprintf("\tBodyLength: %v\n", t.BodyLength))
	builder.WriteString(fmt.Sprintf("\tCacheControl: %v\n", t.CacheControl))
	builder.WriteString(fmt.Sprintf("\tContentType: %v\n", t.ContentType))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tDescription: %v\n", t.Description))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tLastModifiedById: %v\n", t.LastModifiedById))
	builder.WriteString(fmt.Sprintf("\tLastModifiedDate: %v\n", t.LastModifiedDate))
	builder.WriteString(fmt.Sprintf("\tName: %v\n", t.Name))
	builder.WriteString(fmt.Sprintf("\tNamespacePrefix: %v\n", t.NamespacePrefix))
	builder.WriteString(fmt.Sprintf("\tSystemModstamp: %v\n", t.SystemModstamp))

	return builder.String()
}

type StaticResourceQueryResponse struct {
	BaseQuery
	Records []StaticResource `json:"Records" force:"records"`
}
