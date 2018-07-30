// This file was generated for SObject Site, API Version v43.0 at 2018-07-30 03:47:32.825095332 -0400 EDT m=+19.168421021

package sobjects

import (
	"fmt"
	"strings"
)

type Site struct {
	BaseSObject
	AdminId                                    string `force:",omitempty"`
	AnalyticsTrackingCode                      string `force:",omitempty"`
	ClickjackProtectionLevel                   string `force:",omitempty"`
	CreatedById                                string `force:",omitempty"`
	CreatedDate                                string `force:",omitempty"`
	DailyBandwidthLimit                        int    `force:",omitempty"`
	DailyBandwidthUsed                         int    `force:",omitempty"`
	DailyRequestTimeLimit                      int    `force:",omitempty"`
	DailyRequestTimeUsed                       int    `force:",omitempty"`
	Description                                string `force:",omitempty"`
	GuestUserId                                string `force:",omitempty"`
	Id                                         string `force:",omitempty"`
	LastModifiedById                           string `force:",omitempty"`
	LastModifiedDate                           string `force:",omitempty"`
	MasterLabel                                string `force:",omitempty"`
	MonthlyPageViewsEntitlement                int    `force:",omitempty"`
	Name                                       string `force:",omitempty"`
	OptionsAllowGuestSupportApi                bool   `force:",omitempty"`
	OptionsAllowHomePage                       bool   `force:",omitempty"`
	OptionsAllowStandardAnswersPages           bool   `force:",omitempty"`
	OptionsAllowStandardIdeasPages             bool   `force:",omitempty"`
	OptionsAllowStandardLookups                bool   `force:",omitempty"`
	OptionsAllowStandardPortalPages            bool   `force:",omitempty"`
	OptionsAllowStandardSearch                 bool   `force:",omitempty"`
	OptionsBrowserXssProtection                bool   `force:",omitempty"`
	OptionsContentSniffingProtection           bool   `force:",omitempty"`
	OptionsCspUpgradeInsecureRequests          bool   `force:",omitempty"`
	OptionsEnableFeeds                         bool   `force:",omitempty"`
	OptionsReferrerPolicyOriginWhenCrossOrigin bool   `force:",omitempty"`
	OptionsRequireHttps                        bool   `force:",omitempty"`
	SiteType                                   string `force:",omitempty"`
	Status                                     string `force:",omitempty"`
	Subdomain                                  string `force:",omitempty"`
	SystemModstamp                             string `force:",omitempty"`
	UrlPathPrefix                              string `force:",omitempty"`
}

func (t *Site) ApiName() string {
	return "Site"
}

func (t *Site) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("Site #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tAdminId: %v\n", t.AdminId))
	builder.WriteString(fmt.Sprintf("\tAnalyticsTrackingCode: %v\n", t.AnalyticsTrackingCode))
	builder.WriteString(fmt.Sprintf("\tClickjackProtectionLevel: %v\n", t.ClickjackProtectionLevel))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tDailyBandwidthLimit: %v\n", t.DailyBandwidthLimit))
	builder.WriteString(fmt.Sprintf("\tDailyBandwidthUsed: %v\n", t.DailyBandwidthUsed))
	builder.WriteString(fmt.Sprintf("\tDailyRequestTimeLimit: %v\n", t.DailyRequestTimeLimit))
	builder.WriteString(fmt.Sprintf("\tDailyRequestTimeUsed: %v\n", t.DailyRequestTimeUsed))
	builder.WriteString(fmt.Sprintf("\tDescription: %v\n", t.Description))
	builder.WriteString(fmt.Sprintf("\tGuestUserId: %v\n", t.GuestUserId))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tLastModifiedById: %v\n", t.LastModifiedById))
	builder.WriteString(fmt.Sprintf("\tLastModifiedDate: %v\n", t.LastModifiedDate))
	builder.WriteString(fmt.Sprintf("\tMasterLabel: %v\n", t.MasterLabel))
	builder.WriteString(fmt.Sprintf("\tMonthlyPageViewsEntitlement: %v\n", t.MonthlyPageViewsEntitlement))
	builder.WriteString(fmt.Sprintf("\tName: %v\n", t.Name))
	builder.WriteString(fmt.Sprintf("\tOptionsAllowGuestSupportApi: %v\n", t.OptionsAllowGuestSupportApi))
	builder.WriteString(fmt.Sprintf("\tOptionsAllowHomePage: %v\n", t.OptionsAllowHomePage))
	builder.WriteString(fmt.Sprintf("\tOptionsAllowStandardAnswersPages: %v\n", t.OptionsAllowStandardAnswersPages))
	builder.WriteString(fmt.Sprintf("\tOptionsAllowStandardIdeasPages: %v\n", t.OptionsAllowStandardIdeasPages))
	builder.WriteString(fmt.Sprintf("\tOptionsAllowStandardLookups: %v\n", t.OptionsAllowStandardLookups))
	builder.WriteString(fmt.Sprintf("\tOptionsAllowStandardPortalPages: %v\n", t.OptionsAllowStandardPortalPages))
	builder.WriteString(fmt.Sprintf("\tOptionsAllowStandardSearch: %v\n", t.OptionsAllowStandardSearch))
	builder.WriteString(fmt.Sprintf("\tOptionsBrowserXssProtection: %v\n", t.OptionsBrowserXssProtection))
	builder.WriteString(fmt.Sprintf("\tOptionsContentSniffingProtection: %v\n", t.OptionsContentSniffingProtection))
	builder.WriteString(fmt.Sprintf("\tOptionsCspUpgradeInsecureRequests: %v\n", t.OptionsCspUpgradeInsecureRequests))
	builder.WriteString(fmt.Sprintf("\tOptionsEnableFeeds: %v\n", t.OptionsEnableFeeds))
	builder.WriteString(fmt.Sprintf("\tOptionsReferrerPolicyOriginWhenCrossOrigin: %v\n", t.OptionsReferrerPolicyOriginWhenCrossOrigin))
	builder.WriteString(fmt.Sprintf("\tOptionsRequireHttps: %v\n", t.OptionsRequireHttps))
	builder.WriteString(fmt.Sprintf("\tSiteType: %v\n", t.SiteType))
	builder.WriteString(fmt.Sprintf("\tStatus: %v\n", t.Status))
	builder.WriteString(fmt.Sprintf("\tSubdomain: %v\n", t.Subdomain))
	builder.WriteString(fmt.Sprintf("\tSystemModstamp: %v\n", t.SystemModstamp))
	builder.WriteString(fmt.Sprintf("\tUrlPathPrefix: %v\n", t.UrlPathPrefix))

	return builder.String()
}

type SiteQueryResponse struct {
	BaseQuery
	Records []Site `json:"Records" force:"records"`
}
