// This file was generated for SObject TodayGoal, API Version v43.0 at 2018-07-30 03:47:43.885945062 -0400 EDT m=+30.229685798

package sobjects

import (
	"fmt"
	"strings"
)

type TodayGoal struct {
	BaseSObject
	CreatedById      string `force:",omitempty"`
	CreatedDate      string `force:",omitempty"`
	Id               string `force:",omitempty"`
	IsDeleted        bool   `force:",omitempty"`
	LastModifiedById string `force:",omitempty"`
	LastModifiedDate string `force:",omitempty"`
	Name             string `force:",omitempty"`
	OwnerId          string `force:",omitempty"`
	SystemModstamp   string `force:",omitempty"`
	UserId           string `force:",omitempty"`
	Value            string `force:",omitempty"`
}

func (t *TodayGoal) ApiName() string {
	return "TodayGoal"
}

func (t *TodayGoal) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("TodayGoal #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tIsDeleted: %v\n", t.IsDeleted))
	builder.WriteString(fmt.Sprintf("\tLastModifiedById: %v\n", t.LastModifiedById))
	builder.WriteString(fmt.Sprintf("\tLastModifiedDate: %v\n", t.LastModifiedDate))
	builder.WriteString(fmt.Sprintf("\tName: %v\n", t.Name))
	builder.WriteString(fmt.Sprintf("\tOwnerId: %v\n", t.OwnerId))
	builder.WriteString(fmt.Sprintf("\tSystemModstamp: %v\n", t.SystemModstamp))
	builder.WriteString(fmt.Sprintf("\tUserId: %v\n", t.UserId))
	builder.WriteString(fmt.Sprintf("\tValue: %v\n", t.Value))

	return builder.String()
}

type TodayGoalQueryResponse struct {
	BaseQuery
	Records []TodayGoal `json:"Records" force:"records"`
}
