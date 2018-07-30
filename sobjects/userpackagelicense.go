// This file was generated for SObject UserPackageLicense, API Version v43.0 at 2018-07-30 03:47:45.389114755 -0400 EDT m=+31.732911896

package sobjects

import (
	"fmt"
	"strings"
)

type UserPackageLicense struct {
	BaseSObject
	CreatedById      string `force:",omitempty"`
	CreatedDate      string `force:",omitempty"`
	Id               string `force:",omitempty"`
	LastModifiedById string `force:",omitempty"`
	LastModifiedDate string `force:",omitempty"`
	PackageLicenseId string `force:",omitempty"`
	SystemModstamp   string `force:",omitempty"`
	UserId           string `force:",omitempty"`
}

func (t *UserPackageLicense) ApiName() string {
	return "UserPackageLicense"
}

func (t *UserPackageLicense) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("UserPackageLicense #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tLastModifiedById: %v\n", t.LastModifiedById))
	builder.WriteString(fmt.Sprintf("\tLastModifiedDate: %v\n", t.LastModifiedDate))
	builder.WriteString(fmt.Sprintf("\tPackageLicenseId: %v\n", t.PackageLicenseId))
	builder.WriteString(fmt.Sprintf("\tSystemModstamp: %v\n", t.SystemModstamp))
	builder.WriteString(fmt.Sprintf("\tUserId: %v\n", t.UserId))

	return builder.String()
}

type UserPackageLicenseQueryResponse struct {
	BaseQuery
	Records []UserPackageLicense `json:"Records" force:"records"`
}
