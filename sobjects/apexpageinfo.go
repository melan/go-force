// This file was generated for SObject ApexPageInfo, API Version v43.0 at 2018-07-30 03:47:33.232195726 -0400 EDT m=+19.575536691

package sobjects

import (
	"fmt"
	"strings"
)

type ApexPageInfo struct {
	BaseSObject
	ApexPageId         string  `force:",omitempty"`
	ApiVersion         float64 `force:",omitempty"`
	Description        string  `force:",omitempty"`
	DurableId          string  `force:",omitempty"`
	Id                 string  `force:",omitempty"`
	IsAvailableInTouch bool    `force:",omitempty"`
	IsShowHeader       string  `force:",omitempty"`
	MasterLabel        string  `force:",omitempty"`
	Name               string  `force:",omitempty"`
	NameSpacePrefix    string  `force:",omitempty"`
}

func (t *ApexPageInfo) ApiName() string {
	return "ApexPageInfo"
}

func (t *ApexPageInfo) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("ApexPageInfo #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tApexPageId: %v\n", t.ApexPageId))
	builder.WriteString(fmt.Sprintf("\tApiVersion: %v\n", t.ApiVersion))
	builder.WriteString(fmt.Sprintf("\tDescription: %v\n", t.Description))
	builder.WriteString(fmt.Sprintf("\tDurableId: %v\n", t.DurableId))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tIsAvailableInTouch: %v\n", t.IsAvailableInTouch))
	builder.WriteString(fmt.Sprintf("\tIsShowHeader: %v\n", t.IsShowHeader))
	builder.WriteString(fmt.Sprintf("\tMasterLabel: %v\n", t.MasterLabel))
	builder.WriteString(fmt.Sprintf("\tName: %v\n", t.Name))
	builder.WriteString(fmt.Sprintf("\tNameSpacePrefix: %v\n", t.NameSpacePrefix))

	return builder.String()
}

type ApexPageInfoQueryResponse struct {
	BaseQuery
	Records []ApexPageInfo `json:"Records" force:"records"`
}
