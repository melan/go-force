// This file was generated for SObject LeadCleanInfo, API Version v43.0 at 2018-07-30 03:47:40.614308915 -0400 EDT m=+26.957926886

package sobjects

import (
	"fmt"
	"strings"
)

type LeadCleanInfo struct {
	BaseSObject
	Address                           *Address `force:",omitempty"`
	AnnualRevenue                     string   `force:",omitempty"`
	City                              string   `force:",omitempty"`
	CleanedByJob                      bool     `force:",omitempty"`
	CleanedByUser                     bool     `force:",omitempty"`
	CompanyDunsNumber                 string   `force:",omitempty"`
	CompanyName                       string   `force:",omitempty"`
	ContactStatusDataDotCom           string   `force:",omitempty"`
	Country                           string   `force:",omitempty"`
	CreatedById                       string   `force:",omitempty"`
	CreatedDate                       string   `force:",omitempty"`
	DandBCompanyDunsNumber            string   `force:",omitempty"`
	DataDotComCompanyId               string   `force:",omitempty"`
	DataDotComId                      string   `force:",omitempty"`
	Email                             string   `force:",omitempty"`
	FirstName                         string   `force:",omitempty"`
	GeocodeAccuracy                   string   `force:",omitempty"`
	Id                                string   `force:",omitempty"`
	Industry                          string   `force:",omitempty"`
	IsDeleted                         bool     `force:",omitempty"`
	IsDifferentAnnualRevenue          bool     `force:",omitempty"`
	IsDifferentCity                   bool     `force:",omitempty"`
	IsDifferentCompanyDunsNumber      bool     `force:",omitempty"`
	IsDifferentCompanyName            bool     `force:",omitempty"`
	IsDifferentCountry                bool     `force:",omitempty"`
	IsDifferentCountryCode            bool     `force:",omitempty"`
	IsDifferentDandBCompanyDunsNumber bool     `force:",omitempty"`
	IsDifferentEmail                  bool     `force:",omitempty"`
	IsDifferentFirstName              bool     `force:",omitempty"`
	IsDifferentIndustry               bool     `force:",omitempty"`
	IsDifferentLastName               bool     `force:",omitempty"`
	IsDifferentNumberOfEmployees      bool     `force:",omitempty"`
	IsDifferentPhone                  bool     `force:",omitempty"`
	IsDifferentPostalCode             bool     `force:",omitempty"`
	IsDifferentState                  bool     `force:",omitempty"`
	IsDifferentStateCode              bool     `force:",omitempty"`
	IsDifferentStreet                 bool     `force:",omitempty"`
	IsDifferentTitle                  bool     `force:",omitempty"`
	IsFlaggedWrongAddress             bool     `force:",omitempty"`
	IsFlaggedWrongAnnualRevenue       bool     `force:",omitempty"`
	IsFlaggedWrongCompanyDunsNumber   bool     `force:",omitempty"`
	IsFlaggedWrongCompanyName         bool     `force:",omitempty"`
	IsFlaggedWrongEmail               bool     `force:",omitempty"`
	IsFlaggedWrongIndustry            bool     `force:",omitempty"`
	IsFlaggedWrongName                bool     `force:",omitempty"`
	IsFlaggedWrongNumberOfEmployees   bool     `force:",omitempty"`
	IsFlaggedWrongPhone               bool     `force:",omitempty"`
	IsFlaggedWrongTitle               bool     `force:",omitempty"`
	IsInactive                        bool     `force:",omitempty"`
	IsReviewedAddress                 bool     `force:",omitempty"`
	IsReviewedAnnualRevenue           bool     `force:",omitempty"`
	IsReviewedCompanyDunsNumber       bool     `force:",omitempty"`
	IsReviewedCompanyName             bool     `force:",omitempty"`
	IsReviewedDandBCompanyDunsNumber  bool     `force:",omitempty"`
	IsReviewedEmail                   bool     `force:",omitempty"`
	IsReviewedIndustry                bool     `force:",omitempty"`
	IsReviewedName                    bool     `force:",omitempty"`
	IsReviewedNumberOfEmployees       bool     `force:",omitempty"`
	IsReviewedPhone                   bool     `force:",omitempty"`
	IsReviewedTitle                   bool     `force:",omitempty"`
	LastMatchedDate                   string   `force:",omitempty"`
	LastModifiedById                  string   `force:",omitempty"`
	LastModifiedDate                  string   `force:",omitempty"`
	LastName                          string   `force:",omitempty"`
	LastStatusChangedById             string   `force:",omitempty"`
	LastStatusChangedDate             string   `force:",omitempty"`
	Latitude                          float64  `force:",omitempty"`
	LeadId                            string   `force:",omitempty"`
	Longitude                         float64  `force:",omitempty"`
	Name                              string   `force:",omitempty"`
	NumberOfEmployees                 int      `force:",omitempty"`
	Phone                             string   `force:",omitempty"`
	PostalCode                        string   `force:",omitempty"`
	State                             string   `force:",omitempty"`
	Street                            string   `force:",omitempty"`
	SystemModstamp                    string   `force:",omitempty"`
	Title                             string   `force:",omitempty"`
}

func (t *LeadCleanInfo) ApiName() string {
	return "LeadCleanInfo"
}

func (t *LeadCleanInfo) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("LeadCleanInfo #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tAddress: %v\n", t.Address))
	builder.WriteString(fmt.Sprintf("\tAnnualRevenue: %v\n", t.AnnualRevenue))
	builder.WriteString(fmt.Sprintf("\tCity: %v\n", t.City))
	builder.WriteString(fmt.Sprintf("\tCleanedByJob: %v\n", t.CleanedByJob))
	builder.WriteString(fmt.Sprintf("\tCleanedByUser: %v\n", t.CleanedByUser))
	builder.WriteString(fmt.Sprintf("\tCompanyDunsNumber: %v\n", t.CompanyDunsNumber))
	builder.WriteString(fmt.Sprintf("\tCompanyName: %v\n", t.CompanyName))
	builder.WriteString(fmt.Sprintf("\tContactStatusDataDotCom: %v\n", t.ContactStatusDataDotCom))
	builder.WriteString(fmt.Sprintf("\tCountry: %v\n", t.Country))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tDandBCompanyDunsNumber: %v\n", t.DandBCompanyDunsNumber))
	builder.WriteString(fmt.Sprintf("\tDataDotComCompanyId: %v\n", t.DataDotComCompanyId))
	builder.WriteString(fmt.Sprintf("\tDataDotComId: %v\n", t.DataDotComId))
	builder.WriteString(fmt.Sprintf("\tEmail: %v\n", t.Email))
	builder.WriteString(fmt.Sprintf("\tFirstName: %v\n", t.FirstName))
	builder.WriteString(fmt.Sprintf("\tGeocodeAccuracy: %v\n", t.GeocodeAccuracy))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tIndustry: %v\n", t.Industry))
	builder.WriteString(fmt.Sprintf("\tIsDeleted: %v\n", t.IsDeleted))
	builder.WriteString(fmt.Sprintf("\tIsDifferentAnnualRevenue: %v\n", t.IsDifferentAnnualRevenue))
	builder.WriteString(fmt.Sprintf("\tIsDifferentCity: %v\n", t.IsDifferentCity))
	builder.WriteString(fmt.Sprintf("\tIsDifferentCompanyDunsNumber: %v\n", t.IsDifferentCompanyDunsNumber))
	builder.WriteString(fmt.Sprintf("\tIsDifferentCompanyName: %v\n", t.IsDifferentCompanyName))
	builder.WriteString(fmt.Sprintf("\tIsDifferentCountry: %v\n", t.IsDifferentCountry))
	builder.WriteString(fmt.Sprintf("\tIsDifferentCountryCode: %v\n", t.IsDifferentCountryCode))
	builder.WriteString(fmt.Sprintf("\tIsDifferentDandBCompanyDunsNumber: %v\n", t.IsDifferentDandBCompanyDunsNumber))
	builder.WriteString(fmt.Sprintf("\tIsDifferentEmail: %v\n", t.IsDifferentEmail))
	builder.WriteString(fmt.Sprintf("\tIsDifferentFirstName: %v\n", t.IsDifferentFirstName))
	builder.WriteString(fmt.Sprintf("\tIsDifferentIndustry: %v\n", t.IsDifferentIndustry))
	builder.WriteString(fmt.Sprintf("\tIsDifferentLastName: %v\n", t.IsDifferentLastName))
	builder.WriteString(fmt.Sprintf("\tIsDifferentNumberOfEmployees: %v\n", t.IsDifferentNumberOfEmployees))
	builder.WriteString(fmt.Sprintf("\tIsDifferentPhone: %v\n", t.IsDifferentPhone))
	builder.WriteString(fmt.Sprintf("\tIsDifferentPostalCode: %v\n", t.IsDifferentPostalCode))
	builder.WriteString(fmt.Sprintf("\tIsDifferentState: %v\n", t.IsDifferentState))
	builder.WriteString(fmt.Sprintf("\tIsDifferentStateCode: %v\n", t.IsDifferentStateCode))
	builder.WriteString(fmt.Sprintf("\tIsDifferentStreet: %v\n", t.IsDifferentStreet))
	builder.WriteString(fmt.Sprintf("\tIsDifferentTitle: %v\n", t.IsDifferentTitle))
	builder.WriteString(fmt.Sprintf("\tIsFlaggedWrongAddress: %v\n", t.IsFlaggedWrongAddress))
	builder.WriteString(fmt.Sprintf("\tIsFlaggedWrongAnnualRevenue: %v\n", t.IsFlaggedWrongAnnualRevenue))
	builder.WriteString(fmt.Sprintf("\tIsFlaggedWrongCompanyDunsNumber: %v\n", t.IsFlaggedWrongCompanyDunsNumber))
	builder.WriteString(fmt.Sprintf("\tIsFlaggedWrongCompanyName: %v\n", t.IsFlaggedWrongCompanyName))
	builder.WriteString(fmt.Sprintf("\tIsFlaggedWrongEmail: %v\n", t.IsFlaggedWrongEmail))
	builder.WriteString(fmt.Sprintf("\tIsFlaggedWrongIndustry: %v\n", t.IsFlaggedWrongIndustry))
	builder.WriteString(fmt.Sprintf("\tIsFlaggedWrongName: %v\n", t.IsFlaggedWrongName))
	builder.WriteString(fmt.Sprintf("\tIsFlaggedWrongNumberOfEmployees: %v\n", t.IsFlaggedWrongNumberOfEmployees))
	builder.WriteString(fmt.Sprintf("\tIsFlaggedWrongPhone: %v\n", t.IsFlaggedWrongPhone))
	builder.WriteString(fmt.Sprintf("\tIsFlaggedWrongTitle: %v\n", t.IsFlaggedWrongTitle))
	builder.WriteString(fmt.Sprintf("\tIsInactive: %v\n", t.IsInactive))
	builder.WriteString(fmt.Sprintf("\tIsReviewedAddress: %v\n", t.IsReviewedAddress))
	builder.WriteString(fmt.Sprintf("\tIsReviewedAnnualRevenue: %v\n", t.IsReviewedAnnualRevenue))
	builder.WriteString(fmt.Sprintf("\tIsReviewedCompanyDunsNumber: %v\n", t.IsReviewedCompanyDunsNumber))
	builder.WriteString(fmt.Sprintf("\tIsReviewedCompanyName: %v\n", t.IsReviewedCompanyName))
	builder.WriteString(fmt.Sprintf("\tIsReviewedDandBCompanyDunsNumber: %v\n", t.IsReviewedDandBCompanyDunsNumber))
	builder.WriteString(fmt.Sprintf("\tIsReviewedEmail: %v\n", t.IsReviewedEmail))
	builder.WriteString(fmt.Sprintf("\tIsReviewedIndustry: %v\n", t.IsReviewedIndustry))
	builder.WriteString(fmt.Sprintf("\tIsReviewedName: %v\n", t.IsReviewedName))
	builder.WriteString(fmt.Sprintf("\tIsReviewedNumberOfEmployees: %v\n", t.IsReviewedNumberOfEmployees))
	builder.WriteString(fmt.Sprintf("\tIsReviewedPhone: %v\n", t.IsReviewedPhone))
	builder.WriteString(fmt.Sprintf("\tIsReviewedTitle: %v\n", t.IsReviewedTitle))
	builder.WriteString(fmt.Sprintf("\tLastMatchedDate: %v\n", t.LastMatchedDate))
	builder.WriteString(fmt.Sprintf("\tLastModifiedById: %v\n", t.LastModifiedById))
	builder.WriteString(fmt.Sprintf("\tLastModifiedDate: %v\n", t.LastModifiedDate))
	builder.WriteString(fmt.Sprintf("\tLastName: %v\n", t.LastName))
	builder.WriteString(fmt.Sprintf("\tLastStatusChangedById: %v\n", t.LastStatusChangedById))
	builder.WriteString(fmt.Sprintf("\tLastStatusChangedDate: %v\n", t.LastStatusChangedDate))
	builder.WriteString(fmt.Sprintf("\tLatitude: %v\n", t.Latitude))
	builder.WriteString(fmt.Sprintf("\tLeadId: %v\n", t.LeadId))
	builder.WriteString(fmt.Sprintf("\tLongitude: %v\n", t.Longitude))
	builder.WriteString(fmt.Sprintf("\tName: %v\n", t.Name))
	builder.WriteString(fmt.Sprintf("\tNumberOfEmployees: %v\n", t.NumberOfEmployees))
	builder.WriteString(fmt.Sprintf("\tPhone: %v\n", t.Phone))
	builder.WriteString(fmt.Sprintf("\tPostalCode: %v\n", t.PostalCode))
	builder.WriteString(fmt.Sprintf("\tState: %v\n", t.State))
	builder.WriteString(fmt.Sprintf("\tStreet: %v\n", t.Street))
	builder.WriteString(fmt.Sprintf("\tSystemModstamp: %v\n", t.SystemModstamp))
	builder.WriteString(fmt.Sprintf("\tTitle: %v\n", t.Title))

	return builder.String()
}

type LeadCleanInfoQueryResponse struct {
	BaseQuery
	Records []LeadCleanInfo `json:"Records" force:"records"`
}
