package HTMLPageRenderer

import (
	//	"reflect"
	"testing"
)

func TestRender(t *testing.T) {
	var pageExample HTMLPage
	files := []string{"hello", "abc"}
	pageExample = Page{
		JSFiles: files,
	}
	/*
		if reflect.DeepEqual(pageExample.GetJSFiles(), []string{"hello", "abc"}) {
			t.Error("Did not receive expected JSFiles")
		}
	*/
	_ = Render(pageExample)
}

/*
type Page struct {
	JSFiles []string
}

func (r Page) GetJSFiles() []string {
	return ["abc"]
}
func (r Page) GetCSSFiles() []string {
	return nil
}

func (r Page) Render() string {
	type Item struct {
		ID    int
		Value string
	}
	pageData := struct {
		Title string
		Items []Item
	}{Title: title}

	// Now render the page template
	pageTemplate, err := template.ParseFiles("templates/main.html")
	if err != nil {
		panic("Couldn't parse template file")
	}

	// Render the template to a bytes Buffer, so that we can return
	// the rendered string. We don't have access to the ResponseWriter
	// here.
	pageBuffer := new(bytes.Buffer)
	err = pageTemplate.Execute(pageBuffer, pageData)
	if err != nil {
		panic("Could not execute main template")
	}
	return pageBuffer.String(), nil
}
*/
