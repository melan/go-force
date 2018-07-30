// This file was generated for SObject LightningComponentResource, API Version v43.0 at 2018-07-30 03:48:05.787449968 -0400 EDT m=+52.132012535

package sobjects

import (
	"fmt"
	"strings"
)

type LightningComponentResource struct {
	BaseSObject
	CreatedById                string `force:",omitempty"`
	CreatedDate                string `force:",omitempty"`
	FilePath                   string `force:",omitempty"`
	Format                     string `force:",omitempty"`
	Id                         string `force:",omitempty"`
	IsDeleted                  bool   `force:",omitempty"`
	LastModifiedById           string `force:",omitempty"`
	LastModifiedDate           string `force:",omitempty"`
	LightningComponentBundleId string `force:",omitempty"`
	Source                     string `force:",omitempty"`
	SystemModstamp             string `force:",omitempty"`
}

func (t *LightningComponentResource) ApiName() string {
	return "LightningComponentResource"
}

func (t *LightningComponentResource) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("LightningComponentResource #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tFilePath: %v\n", t.FilePath))
	builder.WriteString(fmt.Sprintf("\tFormat: %v\n", t.Format))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tIsDeleted: %v\n", t.IsDeleted))
	builder.WriteString(fmt.Sprintf("\tLastModifiedById: %v\n", t.LastModifiedById))
	builder.WriteString(fmt.Sprintf("\tLastModifiedDate: %v\n", t.LastModifiedDate))
	builder.WriteString(fmt.Sprintf("\tLightningComponentBundleId: %v\n", t.LightningComponentBundleId))
	builder.WriteString(fmt.Sprintf("\tSource: %v\n", t.Source))
	builder.WriteString(fmt.Sprintf("\tSystemModstamp: %v\n", t.SystemModstamp))

	return builder.String()
}

type LightningComponentResourceQueryResponse struct {
	BaseQuery
	Records []LightningComponentResource `json:"Records" force:"records"`
}
