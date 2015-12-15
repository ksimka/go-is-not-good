package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"text/template"
)

type Entry struct {
	URL        string
	Author     string
	Year       int
	Complaints []string
}

type ByYear []*Entry

func (a ByYear) Len() int {
	return len(a)
}
func (a ByYear) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
func (a ByYear) Less(i, j int) bool {
	return a[i].Year < a[j].Year
}

func main() {
	entries := []*Entry{}
	entryBytes, err := ioutil.ReadFile("entries.json")
	if err != nil {
		fmt.Println("Error reading entries.json:", err)
		fmt.Println(string(entryBytes))
		os.Exit(1)
	}

	err = json.Unmarshal(entryBytes, &entries)
	if err != nil {
		fmt.Println("Error unmarshalling entries:", err)
		fmt.Println(string(entryBytes))
		os.Exit(1)
	}

	sort.Sort(ByYear(entries))

	entryTmpl, err := template.New("entry").Parse(`
# The List

{{ range . }}+ {{ .URL }} ({{ .Author }}, {{ .Year }})
{{ range .Complaints }}  - {{ . }}
{{ end }}{{ end }}
`)
	if err != nil {
		fmt.Println("Error parsing entry template:", err)
		os.Exit(1)
	}
	complaintMap := map[string][]*Entry{}
	for _, e := range entries {
		for _, c := range e.Complaints {
			complaintMap[c] = append(complaintMap[c], e)
		}
	}

	complaintTmpl, err := template.New("complaint").Parse(`
# Complaints

{{ range $complaint, $entries := . }}+ {{ $complaint }}
{{ range $entries }}  - {{ .URL }} ({{ .Author }} {{ .Year }})
{{ end }}{{ end }}
`)
	if err != nil {
		fmt.Println("Error parsing complaint template:", err)
		os.Exit(1)
	}

	f, err := os.Create("README.md")
	if err != nil {
		fmt.Println("Error opening README.md:", err)
		os.Exit(1)
	}
	defer f.Close()

	err = copyContents("HEADER.md", f)
	entryTmpl.Execute(f, entries)
	complaintTmpl.Execute(f, complaintMap)
	err = copyContents("FOOTER.md", f)

}

func copyContents(origin string, target *os.File) error {
	f, err := os.Open(origin)
	if err != nil {
		fmt.Println("Error opening "+origin+":", err)
		os.Exit(1)
	}
	defer f.Close()

	contents, err := ioutil.ReadAll(f)
	if err != nil {
		return err
	}

	_, err = target.Write(contents)
	if err != nil {
		return err
	}

	return nil
}
