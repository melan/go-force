// This file was generated for SObject ContactCleanInfo, API Version v43.0 at 2018-07-30 03:47:40.22243352 -0400 EDT m=+26.566036786

package sobjects

import (
	"fmt"
	"strings"
)

type ContactCleanInfo struct {
	BaseSObject
	Address                 *Address `force:",omitempty"`
	City                    string   `force:",omitempty"`
	CleanedByJob            bool     `force:",omitempty"`
	CleanedByUser           bool     `force:",omitempty"`
	ContactId               string   `force:",omitempty"`
	ContactStatusDataDotCom string   `force:",omitempty"`
	Country                 string   `force:",omitempty"`
	CreatedById             string   `force:",omitempty"`
	CreatedDate             string   `force:",omitempty"`
	DataDotComId            string   `force:",omitempty"`
	Email                   string   `force:",omitempty"`
	FirstName               string   `force:",omitempty"`
	GeocodeAccuracy         string   `force:",omitempty"`
	Id                      string   `force:",omitempty"`
	IsDeleted               bool     `force:",omitempty"`
	IsDifferentCity         bool     `force:",omitempty"`
	IsDifferentCountry      bool     `force:",omitempty"`
	IsDifferentCountryCode  bool     `force:",omitempty"`
	IsDifferentEmail        bool     `force:",omitempty"`
	IsDifferentFirstName    bool     `force:",omitempty"`
	IsDifferentLastName     bool     `force:",omitempty"`
	IsDifferentPhone        bool     `force:",omitempty"`
	IsDifferentPostalCode   bool     `force:",omitempty"`
	IsDifferentState        bool     `force:",omitempty"`
	IsDifferentStateCode    bool     `force:",omitempty"`
	IsDifferentStreet       bool     `force:",omitempty"`
	IsDifferentTitle        bool     `force:",omitempty"`
	IsFlaggedWrongAddress   bool     `force:",omitempty"`
	IsFlaggedWrongEmail     bool     `force:",omitempty"`
	IsFlaggedWrongName      bool     `force:",omitempty"`
	IsFlaggedWrongPhone     bool     `force:",omitempty"`
	IsFlaggedWrongTitle     bool     `force:",omitempty"`
	IsInactive              bool     `force:",omitempty"`
	IsReviewedAddress       bool     `force:",omitempty"`
	IsReviewedEmail         bool     `force:",omitempty"`
	IsReviewedName          bool     `force:",omitempty"`
	IsReviewedPhone         bool     `force:",omitempty"`
	IsReviewedTitle         bool     `force:",omitempty"`
	LastMatchedDate         string   `force:",omitempty"`
	LastModifiedById        string   `force:",omitempty"`
	LastModifiedDate        string   `force:",omitempty"`
	LastName                string   `force:",omitempty"`
	LastStatusChangedById   string   `force:",omitempty"`
	LastStatusChangedDate   string   `force:",omitempty"`
	Latitude                float64  `force:",omitempty"`
	Longitude               float64  `force:",omitempty"`
	Name                    string   `force:",omitempty"`
	Phone                   string   `force:",omitempty"`
	PostalCode              string   `force:",omitempty"`
	State                   string   `force:",omitempty"`
	Street                  string   `force:",omitempty"`
	SystemModstamp          string   `force:",omitempty"`
	Title                   string   `force:",omitempty"`
}

func (t *ContactCleanInfo) ApiName() string {
	return "ContactCleanInfo"
}

func (t *ContactCleanInfo) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("ContactCleanInfo #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tAddress: %v\n", t.Address))
	builder.WriteString(fmt.Sprintf("\tCity: %v\n", t.City))
	builder.WriteString(fmt.Sprintf("\tCleanedByJob: %v\n", t.CleanedByJob))
	builder.WriteString(fmt.Sprintf("\tCleanedByUser: %v\n", t.CleanedByUser))
	builder.WriteString(fmt.Sprintf("\tContactId: %v\n", t.ContactId))
	builder.WriteString(fmt.Sprintf("\tContactStatusDataDotCom: %v\n", t.ContactStatusDataDotCom))
	builder.WriteString(fmt.Sprintf("\tCountry: %v\n", t.Country))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tDataDotComId: %v\n", t.DataDotComId))
	builder.WriteString(fmt.Sprintf("\tEmail: %v\n", t.Email))
	builder.WriteString(fmt.Sprintf("\tFirstName: %v\n", t.FirstName))
	builder.WriteString(fmt.Sprintf("\tGeocodeAccuracy: %v\n", t.GeocodeAccuracy))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tIsDeleted: %v\n", t.IsDeleted))
	builder.WriteString(fmt.Sprintf("\tIsDifferentCity: %v\n", t.IsDifferentCity))
	builder.WriteString(fmt.Sprintf("\tIsDifferentCountry: %v\n", t.IsDifferentCountry))
	builder.WriteString(fmt.Sprintf("\tIsDifferentCountryCode: %v\n", t.IsDifferentCountryCode))
	builder.WriteString(fmt.Sprintf("\tIsDifferentEmail: %v\n", t.IsDifferentEmail))
	builder.WriteString(fmt.Sprintf("\tIsDifferentFirstName: %v\n", t.IsDifferentFirstName))
	builder.WriteString(fmt.Sprintf("\tIsDifferentLastName: %v\n", t.IsDifferentLastName))
	builder.WriteString(fmt.Sprintf("\tIsDifferentPhone: %v\n", t.IsDifferentPhone))
	builder.WriteString(fmt.Sprintf("\tIsDifferentPostalCode: %v\n", t.IsDifferentPostalCode))
	builder.WriteString(fmt.Sprintf("\tIsDifferentState: %v\n", t.IsDifferentState))
	builder.WriteString(fmt.Sprintf("\tIsDifferentStateCode: %v\n", t.IsDifferentStateCode))
	builder.WriteString(fmt.Sprintf("\tIsDifferentStreet: %v\n", t.IsDifferentStreet))
	builder.WriteString(fmt.Sprintf("\tIsDifferentTitle: %v\n", t.IsDifferentTitle))
	builder.WriteString(fmt.Sprintf("\tIsFlaggedWrongAddress: %v\n", t.IsFlaggedWrongAddress))
	builder.WriteString(fmt.Sprintf("\tIsFlaggedWrongEmail: %v\n", t.IsFlaggedWrongEmail))
	builder.WriteString(fmt.Sprintf("\tIsFlaggedWrongName: %v\n", t.IsFlaggedWrongName))
	builder.WriteString(fmt.Sprintf("\tIsFlaggedWrongPhone: %v\n", t.IsFlaggedWrongPhone))
	builder.WriteString(fmt.Sprintf("\tIsFlaggedWrongTitle: %v\n", t.IsFlaggedWrongTitle))
	builder.WriteString(fmt.Sprintf("\tIsInactive: %v\n", t.IsInactive))
	builder.WriteString(fmt.Sprintf("\tIsReviewedAddress: %v\n", t.IsReviewedAddress))
	builder.WriteString(fmt.Sprintf("\tIsReviewedEmail: %v\n", t.IsReviewedEmail))
	builder.WriteString(fmt.Sprintf("\tIsReviewedName: %v\n", t.IsReviewedName))
	builder.WriteString(fmt.Sprintf("\tIsReviewedPhone: %v\n", t.IsReviewedPhone))
	builder.WriteString(fmt.Sprintf("\tIsReviewedTitle: %v\n", t.IsReviewedTitle))
	builder.WriteString(fmt.Sprintf("\tLastMatchedDate: %v\n", t.LastMatchedDate))
	builder.WriteString(fmt.Sprintf("\tLastModifiedById: %v\n", t.LastModifiedById))
	builder.WriteString(fmt.Sprintf("\tLastModifiedDate: %v\n", t.LastModifiedDate))
	builder.WriteString(fmt.Sprintf("\tLastName: %v\n", t.LastName))
	builder.WriteString(fmt.Sprintf("\tLastStatusChangedById: %v\n", t.LastStatusChangedById))
	builder.WriteString(fmt.Sprintf("\tLastStatusChangedDate: %v\n", t.LastStatusChangedDate))
	builder.WriteString(fmt.Sprintf("\tLatitude: %v\n", t.Latitude))
	builder.WriteString(fmt.Sprintf("\tLongitude: %v\n", t.Longitude))
	builder.WriteString(fmt.Sprintf("\tName: %v\n", t.Name))
	builder.WriteString(fmt.Sprintf("\tPhone: %v\n", t.Phone))
	builder.WriteString(fmt.Sprintf("\tPostalCode: %v\n", t.PostalCode))
	builder.WriteString(fmt.Sprintf("\tState: %v\n", t.State))
	builder.WriteString(fmt.Sprintf("\tStreet: %v\n", t.Street))
	builder.WriteString(fmt.Sprintf("\tSystemModstamp: %v\n", t.SystemModstamp))
	builder.WriteString(fmt.Sprintf("\tTitle: %v\n", t.Title))

	return builder.String()
}

type ContactCleanInfoQueryResponse struct {
	BaseQuery
	Records []ContactCleanInfo `json:"Records" force:"records"`
}
