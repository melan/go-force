// This file was generated for SObject FileSearchActivity, API Version v43.0 at 2018-07-30 03:48:06.735124511 -0400 EDT m=+53.079722639

package sobjects

import (
	"fmt"
	"strings"
)

type FileSearchActivity struct {
	BaseSObject
	AvgNumResults    float64 `force:",omitempty"`
	ClickRank        float64 `force:",omitempty"`
	CountQueries     int     `force:",omitempty"`
	CountUsers       int     `force:",omitempty"`
	CreatedById      string  `force:",omitempty"`
	CreatedDate      string  `force:",omitempty"`
	Id               string  `force:",omitempty"`
	IsDeleted        bool    `force:",omitempty"`
	LastModifiedById string  `force:",omitempty"`
	LastModifiedDate string  `force:",omitempty"`
	Name             string  `force:",omitempty"`
	Period           string  `force:",omitempty"`
	QueryDate        string  `force:",omitempty"`
	QueryLanguage    string  `force:",omitempty"`
	SearchTerm       string  `force:",omitempty"`
	SystemModstamp   string  `force:",omitempty"`
}

func (t *FileSearchActivity) ApiName() string {
	return "FileSearchActivity"
}

func (t *FileSearchActivity) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("FileSearchActivity #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tAvgNumResults: %v\n", t.AvgNumResults))
	builder.WriteString(fmt.Sprintf("\tClickRank: %v\n", t.ClickRank))
	builder.WriteString(fmt.Sprintf("\tCountQueries: %v\n", t.CountQueries))
	builder.WriteString(fmt.Sprintf("\tCountUsers: %v\n", t.CountUsers))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tIsDeleted: %v\n", t.IsDeleted))
	builder.WriteString(fmt.Sprintf("\tLastModifiedById: %v\n", t.LastModifiedById))
	builder.WriteString(fmt.Sprintf("\tLastModifiedDate: %v\n", t.LastModifiedDate))
	builder.WriteString(fmt.Sprintf("\tName: %v\n", t.Name))
	builder.WriteString(fmt.Sprintf("\tPeriod: %v\n", t.Period))
	builder.WriteString(fmt.Sprintf("\tQueryDate: %v\n", t.QueryDate))
	builder.WriteString(fmt.Sprintf("\tQueryLanguage: %v\n", t.QueryLanguage))
	builder.WriteString(fmt.Sprintf("\tSearchTerm: %v\n", t.SearchTerm))
	builder.WriteString(fmt.Sprintf("\tSystemModstamp: %v\n", t.SystemModstamp))

	return builder.String()
}

type FileSearchActivityQueryResponse struct {
	BaseQuery
	Records []FileSearchActivity `json:"Records" force:"records"`
}
