// This file was generated for SObject FieldDefinition, API Version v43.0 at 2018-07-30 03:47:31.444755595 -0400 EDT m=+17.788029488

package sobjects

import (
	"fmt"
	"strings"
)

type FieldDefinition struct {
	BaseSObject
	BusinessOwnerId              string `force:",omitempty"`
	BusinessStatus               string `force:",omitempty"`
	ControllingFieldDefinitionId string `force:",omitempty"`
	DataType                     string `force:",omitempty"`
	Description                  string `force:",omitempty"`
	DeveloperName                string `force:",omitempty"`
	DurableId                    string `force:",omitempty"`
	EntityDefinitionId           string `force:",omitempty"`
	ExtraTypeInfo                string `force:",omitempty"`
	Id                           string `force:",omitempty"`
	IsApiFilterable              bool   `force:",omitempty"`
	IsApiGroupable               bool   `force:",omitempty"`
	IsApiSortable                bool   `force:",omitempty"`
	IsCalculated                 bool   `force:",omitempty"`
	IsCompactLayoutable          bool   `force:",omitempty"`
	IsCompound                   bool   `force:",omitempty"`
	IsFieldHistoryTracked        bool   `force:",omitempty"`
	IsHighScaleNumber            bool   `force:",omitempty"`
	IsHtmlFormatted              bool   `force:",omitempty"`
	IsIndexed                    bool   `force:",omitempty"`
	IsListFilterable             bool   `force:",omitempty"`
	IsListSortable               bool   `force:",omitempty"`
	IsListVisible                bool   `force:",omitempty"`
	IsNameField                  bool   `force:",omitempty"`
	IsNillable                   bool   `force:",omitempty"`
	IsPolymorphicForeignKey      bool   `force:",omitempty"`
	IsSearchPrefilterable        bool   `force:",omitempty"`
	IsWorkflowFilterable         bool   `force:",omitempty"`
	Label                        string `force:",omitempty"`
	LastModifiedById             string `force:",omitempty"`
	LastModifiedDate             string `force:",omitempty"`
	Length                       int    `force:",omitempty"`
	MasterLabel                  string `force:",omitempty"`
	NamespacePrefix              string `force:",omitempty"`
	Precision                    int    `force:",omitempty"`
	PublisherId                  string `force:",omitempty"`
	QualifiedApiName             string `force:",omitempty"`
	ReferenceTargetField         string `force:",omitempty"`
	ReferenceTo                  string `force:",omitempty"`
	RelationshipName             string `force:",omitempty"`
	RunningUserFieldAccessId     string `force:",omitempty"`
	Scale                        int    `force:",omitempty"`
	SecurityClassification       string `force:",omitempty"`
	ServiceDataTypeId            string `force:",omitempty"`
	ValueTypeId                  string `force:",omitempty"`
}

func (t *FieldDefinition) ApiName() string {
	return "FieldDefinition"
}

func (t *FieldDefinition) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("FieldDefinition #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tBusinessOwnerId: %v\n", t.BusinessOwnerId))
	builder.WriteString(fmt.Sprintf("\tBusinessStatus: %v\n", t.BusinessStatus))
	builder.WriteString(fmt.Sprintf("\tControllingFieldDefinitionId: %v\n", t.ControllingFieldDefinitionId))
	builder.WriteString(fmt.Sprintf("\tDataType: %v\n", t.DataType))
	builder.WriteString(fmt.Sprintf("\tDescription: %v\n", t.Description))
	builder.WriteString(fmt.Sprintf("\tDeveloperName: %v\n", t.DeveloperName))
	builder.WriteString(fmt.Sprintf("\tDurableId: %v\n", t.DurableId))
	builder.WriteString(fmt.Sprintf("\tEntityDefinitionId: %v\n", t.EntityDefinitionId))
	builder.WriteString(fmt.Sprintf("\tExtraTypeInfo: %v\n", t.ExtraTypeInfo))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tIsApiFilterable: %v\n", t.IsApiFilterable))
	builder.WriteString(fmt.Sprintf("\tIsApiGroupable: %v\n", t.IsApiGroupable))
	builder.WriteString(fmt.Sprintf("\tIsApiSortable: %v\n", t.IsApiSortable))
	builder.WriteString(fmt.Sprintf("\tIsCalculated: %v\n", t.IsCalculated))
	builder.WriteString(fmt.Sprintf("\tIsCompactLayoutable: %v\n", t.IsCompactLayoutable))
	builder.WriteString(fmt.Sprintf("\tIsCompound: %v\n", t.IsCompound))
	builder.WriteString(fmt.Sprintf("\tIsFieldHistoryTracked: %v\n", t.IsFieldHistoryTracked))
	builder.WriteString(fmt.Sprintf("\tIsHighScaleNumber: %v\n", t.IsHighScaleNumber))
	builder.WriteString(fmt.Sprintf("\tIsHtmlFormatted: %v\n", t.IsHtmlFormatted))
	builder.WriteString(fmt.Sprintf("\tIsIndexed: %v\n", t.IsIndexed))
	builder.WriteString(fmt.Sprintf("\tIsListFilterable: %v\n", t.IsListFilterable))
	builder.WriteString(fmt.Sprintf("\tIsListSortable: %v\n", t.IsListSortable))
	builder.WriteString(fmt.Sprintf("\tIsListVisible: %v\n", t.IsListVisible))
	builder.WriteString(fmt.Sprintf("\tIsNameField: %v\n", t.IsNameField))
	builder.WriteString(fmt.Sprintf("\tIsNillable: %v\n", t.IsNillable))
	builder.WriteString(fmt.Sprintf("\tIsPolymorphicForeignKey: %v\n", t.IsPolymorphicForeignKey))
	builder.WriteString(fmt.Sprintf("\tIsSearchPrefilterable: %v\n", t.IsSearchPrefilterable))
	builder.WriteString(fmt.Sprintf("\tIsWorkflowFilterable: %v\n", t.IsWorkflowFilterable))
	builder.WriteString(fmt.Sprintf("\tLabel: %v\n", t.Label))
	builder.WriteString(fmt.Sprintf("\tLastModifiedById: %v\n", t.LastModifiedById))
	builder.WriteString(fmt.Sprintf("\tLastModifiedDate: %v\n", t.LastModifiedDate))
	builder.WriteString(fmt.Sprintf("\tLength: %v\n", t.Length))
	builder.WriteString(fmt.Sprintf("\tMasterLabel: %v\n", t.MasterLabel))
	builder.WriteString(fmt.Sprintf("\tNamespacePrefix: %v\n", t.NamespacePrefix))
	builder.WriteString(fmt.Sprintf("\tPrecision: %v\n", t.Precision))
	builder.WriteString(fmt.Sprintf("\tPublisherId: %v\n", t.PublisherId))
	builder.WriteString(fmt.Sprintf("\tQualifiedApiName: %v\n", t.QualifiedApiName))
	builder.WriteString(fmt.Sprintf("\tReferenceTargetField: %v\n", t.ReferenceTargetField))
	builder.WriteString(fmt.Sprintf("\tReferenceTo: %v\n", t.ReferenceTo))
	builder.WriteString(fmt.Sprintf("\tRelationshipName: %v\n", t.RelationshipName))
	builder.WriteString(fmt.Sprintf("\tRunningUserFieldAccessId: %v\n", t.RunningUserFieldAccessId))
	builder.WriteString(fmt.Sprintf("\tScale: %v\n", t.Scale))
	builder.WriteString(fmt.Sprintf("\tSecurityClassification: %v\n", t.SecurityClassification))
	builder.WriteString(fmt.Sprintf("\tServiceDataTypeId: %v\n", t.ServiceDataTypeId))
	builder.WriteString(fmt.Sprintf("\tValueTypeId: %v\n", t.ValueTypeId))

	return builder.String()
}

type FieldDefinitionQueryResponse struct {
	BaseQuery
	Records []FieldDefinition `json:"Records" force:"records"`
}
