// This file was generated for SObject EmbeddedServiceDetail, API Version v43.0 at 2018-07-30 03:47:23.375298708 -0400 EDT m=+9.718269803

package sobjects

import (
	"fmt"
	"strings"
)

type EmbeddedServiceDetail struct {
	BaseSObject
	AvatarImg                     string `force:",omitempty"`
	CancelApptBookingFlowName     string `force:",omitempty"`
	ContrastInvertedColor         string `force:",omitempty"`
	ContrastPrimaryColor          string `force:",omitempty"`
	CustomMinimizedComponent      string `force:",omitempty"`
	CustomPrechatComponent        string `force:",omitempty"`
	DurableId                     string `force:",omitempty"`
	FieldServiceConfirmCardImg    string `force:",omitempty"`
	FieldServiceHomeImg           string `force:",omitempty"`
	FieldServiceLogoImg           string `force:",omitempty"`
	FlowDeveloperName             string `force:",omitempty"`
	Font                          string `force:",omitempty"`
	FontSize                      string `force:",omitempty"`
	HeaderBackgroundImg           string `force:",omitempty"`
	Height                        int    `force:",omitempty"`
	Id                            string `force:",omitempty"`
	IsFieldServiceEnabled         bool   `force:",omitempty"`
	IsLiveAgentEnabled            bool   `force:",omitempty"`
	IsOfflineCaseEnabled          bool   `force:",omitempty"`
	IsPrechatEnabled              bool   `force:",omitempty"`
	IsQueuePositionEnabled        bool   `force:",omitempty"`
	ModifyApptBookingFlowName     string `force:",omitempty"`
	NavBarColor                   string `force:",omitempty"`
	OfflineCaseBackgroundImg      string `force:",omitempty"`
	PrechatBackgroundImg          string `force:",omitempty"`
	PrimaryColor                  string `force:",omitempty"`
	SecondaryColor                string `force:",omitempty"`
	ShouldHideAuthDialog          bool   `force:",omitempty"`
	ShouldShowExistingAppointment bool   `force:",omitempty"`
	ShouldShowNewAppointment      bool   `force:",omitempty"`
	Site                          string `force:",omitempty"`
	SmallCompanyLogoImg           string `force:",omitempty"`
	WaitingStateBackgroundImg     string `force:",omitempty"`
	Width                         int    `force:",omitempty"`
}

func (t *EmbeddedServiceDetail) ApiName() string {
	return "EmbeddedServiceDetail"
}

func (t *EmbeddedServiceDetail) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("EmbeddedServiceDetail #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tAvatarImg: %v\n", t.AvatarImg))
	builder.WriteString(fmt.Sprintf("\tCancelApptBookingFlowName: %v\n", t.CancelApptBookingFlowName))
	builder.WriteString(fmt.Sprintf("\tContrastInvertedColor: %v\n", t.ContrastInvertedColor))
	builder.WriteString(fmt.Sprintf("\tContrastPrimaryColor: %v\n", t.ContrastPrimaryColor))
	builder.WriteString(fmt.Sprintf("\tCustomMinimizedComponent: %v\n", t.CustomMinimizedComponent))
	builder.WriteString(fmt.Sprintf("\tCustomPrechatComponent: %v\n", t.CustomPrechatComponent))
	builder.WriteString(fmt.Sprintf("\tDurableId: %v\n", t.DurableId))
	builder.WriteString(fmt.Sprintf("\tFieldServiceConfirmCardImg: %v\n", t.FieldServiceConfirmCardImg))
	builder.WriteString(fmt.Sprintf("\tFieldServiceHomeImg: %v\n", t.FieldServiceHomeImg))
	builder.WriteString(fmt.Sprintf("\tFieldServiceLogoImg: %v\n", t.FieldServiceLogoImg))
	builder.WriteString(fmt.Sprintf("\tFlowDeveloperName: %v\n", t.FlowDeveloperName))
	builder.WriteString(fmt.Sprintf("\tFont: %v\n", t.Font))
	builder.WriteString(fmt.Sprintf("\tFontSize: %v\n", t.FontSize))
	builder.WriteString(fmt.Sprintf("\tHeaderBackgroundImg: %v\n", t.HeaderBackgroundImg))
	builder.WriteString(fmt.Sprintf("\tHeight: %v\n", t.Height))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tIsFieldServiceEnabled: %v\n", t.IsFieldServiceEnabled))
	builder.WriteString(fmt.Sprintf("\tIsLiveAgentEnabled: %v\n", t.IsLiveAgentEnabled))
	builder.WriteString(fmt.Sprintf("\tIsOfflineCaseEnabled: %v\n", t.IsOfflineCaseEnabled))
	builder.WriteString(fmt.Sprintf("\tIsPrechatEnabled: %v\n", t.IsPrechatEnabled))
	builder.WriteString(fmt.Sprintf("\tIsQueuePositionEnabled: %v\n", t.IsQueuePositionEnabled))
	builder.WriteString(fmt.Sprintf("\tModifyApptBookingFlowName: %v\n", t.ModifyApptBookingFlowName))
	builder.WriteString(fmt.Sprintf("\tNavBarColor: %v\n", t.NavBarColor))
	builder.WriteString(fmt.Sprintf("\tOfflineCaseBackgroundImg: %v\n", t.OfflineCaseBackgroundImg))
	builder.WriteString(fmt.Sprintf("\tPrechatBackgroundImg: %v\n", t.PrechatBackgroundImg))
	builder.WriteString(fmt.Sprintf("\tPrimaryColor: %v\n", t.PrimaryColor))
	builder.WriteString(fmt.Sprintf("\tSecondaryColor: %v\n", t.SecondaryColor))
	builder.WriteString(fmt.Sprintf("\tShouldHideAuthDialog: %v\n", t.ShouldHideAuthDialog))
	builder.WriteString(fmt.Sprintf("\tShouldShowExistingAppointment: %v\n", t.ShouldShowExistingAppointment))
	builder.WriteString(fmt.Sprintf("\tShouldShowNewAppointment: %v\n", t.ShouldShowNewAppointment))
	builder.WriteString(fmt.Sprintf("\tSite: %v\n", t.Site))
	builder.WriteString(fmt.Sprintf("\tSmallCompanyLogoImg: %v\n", t.SmallCompanyLogoImg))
	builder.WriteString(fmt.Sprintf("\tWaitingStateBackgroundImg: %v\n", t.WaitingStateBackgroundImg))
	builder.WriteString(fmt.Sprintf("\tWidth: %v\n", t.Width))

	return builder.String()
}

type EmbeddedServiceDetailQueryResponse struct {
	BaseQuery
	Records []EmbeddedServiceDetail `json:"Records" force:"records"`
}
