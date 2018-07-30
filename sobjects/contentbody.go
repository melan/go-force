// This file was generated for SObject ContentBody, API Version v43.0 at 2018-07-30 03:48:09.148505569 -0400 EDT m=+55.493194256

package sobjects

import (
	"fmt"
	"strings"
)

type ContentBody struct {
	BaseSObject
	Id string `force:",omitempty"`
}

func (t *ContentBody) ApiName() string {
	return "ContentBody"
}

func (t *ContentBody) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("ContentBody #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))

	return builder.String()
}

type ContentBodyQueryResponse struct {
	BaseQuery
	Records []ContentBody `json:"Records" force:"records"`
}
