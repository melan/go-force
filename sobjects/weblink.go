// This file was generated for SObject WebLink, API Version v43.0 at 2018-07-30 03:47:53.585752942 -0400 EDT m=+39.929857653

package sobjects

import (
	"fmt"
	"strings"
)

type WebLink struct {
	BaseSObject
	CreatedById         string `force:",omitempty"`
	CreatedDate         string `force:",omitempty"`
	Description         string `force:",omitempty"`
	DisplayType         string `force:",omitempty"`
	EncodingKey         string `force:",omitempty"`
	HasMenubar          bool   `force:",omitempty"`
	HasScrollbars       bool   `force:",omitempty"`
	HasToolbar          bool   `force:",omitempty"`
	Height              int    `force:",omitempty"`
	Id                  string `force:",omitempty"`
	IsProtected         bool   `force:",omitempty"`
	IsResizable         bool   `force:",omitempty"`
	LastModifiedById    string `force:",omitempty"`
	LastModifiedDate    string `force:",omitempty"`
	LinkType            string `force:",omitempty"`
	MasterLabel         string `force:",omitempty"`
	Name                string `force:",omitempty"`
	NamespacePrefix     string `force:",omitempty"`
	OpenType            string `force:",omitempty"`
	PageOrSobjectType   string `force:",omitempty"`
	Position            string `force:",omitempty"`
	RequireRowSelection bool   `force:",omitempty"`
	ScontrolId          string `force:",omitempty"`
	ShowsLocation       bool   `force:",omitempty"`
	ShowsStatus         bool   `force:",omitempty"`
	SystemModstamp      string `force:",omitempty"`
	Url                 string `force:",omitempty"`
	Width               int    `force:",omitempty"`
}

func (t *WebLink) ApiName() string {
	return "WebLink"
}

func (t *WebLink) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("WebLink #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tDescription: %v\n", t.Description))
	builder.WriteString(fmt.Sprintf("\tDisplayType: %v\n", t.DisplayType))
	builder.WriteString(fmt.Sprintf("\tEncodingKey: %v\n", t.EncodingKey))
	builder.WriteString(fmt.Sprintf("\tHasMenubar: %v\n", t.HasMenubar))
	builder.WriteString(fmt.Sprintf("\tHasScrollbars: %v\n", t.HasScrollbars))
	builder.WriteString(fmt.Sprintf("\tHasToolbar: %v\n", t.HasToolbar))
	builder.WriteString(fmt.Sprintf("\tHeight: %v\n", t.Height))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tIsProtected: %v\n", t.IsProtected))
	builder.WriteString(fmt.Sprintf("\tIsResizable: %v\n", t.IsResizable))
	builder.WriteString(fmt.Sprintf("\tLastModifiedById: %v\n", t.LastModifiedById))
	builder.WriteString(fmt.Sprintf("\tLastModifiedDate: %v\n", t.LastModifiedDate))
	builder.WriteString(fmt.Sprintf("\tLinkType: %v\n", t.LinkType))
	builder.WriteString(fmt.Sprintf("\tMasterLabel: %v\n", t.MasterLabel))
	builder.WriteString(fmt.Sprintf("\tName: %v\n", t.Name))
	builder.WriteString(fmt.Sprintf("\tNamespacePrefix: %v\n", t.NamespacePrefix))
	builder.WriteString(fmt.Sprintf("\tOpenType: %v\n", t.OpenType))
	builder.WriteString(fmt.Sprintf("\tPageOrSobjectType: %v\n", t.PageOrSobjectType))
	builder.WriteString(fmt.Sprintf("\tPosition: %v\n", t.Position))
	builder.WriteString(fmt.Sprintf("\tRequireRowSelection: %v\n", t.RequireRowSelection))
	builder.WriteString(fmt.Sprintf("\tScontrolId: %v\n", t.ScontrolId))
	builder.WriteString(fmt.Sprintf("\tShowsLocation: %v\n", t.ShowsLocation))
	builder.WriteString(fmt.Sprintf("\tShowsStatus: %v\n", t.ShowsStatus))
	builder.WriteString(fmt.Sprintf("\tSystemModstamp: %v\n", t.SystemModstamp))
	builder.WriteString(fmt.Sprintf("\tUrl: %v\n", t.Url))
	builder.WriteString(fmt.Sprintf("\tWidth: %v\n", t.Width))

	return builder.String()
}

type WebLinkQueryResponse struct {
	BaseQuery
	Records []WebLink `json:"Records" force:"records"`
}
