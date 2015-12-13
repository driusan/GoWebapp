package HTMLPageRenderer

import (
	"bytes"
	"html/template"
)

type Page struct {
	JSFiles []string
}

func (r Page) GetJSFiles() []string {
	return []string{ "abc", "def" }
}
func (r Page) GetCSSFiles() []string {
	return nil
}

func (r Page) Render() string {
	// Now render the page template
	pageTemplate, err := template.ParseFiles("main.html")
	if err != nil {
		panic("Couldn't parse template file")
	}

	// Render the template to a bytes Buffer, so that we can return
	// the rendered string. We don't have access to the ResponseWriter
	// here.
	pageBuffer := new(bytes.Buffer)
	err = pageTemplate.Execute(pageBuffer, r)
	if err != nil {
		panic("Could not execute main template")
	}
	return pageBuffer.String()
}
