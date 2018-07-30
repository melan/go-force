// This file was generated for SObject FiscalYearSettings, API Version v43.0 at 2018-07-30 03:47:50.328372034 -0400 EDT m=+36.672354515

package sobjects

import (
	"fmt"
	"strings"
)

type FiscalYearSettings struct {
	BaseSObject
	Description        string `force:",omitempty"`
	EndDate            string `force:",omitempty"`
	Id                 string `force:",omitempty"`
	IsStandardYear     bool   `force:",omitempty"`
	Name               string `force:",omitempty"`
	PeriodId           string `force:",omitempty"`
	PeriodLabelScheme  string `force:",omitempty"`
	PeriodPrefix       string `force:",omitempty"`
	QuarterLabelScheme string `force:",omitempty"`
	QuarterPrefix      string `force:",omitempty"`
	StartDate          string `force:",omitempty"`
	SystemModstamp     string `force:",omitempty"`
	WeekLabelScheme    string `force:",omitempty"`
	WeekStartDay       int    `force:",omitempty"`
	YearType           string `force:",omitempty"`
}

func (t *FiscalYearSettings) ApiName() string {
	return "FiscalYearSettings"
}

func (t *FiscalYearSettings) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("FiscalYearSettings #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tDescription: %v\n", t.Description))
	builder.WriteString(fmt.Sprintf("\tEndDate: %v\n", t.EndDate))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tIsStandardYear: %v\n", t.IsStandardYear))
	builder.WriteString(fmt.Sprintf("\tName: %v\n", t.Name))
	builder.WriteString(fmt.Sprintf("\tPeriodId: %v\n", t.PeriodId))
	builder.WriteString(fmt.Sprintf("\tPeriodLabelScheme: %v\n", t.PeriodLabelScheme))
	builder.WriteString(fmt.Sprintf("\tPeriodPrefix: %v\n", t.PeriodPrefix))
	builder.WriteString(fmt.Sprintf("\tQuarterLabelScheme: %v\n", t.QuarterLabelScheme))
	builder.WriteString(fmt.Sprintf("\tQuarterPrefix: %v\n", t.QuarterPrefix))
	builder.WriteString(fmt.Sprintf("\tStartDate: %v\n", t.StartDate))
	builder.WriteString(fmt.Sprintf("\tSystemModstamp: %v\n", t.SystemModstamp))
	builder.WriteString(fmt.Sprintf("\tWeekLabelScheme: %v\n", t.WeekLabelScheme))
	builder.WriteString(fmt.Sprintf("\tWeekStartDay: %v\n", t.WeekStartDay))
	builder.WriteString(fmt.Sprintf("\tYearType: %v\n", t.YearType))

	return builder.String()
}

type FiscalYearSettingsQueryResponse struct {
	BaseQuery
	Records []FiscalYearSettings `json:"Records" force:"records"`
}
