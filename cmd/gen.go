package cmd

import (
	"fmt"
	"go/format"
	"os"
	"path"
	"reflect"
	"sort"
	"strings"
	"time"

	"github.com/melan/go-force/force"
	"github.com/melan/go-force/sobjects"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	outputFolder, packageName string
	debug                     bool
	structExists              = struct{}{}
)

type describableSObject struct {
	apiName         string
	externalApiName string
}

func (dso *describableSObject) ApiName() string {
	return dso.apiName
}

func (dso *describableSObject) ExternalIdApiName() string {
	return dso.externalApiName
}

var genCmd = &cobra.Command{
	Use: "gen",
	Run: func(cmd *cobra.Command, args []string) {
		client, err := initClient()
		if err != nil {
			fmt.Println(fmt.Errorf("can't initialize client for Force.com: %v\n", err))
			os.Exit(1)
		}

		var knownObjects map[string]struct{}

		if len(args) == 0 {
			sobjects, err := client.DescribeSObjects()
			if err != nil {
				panic(err)
			}

			knownObjects = make(map[string]struct{}, len(sobjects))
			i := 0
			for _, o := range sobjects {
				knownObjects[o.Name] = structExists
				i++
			}

		} else {
			knownObjects = make(map[string]struct{}, len(args))
			for _, so := range args {
				knownObjects[so] = structExists
			}
		}

		for sObjectName := range knownObjects {
			sod, err := describeSObject(client, sObjectName)
			if err != nil {
				fmt.Printf("Skipping %s - can't get metadata: %v\n", sObjectName, err)
				continue
			}

			if err := generateFile(sod, viper.GetString(ApiVersionKey), outputFolder, packageName); err != nil {
				fmt.Printf("Skipping %s - can't generate file: %v\n", sObjectName, err)
				continue
			}

			fmt.Printf("Successfully generated code for %s\n", sObjectName)
		}
	},
}

func describeSObject(client *force.Api, apiName string) (*force.SObjectDescription, error) {
	sobject := &describableSObject{
		apiName: apiName,
	}

	resp, err := client.DescribeSObject(sobject)
	return resp, err
}

func descriptionToStr(so *force.SObjectDescription) string {
	if so == nil {
		return "nil"
	}

	return fmt.Sprintf("Name: %s\n"+
		"\tKey prefix: %s\n"+
		"\tLabel: %s\n"+
		"\tLabel plural: %s\n"+
		"\tCustom: %v\n"+
		"\tCustom setting: %v\n"+
		"\tActivateable: %t\n"+
		"\tSearchable: %t\n"+
		"\tQueryable: %t\n"+
		"\tDeletable: %t\n"+
		"\tUpdateable: %t\n"+
		"\tCreateable: %t\n"+
		"\tUndeletable: %t\n"+
		"\tMergeable: %t\n"+
		"\tReplicateable: %t\n"+
		"\tTriggerable: %t\n"+
		"\tFeedEnabled: %t\n"+
		"\tRetrievable: %t\n"+
		"\tListviewable: %t\n"+
		"\tLayoutable: %t\n"+
		"\tCompact layoutable: %t\n"+
		"\tSearch layoutable: %t\n"+
		"\tLookup layoutable: %t\n"+
		"\tDeprecatedAndHidden: %t\n"+
		"\tFields:\n%v\n"+
		"\tRecord Type Infos:\n%v\n"+
		"\tChild Relationships:\n%v\n"+
		"\tURLs:\n%v\n"+
		"",
		so.Name,
		so.KeyPrefix,
		so.Label,
		so.LabelPlural,
		so.Custom,
		so.CustomSetting,
		so.Activateable,
		so.Searchable,
		so.Queryable,
		so.Deletable,
		so.Updateable,
		so.Createable,
		so.Undeletable,
		so.Mergeable,
		so.Replicateable,
		so.Triggerable,
		so.FeedEnabled,
		so.Retrievable,
		so.Listviewable,
		so.Layoutable,
		so.CompactLayoutable,
		so.SearchLayoutable,
		so.LookupLayoutable,
		so.DeprecatedAndHidden,
		fieldsToStr(so.Fields, "\t\t"),
		recordTypeInfosToStr(so.RecordTypeInfos, "\t\t"),
		childRelsToStr(so.ChildRelationsips, "\t\t"),
		urlsToStr(so.URLs, "\t\t"),
	)
}

func fieldsToStr(fields []*force.SObjectField, indent string) string {
	serialized := make([]string, len(fields))
	for i, field := range fields {
		str := fmt.Sprintf(indent+"%s (%s:%.2f) = %v", field.Name, field.Type, field.Length, field.DefaultValue)
		serialized[i] = str
	}

	return strings.Join(serialized, "\n")
}

func recordTypeInfosToStr(infos []*force.RecordTypeInfo, indent string) string {
	serialized := make([]string, len(infos))
	for i, info := range infos {
		str := fmt.Sprintf(indent+"%s, available: %t", info.Name, info.Available)
		serialized[i] = str
	}
	return strings.Join(serialized, "\n")
}

func childRelsToStr(childRels []*force.ChildRelationship, indent string) string {
	serialized := make([]string, len(childRels))
	for i, cr := range childRels {
		str := fmt.Sprintf(indent+"%s -> %s (%s)", cr.Field, cr.ChildSObject, cr.RelationshipName)
		serialized[i] = str
	}
	return strings.Join(serialized, "\n")
}

func urlsToStr(urls map[string]string, indent string) string {
	serialized := make([]string, len(urls))
	i := 0
	for k, v := range urls {
		str := fmt.Sprintf(indent+"%s: %s", k, v)
		serialized[i] = str
		i++
	}
	return strings.Join(serialized, "\n")
}

func getGoType(sobjectFieldType string) string {
	switch sobjectFieldType {
	case "base64":
		return "string"
	case "boolean":
		return "bool"
	case "byte":
		return "byte"
	case "date":
		return "string"
	case "dateTime":
		return "string"
	case "double":
		return "float64"
	case "int":
		return "int"
	case "string":
		return "string"
	case "time":
		return "string"
	case "id":
		return "string"
	case "email":
		return "string"
	case "encryptedstring":
		return "string"
	case "reference":
		return "string"
	case "textarea":
		return "string"
	case "url":
		return "string"
	case "address":
		return "*Address"
	default:
		return "string"
	}
}

func generateFile(sod *force.SObjectDescription, apiVersion, outputFolder, packageName string) error {
	if _, err := os.Stat(outputFolder); os.IsNotExist(err) {
		return fmt.Errorf("output directory %s doesn't exists", outputFolder)
	} else if err != nil {
		return fmt.Errorf("error accessing output directory %s: %v", outputFolder, err)
	}

	baseSOType := reflect.TypeOf(sobjects.BaseSObject{})

	fields := make([]*sobjects.SObjectTplFieldData, 0, len(sod.Fields))
	for _, field := range sod.Fields {
		if _, found := baseSOType.FieldByName(field.Name); found {
			// Skip field from base type
			if debug {
				fmt.Printf("skipping %s because it exists in the base class\n", field.Name)
				continue
			}
		}

		fields = append(fields, &sobjects.SObjectTplFieldData{
			Name: field.Name,
			Type: getGoType(field.Type),
		})
	}

	sort.Slice(fields, func(i, j int) bool {
		return fields[i].Name <= fields[j].Name
	})

	tplData := sobjects.SObjectTplData{
		Name:        sod.Name,
		APIVersion:  apiVersion,
		PackageName: packageName,
		Timestamp:   time.Now(),
		Fields:      fields,
	}

	code, err := tplData.ToCode()
	if err != nil {
		return fmt.Errorf("can't generate code: %v", err)
	}

	if debug {
		fmt.Printf("Pre-formatting code for %s:\n%s\n", sod.Name, code)
	}

	res, err := format.Source([]byte(code))
	if err != nil {
		return fmt.Errorf("can't reformat generated code: %v", err)
	}

	outputPath := path.Join(outputFolder, strings.ToLower(sod.Name)+".go")
	if f, err := os.Create(outputPath); err != nil {
		return fmt.Errorf("can't create file %s: %v", outputPath, err)
	} else {
		defer f.Close()

		if _, err := f.Write(res); err != nil {
			return fmt.Errorf("can't write code to file %s: %v", outputPath, err)
		}
	}

	return nil
}

func init() {
	genCmd.Flags().StringVar(&outputFolder, "output", ".", "Output folder for the newly generated types")
	genCmd.Flags().StringVar(&packageName, "package", "generated", "Name of a package to put into the generated files")
	genCmd.Flags().BoolVar(&debug, "debug", false, "Enable debug")
	rootCmd.AddCommand(genCmd)
}
