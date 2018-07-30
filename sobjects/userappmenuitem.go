// This file was generated for SObject UserAppMenuItem, API Version v43.0 at 2018-07-30 03:47:34.623499103 -0400 EDT m=+20.966892275

package sobjects

import (
	"fmt"
	"strings"
)

type UserAppMenuItem struct {
	BaseSObject
	AppMenuItemId             string `force:",omitempty"`
	ApplicationId             string `force:",omitempty"`
	Description               string `force:",omitempty"`
	IconUrl                   string `force:",omitempty"`
	Id                        string `force:",omitempty"`
	InfoUrl                   string `force:",omitempty"`
	IsUsingAdminAuthorization bool   `force:",omitempty"`
	IsVisible                 bool   `force:",omitempty"`
	Label                     string `force:",omitempty"`
	LogoUrl                   string `force:",omitempty"`
	MobileStartUrl            string `force:",omitempty"`
	Name                      string `force:",omitempty"`
	SortOrder                 int    `force:",omitempty"`
	StartUrl                  string `force:",omitempty"`
	Type                      string `force:",omitempty"`
	UserSortOrder             int    `force:",omitempty"`
}

func (t *UserAppMenuItem) ApiName() string {
	return "UserAppMenuItem"
}

func (t *UserAppMenuItem) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("UserAppMenuItem #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tAppMenuItemId: %v\n", t.AppMenuItemId))
	builder.WriteString(fmt.Sprintf("\tApplicationId: %v\n", t.ApplicationId))
	builder.WriteString(fmt.Sprintf("\tDescription: %v\n", t.Description))
	builder.WriteString(fmt.Sprintf("\tIconUrl: %v\n", t.IconUrl))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tInfoUrl: %v\n", t.InfoUrl))
	builder.WriteString(fmt.Sprintf("\tIsUsingAdminAuthorization: %v\n", t.IsUsingAdminAuthorization))
	builder.WriteString(fmt.Sprintf("\tIsVisible: %v\n", t.IsVisible))
	builder.WriteString(fmt.Sprintf("\tLabel: %v\n", t.Label))
	builder.WriteString(fmt.Sprintf("\tLogoUrl: %v\n", t.LogoUrl))
	builder.WriteString(fmt.Sprintf("\tMobileStartUrl: %v\n", t.MobileStartUrl))
	builder.WriteString(fmt.Sprintf("\tName: %v\n", t.Name))
	builder.WriteString(fmt.Sprintf("\tSortOrder: %v\n", t.SortOrder))
	builder.WriteString(fmt.Sprintf("\tStartUrl: %v\n", t.StartUrl))
	builder.WriteString(fmt.Sprintf("\tType: %v\n", t.Type))
	builder.WriteString(fmt.Sprintf("\tUserSortOrder: %v\n", t.UserSortOrder))

	return builder.String()
}

type UserAppMenuItemQueryResponse struct {
	BaseQuery
	Records []UserAppMenuItem `json:"Records" force:"records"`
}
