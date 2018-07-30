// This file was generated for SObject ClientBrowser, API Version v43.0 at 2018-07-30 03:47:47.638587708 -0400 EDT m=+33.982469258

package sobjects

import (
	"fmt"
	"strings"
)

type ClientBrowser struct {
	BaseSObject
	CreatedDate   string `force:",omitempty"`
	FullUserAgent string `force:",omitempty"`
	Id            string `force:",omitempty"`
	LastUpdate    string `force:",omitempty"`
	ProxyInfo     string `force:",omitempty"`
	UsersId       string `force:",omitempty"`
}

func (t *ClientBrowser) ApiName() string {
	return "ClientBrowser"
}

func (t *ClientBrowser) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("ClientBrowser #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tFullUserAgent: %v\n", t.FullUserAgent))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tLastUpdate: %v\n", t.LastUpdate))
	builder.WriteString(fmt.Sprintf("\tProxyInfo: %v\n", t.ProxyInfo))
	builder.WriteString(fmt.Sprintf("\tUsersId: %v\n", t.UsersId))

	return builder.String()
}

type ClientBrowserQueryResponse struct {
	BaseQuery
	Records []ClientBrowser `json:"Records" force:"records"`
}
