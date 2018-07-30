// This file was generated for SObject EntityParticle, API Version v43.0 at 2018-07-30 03:47:30.900090482 -0400 EDT m=+17.243343938

package sobjects

import (
	"fmt"
	"strings"
)

type EntityParticle struct {
	BaseSObject
	ByteLength                 int    `force:",omitempty"`
	DataType                   string `force:",omitempty"`
	DefaultValueFormula        string `force:",omitempty"`
	DeveloperName              string `force:",omitempty"`
	Digits                     int    `force:",omitempty"`
	DurableId                  string `force:",omitempty"`
	EntityDefinitionId         string `force:",omitempty"`
	ExtraTypeInfo              string `force:",omitempty"`
	FieldDefinitionId          string `force:",omitempty"`
	Id                         string `force:",omitempty"`
	InlineHelpText             string `force:",omitempty"`
	IsApiFilterable            bool   `force:",omitempty"`
	IsApiGroupable             bool   `force:",omitempty"`
	IsApiSortable              bool   `force:",omitempty"`
	IsAutonumber               bool   `force:",omitempty"`
	IsCalculated               bool   `force:",omitempty"`
	IsCaseSensitive            bool   `force:",omitempty"`
	IsCompactLayoutable        bool   `force:",omitempty"`
	IsComponent                bool   `force:",omitempty"`
	IsCompound                 bool   `force:",omitempty"`
	IsCreatable                bool   `force:",omitempty"`
	IsDefaultedOnCreate        bool   `force:",omitempty"`
	IsDependentPicklist        bool   `force:",omitempty"`
	IsDeprecatedAndHidden      bool   `force:",omitempty"`
	IsDisplayLocationInDecimal bool   `force:",omitempty"`
	IsEncrypted                bool   `force:",omitempty"`
	IsFieldHistoryTracked      bool   `force:",omitempty"`
	IsHighScaleNumber          bool   `force:",omitempty"`
	IsHtmlFormatted            bool   `force:",omitempty"`
	IsIdLookup                 bool   `force:",omitempty"`
	IsLayoutable               bool   `force:",omitempty"`
	IsListVisible              bool   `force:",omitempty"`
	IsNameField                bool   `force:",omitempty"`
	IsNamePointing             bool   `force:",omitempty"`
	IsNillable                 bool   `force:",omitempty"`
	IsPermissionable           bool   `force:",omitempty"`
	IsUnique                   bool   `force:",omitempty"`
	IsUpdatable                bool   `force:",omitempty"`
	IsWorkflowFilterable       bool   `force:",omitempty"`
	IsWriteRequiresMasterRead  bool   `force:",omitempty"`
	Label                      string `force:",omitempty"`
	Length                     int    `force:",omitempty"`
	Mask                       string `force:",omitempty"`
	MaskType                   string `force:",omitempty"`
	MasterLabel                string `force:",omitempty"`
	Name                       string `force:",omitempty"`
	NamespacePrefix            string `force:",omitempty"`
	Precision                  int    `force:",omitempty"`
	QualifiedApiName           string `force:",omitempty"`
	ReferenceTargetField       string `force:",omitempty"`
	ReferenceTo                string `force:",omitempty"`
	RelationshipName           string `force:",omitempty"`
	RelationshipOrder          int    `force:",omitempty"`
	Scale                      int    `force:",omitempty"`
	ServiceDataTypeId          string `force:",omitempty"`
	ValueTypeId                string `force:",omitempty"`
}

func (t *EntityParticle) ApiName() string {
	return "EntityParticle"
}

func (t *EntityParticle) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("EntityParticle #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tByteLength: %v\n", t.ByteLength))
	builder.WriteString(fmt.Sprintf("\tDataType: %v\n", t.DataType))
	builder.WriteString(fmt.Sprintf("\tDefaultValueFormula: %v\n", t.DefaultValueFormula))
	builder.WriteString(fmt.Sprintf("\tDeveloperName: %v\n", t.DeveloperName))
	builder.WriteString(fmt.Sprintf("\tDigits: %v\n", t.Digits))
	builder.WriteString(fmt.Sprintf("\tDurableId: %v\n", t.DurableId))
	builder.WriteString(fmt.Sprintf("\tEntityDefinitionId: %v\n", t.EntityDefinitionId))
	builder.WriteString(fmt.Sprintf("\tExtraTypeInfo: %v\n", t.ExtraTypeInfo))
	builder.WriteString(fmt.Sprintf("\tFieldDefinitionId: %v\n", t.FieldDefinitionId))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tInlineHelpText: %v\n", t.InlineHelpText))
	builder.WriteString(fmt.Sprintf("\tIsApiFilterable: %v\n", t.IsApiFilterable))
	builder.WriteString(fmt.Sprintf("\tIsApiGroupable: %v\n", t.IsApiGroupable))
	builder.WriteString(fmt.Sprintf("\tIsApiSortable: %v\n", t.IsApiSortable))
	builder.WriteString(fmt.Sprintf("\tIsAutonumber: %v\n", t.IsAutonumber))
	builder.WriteString(fmt.Sprintf("\tIsCalculated: %v\n", t.IsCalculated))
	builder.WriteString(fmt.Sprintf("\tIsCaseSensitive: %v\n", t.IsCaseSensitive))
	builder.WriteString(fmt.Sprintf("\tIsCompactLayoutable: %v\n", t.IsCompactLayoutable))
	builder.WriteString(fmt.Sprintf("\tIsComponent: %v\n", t.IsComponent))
	builder.WriteString(fmt.Sprintf("\tIsCompound: %v\n", t.IsCompound))
	builder.WriteString(fmt.Sprintf("\tIsCreatable: %v\n", t.IsCreatable))
	builder.WriteString(fmt.Sprintf("\tIsDefaultedOnCreate: %v\n", t.IsDefaultedOnCreate))
	builder.WriteString(fmt.Sprintf("\tIsDependentPicklist: %v\n", t.IsDependentPicklist))
	builder.WriteString(fmt.Sprintf("\tIsDeprecatedAndHidden: %v\n", t.IsDeprecatedAndHidden))
	builder.WriteString(fmt.Sprintf("\tIsDisplayLocationInDecimal: %v\n", t.IsDisplayLocationInDecimal))
	builder.WriteString(fmt.Sprintf("\tIsEncrypted: %v\n", t.IsEncrypted))
	builder.WriteString(fmt.Sprintf("\tIsFieldHistoryTracked: %v\n", t.IsFieldHistoryTracked))
	builder.WriteString(fmt.Sprintf("\tIsHighScaleNumber: %v\n", t.IsHighScaleNumber))
	builder.WriteString(fmt.Sprintf("\tIsHtmlFormatted: %v\n", t.IsHtmlFormatted))
	builder.WriteString(fmt.Sprintf("\tIsIdLookup: %v\n", t.IsIdLookup))
	builder.WriteString(fmt.Sprintf("\tIsLayoutable: %v\n", t.IsLayoutable))
	builder.WriteString(fmt.Sprintf("\tIsListVisible: %v\n", t.IsListVisible))
	builder.WriteString(fmt.Sprintf("\tIsNameField: %v\n", t.IsNameField))
	builder.WriteString(fmt.Sprintf("\tIsNamePointing: %v\n", t.IsNamePointing))
	builder.WriteString(fmt.Sprintf("\tIsNillable: %v\n", t.IsNillable))
	builder.WriteString(fmt.Sprintf("\tIsPermissionable: %v\n", t.IsPermissionable))
	builder.WriteString(fmt.Sprintf("\tIsUnique: %v\n", t.IsUnique))
	builder.WriteString(fmt.Sprintf("\tIsUpdatable: %v\n", t.IsUpdatable))
	builder.WriteString(fmt.Sprintf("\tIsWorkflowFilterable: %v\n", t.IsWorkflowFilterable))
	builder.WriteString(fmt.Sprintf("\tIsWriteRequiresMasterRead: %v\n", t.IsWriteRequiresMasterRead))
	builder.WriteString(fmt.Sprintf("\tLabel: %v\n", t.Label))
	builder.WriteString(fmt.Sprintf("\tLength: %v\n", t.Length))
	builder.WriteString(fmt.Sprintf("\tMask: %v\n", t.Mask))
	builder.WriteString(fmt.Sprintf("\tMaskType: %v\n", t.MaskType))
	builder.WriteString(fmt.Sprintf("\tMasterLabel: %v\n", t.MasterLabel))
	builder.WriteString(fmt.Sprintf("\tName: %v\n", t.Name))
	builder.WriteString(fmt.Sprintf("\tNamespacePrefix: %v\n", t.NamespacePrefix))
	builder.WriteString(fmt.Sprintf("\tPrecision: %v\n", t.Precision))
	builder.WriteString(fmt.Sprintf("\tQualifiedApiName: %v\n", t.QualifiedApiName))
	builder.WriteString(fmt.Sprintf("\tReferenceTargetField: %v\n", t.ReferenceTargetField))
	builder.WriteString(fmt.Sprintf("\tReferenceTo: %v\n", t.ReferenceTo))
	builder.WriteString(fmt.Sprintf("\tRelationshipName: %v\n", t.RelationshipName))
	builder.WriteString(fmt.Sprintf("\tRelationshipOrder: %v\n", t.RelationshipOrder))
	builder.WriteString(fmt.Sprintf("\tScale: %v\n", t.Scale))
	builder.WriteString(fmt.Sprintf("\tServiceDataTypeId: %v\n", t.ServiceDataTypeId))
	builder.WriteString(fmt.Sprintf("\tValueTypeId: %v\n", t.ValueTypeId))

	return builder.String()
}

type EntityParticleQueryResponse struct {
	BaseQuery
	Records []EntityParticle `json:"Records" force:"records"`
}
