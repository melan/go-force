// This file was generated for SObject EntityDefinition, API Version v43.0 at 2018-07-30 03:47:50.05474474 -0400 EDT m=+36.398716954

package sobjects

import (
	"fmt"
	"strings"
)

type EntityDefinition struct {
	BaseSObject
	DataStewardId                string `force:",omitempty"`
	DefaultCompactLayoutId       string `force:",omitempty"`
	DetailUrl                    string `force:",omitempty"`
	DeveloperName                string `force:",omitempty"`
	DurableId                    string `force:",omitempty"`
	EditDefinitionUrl            string `force:",omitempty"`
	EditUrl                      string `force:",omitempty"`
	ExternalSharingModel         string `force:",omitempty"`
	HasSubtypes                  bool   `force:",omitempty"`
	HelpSettingPageName          string `force:",omitempty"`
	HelpSettingPageUrl           string `force:",omitempty"`
	Id                           string `force:",omitempty"`
	InternalSharingModel         string `force:",omitempty"`
	IsApexTriggerable            bool   `force:",omitempty"`
	IsAutoActivityCaptureEnabled bool   `force:",omitempty"`
	IsCompactLayoutable          bool   `force:",omitempty"`
	IsCustomSetting              bool   `force:",omitempty"`
	IsCustomizable               bool   `force:",omitempty"`
	IsDeprecatedAndHidden        bool   `force:",omitempty"`
	IsEverCreatable              bool   `force:",omitempty"`
	IsEverDeletable              bool   `force:",omitempty"`
	IsEverUpdatable              bool   `force:",omitempty"`
	IsFeedEnabled                bool   `force:",omitempty"`
	IsIdEnabled                  bool   `force:",omitempty"`
	IsLayoutable                 bool   `force:",omitempty"`
	IsMruEnabled                 bool   `force:",omitempty"`
	IsProcessEnabled             bool   `force:",omitempty"`
	IsQueryable                  bool   `force:",omitempty"`
	IsReplicateable              bool   `force:",omitempty"`
	IsRetrieveable               bool   `force:",omitempty"`
	IsSearchLayoutable           bool   `force:",omitempty"`
	IsSearchable                 bool   `force:",omitempty"`
	IsSubtype                    bool   `force:",omitempty"`
	IsTriggerable                bool   `force:",omitempty"`
	IsWorkflowEnabled            bool   `force:",omitempty"`
	KeyPrefix                    string `force:",omitempty"`
	Label                        string `force:",omitempty"`
	LastModifiedById             string `force:",omitempty"`
	LastModifiedDate             string `force:",omitempty"`
	MasterLabel                  string `force:",omitempty"`
	NamespacePrefix              string `force:",omitempty"`
	NewUrl                       string `force:",omitempty"`
	PluralLabel                  string `force:",omitempty"`
	PublisherId                  string `force:",omitempty"`
	QualifiedApiName             string `force:",omitempty"`
	RecordTypesSupported         string `force:",omitempty"`
	RunningUserEntityAccessId    string `force:",omitempty"`
}

func (t *EntityDefinition) ApiName() string {
	return "EntityDefinition"
}

func (t *EntityDefinition) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("EntityDefinition #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tDataStewardId: %v\n", t.DataStewardId))
	builder.WriteString(fmt.Sprintf("\tDefaultCompactLayoutId: %v\n", t.DefaultCompactLayoutId))
	builder.WriteString(fmt.Sprintf("\tDetailUrl: %v\n", t.DetailUrl))
	builder.WriteString(fmt.Sprintf("\tDeveloperName: %v\n", t.DeveloperName))
	builder.WriteString(fmt.Sprintf("\tDurableId: %v\n", t.DurableId))
	builder.WriteString(fmt.Sprintf("\tEditDefinitionUrl: %v\n", t.EditDefinitionUrl))
	builder.WriteString(fmt.Sprintf("\tEditUrl: %v\n", t.EditUrl))
	builder.WriteString(fmt.Sprintf("\tExternalSharingModel: %v\n", t.ExternalSharingModel))
	builder.WriteString(fmt.Sprintf("\tHasSubtypes: %v\n", t.HasSubtypes))
	builder.WriteString(fmt.Sprintf("\tHelpSettingPageName: %v\n", t.HelpSettingPageName))
	builder.WriteString(fmt.Sprintf("\tHelpSettingPageUrl: %v\n", t.HelpSettingPageUrl))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tInternalSharingModel: %v\n", t.InternalSharingModel))
	builder.WriteString(fmt.Sprintf("\tIsApexTriggerable: %v\n", t.IsApexTriggerable))
	builder.WriteString(fmt.Sprintf("\tIsAutoActivityCaptureEnabled: %v\n", t.IsAutoActivityCaptureEnabled))
	builder.WriteString(fmt.Sprintf("\tIsCompactLayoutable: %v\n", t.IsCompactLayoutable))
	builder.WriteString(fmt.Sprintf("\tIsCustomSetting: %v\n", t.IsCustomSetting))
	builder.WriteString(fmt.Sprintf("\tIsCustomizable: %v\n", t.IsCustomizable))
	builder.WriteString(fmt.Sprintf("\tIsDeprecatedAndHidden: %v\n", t.IsDeprecatedAndHidden))
	builder.WriteString(fmt.Sprintf("\tIsEverCreatable: %v\n", t.IsEverCreatable))
	builder.WriteString(fmt.Sprintf("\tIsEverDeletable: %v\n", t.IsEverDeletable))
	builder.WriteString(fmt.Sprintf("\tIsEverUpdatable: %v\n", t.IsEverUpdatable))
	builder.WriteString(fmt.Sprintf("\tIsFeedEnabled: %v\n", t.IsFeedEnabled))
	builder.WriteString(fmt.Sprintf("\tIsIdEnabled: %v\n", t.IsIdEnabled))
	builder.WriteString(fmt.Sprintf("\tIsLayoutable: %v\n", t.IsLayoutable))
	builder.WriteString(fmt.Sprintf("\tIsMruEnabled: %v\n", t.IsMruEnabled))
	builder.WriteString(fmt.Sprintf("\tIsProcessEnabled: %v\n", t.IsProcessEnabled))
	builder.WriteString(fmt.Sprintf("\tIsQueryable: %v\n", t.IsQueryable))
	builder.WriteString(fmt.Sprintf("\tIsReplicateable: %v\n", t.IsReplicateable))
	builder.WriteString(fmt.Sprintf("\tIsRetrieveable: %v\n", t.IsRetrieveable))
	builder.WriteString(fmt.Sprintf("\tIsSearchLayoutable: %v\n", t.IsSearchLayoutable))
	builder.WriteString(fmt.Sprintf("\tIsSearchable: %v\n", t.IsSearchable))
	builder.WriteString(fmt.Sprintf("\tIsSubtype: %v\n", t.IsSubtype))
	builder.WriteString(fmt.Sprintf("\tIsTriggerable: %v\n", t.IsTriggerable))
	builder.WriteString(fmt.Sprintf("\tIsWorkflowEnabled: %v\n", t.IsWorkflowEnabled))
	builder.WriteString(fmt.Sprintf("\tKeyPrefix: %v\n", t.KeyPrefix))
	builder.WriteString(fmt.Sprintf("\tLabel: %v\n", t.Label))
	builder.WriteString(fmt.Sprintf("\tLastModifiedById: %v\n", t.LastModifiedById))
	builder.WriteString(fmt.Sprintf("\tLastModifiedDate: %v\n", t.LastModifiedDate))
	builder.WriteString(fmt.Sprintf("\tMasterLabel: %v\n", t.MasterLabel))
	builder.WriteString(fmt.Sprintf("\tNamespacePrefix: %v\n", t.NamespacePrefix))
	builder.WriteString(fmt.Sprintf("\tNewUrl: %v\n", t.NewUrl))
	builder.WriteString(fmt.Sprintf("\tPluralLabel: %v\n", t.PluralLabel))
	builder.WriteString(fmt.Sprintf("\tPublisherId: %v\n", t.PublisherId))
	builder.WriteString(fmt.Sprintf("\tQualifiedApiName: %v\n", t.QualifiedApiName))
	builder.WriteString(fmt.Sprintf("\tRecordTypesSupported: %v\n", t.RecordTypesSupported))
	builder.WriteString(fmt.Sprintf("\tRunningUserEntityAccessId: %v\n", t.RunningUserEntityAccessId))

	return builder.String()
}

type EntityDefinitionQueryResponse struct {
	BaseQuery
	Records []EntityDefinition `json:"Records" force:"records"`
}
