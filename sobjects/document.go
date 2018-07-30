// This file was generated for SObject Document, API Version v43.0 at 2018-07-30 03:47:17.220192274 -0400 EDT m=+3.562932403

package sobjects

import (
	"fmt"
	"strings"
)

type Document struct {
	BaseSObject
	AuthorId           string `force:",omitempty"`
	Body               string `force:",omitempty"`
	BodyLength         int    `force:",omitempty"`
	ContentType        string `force:",omitempty"`
	CreatedById        string `force:",omitempty"`
	CreatedDate        string `force:",omitempty"`
	Description        string `force:",omitempty"`
	DeveloperName      string `force:",omitempty"`
	FolderId           string `force:",omitempty"`
	Id                 string `force:",omitempty"`
	IsBodySearchable   bool   `force:",omitempty"`
	IsDeleted          bool   `force:",omitempty"`
	IsInternalUseOnly  bool   `force:",omitempty"`
	IsPublic           bool   `force:",omitempty"`
	Keywords           string `force:",omitempty"`
	LastModifiedById   string `force:",omitempty"`
	LastModifiedDate   string `force:",omitempty"`
	LastReferencedDate string `force:",omitempty"`
	LastViewedDate     string `force:",omitempty"`
	Name               string `force:",omitempty"`
	NamespacePrefix    string `force:",omitempty"`
	SystemModstamp     string `force:",omitempty"`
	Type               string `force:",omitempty"`
	Url                string `force:",omitempty"`
}

func (t *Document) ApiName() string {
	return "Document"
}

func (t *Document) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("Document #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tAuthorId: %v\n", t.AuthorId))
	builder.WriteString(fmt.Sprintf("\tBody: %v\n", t.Body))
	builder.WriteString(fmt.Sprintf("\tBodyLength: %v\n", t.BodyLength))
	builder.WriteString(fmt.Sprintf("\tContentType: %v\n", t.ContentType))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tDescription: %v\n", t.Description))
	builder.WriteString(fmt.Sprintf("\tDeveloperName: %v\n", t.DeveloperName))
	builder.WriteString(fmt.Sprintf("\tFolderId: %v\n", t.FolderId))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tIsBodySearchable: %v\n", t.IsBodySearchable))
	builder.WriteString(fmt.Sprintf("\tIsDeleted: %v\n", t.IsDeleted))
	builder.WriteString(fmt.Sprintf("\tIsInternalUseOnly: %v\n", t.IsInternalUseOnly))
	builder.WriteString(fmt.Sprintf("\tIsPublic: %v\n", t.IsPublic))
	builder.WriteString(fmt.Sprintf("\tKeywords: %v\n", t.Keywords))
	builder.WriteString(fmt.Sprintf("\tLastModifiedById: %v\n", t.LastModifiedById))
	builder.WriteString(fmt.Sprintf("\tLastModifiedDate: %v\n", t.LastModifiedDate))
	builder.WriteString(fmt.Sprintf("\tLastReferencedDate: %v\n", t.LastReferencedDate))
	builder.WriteString(fmt.Sprintf("\tLastViewedDate: %v\n", t.LastViewedDate))
	builder.WriteString(fmt.Sprintf("\tName: %v\n", t.Name))
	builder.WriteString(fmt.Sprintf("\tNamespacePrefix: %v\n", t.NamespacePrefix))
	builder.WriteString(fmt.Sprintf("\tSystemModstamp: %v\n", t.SystemModstamp))
	builder.WriteString(fmt.Sprintf("\tType: %v\n", t.Type))
	builder.WriteString(fmt.Sprintf("\tUrl: %v\n", t.Url))

	return builder.String()
}

type DocumentQueryResponse struct {
	BaseQuery
	Records []Document `json:"Records" force:"records"`
}
