package main

import (
	"fmt"
	"os"
	"strings"
	"text/template"
)

func main() {
	fmt.Println("Welcome to Template")

	name := "Ridhdhish"

	templateString := "Hello {{.}}\n"
	templ, err := template.New("test").Parse(templateString)

	if err != nil {
		panic(err)
	}

	templ.Execute(os.Stdout, name)

	templateString = "{{if .}}Hello\n{{end}}"
	templ, err = template.New("condition").Parse(templateString)
	templ.Execute(os.Stdout, 1)

	templateString = "{{range .}}Hello {{.}}\n{{end}}"
	templ, err = template.New("range").Parse(templateString)
	templ.Execute(os.Stdout, []string{"Naruto", "Sasuke", "Kakashi"})

	templateString = "{{range .}}Hello {{toUpper .}}\n{{end}}"
	funcMap := map[string]interface{}{
		"toUpper": strings.ToUpper,
	}
	templ = template.Must(template.New("funcs").Funcs(funcMap).Parse(templateString))
	templ.Execute(os.Stdout, []string{"Naruto", "Sasuke", "Kakashi"})

}
