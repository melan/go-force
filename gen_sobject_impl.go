//+build ignore

// This code generates a list of known implementations of SObject interface. It scans files only in the sobjects package
package main

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"sort"
	"text/template"
)

var (
	re = regexp.MustCompile(`^func \(\w+ \*?([A-Z][^)]+)\) ApiName\(\)`)
)

// inspired by: https://stackoverflow.com/questions/14668850/list-directory-in-go
func walkDir(root string) ([]string, error) {
	var files []string
	err := filepath.Walk(root, func(filePath string, info os.FileInfo, err error) error {
		if !info.IsDir() && path.Ext(filePath) == ".go" {
			fmt.Printf("Found file %s\n", filePath)
			files = append(files, filePath)
		}
		return nil
	})

	return files, err
}

func scanFile(re *regexp.Regexp, filePath string) ([]string, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return []string{}, fmt.Errorf("can't open: %v", err)
	}
	defer f.Close()

	foundClasses := make([]string, 0)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		match := re.FindStringSubmatch(line)
		if match != nil && len(match) > 1 {
			fmt.Printf("Found implementation %s in file %s\n", match[1], filePath)
			foundClasses = append(foundClasses, match[1])
		}
	}
	if err := scanner.Err(); err != nil {
		return []string{}, fmt.Errorf("error while scanning: %v", err)
	}

	return foundClasses, nil
}

func main() {
	basicFile := os.Getenv("GOFILE")
	if basicFile == "" {
		panic("GOFILE env variable must be set. Please run this command by calling `go generate ./...`")
	}

	baseDir := path.Dir(basicFile)
	files, err := walkDir(baseDir)
	if err != nil {
		panic(fmt.Errorf("can't scan %s for files: %v", baseDir, err))
	}

	foundClasses := make(map[string]struct{}, 0)
	for _, filePath := range files {
		newClasses, err := scanFile(re, filePath)
		if err != nil {
			panic(fmt.Sprintf("can't parse file %s: %v", filePath, err))
		}

		for _, nc := range newClasses {
			foundClasses[nc] = struct{}{}
		}
	}

	classesList := make([]string, 0, len(foundClasses))
	for fc, _ := range foundClasses {
		classesList = append(classesList, fc)
	}

	sort.Slice(classesList, func(i, j int) bool {
		return classesList[i] < classesList[j]
	})

	outputPath := path.Join(baseDir, "sobject_impls.go")
	f, err := os.Create(outputPath)
	if err != nil {
		panic(fmt.Sprintf("can't open file %s to write generated list: %v", outputPath, err))
	}
	defer f.Close()

	implementationsTemplate.Execute(f, struct {
		Impl []string
	}{
		Impl: classesList,
	})
}

var implementationsTemplate = template.Must(template.New("").Parse(`
package sobjects

var SObjectsImplementations = map[string]SObject {
	{{- range .Impl }}
	{{ printf "%q" . }}: &{{ . }}{},
	{{- end }}
}
`))
