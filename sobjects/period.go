// This file was generated for SObject Period, API Version v43.0 at 2018-07-30 03:48:09.412791435 -0400 EDT m=+55.757490039

package sobjects

import (
	"fmt"
	"strings"
)

type Period struct {
	BaseSObject
	EndDate              string `force:",omitempty"`
	FiscalYearSettingsId string `force:",omitempty"`
	FullyQualifiedLabel  string `force:",omitempty"`
	Id                   string `force:",omitempty"`
	IsForecastPeriod     bool   `force:",omitempty"`
	Number               int    `force:",omitempty"`
	PeriodLabel          string `force:",omitempty"`
	QuarterLabel         string `force:",omitempty"`
	StartDate            string `force:",omitempty"`
	SystemModstamp       string `force:",omitempty"`
	Type                 string `force:",omitempty"`
}

func (t *Period) ApiName() string {
	return "Period"
}

func (t *Period) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("Period #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tEndDate: %v\n", t.EndDate))
	builder.WriteString(fmt.Sprintf("\tFiscalYearSettingsId: %v\n", t.FiscalYearSettingsId))
	builder.WriteString(fmt.Sprintf("\tFullyQualifiedLabel: %v\n", t.FullyQualifiedLabel))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tIsForecastPeriod: %v\n", t.IsForecastPeriod))
	builder.WriteString(fmt.Sprintf("\tNumber: %v\n", t.Number))
	builder.WriteString(fmt.Sprintf("\tPeriodLabel: %v\n", t.PeriodLabel))
	builder.WriteString(fmt.Sprintf("\tQuarterLabel: %v\n", t.QuarterLabel))
	builder.WriteString(fmt.Sprintf("\tStartDate: %v\n", t.StartDate))
	builder.WriteString(fmt.Sprintf("\tSystemModstamp: %v\n", t.SystemModstamp))
	builder.WriteString(fmt.Sprintf("\tType: %v\n", t.Type))

	return builder.String()
}

type PeriodQueryResponse struct {
	BaseQuery
	Records []Period `json:"Records" force:"records"`
}
