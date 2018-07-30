// This file was generated for SObject DatacloudCompany, API Version v43.0 at 2018-07-30 03:47:47.000733633 -0400 EDT m=+33.344591248

package sobjects

import (
	"fmt"
	"strings"
)

type DatacloudCompany struct {
	BaseSObject
	ActiveContacts             int     `force:",omitempty"`
	AnnualRevenue              string  `force:",omitempty"`
	City                       string  `force:",omitempty"`
	CompanyId                  string  `force:",omitempty"`
	Country                    string  `force:",omitempty"`
	CountryCode                string  `force:",omitempty"`
	Description                string  `force:",omitempty"`
	DunsNumber                 string  `force:",omitempty"`
	EmployeeQuantityGrowthRate float64 `force:",omitempty"`
	ExternalId                 string  `force:",omitempty"`
	Fax                        string  `force:",omitempty"`
	FortuneRank                int     `force:",omitempty"`
	FullAddress                string  `force:",omitempty"`
	Id                         string  `force:",omitempty"`
	IncludedInSnP500           string  `force:",omitempty"`
	Industry                   string  `force:",omitempty"`
	IsInCrm                    bool    `force:",omitempty"`
	IsInactive                 bool    `force:",omitempty"`
	IsOwned                    bool    `force:",omitempty"`
	NaicsCode                  string  `force:",omitempty"`
	NaicsDesc                  string  `force:",omitempty"`
	Name                       string  `force:",omitempty"`
	NumberOfEmployees          int     `force:",omitempty"`
	Ownership                  string  `force:",omitempty"`
	Phone                      string  `force:",omitempty"`
	PremisesMeasure            int     `force:",omitempty"`
	PremisesMeasureReliability string  `force:",omitempty"`
	PremisesMeasureUnit        string  `force:",omitempty"`
	PriorYearEmployees         int     `force:",omitempty"`
	PriorYearRevenue           string  `force:",omitempty"`
	SalesTurnoverGrowthRate    float64 `force:",omitempty"`
	Sic                        string  `force:",omitempty"`
	SicCodeDesc                string  `force:",omitempty"`
	SicDesc                    string  `force:",omitempty"`
	Site                       string  `force:",omitempty"`
	State                      string  `force:",omitempty"`
	StateCode                  string  `force:",omitempty"`
	Street                     string  `force:",omitempty"`
	TickerSymbol               string  `force:",omitempty"`
	TradeStyle                 string  `force:",omitempty"`
	UpdatedDate                string  `force:",omitempty"`
	Website                    string  `force:",omitempty"`
	YearStarted                string  `force:",omitempty"`
	Zip                        string  `force:",omitempty"`
}

func (t *DatacloudCompany) ApiName() string {
	return "DatacloudCompany"
}

func (t *DatacloudCompany) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("DatacloudCompany #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tActiveContacts: %v\n", t.ActiveContacts))
	builder.WriteString(fmt.Sprintf("\tAnnualRevenue: %v\n", t.AnnualRevenue))
	builder.WriteString(fmt.Sprintf("\tCity: %v\n", t.City))
	builder.WriteString(fmt.Sprintf("\tCompanyId: %v\n", t.CompanyId))
	builder.WriteString(fmt.Sprintf("\tCountry: %v\n", t.Country))
	builder.WriteString(fmt.Sprintf("\tCountryCode: %v\n", t.CountryCode))
	builder.WriteString(fmt.Sprintf("\tDescription: %v\n", t.Description))
	builder.WriteString(fmt.Sprintf("\tDunsNumber: %v\n", t.DunsNumber))
	builder.WriteString(fmt.Sprintf("\tEmployeeQuantityGrowthRate: %v\n", t.EmployeeQuantityGrowthRate))
	builder.WriteString(fmt.Sprintf("\tExternalId: %v\n", t.ExternalId))
	builder.WriteString(fmt.Sprintf("\tFax: %v\n", t.Fax))
	builder.WriteString(fmt.Sprintf("\tFortuneRank: %v\n", t.FortuneRank))
	builder.WriteString(fmt.Sprintf("\tFullAddress: %v\n", t.FullAddress))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tIncludedInSnP500: %v\n", t.IncludedInSnP500))
	builder.WriteString(fmt.Sprintf("\tIndustry: %v\n", t.Industry))
	builder.WriteString(fmt.Sprintf("\tIsInCrm: %v\n", t.IsInCrm))
	builder.WriteString(fmt.Sprintf("\tIsInactive: %v\n", t.IsInactive))
	builder.WriteString(fmt.Sprintf("\tIsOwned: %v\n", t.IsOwned))
	builder.WriteString(fmt.Sprintf("\tNaicsCode: %v\n", t.NaicsCode))
	builder.WriteString(fmt.Sprintf("\tNaicsDesc: %v\n", t.NaicsDesc))
	builder.WriteString(fmt.Sprintf("\tName: %v\n", t.Name))
	builder.WriteString(fmt.Sprintf("\tNumberOfEmployees: %v\n", t.NumberOfEmployees))
	builder.WriteString(fmt.Sprintf("\tOwnership: %v\n", t.Ownership))
	builder.WriteString(fmt.Sprintf("\tPhone: %v\n", t.Phone))
	builder.WriteString(fmt.Sprintf("\tPremisesMeasure: %v\n", t.PremisesMeasure))
	builder.WriteString(fmt.Sprintf("\tPremisesMeasureReliability: %v\n", t.PremisesMeasureReliability))
	builder.WriteString(fmt.Sprintf("\tPremisesMeasureUnit: %v\n", t.PremisesMeasureUnit))
	builder.WriteString(fmt.Sprintf("\tPriorYearEmployees: %v\n", t.PriorYearEmployees))
	builder.WriteString(fmt.Sprintf("\tPriorYearRevenue: %v\n", t.PriorYearRevenue))
	builder.WriteString(fmt.Sprintf("\tSalesTurnoverGrowthRate: %v\n", t.SalesTurnoverGrowthRate))
	builder.WriteString(fmt.Sprintf("\tSic: %v\n", t.Sic))
	builder.WriteString(fmt.Sprintf("\tSicCodeDesc: %v\n", t.SicCodeDesc))
	builder.WriteString(fmt.Sprintf("\tSicDesc: %v\n", t.SicDesc))
	builder.WriteString(fmt.Sprintf("\tSite: %v\n", t.Site))
	builder.WriteString(fmt.Sprintf("\tState: %v\n", t.State))
	builder.WriteString(fmt.Sprintf("\tStateCode: %v\n", t.StateCode))
	builder.WriteString(fmt.Sprintf("\tStreet: %v\n", t.Street))
	builder.WriteString(fmt.Sprintf("\tTickerSymbol: %v\n", t.TickerSymbol))
	builder.WriteString(fmt.Sprintf("\tTradeStyle: %v\n", t.TradeStyle))
	builder.WriteString(fmt.Sprintf("\tUpdatedDate: %v\n", t.UpdatedDate))
	builder.WriteString(fmt.Sprintf("\tWebsite: %v\n", t.Website))
	builder.WriteString(fmt.Sprintf("\tYearStarted: %v\n", t.YearStarted))
	builder.WriteString(fmt.Sprintf("\tZip: %v\n", t.Zip))

	return builder.String()
}

type DatacloudCompanyQueryResponse struct {
	BaseQuery
	Records []DatacloudCompany `json:"Records" force:"records"`
}
