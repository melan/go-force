package sobjects

import (
	"fmt"
	"strings"
	"text/template"
	"time"
)

type SObjectTplFieldData struct {
	Name string
	Type string
}

type SObjectTplData struct {
	Name        string
	APIVersion  string
	Timestamp   time.Time
	PackageName string
	Fields      []*SObjectTplFieldData
}

func (sofd *SObjectTplFieldData) ToCode() string {
	return fmt.Sprintf("%s %s `force:\",omitempty\"`", sofd.Name, sofd.Type)
}

func (sod *SObjectTplData) ToCode() (string, error) {
	writer := &strings.Builder{}

	if err := sObjectTemplate.Execute(writer, sod); err != nil {
		return "", fmt.Errorf("can't fill template: %v", err)
	}
	return writer.String(), nil
}

var sObjectTemplate = template.Must(template.New("").Parse(`
// This file was generated for SObject {{ .Name }}, API Version {{ .APIVersion }} at {{ .Timestamp }}

package {{ .PackageName }}

import (
	"fmt"
	"strings"
)

type {{ .Name }} struct {
	BaseSObject
	{{- range .Fields }}
	{{ printf "%s" .ToCode }}
	{{- end }}
}

func (t *{{ .Name }}) ApiName() string {
	return "{{ .Name }}"
}

func (t *{{ .Name }}) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("{{ .Name }} #%s - %s\n", t.Id, t.Name))
	{{- range .Fields }}
	builder.WriteString(fmt.Sprintf("\t{{ .Name }}: %v\n", t.{{ .Name }}))
	{{- end }}
	
	return builder.String()
}

type {{ .Name }}QueryResponse struct {
	BaseQuery
	Records []{{ .Name }} ` + "`" + `json:"Records" force:"records"` + "`" + `
}
`))
