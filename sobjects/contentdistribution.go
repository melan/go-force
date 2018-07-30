// This file was generated for SObject ContentDistribution, API Version v43.0 at 2018-07-30 03:47:57.944795889 -0400 EDT m=+44.289064169

package sobjects

import (
	"fmt"
	"strings"
)

type ContentDistribution struct {
	BaseSObject
	ContentDocumentId                string `force:",omitempty"`
	ContentDownloadUrl               string `force:",omitempty"`
	ContentVersionId                 string `force:",omitempty"`
	CreatedById                      string `force:",omitempty"`
	CreatedDate                      string `force:",omitempty"`
	DistributionPublicUrl            string `force:",omitempty"`
	ExpiryDate                       string `force:",omitempty"`
	FirstViewDate                    string `force:",omitempty"`
	Id                               string `force:",omitempty"`
	IsDeleted                        bool   `force:",omitempty"`
	LastModifiedById                 string `force:",omitempty"`
	LastModifiedDate                 string `force:",omitempty"`
	LastViewDate                     string `force:",omitempty"`
	Name                             string `force:",omitempty"`
	OwnerId                          string `force:",omitempty"`
	Password                         string `force:",omitempty"`
	PdfDownloadUrl                   string `force:",omitempty"`
	PreferencesAllowOriginalDownload bool   `force:",omitempty"`
	PreferencesAllowPDFDownload      bool   `force:",omitempty"`
	PreferencesAllowViewInBrowser    bool   `force:",omitempty"`
	PreferencesExpires               bool   `force:",omitempty"`
	PreferencesLinkLatestVersion     bool   `force:",omitempty"`
	PreferencesNotifyOnVisit         bool   `force:",omitempty"`
	PreferencesNotifyRndtnComplete   bool   `force:",omitempty"`
	PreferencesPasswordRequired      bool   `force:",omitempty"`
	RelatedRecordId                  string `force:",omitempty"`
	SystemModstamp                   string `force:",omitempty"`
	ViewCount                        int    `force:",omitempty"`
}

func (t *ContentDistribution) ApiName() string {
	return "ContentDistribution"
}

func (t *ContentDistribution) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("ContentDistribution #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tContentDocumentId: %v\n", t.ContentDocumentId))
	builder.WriteString(fmt.Sprintf("\tContentDownloadUrl: %v\n", t.ContentDownloadUrl))
	builder.WriteString(fmt.Sprintf("\tContentVersionId: %v\n", t.ContentVersionId))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tDistributionPublicUrl: %v\n", t.DistributionPublicUrl))
	builder.WriteString(fmt.Sprintf("\tExpiryDate: %v\n", t.ExpiryDate))
	builder.WriteString(fmt.Sprintf("\tFirstViewDate: %v\n", t.FirstViewDate))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tIsDeleted: %v\n", t.IsDeleted))
	builder.WriteString(fmt.Sprintf("\tLastModifiedById: %v\n", t.LastModifiedById))
	builder.WriteString(fmt.Sprintf("\tLastModifiedDate: %v\n", t.LastModifiedDate))
	builder.WriteString(fmt.Sprintf("\tLastViewDate: %v\n", t.LastViewDate))
	builder.WriteString(fmt.Sprintf("\tName: %v\n", t.Name))
	builder.WriteString(fmt.Sprintf("\tOwnerId: %v\n", t.OwnerId))
	builder.WriteString(fmt.Sprintf("\tPassword: %v\n", t.Password))
	builder.WriteString(fmt.Sprintf("\tPdfDownloadUrl: %v\n", t.PdfDownloadUrl))
	builder.WriteString(fmt.Sprintf("\tPreferencesAllowOriginalDownload: %v\n", t.PreferencesAllowOriginalDownload))
	builder.WriteString(fmt.Sprintf("\tPreferencesAllowPDFDownload: %v\n", t.PreferencesAllowPDFDownload))
	builder.WriteString(fmt.Sprintf("\tPreferencesAllowViewInBrowser: %v\n", t.PreferencesAllowViewInBrowser))
	builder.WriteString(fmt.Sprintf("\tPreferencesExpires: %v\n", t.PreferencesExpires))
	builder.WriteString(fmt.Sprintf("\tPreferencesLinkLatestVersion: %v\n", t.PreferencesLinkLatestVersion))
	builder.WriteString(fmt.Sprintf("\tPreferencesNotifyOnVisit: %v\n", t.PreferencesNotifyOnVisit))
	builder.WriteString(fmt.Sprintf("\tPreferencesNotifyRndtnComplete: %v\n", t.PreferencesNotifyRndtnComplete))
	builder.WriteString(fmt.Sprintf("\tPreferencesPasswordRequired: %v\n", t.PreferencesPasswordRequired))
	builder.WriteString(fmt.Sprintf("\tRelatedRecordId: %v\n", t.RelatedRecordId))
	builder.WriteString(fmt.Sprintf("\tSystemModstamp: %v\n", t.SystemModstamp))
	builder.WriteString(fmt.Sprintf("\tViewCount: %v\n", t.ViewCount))

	return builder.String()
}

type ContentDistributionQueryResponse struct {
	BaseQuery
	Records []ContentDistribution `json:"Records" force:"records"`
}
