// This file was generated for SObject Domain, API Version v43.0 at 2018-07-30 03:47:22.997268812 -0400 EDT m=+9.340225722

package sobjects

import (
	"fmt"
	"strings"
)

type Domain struct {
	BaseSObject
	CnameTarget          string `force:",omitempty"`
	CreatedById          string `force:",omitempty"`
	CreatedDate          string `force:",omitempty"`
	Domain               string `force:",omitempty"`
	DomainType           string `force:",omitempty"`
	Id                   string `force:",omitempty"`
	LastModifiedById     string `force:",omitempty"`
	LastModifiedDate     string `force:",omitempty"`
	OptionsExternalHttps bool   `force:",omitempty"`
	SystemModstamp       string `force:",omitempty"`
}

func (t *Domain) ApiName() string {
	return "Domain"
}

func (t *Domain) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("Domain #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tCnameTarget: %v\n", t.CnameTarget))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tDomain: %v\n", t.Domain))
	builder.WriteString(fmt.Sprintf("\tDomainType: %v\n", t.DomainType))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tLastModifiedById: %v\n", t.LastModifiedById))
	builder.WriteString(fmt.Sprintf("\tLastModifiedDate: %v\n", t.LastModifiedDate))
	builder.WriteString(fmt.Sprintf("\tOptionsExternalHttps: %v\n", t.OptionsExternalHttps))
	builder.WriteString(fmt.Sprintf("\tSystemModstamp: %v\n", t.SystemModstamp))

	return builder.String()
}

type DomainQueryResponse struct {
	BaseQuery
	Records []Domain `json:"Records" force:"records"`
}
