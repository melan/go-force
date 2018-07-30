// This file was generated for SObject AppMenuItem, API Version v43.0 at 2018-07-30 03:47:19.129672175 -0400 EDT m=+5.472483956

package sobjects

import (
	"fmt"
	"strings"
)

type AppMenuItem struct {
	BaseSObject
	ApplicationId             string `force:",omitempty"`
	CanvasAccessMethod        string `force:",omitempty"`
	CanvasEnabled             bool   `force:",omitempty"`
	CanvasOptions             string `force:",omitempty"`
	CanvasReferenceId         string `force:",omitempty"`
	CanvasSelectedLocations   string `force:",omitempty"`
	CanvasUrl                 string `force:",omitempty"`
	CreatedById               string `force:",omitempty"`
	CreatedDate               string `force:",omitempty"`
	Description               string `force:",omitempty"`
	IconUrl                   string `force:",omitempty"`
	Id                        string `force:",omitempty"`
	InfoUrl                   string `force:",omitempty"`
	IsAccessible              bool   `force:",omitempty"`
	IsDeleted                 bool   `force:",omitempty"`
	IsRegisteredDeviceOnly    bool   `force:",omitempty"`
	IsUsingAdminAuthorization bool   `force:",omitempty"`
	IsVisible                 bool   `force:",omitempty"`
	Label                     string `force:",omitempty"`
	LastModifiedById          string `force:",omitempty"`
	LastModifiedDate          string `force:",omitempty"`
	LogoUrl                   string `force:",omitempty"`
	MobileAppBinaryId         string `force:",omitempty"`
	MobileAppInstallUrl       string `force:",omitempty"`
	MobileAppInstalledDate    string `force:",omitempty"`
	MobileAppInstalledVersion string `force:",omitempty"`
	MobileAppVer              string `force:",omitempty"`
	MobileDeviceType          string `force:",omitempty"`
	MobileMinOsVer            string `force:",omitempty"`
	MobilePlatform            string `force:",omitempty"`
	MobileStartUrl            string `force:",omitempty"`
	Name                      string `force:",omitempty"`
	NamespacePrefix           string `force:",omitempty"`
	SortOrder                 int    `force:",omitempty"`
	StartUrl                  string `force:",omitempty"`
	SystemModstamp            string `force:",omitempty"`
	Type                      string `force:",omitempty"`
	UserSortOrder             int    `force:",omitempty"`
}

func (t *AppMenuItem) ApiName() string {
	return "AppMenuItem"
}

func (t *AppMenuItem) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("AppMenuItem #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tApplicationId: %v\n", t.ApplicationId))
	builder.WriteString(fmt.Sprintf("\tCanvasAccessMethod: %v\n", t.CanvasAccessMethod))
	builder.WriteString(fmt.Sprintf("\tCanvasEnabled: %v\n", t.CanvasEnabled))
	builder.WriteString(fmt.Sprintf("\tCanvasOptions: %v\n", t.CanvasOptions))
	builder.WriteString(fmt.Sprintf("\tCanvasReferenceId: %v\n", t.CanvasReferenceId))
	builder.WriteString(fmt.Sprintf("\tCanvasSelectedLocations: %v\n", t.CanvasSelectedLocations))
	builder.WriteString(fmt.Sprintf("\tCanvasUrl: %v\n", t.CanvasUrl))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tDescription: %v\n", t.Description))
	builder.WriteString(fmt.Sprintf("\tIconUrl: %v\n", t.IconUrl))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tInfoUrl: %v\n", t.InfoUrl))
	builder.WriteString(fmt.Sprintf("\tIsAccessible: %v\n", t.IsAccessible))
	builder.WriteString(fmt.Sprintf("\tIsDeleted: %v\n", t.IsDeleted))
	builder.WriteString(fmt.Sprintf("\tIsRegisteredDeviceOnly: %v\n", t.IsRegisteredDeviceOnly))
	builder.WriteString(fmt.Sprintf("\tIsUsingAdminAuthorization: %v\n", t.IsUsingAdminAuthorization))
	builder.WriteString(fmt.Sprintf("\tIsVisible: %v\n", t.IsVisible))
	builder.WriteString(fmt.Sprintf("\tLabel: %v\n", t.Label))
	builder.WriteString(fmt.Sprintf("\tLastModifiedById: %v\n", t.LastModifiedById))
	builder.WriteString(fmt.Sprintf("\tLastModifiedDate: %v\n", t.LastModifiedDate))
	builder.WriteString(fmt.Sprintf("\tLogoUrl: %v\n", t.LogoUrl))
	builder.WriteString(fmt.Sprintf("\tMobileAppBinaryId: %v\n", t.MobileAppBinaryId))
	builder.WriteString(fmt.Sprintf("\tMobileAppInstallUrl: %v\n", t.MobileAppInstallUrl))
	builder.WriteString(fmt.Sprintf("\tMobileAppInstalledDate: %v\n", t.MobileAppInstalledDate))
	builder.WriteString(fmt.Sprintf("\tMobileAppInstalledVersion: %v\n", t.MobileAppInstalledVersion))
	builder.WriteString(fmt.Sprintf("\tMobileAppVer: %v\n", t.MobileAppVer))
	builder.WriteString(fmt.Sprintf("\tMobileDeviceType: %v\n", t.MobileDeviceType))
	builder.WriteString(fmt.Sprintf("\tMobileMinOsVer: %v\n", t.MobileMinOsVer))
	builder.WriteString(fmt.Sprintf("\tMobilePlatform: %v\n", t.MobilePlatform))
	builder.WriteString(fmt.Sprintf("\tMobileStartUrl: %v\n", t.MobileStartUrl))
	builder.WriteString(fmt.Sprintf("\tName: %v\n", t.Name))
	builder.WriteString(fmt.Sprintf("\tNamespacePrefix: %v\n", t.NamespacePrefix))
	builder.WriteString(fmt.Sprintf("\tSortOrder: %v\n", t.SortOrder))
	builder.WriteString(fmt.Sprintf("\tStartUrl: %v\n", t.StartUrl))
	builder.WriteString(fmt.Sprintf("\tSystemModstamp: %v\n", t.SystemModstamp))
	builder.WriteString(fmt.Sprintf("\tType: %v\n", t.Type))
	builder.WriteString(fmt.Sprintf("\tUserSortOrder: %v\n", t.UserSortOrder))

	return builder.String()
}

type AppMenuItemQueryResponse struct {
	BaseQuery
	Records []AppMenuItem `json:"Records" force:"records"`
}
