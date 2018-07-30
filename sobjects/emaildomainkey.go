// This file was generated for SObject EmailDomainKey, API Version v43.0 at 2018-07-30 03:47:33.367578621 -0400 EDT m=+19.710924666

package sobjects

import (
	"fmt"
	"strings"
)

type EmailDomainKey struct {
	BaseSObject
	CreatedById      string `force:",omitempty"`
	CreatedDate      string `force:",omitempty"`
	Domain           string `force:",omitempty"`
	DomainMatch      string `force:",omitempty"`
	Id               string `force:",omitempty"`
	IsActive         bool   `force:",omitempty"`
	IsDeleted        bool   `force:",omitempty"`
	LastModifiedById string `force:",omitempty"`
	LastModifiedDate string `force:",omitempty"`
	PrivateKey       string `force:",omitempty"`
	PublicKey        string `force:",omitempty"`
	Selector         string `force:",omitempty"`
	SystemModstamp   string `force:",omitempty"`
}

func (t *EmailDomainKey) ApiName() string {
	return "EmailDomainKey"
}

func (t *EmailDomainKey) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("EmailDomainKey #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tDomain: %v\n", t.Domain))
	builder.WriteString(fmt.Sprintf("\tDomainMatch: %v\n", t.DomainMatch))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tIsActive: %v\n", t.IsActive))
	builder.WriteString(fmt.Sprintf("\tIsDeleted: %v\n", t.IsDeleted))
	builder.WriteString(fmt.Sprintf("\tLastModifiedById: %v\n", t.LastModifiedById))
	builder.WriteString(fmt.Sprintf("\tLastModifiedDate: %v\n", t.LastModifiedDate))
	builder.WriteString(fmt.Sprintf("\tPrivateKey: %v\n", t.PrivateKey))
	builder.WriteString(fmt.Sprintf("\tPublicKey: %v\n", t.PublicKey))
	builder.WriteString(fmt.Sprintf("\tSelector: %v\n", t.Selector))
	builder.WriteString(fmt.Sprintf("\tSystemModstamp: %v\n", t.SystemModstamp))

	return builder.String()
}

type EmailDomainKeyQueryResponse struct {
	BaseQuery
	Records []EmailDomainKey `json:"Records" force:"records"`
}
