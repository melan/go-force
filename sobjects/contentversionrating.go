// This file was generated for SObject ContentVersionRating, API Version v43.0 at 2018-07-30 03:47:57.072601509 -0400 EDT m=+43.416837061

package sobjects

import (
	"fmt"
	"strings"
)

type ContentVersionRating struct {
	BaseSObject
	ContentVersionId string `force:",omitempty"`
	Id               string `force:",omitempty"`
	LastModifiedDate string `force:",omitempty"`
	Rating           int    `force:",omitempty"`
	UserComment      string `force:",omitempty"`
	UserId           string `force:",omitempty"`
}

func (t *ContentVersionRating) ApiName() string {
	return "ContentVersionRating"
}

func (t *ContentVersionRating) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("ContentVersionRating #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tContentVersionId: %v\n", t.ContentVersionId))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tLastModifiedDate: %v\n", t.LastModifiedDate))
	builder.WriteString(fmt.Sprintf("\tRating: %v\n", t.Rating))
	builder.WriteString(fmt.Sprintf("\tUserComment: %v\n", t.UserComment))
	builder.WriteString(fmt.Sprintf("\tUserId: %v\n", t.UserId))

	return builder.String()
}

type ContentVersionRatingQueryResponse struct {
	BaseQuery
	Records []ContentVersionRating `json:"Records" force:"records"`
}
