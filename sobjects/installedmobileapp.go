// This file was generated for SObject InstalledMobileApp, API Version v43.0 at 2018-07-30 03:47:48.163291629 -0400 EDT m=+34.507192868

package sobjects

import (
	"fmt"
	"strings"
)

type InstalledMobileApp struct {
	BaseSObject
	ConnectedApplicationId string `force:",omitempty"`
	CreatedById            string `force:",omitempty"`
	CreatedDate            string `force:",omitempty"`
	Id                     string `force:",omitempty"`
	IsDeleted              bool   `force:",omitempty"`
	LastModifiedById       string `force:",omitempty"`
	LastModifiedDate       string `force:",omitempty"`
	Name                   string `force:",omitempty"`
	Status                 string `force:",omitempty"`
	SystemModstamp         string `force:",omitempty"`
	UserId                 string `force:",omitempty"`
	Version                string `force:",omitempty"`
}

func (t *InstalledMobileApp) ApiName() string {
	return "InstalledMobileApp"
}

func (t *InstalledMobileApp) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("InstalledMobileApp #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tConnectedApplicationId: %v\n", t.ConnectedApplicationId))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tIsDeleted: %v\n", t.IsDeleted))
	builder.WriteString(fmt.Sprintf("\tLastModifiedById: %v\n", t.LastModifiedById))
	builder.WriteString(fmt.Sprintf("\tLastModifiedDate: %v\n", t.LastModifiedDate))
	builder.WriteString(fmt.Sprintf("\tName: %v\n", t.Name))
	builder.WriteString(fmt.Sprintf("\tStatus: %v\n", t.Status))
	builder.WriteString(fmt.Sprintf("\tSystemModstamp: %v\n", t.SystemModstamp))
	builder.WriteString(fmt.Sprintf("\tUserId: %v\n", t.UserId))
	builder.WriteString(fmt.Sprintf("\tVersion: %v\n", t.Version))

	return builder.String()
}

type InstalledMobileAppQueryResponse struct {
	BaseQuery
	Records []InstalledMobileApp `json:"Records" force:"records"`
}
