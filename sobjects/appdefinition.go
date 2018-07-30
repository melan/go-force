// This file was generated for SObject AppDefinition, API Version v43.0 at 2018-07-30 03:47:44.624862783 -0400 EDT m=+30.968631246

package sobjects

import (
	"fmt"
	"strings"
)

type AppDefinition struct {
	BaseSObject
	Description                  string `force:",omitempty"`
	DeveloperName                string `force:",omitempty"`
	DurableId                    string `force:",omitempty"`
	HeaderColor                  string `force:",omitempty"`
	Id                           string `force:",omitempty"`
	IsLargeFormFactorSupported   bool   `force:",omitempty"`
	IsMediumFormFactorSupported  bool   `force:",omitempty"`
	IsNavAutoTempTabsDisabled    bool   `force:",omitempty"`
	IsNavPersonalizationDisabled bool   `force:",omitempty"`
	IsOverrideOrgTheme           bool   `force:",omitempty"`
	IsSmallFormFactorSupported   bool   `force:",omitempty"`
	Label                        string `force:",omitempty"`
	LogoUrl                      string `force:",omitempty"`
	MasterLabel                  string `force:",omitempty"`
	NamespacePrefix              string `force:",omitempty"`
	NavType                      string `force:",omitempty"`
	UiType                       string `force:",omitempty"`
	UtilityBar                   string `force:",omitempty"`
}

func (t *AppDefinition) ApiName() string {
	return "AppDefinition"
}

func (t *AppDefinition) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("AppDefinition #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tDescription: %v\n", t.Description))
	builder.WriteString(fmt.Sprintf("\tDeveloperName: %v\n", t.DeveloperName))
	builder.WriteString(fmt.Sprintf("\tDurableId: %v\n", t.DurableId))
	builder.WriteString(fmt.Sprintf("\tHeaderColor: %v\n", t.HeaderColor))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tIsLargeFormFactorSupported: %v\n", t.IsLargeFormFactorSupported))
	builder.WriteString(fmt.Sprintf("\tIsMediumFormFactorSupported: %v\n", t.IsMediumFormFactorSupported))
	builder.WriteString(fmt.Sprintf("\tIsNavAutoTempTabsDisabled: %v\n", t.IsNavAutoTempTabsDisabled))
	builder.WriteString(fmt.Sprintf("\tIsNavPersonalizationDisabled: %v\n", t.IsNavPersonalizationDisabled))
	builder.WriteString(fmt.Sprintf("\tIsOverrideOrgTheme: %v\n", t.IsOverrideOrgTheme))
	builder.WriteString(fmt.Sprintf("\tIsSmallFormFactorSupported: %v\n", t.IsSmallFormFactorSupported))
	builder.WriteString(fmt.Sprintf("\tLabel: %v\n", t.Label))
	builder.WriteString(fmt.Sprintf("\tLogoUrl: %v\n", t.LogoUrl))
	builder.WriteString(fmt.Sprintf("\tMasterLabel: %v\n", t.MasterLabel))
	builder.WriteString(fmt.Sprintf("\tNamespacePrefix: %v\n", t.NamespacePrefix))
	builder.WriteString(fmt.Sprintf("\tNavType: %v\n", t.NavType))
	builder.WriteString(fmt.Sprintf("\tUiType: %v\n", t.UiType))
	builder.WriteString(fmt.Sprintf("\tUtilityBar: %v\n", t.UtilityBar))

	return builder.String()
}

type AppDefinitionQueryResponse struct {
	BaseQuery
	Records []AppDefinition `json:"Records" force:"records"`
}
