package HTMLPageRenderer

import (
	"bytes"
	"html/template"
	"log"
)

type Page struct {
	Title    string
	JSFiles  []string
	CSSFiles []string
	Body     string
}

func (r Page) GetTitle() string {
	return r.Title
}
func (r Page) GetJSFiles() []string {
	return r.JSFiles
}
func (r Page) GetCSSFiles() []string {
	return r.CSSFiles
}

func (r Page) GetBody() string {
	return r.Body
}

func Render(p HTMLPage) string {
	// Now render the page template
	const tpl = `
<!DOCTYPE html>
<html>
<head>
	<title>{{ .Title}} </title>
	{{range .JSFiles}}
		<script src="{{.}}"></script>{{end}}
	{{range .CSSFiles}}
		<link rel="stylesheet" href="{{.}}"></script>{{end}}
</head>
<body>
{{ .Body }} 
</body>
</html>`
	pageTemplate, err := template.New("main").Parse(tpl)
	if err != nil {
		log.Print(err)
		//panic("Couldn't parse template file")
	}

	pageData := struct {
		Title    string
		JSFiles  []string
		CSSFiles []string
		Body     template.HTML
	}{Title: p.GetTitle(),
		JSFiles:  p.GetJSFiles(),
		CSSFiles: p.GetCSSFiles(),
		Body:     template.HTML(p.GetBody()),
	}
	// Render the template to a bytes Buffer, so that we can return
	// the rendered string. We don't have access to the ResponseWriter
	// here.
	pageBuffer := new(bytes.Buffer)
	err = pageTemplate.Execute(pageBuffer, pageData)
	if err != nil {
		log.Print(err)
		//panic("Could not execute main template")
	}
	return pageBuffer.String()
}
