// This file was generated for SObject Case, API Version v43.0 at 2018-07-30 03:47:50.19712096 -0400 EDT m=+36.541098516

package sobjects

import (
	"fmt"
	"strings"
)

type Case struct {
	BaseSObject
	AccountId               string `force:",omitempty"`
	AssetId                 string `force:",omitempty"`
	CaseNumber              string `force:",omitempty"`
	ClosedDate              string `force:",omitempty"`
	Comments                string `force:",omitempty"`
	ContactEmail            string `force:",omitempty"`
	ContactFax              string `force:",omitempty"`
	ContactId               string `force:",omitempty"`
	ContactMobile           string `force:",omitempty"`
	ContactPhone            string `force:",omitempty"`
	CreatedById             string `force:",omitempty"`
	CreatedDate             string `force:",omitempty"`
	Description             string `force:",omitempty"`
	EngineeringReqNumber__c string `force:",omitempty"`
	Id                      string `force:",omitempty"`
	IsClosed                bool   `force:",omitempty"`
	IsDeleted               bool   `force:",omitempty"`
	IsEscalated             bool   `force:",omitempty"`
	LastModifiedById        string `force:",omitempty"`
	LastModifiedDate        string `force:",omitempty"`
	LastReferencedDate      string `force:",omitempty"`
	LastViewedDate          string `force:",omitempty"`
	Origin                  string `force:",omitempty"`
	OwnerId                 string `force:",omitempty"`
	ParentId                string `force:",omitempty"`
	PotentialLiability__c   string `force:",omitempty"`
	Priority                string `force:",omitempty"`
	Product__c              string `force:",omitempty"`
	Reason                  string `force:",omitempty"`
	SLAViolation__c         string `force:",omitempty"`
	Status                  string `force:",omitempty"`
	Subject                 string `force:",omitempty"`
	SuppliedCompany         string `force:",omitempty"`
	SuppliedEmail           string `force:",omitempty"`
	SuppliedName            string `force:",omitempty"`
	SuppliedPhone           string `force:",omitempty"`
	SystemModstamp          string `force:",omitempty"`
	Type                    string `force:",omitempty"`
}

func (t *Case) ApiName() string {
	return "Case"
}

func (t *Case) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("Case #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tAccountId: %v\n", t.AccountId))
	builder.WriteString(fmt.Sprintf("\tAssetId: %v\n", t.AssetId))
	builder.WriteString(fmt.Sprintf("\tCaseNumber: %v\n", t.CaseNumber))
	builder.WriteString(fmt.Sprintf("\tClosedDate: %v\n", t.ClosedDate))
	builder.WriteString(fmt.Sprintf("\tComments: %v\n", t.Comments))
	builder.WriteString(fmt.Sprintf("\tContactEmail: %v\n", t.ContactEmail))
	builder.WriteString(fmt.Sprintf("\tContactFax: %v\n", t.ContactFax))
	builder.WriteString(fmt.Sprintf("\tContactId: %v\n", t.ContactId))
	builder.WriteString(fmt.Sprintf("\tContactMobile: %v\n", t.ContactMobile))
	builder.WriteString(fmt.Sprintf("\tContactPhone: %v\n", t.ContactPhone))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tDescription: %v\n", t.Description))
	builder.WriteString(fmt.Sprintf("\tEngineeringReqNumber__c: %v\n", t.EngineeringReqNumber__c))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tIsClosed: %v\n", t.IsClosed))
	builder.WriteString(fmt.Sprintf("\tIsDeleted: %v\n", t.IsDeleted))
	builder.WriteString(fmt.Sprintf("\tIsEscalated: %v\n", t.IsEscalated))
	builder.WriteString(fmt.Sprintf("\tLastModifiedById: %v\n", t.LastModifiedById))
	builder.WriteString(fmt.Sprintf("\tLastModifiedDate: %v\n", t.LastModifiedDate))
	builder.WriteString(fmt.Sprintf("\tLastReferencedDate: %v\n", t.LastReferencedDate))
	builder.WriteString(fmt.Sprintf("\tLastViewedDate: %v\n", t.LastViewedDate))
	builder.WriteString(fmt.Sprintf("\tOrigin: %v\n", t.Origin))
	builder.WriteString(fmt.Sprintf("\tOwnerId: %v\n", t.OwnerId))
	builder.WriteString(fmt.Sprintf("\tParentId: %v\n", t.ParentId))
	builder.WriteString(fmt.Sprintf("\tPotentialLiability__c: %v\n", t.PotentialLiability__c))
	builder.WriteString(fmt.Sprintf("\tPriority: %v\n", t.Priority))
	builder.WriteString(fmt.Sprintf("\tProduct__c: %v\n", t.Product__c))
	builder.WriteString(fmt.Sprintf("\tReason: %v\n", t.Reason))
	builder.WriteString(fmt.Sprintf("\tSLAViolation__c: %v\n", t.SLAViolation__c))
	builder.WriteString(fmt.Sprintf("\tStatus: %v\n", t.Status))
	builder.WriteString(fmt.Sprintf("\tSubject: %v\n", t.Subject))
	builder.WriteString(fmt.Sprintf("\tSuppliedCompany: %v\n", t.SuppliedCompany))
	builder.WriteString(fmt.Sprintf("\tSuppliedEmail: %v\n", t.SuppliedEmail))
	builder.WriteString(fmt.Sprintf("\tSuppliedName: %v\n", t.SuppliedName))
	builder.WriteString(fmt.Sprintf("\tSuppliedPhone: %v\n", t.SuppliedPhone))
	builder.WriteString(fmt.Sprintf("\tSystemModstamp: %v\n", t.SystemModstamp))
	builder.WriteString(fmt.Sprintf("\tType: %v\n", t.Type))

	return builder.String()
}

type CaseQueryResponse struct {
	BaseQuery
	Records []Case `json:"Records" force:"records"`
}
