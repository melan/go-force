// This file was generated for SObject ForecastShare, API Version v43.0 at 2018-07-30 03:47:37.018617875 -0400 EDT m=+23.362100922

package sobjects

import (
	"fmt"
	"strings"
)

type ForecastShare struct {
	BaseSObject
	AccessLevel      string `force:",omitempty"`
	CanSubmit        bool   `force:",omitempty"`
	Id               string `force:",omitempty"`
	LastModifiedById string `force:",omitempty"`
	LastModifiedDate string `force:",omitempty"`
	RowCause         string `force:",omitempty"`
	UserOrGroupId    string `force:",omitempty"`
	UserRoleId       string `force:",omitempty"`
}

func (t *ForecastShare) ApiName() string {
	return "ForecastShare"
}

func (t *ForecastShare) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("ForecastShare #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tAccessLevel: %v\n", t.AccessLevel))
	builder.WriteString(fmt.Sprintf("\tCanSubmit: %v\n", t.CanSubmit))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tLastModifiedById: %v\n", t.LastModifiedById))
	builder.WriteString(fmt.Sprintf("\tLastModifiedDate: %v\n", t.LastModifiedDate))
	builder.WriteString(fmt.Sprintf("\tRowCause: %v\n", t.RowCause))
	builder.WriteString(fmt.Sprintf("\tUserOrGroupId: %v\n", t.UserOrGroupId))
	builder.WriteString(fmt.Sprintf("\tUserRoleId: %v\n", t.UserRoleId))

	return builder.String()
}

type ForecastShareQueryResponse struct {
	BaseQuery
	Records []ForecastShare `json:"Records" force:"records"`
}
