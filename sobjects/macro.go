// This file was generated for SObject Macro, API Version v43.0 at 2018-07-30 03:47:22.725306276 -0400 EDT m=+9.068252981

package sobjects

import (
	"fmt"
	"strings"
)

type Macro struct {
	BaseSObject
	CreatedById          string `force:",omitempty"`
	CreatedDate          string `force:",omitempty"`
	Description          string `force:",omitempty"`
	Id                   string `force:",omitempty"`
	IsAlohaSupported     bool   `force:",omitempty"`
	IsDeleted            bool   `force:",omitempty"`
	IsLightningSupported bool   `force:",omitempty"`
	LastModifiedById     string `force:",omitempty"`
	LastModifiedDate     string `force:",omitempty"`
	LastReferencedDate   string `force:",omitempty"`
	LastViewedDate       string `force:",omitempty"`
	Name                 string `force:",omitempty"`
	OwnerId              string `force:",omitempty"`
	StartingContext      string `force:",omitempty"`
	SystemModstamp       string `force:",omitempty"`
}

func (t *Macro) ApiName() string {
	return "Macro"
}

func (t *Macro) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("Macro #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tDescription: %v\n", t.Description))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tIsAlohaSupported: %v\n", t.IsAlohaSupported))
	builder.WriteString(fmt.Sprintf("\tIsDeleted: %v\n", t.IsDeleted))
	builder.WriteString(fmt.Sprintf("\tIsLightningSupported: %v\n", t.IsLightningSupported))
	builder.WriteString(fmt.Sprintf("\tLastModifiedById: %v\n", t.LastModifiedById))
	builder.WriteString(fmt.Sprintf("\tLastModifiedDate: %v\n", t.LastModifiedDate))
	builder.WriteString(fmt.Sprintf("\tLastReferencedDate: %v\n", t.LastReferencedDate))
	builder.WriteString(fmt.Sprintf("\tLastViewedDate: %v\n", t.LastViewedDate))
	builder.WriteString(fmt.Sprintf("\tName: %v\n", t.Name))
	builder.WriteString(fmt.Sprintf("\tOwnerId: %v\n", t.OwnerId))
	builder.WriteString(fmt.Sprintf("\tStartingContext: %v\n", t.StartingContext))
	builder.WriteString(fmt.Sprintf("\tSystemModstamp: %v\n", t.SystemModstamp))

	return builder.String()
}

type MacroQueryResponse struct {
	BaseQuery
	Records []Macro `json:"Records" force:"records"`
}
