// This file was generated for SObject ContentVersion, API Version v43.0 at 2018-07-30 03:47:31.168644141 -0400 EDT m=+17.511907673

package sobjects

import (
	"fmt"
	"strings"
)

type ContentVersion struct {
	BaseSObject
	Checksum               string `force:",omitempty"`
	ContentBodyId          string `force:",omitempty"`
	ContentDocumentId      string `force:",omitempty"`
	ContentLocation        string `force:",omitempty"`
	ContentModifiedById    string `force:",omitempty"`
	ContentModifiedDate    string `force:",omitempty"`
	ContentSize            int    `force:",omitempty"`
	ContentUrl             string `force:",omitempty"`
	CreatedById            string `force:",omitempty"`
	CreatedDate            string `force:",omitempty"`
	Description            string `force:",omitempty"`
	ExternalDataSourceId   string `force:",omitempty"`
	ExternalDocumentInfo1  string `force:",omitempty"`
	ExternalDocumentInfo2  string `force:",omitempty"`
	FeaturedContentBoost   int    `force:",omitempty"`
	FeaturedContentDate    string `force:",omitempty"`
	FileExtension          string `force:",omitempty"`
	FileType               string `force:",omitempty"`
	FirstPublishLocationId string `force:",omitempty"`
	Id                     string `force:",omitempty"`
	IsAssetEnabled         bool   `force:",omitempty"`
	IsDeleted              bool   `force:",omitempty"`
	IsLatest               bool   `force:",omitempty"`
	IsMajorVersion         bool   `force:",omitempty"`
	LastModifiedById       string `force:",omitempty"`
	LastModifiedDate       string `force:",omitempty"`
	NegativeRatingCount    int    `force:",omitempty"`
	Origin                 string `force:",omitempty"`
	OwnerId                string `force:",omitempty"`
	PathOnClient           string `force:",omitempty"`
	PositiveRatingCount    int    `force:",omitempty"`
	PublishStatus          string `force:",omitempty"`
	RatingCount            int    `force:",omitempty"`
	ReasonForChange        string `force:",omitempty"`
	SharingOption          string `force:",omitempty"`
	SharingPrivacy         string `force:",omitempty"`
	SystemModstamp         string `force:",omitempty"`
	TagCsv                 string `force:",omitempty"`
	TextPreview            string `force:",omitempty"`
	Title                  string `force:",omitempty"`
	VersionData            string `force:",omitempty"`
	VersionNumber          string `force:",omitempty"`
}

func (t *ContentVersion) ApiName() string {
	return "ContentVersion"
}

func (t *ContentVersion) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("ContentVersion #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tChecksum: %v\n", t.Checksum))
	builder.WriteString(fmt.Sprintf("\tContentBodyId: %v\n", t.ContentBodyId))
	builder.WriteString(fmt.Sprintf("\tContentDocumentId: %v\n", t.ContentDocumentId))
	builder.WriteString(fmt.Sprintf("\tContentLocation: %v\n", t.ContentLocation))
	builder.WriteString(fmt.Sprintf("\tContentModifiedById: %v\n", t.ContentModifiedById))
	builder.WriteString(fmt.Sprintf("\tContentModifiedDate: %v\n", t.ContentModifiedDate))
	builder.WriteString(fmt.Sprintf("\tContentSize: %v\n", t.ContentSize))
	builder.WriteString(fmt.Sprintf("\tContentUrl: %v\n", t.ContentUrl))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tDescription: %v\n", t.Description))
	builder.WriteString(fmt.Sprintf("\tExternalDataSourceId: %v\n", t.ExternalDataSourceId))
	builder.WriteString(fmt.Sprintf("\tExternalDocumentInfo1: %v\n", t.ExternalDocumentInfo1))
	builder.WriteString(fmt.Sprintf("\tExternalDocumentInfo2: %v\n", t.ExternalDocumentInfo2))
	builder.WriteString(fmt.Sprintf("\tFeaturedContentBoost: %v\n", t.FeaturedContentBoost))
	builder.WriteString(fmt.Sprintf("\tFeaturedContentDate: %v\n", t.FeaturedContentDate))
	builder.WriteString(fmt.Sprintf("\tFileExtension: %v\n", t.FileExtension))
	builder.WriteString(fmt.Sprintf("\tFileType: %v\n", t.FileType))
	builder.WriteString(fmt.Sprintf("\tFirstPublishLocationId: %v\n", t.FirstPublishLocationId))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tIsAssetEnabled: %v\n", t.IsAssetEnabled))
	builder.WriteString(fmt.Sprintf("\tIsDeleted: %v\n", t.IsDeleted))
	builder.WriteString(fmt.Sprintf("\tIsLatest: %v\n", t.IsLatest))
	builder.WriteString(fmt.Sprintf("\tIsMajorVersion: %v\n", t.IsMajorVersion))
	builder.WriteString(fmt.Sprintf("\tLastModifiedById: %v\n", t.LastModifiedById))
	builder.WriteString(fmt.Sprintf("\tLastModifiedDate: %v\n", t.LastModifiedDate))
	builder.WriteString(fmt.Sprintf("\tNegativeRatingCount: %v\n", t.NegativeRatingCount))
	builder.WriteString(fmt.Sprintf("\tOrigin: %v\n", t.Origin))
	builder.WriteString(fmt.Sprintf("\tOwnerId: %v\n", t.OwnerId))
	builder.WriteString(fmt.Sprintf("\tPathOnClient: %v\n", t.PathOnClient))
	builder.WriteString(fmt.Sprintf("\tPositiveRatingCount: %v\n", t.PositiveRatingCount))
	builder.WriteString(fmt.Sprintf("\tPublishStatus: %v\n", t.PublishStatus))
	builder.WriteString(fmt.Sprintf("\tRatingCount: %v\n", t.RatingCount))
	builder.WriteString(fmt.Sprintf("\tReasonForChange: %v\n", t.ReasonForChange))
	builder.WriteString(fmt.Sprintf("\tSharingOption: %v\n", t.SharingOption))
	builder.WriteString(fmt.Sprintf("\tSharingPrivacy: %v\n", t.SharingPrivacy))
	builder.WriteString(fmt.Sprintf("\tSystemModstamp: %v\n", t.SystemModstamp))
	builder.WriteString(fmt.Sprintf("\tTagCsv: %v\n", t.TagCsv))
	builder.WriteString(fmt.Sprintf("\tTextPreview: %v\n", t.TextPreview))
	builder.WriteString(fmt.Sprintf("\tTitle: %v\n", t.Title))
	builder.WriteString(fmt.Sprintf("\tVersionData: %v\n", t.VersionData))
	builder.WriteString(fmt.Sprintf("\tVersionNumber: %v\n", t.VersionNumber))

	return builder.String()
}

type ContentVersionQueryResponse struct {
	BaseQuery
	Records []ContentVersion `json:"Records" force:"records"`
}
