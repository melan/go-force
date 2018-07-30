// This file was generated for SObject OrgLifecycleNotification, API Version v43.0 at 2018-07-30 03:47:44.131401806 -0400 EDT m=+30.475151753

package sobjects

import (
	"fmt"
	"strings"
)

type OrgLifecycleNotification struct {
	BaseSObject
	CreatedById          string `force:",omitempty"`
	CreatedDate          string `force:",omitempty"`
	LifecycleRequestId   string `force:",omitempty"`
	LifecycleRequestType string `force:",omitempty"`
	OrgId                string `force:",omitempty"`
	ReplayId             string `force:",omitempty"`
	Status               string `force:",omitempty"`
	StatusCode           string `force:",omitempty"`
}

func (t *OrgLifecycleNotification) ApiName() string {
	return "OrgLifecycleNotification"
}

func (t *OrgLifecycleNotification) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("OrgLifecycleNotification #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tLifecycleRequestId: %v\n", t.LifecycleRequestId))
	builder.WriteString(fmt.Sprintf("\tLifecycleRequestType: %v\n", t.LifecycleRequestType))
	builder.WriteString(fmt.Sprintf("\tOrgId: %v\n", t.OrgId))
	builder.WriteString(fmt.Sprintf("\tReplayId: %v\n", t.ReplayId))
	builder.WriteString(fmt.Sprintf("\tStatus: %v\n", t.Status))
	builder.WriteString(fmt.Sprintf("\tStatusCode: %v\n", t.StatusCode))

	return builder.String()
}

type OrgLifecycleNotificationQueryResponse struct {
	BaseQuery
	Records []OrgLifecycleNotification `json:"Records" force:"records"`
}
