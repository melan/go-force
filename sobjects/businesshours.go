// This file was generated for SObject BusinessHours, API Version v43.0 at 2018-07-30 03:47:19.964149023 -0400 EDT m=+6.306992117

package sobjects

import (
	"fmt"
	"strings"
)

type BusinessHours struct {
	BaseSObject
	CreatedById        string `force:",omitempty"`
	CreatedDate        string `force:",omitempty"`
	FridayEndTime      string `force:",omitempty"`
	FridayStartTime    string `force:",omitempty"`
	Id                 string `force:",omitempty"`
	IsActive           bool   `force:",omitempty"`
	IsDefault          bool   `force:",omitempty"`
	LastModifiedById   string `force:",omitempty"`
	LastModifiedDate   string `force:",omitempty"`
	LastViewedDate     string `force:",omitempty"`
	MondayEndTime      string `force:",omitempty"`
	MondayStartTime    string `force:",omitempty"`
	Name               string `force:",omitempty"`
	SaturdayEndTime    string `force:",omitempty"`
	SaturdayStartTime  string `force:",omitempty"`
	SundayEndTime      string `force:",omitempty"`
	SundayStartTime    string `force:",omitempty"`
	SystemModstamp     string `force:",omitempty"`
	ThursdayEndTime    string `force:",omitempty"`
	ThursdayStartTime  string `force:",omitempty"`
	TimeZoneSidKey     string `force:",omitempty"`
	TuesdayEndTime     string `force:",omitempty"`
	TuesdayStartTime   string `force:",omitempty"`
	WednesdayEndTime   string `force:",omitempty"`
	WednesdayStartTime string `force:",omitempty"`
}

func (t *BusinessHours) ApiName() string {
	return "BusinessHours"
}

func (t *BusinessHours) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("BusinessHours #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tFridayEndTime: %v\n", t.FridayEndTime))
	builder.WriteString(fmt.Sprintf("\tFridayStartTime: %v\n", t.FridayStartTime))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tIsActive: %v\n", t.IsActive))
	builder.WriteString(fmt.Sprintf("\tIsDefault: %v\n", t.IsDefault))
	builder.WriteString(fmt.Sprintf("\tLastModifiedById: %v\n", t.LastModifiedById))
	builder.WriteString(fmt.Sprintf("\tLastModifiedDate: %v\n", t.LastModifiedDate))
	builder.WriteString(fmt.Sprintf("\tLastViewedDate: %v\n", t.LastViewedDate))
	builder.WriteString(fmt.Sprintf("\tMondayEndTime: %v\n", t.MondayEndTime))
	builder.WriteString(fmt.Sprintf("\tMondayStartTime: %v\n", t.MondayStartTime))
	builder.WriteString(fmt.Sprintf("\tName: %v\n", t.Name))
	builder.WriteString(fmt.Sprintf("\tSaturdayEndTime: %v\n", t.SaturdayEndTime))
	builder.WriteString(fmt.Sprintf("\tSaturdayStartTime: %v\n", t.SaturdayStartTime))
	builder.WriteString(fmt.Sprintf("\tSundayEndTime: %v\n", t.SundayEndTime))
	builder.WriteString(fmt.Sprintf("\tSundayStartTime: %v\n", t.SundayStartTime))
	builder.WriteString(fmt.Sprintf("\tSystemModstamp: %v\n", t.SystemModstamp))
	builder.WriteString(fmt.Sprintf("\tThursdayEndTime: %v\n", t.ThursdayEndTime))
	builder.WriteString(fmt.Sprintf("\tThursdayStartTime: %v\n", t.ThursdayStartTime))
	builder.WriteString(fmt.Sprintf("\tTimeZoneSidKey: %v\n", t.TimeZoneSidKey))
	builder.WriteString(fmt.Sprintf("\tTuesdayEndTime: %v\n", t.TuesdayEndTime))
	builder.WriteString(fmt.Sprintf("\tTuesdayStartTime: %v\n", t.TuesdayStartTime))
	builder.WriteString(fmt.Sprintf("\tWednesdayEndTime: %v\n", t.WednesdayEndTime))
	builder.WriteString(fmt.Sprintf("\tWednesdayStartTime: %v\n", t.WednesdayStartTime))

	return builder.String()
}

type BusinessHoursQueryResponse struct {
	BaseQuery
	Records []BusinessHours `json:"Records" force:"records"`
}
