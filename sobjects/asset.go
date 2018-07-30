// This file was generated for SObject Asset, API Version v43.0 at 2018-07-30 03:47:51.755791363 -0400 EDT m=+38.099827407

package sobjects

import (
	"fmt"
	"strings"
)

type Asset struct {
	BaseSObject
	AccountId           string  `force:",omitempty"`
	AssetLevel          int     `force:",omitempty"`
	AssetProvidedById   string  `force:",omitempty"`
	AssetServicedById   string  `force:",omitempty"`
	ContactId           string  `force:",omitempty"`
	CreatedById         string  `force:",omitempty"`
	CreatedDate         string  `force:",omitempty"`
	Description         string  `force:",omitempty"`
	Id                  string  `force:",omitempty"`
	InstallDate         string  `force:",omitempty"`
	IsCompetitorProduct bool    `force:",omitempty"`
	IsDeleted           bool    `force:",omitempty"`
	IsInternal          bool    `force:",omitempty"`
	LastModifiedById    string  `force:",omitempty"`
	LastModifiedDate    string  `force:",omitempty"`
	LastReferencedDate  string  `force:",omitempty"`
	LastViewedDate      string  `force:",omitempty"`
	Name                string  `force:",omitempty"`
	OwnerId             string  `force:",omitempty"`
	ParentId            string  `force:",omitempty"`
	Price               string  `force:",omitempty"`
	Product2Id          string  `force:",omitempty"`
	ProductCode         string  `force:",omitempty"`
	PurchaseDate        string  `force:",omitempty"`
	Quantity            float64 `force:",omitempty"`
	RootAssetId         string  `force:",omitempty"`
	SerialNumber        string  `force:",omitempty"`
	Status              string  `force:",omitempty"`
	StockKeepingUnit    string  `force:",omitempty"`
	SystemModstamp      string  `force:",omitempty"`
	UsageEndDate        string  `force:",omitempty"`
}

func (t *Asset) ApiName() string {
	return "Asset"
}

func (t *Asset) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("Asset #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tAccountId: %v\n", t.AccountId))
	builder.WriteString(fmt.Sprintf("\tAssetLevel: %v\n", t.AssetLevel))
	builder.WriteString(fmt.Sprintf("\tAssetProvidedById: %v\n", t.AssetProvidedById))
	builder.WriteString(fmt.Sprintf("\tAssetServicedById: %v\n", t.AssetServicedById))
	builder.WriteString(fmt.Sprintf("\tContactId: %v\n", t.ContactId))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tDescription: %v\n", t.Description))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tInstallDate: %v\n", t.InstallDate))
	builder.WriteString(fmt.Sprintf("\tIsCompetitorProduct: %v\n", t.IsCompetitorProduct))
	builder.WriteString(fmt.Sprintf("\tIsDeleted: %v\n", t.IsDeleted))
	builder.WriteString(fmt.Sprintf("\tIsInternal: %v\n", t.IsInternal))
	builder.WriteString(fmt.Sprintf("\tLastModifiedById: %v\n", t.LastModifiedById))
	builder.WriteString(fmt.Sprintf("\tLastModifiedDate: %v\n", t.LastModifiedDate))
	builder.WriteString(fmt.Sprintf("\tLastReferencedDate: %v\n", t.LastReferencedDate))
	builder.WriteString(fmt.Sprintf("\tLastViewedDate: %v\n", t.LastViewedDate))
	builder.WriteString(fmt.Sprintf("\tName: %v\n", t.Name))
	builder.WriteString(fmt.Sprintf("\tOwnerId: %v\n", t.OwnerId))
	builder.WriteString(fmt.Sprintf("\tParentId: %v\n", t.ParentId))
	builder.WriteString(fmt.Sprintf("\tPrice: %v\n", t.Price))
	builder.WriteString(fmt.Sprintf("\tProduct2Id: %v\n", t.Product2Id))
	builder.WriteString(fmt.Sprintf("\tProductCode: %v\n", t.ProductCode))
	builder.WriteString(fmt.Sprintf("\tPurchaseDate: %v\n", t.PurchaseDate))
	builder.WriteString(fmt.Sprintf("\tQuantity: %v\n", t.Quantity))
	builder.WriteString(fmt.Sprintf("\tRootAssetId: %v\n", t.RootAssetId))
	builder.WriteString(fmt.Sprintf("\tSerialNumber: %v\n", t.SerialNumber))
	builder.WriteString(fmt.Sprintf("\tStatus: %v\n", t.Status))
	builder.WriteString(fmt.Sprintf("\tStockKeepingUnit: %v\n", t.StockKeepingUnit))
	builder.WriteString(fmt.Sprintf("\tSystemModstamp: %v\n", t.SystemModstamp))
	builder.WriteString(fmt.Sprintf("\tUsageEndDate: %v\n", t.UsageEndDate))

	return builder.String()
}

type AssetQueryResponse struct {
	BaseQuery
	Records []Asset `json:"Records" force:"records"`
}
