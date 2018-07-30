// This file was generated for SObject UserPreference, API Version v43.0 at 2018-07-30 03:48:00.634810759 -0400 EDT m=+46.979179979

package sobjects

import (
	"fmt"
	"strings"
)

type UserPreference struct {
	BaseSObject
	Id             string `force:",omitempty"`
	Preference     string `force:",omitempty"`
	SystemModstamp string `force:",omitempty"`
	UserId         string `force:",omitempty"`
	Value          string `force:",omitempty"`
}

func (t *UserPreference) ApiName() string {
	return "UserPreference"
}

func (t *UserPreference) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("UserPreference #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tPreference: %v\n", t.Preference))
	builder.WriteString(fmt.Sprintf("\tSystemModstamp: %v\n", t.SystemModstamp))
	builder.WriteString(fmt.Sprintf("\tUserId: %v\n", t.UserId))
	builder.WriteString(fmt.Sprintf("\tValue: %v\n", t.Value))

	return builder.String()
}

type UserPreferenceQueryResponse struct {
	BaseQuery
	Records []UserPreference `json:"Records" force:"records"`
}
