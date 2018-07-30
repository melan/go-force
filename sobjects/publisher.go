// This file was generated for SObject Publisher, API Version v43.0 at 2018-07-30 03:47:25.940649177 -0400 EDT m=+12.283716535

package sobjects

import (
	"fmt"
	"strings"
)

type Publisher struct {
	BaseSObject
	DurableId       string `force:",omitempty"`
	Id              string `force:",omitempty"`
	IsSalesforce    bool   `force:",omitempty"`
	MajorVersion    int    `force:",omitempty"`
	MinorVersion    int    `force:",omitempty"`
	Name            string `force:",omitempty"`
	NamespacePrefix string `force:",omitempty"`
}

func (t *Publisher) ApiName() string {
	return "Publisher"
}

func (t *Publisher) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("Publisher #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tDurableId: %v\n", t.DurableId))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tIsSalesforce: %v\n", t.IsSalesforce))
	builder.WriteString(fmt.Sprintf("\tMajorVersion: %v\n", t.MajorVersion))
	builder.WriteString(fmt.Sprintf("\tMinorVersion: %v\n", t.MinorVersion))
	builder.WriteString(fmt.Sprintf("\tName: %v\n", t.Name))
	builder.WriteString(fmt.Sprintf("\tNamespacePrefix: %v\n", t.NamespacePrefix))

	return builder.String()
}

type PublisherQueryResponse struct {
	BaseQuery
	Records []Publisher `json:"Records" force:"records"`
}
