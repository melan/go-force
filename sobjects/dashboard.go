// This file was generated for SObject Dashboard, API Version v43.0 at 2018-07-30 03:47:23.697495368 -0400 EDT m=+10.040478554

package sobjects

import (
	"fmt"
	"strings"
)

type Dashboard struct {
	BaseSObject
	BackgroundDirection          string `force:",omitempty"`
	BackgroundEnd                int    `force:",omitempty"`
	BackgroundStart              int    `force:",omitempty"`
	ChartTheme                   string `force:",omitempty"`
	ColorPalette                 string `force:",omitempty"`
	CreatedById                  string `force:",omitempty"`
	CreatedDate                  string `force:",omitempty"`
	DashboardResultRefreshedDate string `force:",omitempty"`
	DashboardResultRunningUser   string `force:",omitempty"`
	Description                  string `force:",omitempty"`
	DeveloperName                string `force:",omitempty"`
	FolderId                     string `force:",omitempty"`
	FolderName                   string `force:",omitempty"`
	Id                           string `force:",omitempty"`
	IsDeleted                    bool   `force:",omitempty"`
	LastModifiedById             string `force:",omitempty"`
	LastModifiedDate             string `force:",omitempty"`
	LastReferencedDate           string `force:",omitempty"`
	LastViewedDate               string `force:",omitempty"`
	LeftSize                     string `force:",omitempty"`
	MiddleSize                   string `force:",omitempty"`
	NamespacePrefix              string `force:",omitempty"`
	RightSize                    string `force:",omitempty"`
	RunningUserId                string `force:",omitempty"`
	SystemModstamp               string `force:",omitempty"`
	TextColor                    int    `force:",omitempty"`
	Title                        string `force:",omitempty"`
	TitleColor                   int    `force:",omitempty"`
	TitleSize                    int    `force:",omitempty"`
	Type                         string `force:",omitempty"`
}

func (t *Dashboard) ApiName() string {
	return "Dashboard"
}

func (t *Dashboard) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("Dashboard #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tBackgroundDirection: %v\n", t.BackgroundDirection))
	builder.WriteString(fmt.Sprintf("\tBackgroundEnd: %v\n", t.BackgroundEnd))
	builder.WriteString(fmt.Sprintf("\tBackgroundStart: %v\n", t.BackgroundStart))
	builder.WriteString(fmt.Sprintf("\tChartTheme: %v\n", t.ChartTheme))
	builder.WriteString(fmt.Sprintf("\tColorPalette: %v\n", t.ColorPalette))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tDashboardResultRefreshedDate: %v\n", t.DashboardResultRefreshedDate))
	builder.WriteString(fmt.Sprintf("\tDashboardResultRunningUser: %v\n", t.DashboardResultRunningUser))
	builder.WriteString(fmt.Sprintf("\tDescription: %v\n", t.Description))
	builder.WriteString(fmt.Sprintf("\tDeveloperName: %v\n", t.DeveloperName))
	builder.WriteString(fmt.Sprintf("\tFolderId: %v\n", t.FolderId))
	builder.WriteString(fmt.Sprintf("\tFolderName: %v\n", t.FolderName))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tIsDeleted: %v\n", t.IsDeleted))
	builder.WriteString(fmt.Sprintf("\tLastModifiedById: %v\n", t.LastModifiedById))
	builder.WriteString(fmt.Sprintf("\tLastModifiedDate: %v\n", t.LastModifiedDate))
	builder.WriteString(fmt.Sprintf("\tLastReferencedDate: %v\n", t.LastReferencedDate))
	builder.WriteString(fmt.Sprintf("\tLastViewedDate: %v\n", t.LastViewedDate))
	builder.WriteString(fmt.Sprintf("\tLeftSize: %v\n", t.LeftSize))
	builder.WriteString(fmt.Sprintf("\tMiddleSize: %v\n", t.MiddleSize))
	builder.WriteString(fmt.Sprintf("\tNamespacePrefix: %v\n", t.NamespacePrefix))
	builder.WriteString(fmt.Sprintf("\tRightSize: %v\n", t.RightSize))
	builder.WriteString(fmt.Sprintf("\tRunningUserId: %v\n", t.RunningUserId))
	builder.WriteString(fmt.Sprintf("\tSystemModstamp: %v\n", t.SystemModstamp))
	builder.WriteString(fmt.Sprintf("\tTextColor: %v\n", t.TextColor))
	builder.WriteString(fmt.Sprintf("\tTitle: %v\n", t.Title))
	builder.WriteString(fmt.Sprintf("\tTitleColor: %v\n", t.TitleColor))
	builder.WriteString(fmt.Sprintf("\tTitleSize: %v\n", t.TitleSize))
	builder.WriteString(fmt.Sprintf("\tType: %v\n", t.Type))

	return builder.String()
}

type DashboardQueryResponse struct {
	BaseQuery
	Records []Dashboard `json:"Records" force:"records"`
}
