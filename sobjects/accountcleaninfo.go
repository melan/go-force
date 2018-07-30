// This file was generated for SObject AccountCleanInfo, API Version v43.0 at 2018-07-30 03:47:56.659680394 -0400 EDT m=+43.003900451

package sobjects

import (
	"fmt"
	"strings"
)

type AccountCleanInfo struct {
	BaseSObject
	AccountId                         string   `force:",omitempty"`
	AccountSite                       string   `force:",omitempty"`
	Address                           *Address `force:",omitempty"`
	AnnualRevenue                     string   `force:",omitempty"`
	City                              string   `force:",omitempty"`
	CleanedByJob                      bool     `force:",omitempty"`
	CleanedByUser                     bool     `force:",omitempty"`
	CompanyName                       string   `force:",omitempty"`
	CompanyStatusDataDotCom           string   `force:",omitempty"`
	Country                           string   `force:",omitempty"`
	CreatedById                       string   `force:",omitempty"`
	CreatedDate                       string   `force:",omitempty"`
	DandBCompanyDunsNumber            string   `force:",omitempty"`
	DataDotComId                      string   `force:",omitempty"`
	Description                       string   `force:",omitempty"`
	DunsNumber                        string   `force:",omitempty"`
	DunsRightMatchConfidence          int      `force:",omitempty"`
	DunsRightMatchGrade               string   `force:",omitempty"`
	Fax                               string   `force:",omitempty"`
	GeocodeAccuracy                   string   `force:",omitempty"`
	Id                                string   `force:",omitempty"`
	Industry                          string   `force:",omitempty"`
	IsDeleted                         bool     `force:",omitempty"`
	IsDifferentAccountSite            bool     `force:",omitempty"`
	IsDifferentAnnualRevenue          bool     `force:",omitempty"`
	IsDifferentCity                   bool     `force:",omitempty"`
	IsDifferentCompanyName            bool     `force:",omitempty"`
	IsDifferentCountry                bool     `force:",omitempty"`
	IsDifferentCountryCode            bool     `force:",omitempty"`
	IsDifferentDandBCompanyDunsNumber bool     `force:",omitempty"`
	IsDifferentDescription            bool     `force:",omitempty"`
	IsDifferentDunsNumber             bool     `force:",omitempty"`
	IsDifferentFax                    bool     `force:",omitempty"`
	IsDifferentIndustry               bool     `force:",omitempty"`
	IsDifferentNaicsCode              bool     `force:",omitempty"`
	IsDifferentNaicsDescription       bool     `force:",omitempty"`
	IsDifferentNumberOfEmployees      bool     `force:",omitempty"`
	IsDifferentOwnership              bool     `force:",omitempty"`
	IsDifferentPhone                  bool     `force:",omitempty"`
	IsDifferentPostalCode             bool     `force:",omitempty"`
	IsDifferentSic                    bool     `force:",omitempty"`
	IsDifferentSicDescription         bool     `force:",omitempty"`
	IsDifferentState                  bool     `force:",omitempty"`
	IsDifferentStateCode              bool     `force:",omitempty"`
	IsDifferentStreet                 bool     `force:",omitempty"`
	IsDifferentTickerSymbol           bool     `force:",omitempty"`
	IsDifferentTradestyle             bool     `force:",omitempty"`
	IsDifferentWebsite                bool     `force:",omitempty"`
	IsDifferentYearStarted            bool     `force:",omitempty"`
	IsFlaggedWrongAccountSite         bool     `force:",omitempty"`
	IsFlaggedWrongAddress             bool     `force:",omitempty"`
	IsFlaggedWrongAnnualRevenue       bool     `force:",omitempty"`
	IsFlaggedWrongCompanyName         bool     `force:",omitempty"`
	IsFlaggedWrongDescription         bool     `force:",omitempty"`
	IsFlaggedWrongDunsNumber          bool     `force:",omitempty"`
	IsFlaggedWrongFax                 bool     `force:",omitempty"`
	IsFlaggedWrongIndustry            bool     `force:",omitempty"`
	IsFlaggedWrongNaicsCode           bool     `force:",omitempty"`
	IsFlaggedWrongNaicsDescription    bool     `force:",omitempty"`
	IsFlaggedWrongNumberOfEmployees   bool     `force:",omitempty"`
	IsFlaggedWrongOwnership           bool     `force:",omitempty"`
	IsFlaggedWrongPhone               bool     `force:",omitempty"`
	IsFlaggedWrongSic                 bool     `force:",omitempty"`
	IsFlaggedWrongSicDescription      bool     `force:",omitempty"`
	IsFlaggedWrongTickerSymbol        bool     `force:",omitempty"`
	IsFlaggedWrongTradestyle          bool     `force:",omitempty"`
	IsFlaggedWrongWebsite             bool     `force:",omitempty"`
	IsFlaggedWrongYearStarted         bool     `force:",omitempty"`
	IsInactive                        bool     `force:",omitempty"`
	IsReviewedAccountSite             bool     `force:",omitempty"`
	IsReviewedAddress                 bool     `force:",omitempty"`
	IsReviewedAnnualRevenue           bool     `force:",omitempty"`
	IsReviewedCompanyName             bool     `force:",omitempty"`
	IsReviewedDandBCompanyDunsNumber  bool     `force:",omitempty"`
	IsReviewedDescription             bool     `force:",omitempty"`
	IsReviewedDunsNumber              bool     `force:",omitempty"`
	IsReviewedFax                     bool     `force:",omitempty"`
	IsReviewedIndustry                bool     `force:",omitempty"`
	IsReviewedNaicsCode               bool     `force:",omitempty"`
	IsReviewedNaicsDescription        bool     `force:",omitempty"`
	IsReviewedNumberOfEmployees       bool     `force:",omitempty"`
	IsReviewedOwnership               bool     `force:",omitempty"`
	IsReviewedPhone                   bool     `force:",omitempty"`
	IsReviewedSic                     bool     `force:",omitempty"`
	IsReviewedSicDescription          bool     `force:",omitempty"`
	IsReviewedTickerSymbol            bool     `force:",omitempty"`
	IsReviewedTradestyle              bool     `force:",omitempty"`
	IsReviewedWebsite                 bool     `force:",omitempty"`
	IsReviewedYearStarted             bool     `force:",omitempty"`
	LastMatchedDate                   string   `force:",omitempty"`
	LastModifiedById                  string   `force:",omitempty"`
	LastModifiedDate                  string   `force:",omitempty"`
	LastStatusChangedById             string   `force:",omitempty"`
	LastStatusChangedDate             string   `force:",omitempty"`
	Latitude                          float64  `force:",omitempty"`
	Longitude                         float64  `force:",omitempty"`
	NaicsCode                         string   `force:",omitempty"`
	NaicsDescription                  string   `force:",omitempty"`
	Name                              string   `force:",omitempty"`
	NumberOfEmployees                 int      `force:",omitempty"`
	Ownership                         string   `force:",omitempty"`
	Phone                             string   `force:",omitempty"`
	PostalCode                        string   `force:",omitempty"`
	Sic                               string   `force:",omitempty"`
	SicDescription                    string   `force:",omitempty"`
	State                             string   `force:",omitempty"`
	Street                            string   `force:",omitempty"`
	SystemModstamp                    string   `force:",omitempty"`
	TickerSymbol                      string   `force:",omitempty"`
	Tradestyle                        string   `force:",omitempty"`
	Website                           string   `force:",omitempty"`
	YearStarted                       string   `force:",omitempty"`
}

func (t *AccountCleanInfo) ApiName() string {
	return "AccountCleanInfo"
}

func (t *AccountCleanInfo) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("AccountCleanInfo #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tAccountId: %v\n", t.AccountId))
	builder.WriteString(fmt.Sprintf("\tAccountSite: %v\n", t.AccountSite))
	builder.WriteString(fmt.Sprintf("\tAddress: %v\n", t.Address))
	builder.WriteString(fmt.Sprintf("\tAnnualRevenue: %v\n", t.AnnualRevenue))
	builder.WriteString(fmt.Sprintf("\tCity: %v\n", t.City))
	builder.WriteString(fmt.Sprintf("\tCleanedByJob: %v\n", t.CleanedByJob))
	builder.WriteString(fmt.Sprintf("\tCleanedByUser: %v\n", t.CleanedByUser))
	builder.WriteString(fmt.Sprintf("\tCompanyName: %v\n", t.CompanyName))
	builder.WriteString(fmt.Sprintf("\tCompanyStatusDataDotCom: %v\n", t.CompanyStatusDataDotCom))
	builder.WriteString(fmt.Sprintf("\tCountry: %v\n", t.Country))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tDandBCompanyDunsNumber: %v\n", t.DandBCompanyDunsNumber))
	builder.WriteString(fmt.Sprintf("\tDataDotComId: %v\n", t.DataDotComId))
	builder.WriteString(fmt.Sprintf("\tDescription: %v\n", t.Description))
	builder.WriteString(fmt.Sprintf("\tDunsNumber: %v\n", t.DunsNumber))
	builder.WriteString(fmt.Sprintf("\tDunsRightMatchConfidence: %v\n", t.DunsRightMatchConfidence))
	builder.WriteString(fmt.Sprintf("\tDunsRightMatchGrade: %v\n", t.DunsRightMatchGrade))
	builder.WriteString(fmt.Sprintf("\tFax: %v\n", t.Fax))
	builder.WriteString(fmt.Sprintf("\tGeocodeAccuracy: %v\n", t.GeocodeAccuracy))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tIndustry: %v\n", t.Industry))
	builder.WriteString(fmt.Sprintf("\tIsDeleted: %v\n", t.IsDeleted))
	builder.WriteString(fmt.Sprintf("\tIsDifferentAccountSite: %v\n", t.IsDifferentAccountSite))
	builder.WriteString(fmt.Sprintf("\tIsDifferentAnnualRevenue: %v\n", t.IsDifferentAnnualRevenue))
	builder.WriteString(fmt.Sprintf("\tIsDifferentCity: %v\n", t.IsDifferentCity))
	builder.WriteString(fmt.Sprintf("\tIsDifferentCompanyName: %v\n", t.IsDifferentCompanyName))
	builder.WriteString(fmt.Sprintf("\tIsDifferentCountry: %v\n", t.IsDifferentCountry))
	builder.WriteString(fmt.Sprintf("\tIsDifferentCountryCode: %v\n", t.IsDifferentCountryCode))
	builder.WriteString(fmt.Sprintf("\tIsDifferentDandBCompanyDunsNumber: %v\n", t.IsDifferentDandBCompanyDunsNumber))
	builder.WriteString(fmt.Sprintf("\tIsDifferentDescription: %v\n", t.IsDifferentDescription))
	builder.WriteString(fmt.Sprintf("\tIsDifferentDunsNumber: %v\n", t.IsDifferentDunsNumber))
	builder.WriteString(fmt.Sprintf("\tIsDifferentFax: %v\n", t.IsDifferentFax))
	builder.WriteString(fmt.Sprintf("\tIsDifferentIndustry: %v\n", t.IsDifferentIndustry))
	builder.WriteString(fmt.Sprintf("\tIsDifferentNaicsCode: %v\n", t.IsDifferentNaicsCode))
	builder.WriteString(fmt.Sprintf("\tIsDifferentNaicsDescription: %v\n", t.IsDifferentNaicsDescription))
	builder.WriteString(fmt.Sprintf("\tIsDifferentNumberOfEmployees: %v\n", t.IsDifferentNumberOfEmployees))
	builder.WriteString(fmt.Sprintf("\tIsDifferentOwnership: %v\n", t.IsDifferentOwnership))
	builder.WriteString(fmt.Sprintf("\tIsDifferentPhone: %v\n", t.IsDifferentPhone))
	builder.WriteString(fmt.Sprintf("\tIsDifferentPostalCode: %v\n", t.IsDifferentPostalCode))
	builder.WriteString(fmt.Sprintf("\tIsDifferentSic: %v\n", t.IsDifferentSic))
	builder.WriteString(fmt.Sprintf("\tIsDifferentSicDescription: %v\n", t.IsDifferentSicDescription))
	builder.WriteString(fmt.Sprintf("\tIsDifferentState: %v\n", t.IsDifferentState))
	builder.WriteString(fmt.Sprintf("\tIsDifferentStateCode: %v\n", t.IsDifferentStateCode))
	builder.WriteString(fmt.Sprintf("\tIsDifferentStreet: %v\n", t.IsDifferentStreet))
	builder.WriteString(fmt.Sprintf("\tIsDifferentTickerSymbol: %v\n", t.IsDifferentTickerSymbol))
	builder.WriteString(fmt.Sprintf("\tIsDifferentTradestyle: %v\n", t.IsDifferentTradestyle))
	builder.WriteString(fmt.Sprintf("\tIsDifferentWebsite: %v\n", t.IsDifferentWebsite))
	builder.WriteString(fmt.Sprintf("\tIsDifferentYearStarted: %v\n", t.IsDifferentYearStarted))
	builder.WriteString(fmt.Sprintf("\tIsFlaggedWrongAccountSite: %v\n", t.IsFlaggedWrongAccountSite))
	builder.WriteString(fmt.Sprintf("\tIsFlaggedWrongAddress: %v\n", t.IsFlaggedWrongAddress))
	builder.WriteString(fmt.Sprintf("\tIsFlaggedWrongAnnualRevenue: %v\n", t.IsFlaggedWrongAnnualRevenue))
	builder.WriteString(fmt.Sprintf("\tIsFlaggedWrongCompanyName: %v\n", t.IsFlaggedWrongCompanyName))
	builder.WriteString(fmt.Sprintf("\tIsFlaggedWrongDescription: %v\n", t.IsFlaggedWrongDescription))
	builder.WriteString(fmt.Sprintf("\tIsFlaggedWrongDunsNumber: %v\n", t.IsFlaggedWrongDunsNumber))
	builder.WriteString(fmt.Sprintf("\tIsFlaggedWrongFax: %v\n", t.IsFlaggedWrongFax))
	builder.WriteString(fmt.Sprintf("\tIsFlaggedWrongIndustry: %v\n", t.IsFlaggedWrongIndustry))
	builder.WriteString(fmt.Sprintf("\tIsFlaggedWrongNaicsCode: %v\n", t.IsFlaggedWrongNaicsCode))
	builder.WriteString(fmt.Sprintf("\tIsFlaggedWrongNaicsDescription: %v\n", t.IsFlaggedWrongNaicsDescription))
	builder.WriteString(fmt.Sprintf("\tIsFlaggedWrongNumberOfEmployees: %v\n", t.IsFlaggedWrongNumberOfEmployees))
	builder.WriteString(fmt.Sprintf("\tIsFlaggedWrongOwnership: %v\n", t.IsFlaggedWrongOwnership))
	builder.WriteString(fmt.Sprintf("\tIsFlaggedWrongPhone: %v\n", t.IsFlaggedWrongPhone))
	builder.WriteString(fmt.Sprintf("\tIsFlaggedWrongSic: %v\n", t.IsFlaggedWrongSic))
	builder.WriteString(fmt.Sprintf("\tIsFlaggedWrongSicDescription: %v\n", t.IsFlaggedWrongSicDescription))
	builder.WriteString(fmt.Sprintf("\tIsFlaggedWrongTickerSymbol: %v\n", t.IsFlaggedWrongTickerSymbol))
	builder.WriteString(fmt.Sprintf("\tIsFlaggedWrongTradestyle: %v\n", t.IsFlaggedWrongTradestyle))
	builder.WriteString(fmt.Sprintf("\tIsFlaggedWrongWebsite: %v\n", t.IsFlaggedWrongWebsite))
	builder.WriteString(fmt.Sprintf("\tIsFlaggedWrongYearStarted: %v\n", t.IsFlaggedWrongYearStarted))
	builder.WriteString(fmt.Sprintf("\tIsInactive: %v\n", t.IsInactive))
	builder.WriteString(fmt.Sprintf("\tIsReviewedAccountSite: %v\n", t.IsReviewedAccountSite))
	builder.WriteString(fmt.Sprintf("\tIsReviewedAddress: %v\n", t.IsReviewedAddress))
	builder.WriteString(fmt.Sprintf("\tIsReviewedAnnualRevenue: %v\n", t.IsReviewedAnnualRevenue))
	builder.WriteString(fmt.Sprintf("\tIsReviewedCompanyName: %v\n", t.IsReviewedCompanyName))
	builder.WriteString(fmt.Sprintf("\tIsReviewedDandBCompanyDunsNumber: %v\n", t.IsReviewedDandBCompanyDunsNumber))
	builder.WriteString(fmt.Sprintf("\tIsReviewedDescription: %v\n", t.IsReviewedDescription))
	builder.WriteString(fmt.Sprintf("\tIsReviewedDunsNumber: %v\n", t.IsReviewedDunsNumber))
	builder.WriteString(fmt.Sprintf("\tIsReviewedFax: %v\n", t.IsReviewedFax))
	builder.WriteString(fmt.Sprintf("\tIsReviewedIndustry: %v\n", t.IsReviewedIndustry))
	builder.WriteString(fmt.Sprintf("\tIsReviewedNaicsCode: %v\n", t.IsReviewedNaicsCode))
	builder.WriteString(fmt.Sprintf("\tIsReviewedNaicsDescription: %v\n", t.IsReviewedNaicsDescription))
	builder.WriteString(fmt.Sprintf("\tIsReviewedNumberOfEmployees: %v\n", t.IsReviewedNumberOfEmployees))
	builder.WriteString(fmt.Sprintf("\tIsReviewedOwnership: %v\n", t.IsReviewedOwnership))
	builder.WriteString(fmt.Sprintf("\tIsReviewedPhone: %v\n", t.IsReviewedPhone))
	builder.WriteString(fmt.Sprintf("\tIsReviewedSic: %v\n", t.IsReviewedSic))
	builder.WriteString(fmt.Sprintf("\tIsReviewedSicDescription: %v\n", t.IsReviewedSicDescription))
	builder.WriteString(fmt.Sprintf("\tIsReviewedTickerSymbol: %v\n", t.IsReviewedTickerSymbol))
	builder.WriteString(fmt.Sprintf("\tIsReviewedTradestyle: %v\n", t.IsReviewedTradestyle))
	builder.WriteString(fmt.Sprintf("\tIsReviewedWebsite: %v\n", t.IsReviewedWebsite))
	builder.WriteString(fmt.Sprintf("\tIsReviewedYearStarted: %v\n", t.IsReviewedYearStarted))
	builder.WriteString(fmt.Sprintf("\tLastMatchedDate: %v\n", t.LastMatchedDate))
	builder.WriteString(fmt.Sprintf("\tLastModifiedById: %v\n", t.LastModifiedById))
	builder.WriteString(fmt.Sprintf("\tLastModifiedDate: %v\n", t.LastModifiedDate))
	builder.WriteString(fmt.Sprintf("\tLastStatusChangedById: %v\n", t.LastStatusChangedById))
	builder.WriteString(fmt.Sprintf("\tLastStatusChangedDate: %v\n", t.LastStatusChangedDate))
	builder.WriteString(fmt.Sprintf("\tLatitude: %v\n", t.Latitude))
	builder.WriteString(fmt.Sprintf("\tLongitude: %v\n", t.Longitude))
	builder.WriteString(fmt.Sprintf("\tNaicsCode: %v\n", t.NaicsCode))
	builder.WriteString(fmt.Sprintf("\tNaicsDescription: %v\n", t.NaicsDescription))
	builder.WriteString(fmt.Sprintf("\tName: %v\n", t.Name))
	builder.WriteString(fmt.Sprintf("\tNumberOfEmployees: %v\n", t.NumberOfEmployees))
	builder.WriteString(fmt.Sprintf("\tOwnership: %v\n", t.Ownership))
	builder.WriteString(fmt.Sprintf("\tPhone: %v\n", t.Phone))
	builder.WriteString(fmt.Sprintf("\tPostalCode: %v\n", t.PostalCode))
	builder.WriteString(fmt.Sprintf("\tSic: %v\n", t.Sic))
	builder.WriteString(fmt.Sprintf("\tSicDescription: %v\n", t.SicDescription))
	builder.WriteString(fmt.Sprintf("\tState: %v\n", t.State))
	builder.WriteString(fmt.Sprintf("\tStreet: %v\n", t.Street))
	builder.WriteString(fmt.Sprintf("\tSystemModstamp: %v\n", t.SystemModstamp))
	builder.WriteString(fmt.Sprintf("\tTickerSymbol: %v\n", t.TickerSymbol))
	builder.WriteString(fmt.Sprintf("\tTradestyle: %v\n", t.Tradestyle))
	builder.WriteString(fmt.Sprintf("\tWebsite: %v\n", t.Website))
	builder.WriteString(fmt.Sprintf("\tYearStarted: %v\n", t.YearStarted))

	return builder.String()
}

type AccountCleanInfoQueryResponse struct {
	BaseQuery
	Records []AccountCleanInfo `json:"Records" force:"records"`
}
