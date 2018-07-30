// This file was generated for SObject DataType, API Version v43.0 at 2018-07-30 03:47:40.350580793 -0400 EDT m=+26.694188868

package sobjects

import (
	"fmt"
	"strings"
)

type DataType struct {
	BaseSObject
	ContextServiceDataTypeId string `force:",omitempty"`
	ContextWsdlDataTypeId    string `force:",omitempty"`
	DeveloperName            string `force:",omitempty"`
	DurableId                string `force:",omitempty"`
	Id                       string `force:",omitempty"`
	IsComplex                bool   `force:",omitempty"`
}

func (t *DataType) ApiName() string {
	return "DataType"
}

func (t *DataType) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("DataType #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tContextServiceDataTypeId: %v\n", t.ContextServiceDataTypeId))
	builder.WriteString(fmt.Sprintf("\tContextWsdlDataTypeId: %v\n", t.ContextWsdlDataTypeId))
	builder.WriteString(fmt.Sprintf("\tDeveloperName: %v\n", t.DeveloperName))
	builder.WriteString(fmt.Sprintf("\tDurableId: %v\n", t.DurableId))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tIsComplex: %v\n", t.IsComplex))

	return builder.String()
}

type DataTypeQueryResponse struct {
	BaseQuery
	Records []DataType `json:"Records" force:"records"`
}
