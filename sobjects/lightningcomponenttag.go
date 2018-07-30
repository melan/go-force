// This file was generated for SObject LightningComponentTag, API Version v43.0 at 2018-07-30 03:47:47.898667139 -0400 EDT m=+34.242558448

package sobjects

import (
	"fmt"
	"strings"
)

type LightningComponentTag struct {
	BaseSObject
	CreatedById                string `force:",omitempty"`
	CreatedDate                string `force:",omitempty"`
	Id                         string `force:",omitempty"`
	IsDeleted                  bool   `force:",omitempty"`
	LastModifiedById           string `force:",omitempty"`
	LastModifiedDate           string `force:",omitempty"`
	LightningComponentBundleId string `force:",omitempty"`
	SystemModstamp             string `force:",omitempty"`
	Tag                        string `force:",omitempty"`
}

func (t *LightningComponentTag) ApiName() string {
	return "LightningComponentTag"
}

func (t *LightningComponentTag) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("LightningComponentTag #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tIsDeleted: %v\n", t.IsDeleted))
	builder.WriteString(fmt.Sprintf("\tLastModifiedById: %v\n", t.LastModifiedById))
	builder.WriteString(fmt.Sprintf("\tLastModifiedDate: %v\n", t.LastModifiedDate))
	builder.WriteString(fmt.Sprintf("\tLightningComponentBundleId: %v\n", t.LightningComponentBundleId))
	builder.WriteString(fmt.Sprintf("\tSystemModstamp: %v\n", t.SystemModstamp))
	builder.WriteString(fmt.Sprintf("\tTag: %v\n", t.Tag))

	return builder.String()
}

type LightningComponentTagQueryResponse struct {
	BaseQuery
	Records []LightningComponentTag `json:"Records" force:"records"`
}
