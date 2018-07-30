// This file was generated for SObject DomainSite, API Version v43.0 at 2018-07-30 03:47:31.015792743 -0400 EDT m=+17.359050540

package sobjects

import (
	"fmt"
	"strings"
)

type DomainSite struct {
	BaseSObject
	CreatedById      string `force:",omitempty"`
	CreatedDate      string `force:",omitempty"`
	DomainId         string `force:",omitempty"`
	Id               string `force:",omitempty"`
	LastModifiedById string `force:",omitempty"`
	LastModifiedDate string `force:",omitempty"`
	PathPrefix       string `force:",omitempty"`
	SiteId           string `force:",omitempty"`
	SystemModstamp   string `force:",omitempty"`
}

func (t *DomainSite) ApiName() string {
	return "DomainSite"
}

func (t *DomainSite) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("DomainSite #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tDomainId: %v\n", t.DomainId))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tLastModifiedById: %v\n", t.LastModifiedById))
	builder.WriteString(fmt.Sprintf("\tLastModifiedDate: %v\n", t.LastModifiedDate))
	builder.WriteString(fmt.Sprintf("\tPathPrefix: %v\n", t.PathPrefix))
	builder.WriteString(fmt.Sprintf("\tSiteId: %v\n", t.SiteId))
	builder.WriteString(fmt.Sprintf("\tSystemModstamp: %v\n", t.SystemModstamp))

	return builder.String()
}

type DomainSiteQueryResponse struct {
	BaseQuery
	Records []DomainSite `json:"Records" force:"records"`
}
