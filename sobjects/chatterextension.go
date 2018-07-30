// This file was generated for SObject ChatterExtension, API Version v43.0 at 2018-07-30 03:47:39.563671683 -0400 EDT m=+25.907250230

package sobjects

import (
	"fmt"
	"strings"
)

type ChatterExtension struct {
	BaseSObject
	CompositionComponentEnumOrId string `force:",omitempty"`
	CreatedById                  string `force:",omitempty"`
	CreatedDate                  string `force:",omitempty"`
	Description                  string `force:",omitempty"`
	DeveloperName                string `force:",omitempty"`
	ExtensionName                string `force:",omitempty"`
	HeaderText                   string `force:",omitempty"`
	HoverText                    string `force:",omitempty"`
	IconId                       string `force:",omitempty"`
	Id                           string `force:",omitempty"`
	IsDeleted                    bool   `force:",omitempty"`
	IsProtected                  bool   `force:",omitempty"`
	Language                     string `force:",omitempty"`
	LastModifiedById             string `force:",omitempty"`
	LastModifiedDate             string `force:",omitempty"`
	MasterLabel                  string `force:",omitempty"`
	NamespacePrefix              string `force:",omitempty"`
	RenderComponentEnumOrId      string `force:",omitempty"`
	SystemModstamp               string `force:",omitempty"`
	Type                         string `force:",omitempty"`
}

func (t *ChatterExtension) ApiName() string {
	return "ChatterExtension"
}

func (t *ChatterExtension) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("ChatterExtension #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tCompositionComponentEnumOrId: %v\n", t.CompositionComponentEnumOrId))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tDescription: %v\n", t.Description))
	builder.WriteString(fmt.Sprintf("\tDeveloperName: %v\n", t.DeveloperName))
	builder.WriteString(fmt.Sprintf("\tExtensionName: %v\n", t.ExtensionName))
	builder.WriteString(fmt.Sprintf("\tHeaderText: %v\n", t.HeaderText))
	builder.WriteString(fmt.Sprintf("\tHoverText: %v\n", t.HoverText))
	builder.WriteString(fmt.Sprintf("\tIconId: %v\n", t.IconId))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tIsDeleted: %v\n", t.IsDeleted))
	builder.WriteString(fmt.Sprintf("\tIsProtected: %v\n", t.IsProtected))
	builder.WriteString(fmt.Sprintf("\tLanguage: %v\n", t.Language))
	builder.WriteString(fmt.Sprintf("\tLastModifiedById: %v\n", t.LastModifiedById))
	builder.WriteString(fmt.Sprintf("\tLastModifiedDate: %v\n", t.LastModifiedDate))
	builder.WriteString(fmt.Sprintf("\tMasterLabel: %v\n", t.MasterLabel))
	builder.WriteString(fmt.Sprintf("\tNamespacePrefix: %v\n", t.NamespacePrefix))
	builder.WriteString(fmt.Sprintf("\tRenderComponentEnumOrId: %v\n", t.RenderComponentEnumOrId))
	builder.WriteString(fmt.Sprintf("\tSystemModstamp: %v\n", t.SystemModstamp))
	builder.WriteString(fmt.Sprintf("\tType: %v\n", t.Type))

	return builder.String()
}

type ChatterExtensionQueryResponse struct {
	BaseQuery
	Records []ChatterExtension `json:"Records" force:"records"`
}
