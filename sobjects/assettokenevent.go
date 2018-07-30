// This file was generated for SObject AssetTokenEvent, API Version v43.0 at 2018-07-30 03:47:48.71121596 -0400 EDT m=+35.055137759

package sobjects

import (
	"fmt"
	"strings"
)

type AssetTokenEvent struct {
	BaseSObject
	ActorTokenPayload string `force:",omitempty"`
	AssetId           string `force:",omitempty"`
	AssetName         string `force:",omitempty"`
	AssetSerialNumber string `force:",omitempty"`
	ConnectedAppId    string `force:",omitempty"`
	CreatedById       string `force:",omitempty"`
	CreatedDate       string `force:",omitempty"`
	DeviceId          string `force:",omitempty"`
	DeviceKey         string `force:",omitempty"`
	Expiration        string `force:",omitempty"`
	Name              string `force:",omitempty"`
	ReplayId          string `force:",omitempty"`
	UserId            string `force:",omitempty"`
}

func (t *AssetTokenEvent) ApiName() string {
	return "AssetTokenEvent"
}

func (t *AssetTokenEvent) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("AssetTokenEvent #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tActorTokenPayload: %v\n", t.ActorTokenPayload))
	builder.WriteString(fmt.Sprintf("\tAssetId: %v\n", t.AssetId))
	builder.WriteString(fmt.Sprintf("\tAssetName: %v\n", t.AssetName))
	builder.WriteString(fmt.Sprintf("\tAssetSerialNumber: %v\n", t.AssetSerialNumber))
	builder.WriteString(fmt.Sprintf("\tConnectedAppId: %v\n", t.ConnectedAppId))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tDeviceId: %v\n", t.DeviceId))
	builder.WriteString(fmt.Sprintf("\tDeviceKey: %v\n", t.DeviceKey))
	builder.WriteString(fmt.Sprintf("\tExpiration: %v\n", t.Expiration))
	builder.WriteString(fmt.Sprintf("\tName: %v\n", t.Name))
	builder.WriteString(fmt.Sprintf("\tReplayId: %v\n", t.ReplayId))
	builder.WriteString(fmt.Sprintf("\tUserId: %v\n", t.UserId))

	return builder.String()
}

type AssetTokenEventQueryResponse struct {
	BaseQuery
	Records []AssetTokenEvent `json:"Records" force:"records"`
}
