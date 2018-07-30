// This file was generated for SObject DatacloudDandBCompany, API Version v43.0 at 2018-07-30 03:47:46.195324381 -0400 EDT m=+32.539151774

package sobjects

import (
	"fmt"
	"strings"
)

type DatacloudDandBCompany struct {
	BaseSObject
	City                         string  `force:",omitempty"`
	CompanyCurrencyIsoCode       string  `force:",omitempty"`
	CompanyId                    string  `force:",omitempty"`
	Country                      string  `force:",omitempty"`
	CountryAccessCode            string  `force:",omitempty"`
	CountryCode                  string  `force:",omitempty"`
	CurrencyCode                 string  `force:",omitempty"`
	Description                  string  `force:",omitempty"`
	DomesticUltimateBusinessName string  `force:",omitempty"`
	DomesticUltimateDunsNumber   string  `force:",omitempty"`
	DunsNumber                   string  `force:",omitempty"`
	EmployeeQuantityGrowthRate   float64 `force:",omitempty"`
	EmployeesHere                float64 `force:",omitempty"`
	EmployeesHereReliability     string  `force:",omitempty"`
	EmployeesTotal               float64 `force:",omitempty"`
	EmployeesTotalReliability    string  `force:",omitempty"`
	ExternalId                   string  `force:",omitempty"`
	FamilyMembers                int     `force:",omitempty"`
	Fax                          string  `force:",omitempty"`
	FifthNaics                   string  `force:",omitempty"`
	FifthNaicsDesc               string  `force:",omitempty"`
	FifthSic                     string  `force:",omitempty"`
	FifthSic8                    string  `force:",omitempty"`
	FifthSic8Desc                string  `force:",omitempty"`
	FifthSicDesc                 string  `force:",omitempty"`
	FipsMsaCode                  string  `force:",omitempty"`
	FipsMsaDesc                  string  `force:",omitempty"`
	FortuneRank                  int     `force:",omitempty"`
	FourthNaics                  string  `force:",omitempty"`
	FourthNaicsDesc              string  `force:",omitempty"`
	FourthSic                    string  `force:",omitempty"`
	FourthSic8                   string  `force:",omitempty"`
	FourthSic8Desc               string  `force:",omitempty"`
	FourthSicDesc                string  `force:",omitempty"`
	FullAddress                  string  `force:",omitempty"`
	GeoCodeAccuracy              string  `force:",omitempty"`
	GlobalUltimateBusinessName   string  `force:",omitempty"`
	GlobalUltimateDunsNumber     string  `force:",omitempty"`
	GlobalUltimateTotalEmployees float64 `force:",omitempty"`
	Id                           string  `force:",omitempty"`
	ImportExportAgent            string  `force:",omitempty"`
	IncludedInSnP500             string  `force:",omitempty"`
	Industry                     string  `force:",omitempty"`
	IsInCrm                      bool    `force:",omitempty"`
	IsOwned                      bool    `force:",omitempty"`
	IsParent                     bool    `force:",omitempty"`
	Latitude                     string  `force:",omitempty"`
	LegalStatus                  string  `force:",omitempty"`
	LocationStatus               string  `force:",omitempty"`
	Longitude                    string  `force:",omitempty"`
	MailingCity                  string  `force:",omitempty"`
	MailingCountry               string  `force:",omitempty"`
	MailingCountryCode           string  `force:",omitempty"`
	MailingState                 string  `force:",omitempty"`
	MailingStateCode             string  `force:",omitempty"`
	MailingStreet                string  `force:",omitempty"`
	MailingZip                   string  `force:",omitempty"`
	MarketingPreScreen           string  `force:",omitempty"`
	MarketingSegmentationCluster string  `force:",omitempty"`
	MinorityOwned                string  `force:",omitempty"`
	Name                         string  `force:",omitempty"`
	NationalId                   string  `force:",omitempty"`
	NationalIdType               string  `force:",omitempty"`
	OutOfBusiness                string  `force:",omitempty"`
	OwnOrRent                    string  `force:",omitempty"`
	ParentOrHqBusinessName       string  `force:",omitempty"`
	ParentOrHqDunsNumber         string  `force:",omitempty"`
	Phone                        string  `force:",omitempty"`
	PremisesMeasure              int     `force:",omitempty"`
	PremisesMeasureReliability   string  `force:",omitempty"`
	PremisesMeasureUnit          string  `force:",omitempty"`
	PrimaryNaics                 string  `force:",omitempty"`
	PrimaryNaicsDesc             string  `force:",omitempty"`
	PrimarySic                   string  `force:",omitempty"`
	PrimarySic8                  string  `force:",omitempty"`
	PrimarySic8Desc              string  `force:",omitempty"`
	PrimarySicDesc               string  `force:",omitempty"`
	PriorYearEmployees           int     `force:",omitempty"`
	PriorYearRevenue             float64 `force:",omitempty"`
	PublicIndicator              string  `force:",omitempty"`
	Revenue                      float64 `force:",omitempty"`
	SalesTurnoverGrowthRate      float64 `force:",omitempty"`
	SalesVolume                  float64 `force:",omitempty"`
	SalesVolumeReliability       string  `force:",omitempty"`
	SecondNaics                  string  `force:",omitempty"`
	SecondNaicsDesc              string  `force:",omitempty"`
	SecondSic                    string  `force:",omitempty"`
	SecondSic8                   string  `force:",omitempty"`
	SecondSic8Desc               string  `force:",omitempty"`
	SecondSicDesc                string  `force:",omitempty"`
	SicCodeDesc                  string  `force:",omitempty"`
	SixthNaics                   string  `force:",omitempty"`
	SixthNaicsDesc               string  `force:",omitempty"`
	SixthSic                     string  `force:",omitempty"`
	SixthSic8                    string  `force:",omitempty"`
	SixthSic8Desc                string  `force:",omitempty"`
	SixthSicDesc                 string  `force:",omitempty"`
	SmallBusiness                string  `force:",omitempty"`
	State                        string  `force:",omitempty"`
	StateCode                    string  `force:",omitempty"`
	StockExchange                string  `force:",omitempty"`
	StockSymbol                  string  `force:",omitempty"`
	Street                       string  `force:",omitempty"`
	Subsidiary                   string  `force:",omitempty"`
	ThirdNaics                   string  `force:",omitempty"`
	ThirdNaicsDesc               string  `force:",omitempty"`
	ThirdSic                     string  `force:",omitempty"`
	ThirdSic8                    string  `force:",omitempty"`
	ThirdSic8Desc                string  `force:",omitempty"`
	ThirdSicDesc                 string  `force:",omitempty"`
	TradeStyle1                  string  `force:",omitempty"`
	TradeStyle2                  string  `force:",omitempty"`
	TradeStyle3                  string  `force:",omitempty"`
	TradeStyle4                  string  `force:",omitempty"`
	TradeStyle5                  string  `force:",omitempty"`
	URL                          string  `force:",omitempty"`
	UsTaxId                      string  `force:",omitempty"`
	WomenOwned                   string  `force:",omitempty"`
	YearStarted                  string  `force:",omitempty"`
	Zip                          string  `force:",omitempty"`
}

func (t *DatacloudDandBCompany) ApiName() string {
	return "DatacloudDandBCompany"
}

func (t *DatacloudDandBCompany) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("DatacloudDandBCompany #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tCity: %v\n", t.City))
	builder.WriteString(fmt.Sprintf("\tCompanyCurrencyIsoCode: %v\n", t.CompanyCurrencyIsoCode))
	builder.WriteString(fmt.Sprintf("\tCompanyId: %v\n", t.CompanyId))
	builder.WriteString(fmt.Sprintf("\tCountry: %v\n", t.Country))
	builder.WriteString(fmt.Sprintf("\tCountryAccessCode: %v\n", t.CountryAccessCode))
	builder.WriteString(fmt.Sprintf("\tCountryCode: %v\n", t.CountryCode))
	builder.WriteString(fmt.Sprintf("\tCurrencyCode: %v\n", t.CurrencyCode))
	builder.WriteString(fmt.Sprintf("\tDescription: %v\n", t.Description))
	builder.WriteString(fmt.Sprintf("\tDomesticUltimateBusinessName: %v\n", t.DomesticUltimateBusinessName))
	builder.WriteString(fmt.Sprintf("\tDomesticUltimateDunsNumber: %v\n", t.DomesticUltimateDunsNumber))
	builder.WriteString(fmt.Sprintf("\tDunsNumber: %v\n", t.DunsNumber))
	builder.WriteString(fmt.Sprintf("\tEmployeeQuantityGrowthRate: %v\n", t.EmployeeQuantityGrowthRate))
	builder.WriteString(fmt.Sprintf("\tEmployeesHere: %v\n", t.EmployeesHere))
	builder.WriteString(fmt.Sprintf("\tEmployeesHereReliability: %v\n", t.EmployeesHereReliability))
	builder.WriteString(fmt.Sprintf("\tEmployeesTotal: %v\n", t.EmployeesTotal))
	builder.WriteString(fmt.Sprintf("\tEmployeesTotalReliability: %v\n", t.EmployeesTotalReliability))
	builder.WriteString(fmt.Sprintf("\tExternalId: %v\n", t.ExternalId))
	builder.WriteString(fmt.Sprintf("\tFamilyMembers: %v\n", t.FamilyMembers))
	builder.WriteString(fmt.Sprintf("\tFax: %v\n", t.Fax))
	builder.WriteString(fmt.Sprintf("\tFifthNaics: %v\n", t.FifthNaics))
	builder.WriteString(fmt.Sprintf("\tFifthNaicsDesc: %v\n", t.FifthNaicsDesc))
	builder.WriteString(fmt.Sprintf("\tFifthSic: %v\n", t.FifthSic))
	builder.WriteString(fmt.Sprintf("\tFifthSic8: %v\n", t.FifthSic8))
	builder.WriteString(fmt.Sprintf("\tFifthSic8Desc: %v\n", t.FifthSic8Desc))
	builder.WriteString(fmt.Sprintf("\tFifthSicDesc: %v\n", t.FifthSicDesc))
	builder.WriteString(fmt.Sprintf("\tFipsMsaCode: %v\n", t.FipsMsaCode))
	builder.WriteString(fmt.Sprintf("\tFipsMsaDesc: %v\n", t.FipsMsaDesc))
	builder.WriteString(fmt.Sprintf("\tFortuneRank: %v\n", t.FortuneRank))
	builder.WriteString(fmt.Sprintf("\tFourthNaics: %v\n", t.FourthNaics))
	builder.WriteString(fmt.Sprintf("\tFourthNaicsDesc: %v\n", t.FourthNaicsDesc))
	builder.WriteString(fmt.Sprintf("\tFourthSic: %v\n", t.FourthSic))
	builder.WriteString(fmt.Sprintf("\tFourthSic8: %v\n", t.FourthSic8))
	builder.WriteString(fmt.Sprintf("\tFourthSic8Desc: %v\n", t.FourthSic8Desc))
	builder.WriteString(fmt.Sprintf("\tFourthSicDesc: %v\n", t.FourthSicDesc))
	builder.WriteString(fmt.Sprintf("\tFullAddress: %v\n", t.FullAddress))
	builder.WriteString(fmt.Sprintf("\tGeoCodeAccuracy: %v\n", t.GeoCodeAccuracy))
	builder.WriteString(fmt.Sprintf("\tGlobalUltimateBusinessName: %v\n", t.GlobalUltimateBusinessName))
	builder.WriteString(fmt.Sprintf("\tGlobalUltimateDunsNumber: %v\n", t.GlobalUltimateDunsNumber))
	builder.WriteString(fmt.Sprintf("\tGlobalUltimateTotalEmployees: %v\n", t.GlobalUltimateTotalEmployees))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tImportExportAgent: %v\n", t.ImportExportAgent))
	builder.WriteString(fmt.Sprintf("\tIncludedInSnP500: %v\n", t.IncludedInSnP500))
	builder.WriteString(fmt.Sprintf("\tIndustry: %v\n", t.Industry))
	builder.WriteString(fmt.Sprintf("\tIsInCrm: %v\n", t.IsInCrm))
	builder.WriteString(fmt.Sprintf("\tIsOwned: %v\n", t.IsOwned))
	builder.WriteString(fmt.Sprintf("\tIsParent: %v\n", t.IsParent))
	builder.WriteString(fmt.Sprintf("\tLatitude: %v\n", t.Latitude))
	builder.WriteString(fmt.Sprintf("\tLegalStatus: %v\n", t.LegalStatus))
	builder.WriteString(fmt.Sprintf("\tLocationStatus: %v\n", t.LocationStatus))
	builder.WriteString(fmt.Sprintf("\tLongitude: %v\n", t.Longitude))
	builder.WriteString(fmt.Sprintf("\tMailingCity: %v\n", t.MailingCity))
	builder.WriteString(fmt.Sprintf("\tMailingCountry: %v\n", t.MailingCountry))
	builder.WriteString(fmt.Sprintf("\tMailingCountryCode: %v\n", t.MailingCountryCode))
	builder.WriteString(fmt.Sprintf("\tMailingState: %v\n", t.MailingState))
	builder.WriteString(fmt.Sprintf("\tMailingStateCode: %v\n", t.MailingStateCode))
	builder.WriteString(fmt.Sprintf("\tMailingStreet: %v\n", t.MailingStreet))
	builder.WriteString(fmt.Sprintf("\tMailingZip: %v\n", t.MailingZip))
	builder.WriteString(fmt.Sprintf("\tMarketingPreScreen: %v\n", t.MarketingPreScreen))
	builder.WriteString(fmt.Sprintf("\tMarketingSegmentationCluster: %v\n", t.MarketingSegmentationCluster))
	builder.WriteString(fmt.Sprintf("\tMinorityOwned: %v\n", t.MinorityOwned))
	builder.WriteString(fmt.Sprintf("\tName: %v\n", t.Name))
	builder.WriteString(fmt.Sprintf("\tNationalId: %v\n", t.NationalId))
	builder.WriteString(fmt.Sprintf("\tNationalIdType: %v\n", t.NationalIdType))
	builder.WriteString(fmt.Sprintf("\tOutOfBusiness: %v\n", t.OutOfBusiness))
	builder.WriteString(fmt.Sprintf("\tOwnOrRent: %v\n", t.OwnOrRent))
	builder.WriteString(fmt.Sprintf("\tParentOrHqBusinessName: %v\n", t.ParentOrHqBusinessName))
	builder.WriteString(fmt.Sprintf("\tParentOrHqDunsNumber: %v\n", t.ParentOrHqDunsNumber))
	builder.WriteString(fmt.Sprintf("\tPhone: %v\n", t.Phone))
	builder.WriteString(fmt.Sprintf("\tPremisesMeasure: %v\n", t.PremisesMeasure))
	builder.WriteString(fmt.Sprintf("\tPremisesMeasureReliability: %v\n", t.PremisesMeasureReliability))
	builder.WriteString(fmt.Sprintf("\tPremisesMeasureUnit: %v\n", t.PremisesMeasureUnit))
	builder.WriteString(fmt.Sprintf("\tPrimaryNaics: %v\n", t.PrimaryNaics))
	builder.WriteString(fmt.Sprintf("\tPrimaryNaicsDesc: %v\n", t.PrimaryNaicsDesc))
	builder.WriteString(fmt.Sprintf("\tPrimarySic: %v\n", t.PrimarySic))
	builder.WriteString(fmt.Sprintf("\tPrimarySic8: %v\n", t.PrimarySic8))
	builder.WriteString(fmt.Sprintf("\tPrimarySic8Desc: %v\n", t.PrimarySic8Desc))
	builder.WriteString(fmt.Sprintf("\tPrimarySicDesc: %v\n", t.PrimarySicDesc))
	builder.WriteString(fmt.Sprintf("\tPriorYearEmployees: %v\n", t.PriorYearEmployees))
	builder.WriteString(fmt.Sprintf("\tPriorYearRevenue: %v\n", t.PriorYearRevenue))
	builder.WriteString(fmt.Sprintf("\tPublicIndicator: %v\n", t.PublicIndicator))
	builder.WriteString(fmt.Sprintf("\tRevenue: %v\n", t.Revenue))
	builder.WriteString(fmt.Sprintf("\tSalesTurnoverGrowthRate: %v\n", t.SalesTurnoverGrowthRate))
	builder.WriteString(fmt.Sprintf("\tSalesVolume: %v\n", t.SalesVolume))
	builder.WriteString(fmt.Sprintf("\tSalesVolumeReliability: %v\n", t.SalesVolumeReliability))
	builder.WriteString(fmt.Sprintf("\tSecondNaics: %v\n", t.SecondNaics))
	builder.WriteString(fmt.Sprintf("\tSecondNaicsDesc: %v\n", t.SecondNaicsDesc))
	builder.WriteString(fmt.Sprintf("\tSecondSic: %v\n", t.SecondSic))
	builder.WriteString(fmt.Sprintf("\tSecondSic8: %v\n", t.SecondSic8))
	builder.WriteString(fmt.Sprintf("\tSecondSic8Desc: %v\n", t.SecondSic8Desc))
	builder.WriteString(fmt.Sprintf("\tSecondSicDesc: %v\n", t.SecondSicDesc))
	builder.WriteString(fmt.Sprintf("\tSicCodeDesc: %v\n", t.SicCodeDesc))
	builder.WriteString(fmt.Sprintf("\tSixthNaics: %v\n", t.SixthNaics))
	builder.WriteString(fmt.Sprintf("\tSixthNaicsDesc: %v\n", t.SixthNaicsDesc))
	builder.WriteString(fmt.Sprintf("\tSixthSic: %v\n", t.SixthSic))
	builder.WriteString(fmt.Sprintf("\tSixthSic8: %v\n", t.SixthSic8))
	builder.WriteString(fmt.Sprintf("\tSixthSic8Desc: %v\n", t.SixthSic8Desc))
	builder.WriteString(fmt.Sprintf("\tSixthSicDesc: %v\n", t.SixthSicDesc))
	builder.WriteString(fmt.Sprintf("\tSmallBusiness: %v\n", t.SmallBusiness))
	builder.WriteString(fmt.Sprintf("\tState: %v\n", t.State))
	builder.WriteString(fmt.Sprintf("\tStateCode: %v\n", t.StateCode))
	builder.WriteString(fmt.Sprintf("\tStockExchange: %v\n", t.StockExchange))
	builder.WriteString(fmt.Sprintf("\tStockSymbol: %v\n", t.StockSymbol))
	builder.WriteString(fmt.Sprintf("\tStreet: %v\n", t.Street))
	builder.WriteString(fmt.Sprintf("\tSubsidiary: %v\n", t.Subsidiary))
	builder.WriteString(fmt.Sprintf("\tThirdNaics: %v\n", t.ThirdNaics))
	builder.WriteString(fmt.Sprintf("\tThirdNaicsDesc: %v\n", t.ThirdNaicsDesc))
	builder.WriteString(fmt.Sprintf("\tThirdSic: %v\n", t.ThirdSic))
	builder.WriteString(fmt.Sprintf("\tThirdSic8: %v\n", t.ThirdSic8))
	builder.WriteString(fmt.Sprintf("\tThirdSic8Desc: %v\n", t.ThirdSic8Desc))
	builder.WriteString(fmt.Sprintf("\tThirdSicDesc: %v\n", t.ThirdSicDesc))
	builder.WriteString(fmt.Sprintf("\tTradeStyle1: %v\n", t.TradeStyle1))
	builder.WriteString(fmt.Sprintf("\tTradeStyle2: %v\n", t.TradeStyle2))
	builder.WriteString(fmt.Sprintf("\tTradeStyle3: %v\n", t.TradeStyle3))
	builder.WriteString(fmt.Sprintf("\tTradeStyle4: %v\n", t.TradeStyle4))
	builder.WriteString(fmt.Sprintf("\tTradeStyle5: %v\n", t.TradeStyle5))
	builder.WriteString(fmt.Sprintf("\tURL: %v\n", t.URL))
	builder.WriteString(fmt.Sprintf("\tUsTaxId: %v\n", t.UsTaxId))
	builder.WriteString(fmt.Sprintf("\tWomenOwned: %v\n", t.WomenOwned))
	builder.WriteString(fmt.Sprintf("\tYearStarted: %v\n", t.YearStarted))
	builder.WriteString(fmt.Sprintf("\tZip: %v\n", t.Zip))

	return builder.String()
}

type DatacloudDandBCompanyQueryResponse struct {
	BaseQuery
	Records []DatacloudDandBCompany `json:"Records" force:"records"`
}
