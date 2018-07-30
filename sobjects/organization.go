// This file was generated for SObject Organization, API Version v43.0 at 2018-07-30 03:47:18.34179307 -0400 EDT m=+4.684575286

package sobjects

import (
	"fmt"
	"strings"
)

type Organization struct {
	BaseSObject
	Address                                *Address `force:",omitempty"`
	City                                   string   `force:",omitempty"`
	ComplianceBccEmail                     string   `force:",omitempty"`
	Country                                string   `force:",omitempty"`
	CreatedById                            string   `force:",omitempty"`
	CreatedDate                            string   `force:",omitempty"`
	DefaultAccountAccess                   string   `force:",omitempty"`
	DefaultCalendarAccess                  string   `force:",omitempty"`
	DefaultCampaignAccess                  string   `force:",omitempty"`
	DefaultCaseAccess                      string   `force:",omitempty"`
	DefaultContactAccess                   string   `force:",omitempty"`
	DefaultLeadAccess                      string   `force:",omitempty"`
	DefaultLocaleSidKey                    string   `force:",omitempty"`
	DefaultOpportunityAccess               string   `force:",omitempty"`
	DefaultPricebookAccess                 string   `force:",omitempty"`
	Division                               string   `force:",omitempty"`
	Fax                                    string   `force:",omitempty"`
	FiscalYearStartMonth                   int      `force:",omitempty"`
	GeocodeAccuracy                        string   `force:",omitempty"`
	Id                                     string   `force:",omitempty"`
	InstanceName                           string   `force:",omitempty"`
	IsReadOnly                             bool     `force:",omitempty"`
	IsSandbox                              bool     `force:",omitempty"`
	LanguageLocaleKey                      string   `force:",omitempty"`
	LastModifiedById                       string   `force:",omitempty"`
	LastModifiedDate                       string   `force:",omitempty"`
	Latitude                               float64  `force:",omitempty"`
	Longitude                              float64  `force:",omitempty"`
	MonthlyPageViewsEntitlement            int      `force:",omitempty"`
	MonthlyPageViewsUsed                   int      `force:",omitempty"`
	Name                                   string   `force:",omitempty"`
	NamespacePrefix                        string   `force:",omitempty"`
	NumKnowledgeService                    int      `force:",omitempty"`
	OrganizationType                       string   `force:",omitempty"`
	Phone                                  string   `force:",omitempty"`
	PostalCode                             string   `force:",omitempty"`
	PreferencesActivityAnalyticsEnabled    bool     `force:",omitempty"`
	PreferencesAutoSelectIndividualOnMerge bool     `force:",omitempty"`
	PreferencesConsentManagementEnabled    bool     `force:",omitempty"`
	PreferencesIndividualAutoCreateEnabled bool     `force:",omitempty"`
	PreferencesLightningLoginEnabled       bool     `force:",omitempty"`
	PreferencesOnlyLLPermUserAllowed       bool     `force:",omitempty"`
	PreferencesRequireOpportunityProducts  bool     `force:",omitempty"`
	PreferencesTerminateOldestSession      bool     `force:",omitempty"`
	PreferencesTransactionSecurityPolicy   bool     `force:",omitempty"`
	PrimaryContact                         string   `force:",omitempty"`
	ReceivesAdminInfoEmails                bool     `force:",omitempty"`
	ReceivesInfoEmails                     bool     `force:",omitempty"`
	SignupCountryIsoCode                   string   `force:",omitempty"`
	State                                  string   `force:",omitempty"`
	Street                                 string   `force:",omitempty"`
	SystemModstamp                         string   `force:",omitempty"`
	TimeZoneSidKey                         string   `force:",omitempty"`
	TrialExpirationDate                    string   `force:",omitempty"`
	UiSkin                                 string   `force:",omitempty"`
	UsesStartDateAsFiscalYearName          bool     `force:",omitempty"`
	WebToCaseDefaultOrigin                 string   `force:",omitempty"`
}

func (t *Organization) ApiName() string {
	return "Organization"
}

func (t *Organization) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("Organization #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tAddress: %v\n", t.Address))
	builder.WriteString(fmt.Sprintf("\tCity: %v\n", t.City))
	builder.WriteString(fmt.Sprintf("\tComplianceBccEmail: %v\n", t.ComplianceBccEmail))
	builder.WriteString(fmt.Sprintf("\tCountry: %v\n", t.Country))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tDefaultAccountAccess: %v\n", t.DefaultAccountAccess))
	builder.WriteString(fmt.Sprintf("\tDefaultCalendarAccess: %v\n", t.DefaultCalendarAccess))
	builder.WriteString(fmt.Sprintf("\tDefaultCampaignAccess: %v\n", t.DefaultCampaignAccess))
	builder.WriteString(fmt.Sprintf("\tDefaultCaseAccess: %v\n", t.DefaultCaseAccess))
	builder.WriteString(fmt.Sprintf("\tDefaultContactAccess: %v\n", t.DefaultContactAccess))
	builder.WriteString(fmt.Sprintf("\tDefaultLeadAccess: %v\n", t.DefaultLeadAccess))
	builder.WriteString(fmt.Sprintf("\tDefaultLocaleSidKey: %v\n", t.DefaultLocaleSidKey))
	builder.WriteString(fmt.Sprintf("\tDefaultOpportunityAccess: %v\n", t.DefaultOpportunityAccess))
	builder.WriteString(fmt.Sprintf("\tDefaultPricebookAccess: %v\n", t.DefaultPricebookAccess))
	builder.WriteString(fmt.Sprintf("\tDivision: %v\n", t.Division))
	builder.WriteString(fmt.Sprintf("\tFax: %v\n", t.Fax))
	builder.WriteString(fmt.Sprintf("\tFiscalYearStartMonth: %v\n", t.FiscalYearStartMonth))
	builder.WriteString(fmt.Sprintf("\tGeocodeAccuracy: %v\n", t.GeocodeAccuracy))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tInstanceName: %v\n", t.InstanceName))
	builder.WriteString(fmt.Sprintf("\tIsReadOnly: %v\n", t.IsReadOnly))
	builder.WriteString(fmt.Sprintf("\tIsSandbox: %v\n", t.IsSandbox))
	builder.WriteString(fmt.Sprintf("\tLanguageLocaleKey: %v\n", t.LanguageLocaleKey))
	builder.WriteString(fmt.Sprintf("\tLastModifiedById: %v\n", t.LastModifiedById))
	builder.WriteString(fmt.Sprintf("\tLastModifiedDate: %v\n", t.LastModifiedDate))
	builder.WriteString(fmt.Sprintf("\tLatitude: %v\n", t.Latitude))
	builder.WriteString(fmt.Sprintf("\tLongitude: %v\n", t.Longitude))
	builder.WriteString(fmt.Sprintf("\tMonthlyPageViewsEntitlement: %v\n", t.MonthlyPageViewsEntitlement))
	builder.WriteString(fmt.Sprintf("\tMonthlyPageViewsUsed: %v\n", t.MonthlyPageViewsUsed))
	builder.WriteString(fmt.Sprintf("\tName: %v\n", t.Name))
	builder.WriteString(fmt.Sprintf("\tNamespacePrefix: %v\n", t.NamespacePrefix))
	builder.WriteString(fmt.Sprintf("\tNumKnowledgeService: %v\n", t.NumKnowledgeService))
	builder.WriteString(fmt.Sprintf("\tOrganizationType: %v\n", t.OrganizationType))
	builder.WriteString(fmt.Sprintf("\tPhone: %v\n", t.Phone))
	builder.WriteString(fmt.Sprintf("\tPostalCode: %v\n", t.PostalCode))
	builder.WriteString(fmt.Sprintf("\tPreferencesActivityAnalyticsEnabled: %v\n", t.PreferencesActivityAnalyticsEnabled))
	builder.WriteString(fmt.Sprintf("\tPreferencesAutoSelectIndividualOnMerge: %v\n", t.PreferencesAutoSelectIndividualOnMerge))
	builder.WriteString(fmt.Sprintf("\tPreferencesConsentManagementEnabled: %v\n", t.PreferencesConsentManagementEnabled))
	builder.WriteString(fmt.Sprintf("\tPreferencesIndividualAutoCreateEnabled: %v\n", t.PreferencesIndividualAutoCreateEnabled))
	builder.WriteString(fmt.Sprintf("\tPreferencesLightningLoginEnabled: %v\n", t.PreferencesLightningLoginEnabled))
	builder.WriteString(fmt.Sprintf("\tPreferencesOnlyLLPermUserAllowed: %v\n", t.PreferencesOnlyLLPermUserAllowed))
	builder.WriteString(fmt.Sprintf("\tPreferencesRequireOpportunityProducts: %v\n", t.PreferencesRequireOpportunityProducts))
	builder.WriteString(fmt.Sprintf("\tPreferencesTerminateOldestSession: %v\n", t.PreferencesTerminateOldestSession))
	builder.WriteString(fmt.Sprintf("\tPreferencesTransactionSecurityPolicy: %v\n", t.PreferencesTransactionSecurityPolicy))
	builder.WriteString(fmt.Sprintf("\tPrimaryContact: %v\n", t.PrimaryContact))
	builder.WriteString(fmt.Sprintf("\tReceivesAdminInfoEmails: %v\n", t.ReceivesAdminInfoEmails))
	builder.WriteString(fmt.Sprintf("\tReceivesInfoEmails: %v\n", t.ReceivesInfoEmails))
	builder.WriteString(fmt.Sprintf("\tSignupCountryIsoCode: %v\n", t.SignupCountryIsoCode))
	builder.WriteString(fmt.Sprintf("\tState: %v\n", t.State))
	builder.WriteString(fmt.Sprintf("\tStreet: %v\n", t.Street))
	builder.WriteString(fmt.Sprintf("\tSystemModstamp: %v\n", t.SystemModstamp))
	builder.WriteString(fmt.Sprintf("\tTimeZoneSidKey: %v\n", t.TimeZoneSidKey))
	builder.WriteString(fmt.Sprintf("\tTrialExpirationDate: %v\n", t.TrialExpirationDate))
	builder.WriteString(fmt.Sprintf("\tUiSkin: %v\n", t.UiSkin))
	builder.WriteString(fmt.Sprintf("\tUsesStartDateAsFiscalYearName: %v\n", t.UsesStartDateAsFiscalYearName))
	builder.WriteString(fmt.Sprintf("\tWebToCaseDefaultOrigin: %v\n", t.WebToCaseDefaultOrigin))

	return builder.String()
}

type OrganizationQueryResponse struct {
	BaseQuery
	Records []Organization `json:"Records" force:"records"`
}
