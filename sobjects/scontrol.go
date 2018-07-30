// This file was generated for SObject Scontrol, API Version v43.0 at 2018-07-30 03:47:36.763644489 -0400 EDT m=+23.107117968

package sobjects

import (
	"fmt"
	"strings"
)

type Scontrol struct {
	BaseSObject
	Binary           string `force:",omitempty"`
	BodyLength       int    `force:",omitempty"`
	ContentSource    string `force:",omitempty"`
	CreatedById      string `force:",omitempty"`
	CreatedDate      string `force:",omitempty"`
	Description      string `force:",omitempty"`
	DeveloperName    string `force:",omitempty"`
	EncodingKey      string `force:",omitempty"`
	Filename         string `force:",omitempty"`
	HtmlWrapper      string `force:",omitempty"`
	Id               string `force:",omitempty"`
	IsDeleted        bool   `force:",omitempty"`
	LastModifiedById string `force:",omitempty"`
	LastModifiedDate string `force:",omitempty"`
	Name             string `force:",omitempty"`
	NamespacePrefix  string `force:",omitempty"`
	SupportsCaching  bool   `force:",omitempty"`
	SystemModstamp   string `force:",omitempty"`
}

func (t *Scontrol) ApiName() string {
	return "Scontrol"
}

func (t *Scontrol) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("Scontrol #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tBinary: %v\n", t.Binary))
	builder.WriteString(fmt.Sprintf("\tBodyLength: %v\n", t.BodyLength))
	builder.WriteString(fmt.Sprintf("\tContentSource: %v\n", t.ContentSource))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tDescription: %v\n", t.Description))
	builder.WriteString(fmt.Sprintf("\tDeveloperName: %v\n", t.DeveloperName))
	builder.WriteString(fmt.Sprintf("\tEncodingKey: %v\n", t.EncodingKey))
	builder.WriteString(fmt.Sprintf("\tFilename: %v\n", t.Filename))
	builder.WriteString(fmt.Sprintf("\tHtmlWrapper: %v\n", t.HtmlWrapper))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tIsDeleted: %v\n", t.IsDeleted))
	builder.WriteString(fmt.Sprintf("\tLastModifiedById: %v\n", t.LastModifiedById))
	builder.WriteString(fmt.Sprintf("\tLastModifiedDate: %v\n", t.LastModifiedDate))
	builder.WriteString(fmt.Sprintf("\tName: %v\n", t.Name))
	builder.WriteString(fmt.Sprintf("\tNamespacePrefix: %v\n", t.NamespacePrefix))
	builder.WriteString(fmt.Sprintf("\tSupportsCaching: %v\n", t.SupportsCaching))
	builder.WriteString(fmt.Sprintf("\tSystemModstamp: %v\n", t.SystemModstamp))

	return builder.String()
}

type ScontrolQueryResponse struct {
	BaseQuery
	Records []Scontrol `json:"Records" force:"records"`
}
