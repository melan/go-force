// This file was generated for SObject PackageLicense, API Version v43.0 at 2018-07-30 03:47:43.765029696 -0400 EDT m=+30.108765895

package sobjects

import (
	"fmt"
	"strings"
)

type PackageLicense struct {
	BaseSObject
	AllowedLicenses  int    `force:",omitempty"`
	CreatedDate      string `force:",omitempty"`
	ExpirationDate   string `force:",omitempty"`
	Id               string `force:",omitempty"`
	IsProvisioned    bool   `force:",omitempty"`
	LastModifiedDate string `force:",omitempty"`
	NamespacePrefix  string `force:",omitempty"`
	Status           string `force:",omitempty"`
	SystemModstamp   string `force:",omitempty"`
	UsedLicenses     int    `force:",omitempty"`
}

func (t *PackageLicense) ApiName() string {
	return "PackageLicense"
}

func (t *PackageLicense) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("PackageLicense #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tAllowedLicenses: %v\n", t.AllowedLicenses))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tExpirationDate: %v\n", t.ExpirationDate))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tIsProvisioned: %v\n", t.IsProvisioned))
	builder.WriteString(fmt.Sprintf("\tLastModifiedDate: %v\n", t.LastModifiedDate))
	builder.WriteString(fmt.Sprintf("\tNamespacePrefix: %v\n", t.NamespacePrefix))
	builder.WriteString(fmt.Sprintf("\tStatus: %v\n", t.Status))
	builder.WriteString(fmt.Sprintf("\tSystemModstamp: %v\n", t.SystemModstamp))
	builder.WriteString(fmt.Sprintf("\tUsedLicenses: %v\n", t.UsedLicenses))

	return builder.String()
}

type PackageLicenseQueryResponse struct {
	BaseQuery
	Records []PackageLicense `json:"Records" force:"records"`
}
