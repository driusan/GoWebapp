package HTMLPageRenderer

type ReactPage struct {
	Title       string
	RootElement string
	JSFiles     []string
	CSSFiles    []string
}

func (r ReactPage) GetTitle() string {
	return r.Title
}
func (r ReactPage) GetJSFiles() []string {
	return r.JSFiles
}
func (r ReactPage) GetCSSFiles() []string {
	return r.CSSFiles
}

func (r ReactPage) GetBody() string {
	return `<div id="workspace"></div>
		<script>
			var R` + r.RootElement + ` = React.createFactory(` + r.RootElement + `);
			ReactDOM.render(R` + r.RootElement + `({}), document.getElementById("workspace"));
		</script>`
}
