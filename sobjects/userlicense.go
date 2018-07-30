// This file was generated for SObject UserLicense, API Version v43.0 at 2018-07-30 03:47:59.824699299 -0400 EDT m=+46.169038120

package sobjects

import (
	"fmt"
	"strings"
)

type UserLicense struct {
	BaseSObject
	CreatedDate             string `force:",omitempty"`
	Id                      string `force:",omitempty"`
	LastModifiedDate        string `force:",omitempty"`
	LicenseDefinitionKey    string `force:",omitempty"`
	MasterLabel             string `force:",omitempty"`
	Name                    string `force:",omitempty"`
	Status                  string `force:",omitempty"`
	SystemModstamp          string `force:",omitempty"`
	TotalLicenses           int    `force:",omitempty"`
	UsedLicenses            int    `force:",omitempty"`
	UsedLicensesLastUpdated string `force:",omitempty"`
}

func (t *UserLicense) ApiName() string {
	return "UserLicense"
}

func (t *UserLicense) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("UserLicense #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tLastModifiedDate: %v\n", t.LastModifiedDate))
	builder.WriteString(fmt.Sprintf("\tLicenseDefinitionKey: %v\n", t.LicenseDefinitionKey))
	builder.WriteString(fmt.Sprintf("\tMasterLabel: %v\n", t.MasterLabel))
	builder.WriteString(fmt.Sprintf("\tName: %v\n", t.Name))
	builder.WriteString(fmt.Sprintf("\tStatus: %v\n", t.Status))
	builder.WriteString(fmt.Sprintf("\tSystemModstamp: %v\n", t.SystemModstamp))
	builder.WriteString(fmt.Sprintf("\tTotalLicenses: %v\n", t.TotalLicenses))
	builder.WriteString(fmt.Sprintf("\tUsedLicenses: %v\n", t.UsedLicenses))
	builder.WriteString(fmt.Sprintf("\tUsedLicensesLastUpdated: %v\n", t.UsedLicensesLastUpdated))

	return builder.String()
}

type UserLicenseQueryResponse struct {
	BaseQuery
	Records []UserLicense `json:"Records" force:"records"`
}
