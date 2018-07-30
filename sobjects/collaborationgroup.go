// This file was generated for SObject CollaborationGroup, API Version v43.0 at 2018-07-30 03:48:03.521256569 -0400 EDT m=+49.865734099

package sobjects

import (
	"fmt"
	"strings"
)

type CollaborationGroup struct {
	BaseSObject
	AnnouncementId         string `force:",omitempty"`
	BannerPhotoUrl         string `force:",omitempty"`
	CanHaveGuests          bool   `force:",omitempty"`
	CollaborationType      string `force:",omitempty"`
	CreatedById            string `force:",omitempty"`
	CreatedDate            string `force:",omitempty"`
	Description            string `force:",omitempty"`
	FullPhotoUrl           string `force:",omitempty"`
	GroupEmail             string `force:",omitempty"`
	HasPrivateFieldsAccess bool   `force:",omitempty"`
	Id                     string `force:",omitempty"`
	InformationBody        string `force:",omitempty"`
	InformationTitle       string `force:",omitempty"`
	IsArchived             bool   `force:",omitempty"`
	IsAutoArchiveDisabled  bool   `force:",omitempty"`
	IsBroadcast            bool   `force:",omitempty"`
	LastFeedModifiedDate   string `force:",omitempty"`
	LastModifiedById       string `force:",omitempty"`
	LastModifiedDate       string `force:",omitempty"`
	LastReferencedDate     string `force:",omitempty"`
	LastViewedDate         string `force:",omitempty"`
	MediumPhotoUrl         string `force:",omitempty"`
	MemberCount            int    `force:",omitempty"`
	Name                   string `force:",omitempty"`
	OwnerId                string `force:",omitempty"`
	SmallPhotoUrl          string `force:",omitempty"`
	SystemModstamp         string `force:",omitempty"`
}

func (t *CollaborationGroup) ApiName() string {
	return "CollaborationGroup"
}

func (t *CollaborationGroup) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("CollaborationGroup #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tAnnouncementId: %v\n", t.AnnouncementId))
	builder.WriteString(fmt.Sprintf("\tBannerPhotoUrl: %v\n", t.BannerPhotoUrl))
	builder.WriteString(fmt.Sprintf("\tCanHaveGuests: %v\n", t.CanHaveGuests))
	builder.WriteString(fmt.Sprintf("\tCollaborationType: %v\n", t.CollaborationType))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tDescription: %v\n", t.Description))
	builder.WriteString(fmt.Sprintf("\tFullPhotoUrl: %v\n", t.FullPhotoUrl))
	builder.WriteString(fmt.Sprintf("\tGroupEmail: %v\n", t.GroupEmail))
	builder.WriteString(fmt.Sprintf("\tHasPrivateFieldsAccess: %v\n", t.HasPrivateFieldsAccess))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tInformationBody: %v\n", t.InformationBody))
	builder.WriteString(fmt.Sprintf("\tInformationTitle: %v\n", t.InformationTitle))
	builder.WriteString(fmt.Sprintf("\tIsArchived: %v\n", t.IsArchived))
	builder.WriteString(fmt.Sprintf("\tIsAutoArchiveDisabled: %v\n", t.IsAutoArchiveDisabled))
	builder.WriteString(fmt.Sprintf("\tIsBroadcast: %v\n", t.IsBroadcast))
	builder.WriteString(fmt.Sprintf("\tLastFeedModifiedDate: %v\n", t.LastFeedModifiedDate))
	builder.WriteString(fmt.Sprintf("\tLastModifiedById: %v\n", t.LastModifiedById))
	builder.WriteString(fmt.Sprintf("\tLastModifiedDate: %v\n", t.LastModifiedDate))
	builder.WriteString(fmt.Sprintf("\tLastReferencedDate: %v\n", t.LastReferencedDate))
	builder.WriteString(fmt.Sprintf("\tLastViewedDate: %v\n", t.LastViewedDate))
	builder.WriteString(fmt.Sprintf("\tMediumPhotoUrl: %v\n", t.MediumPhotoUrl))
	builder.WriteString(fmt.Sprintf("\tMemberCount: %v\n", t.MemberCount))
	builder.WriteString(fmt.Sprintf("\tName: %v\n", t.Name))
	builder.WriteString(fmt.Sprintf("\tOwnerId: %v\n", t.OwnerId))
	builder.WriteString(fmt.Sprintf("\tSmallPhotoUrl: %v\n", t.SmallPhotoUrl))
	builder.WriteString(fmt.Sprintf("\tSystemModstamp: %v\n", t.SystemModstamp))

	return builder.String()
}

type CollaborationGroupQueryResponse struct {
	BaseQuery
	Records []CollaborationGroup `json:"Records" force:"records"`
}
